package main

import (
	// "fmt"
	"flag"
	"log"
	"os"
	"time"
	// "path/filepath"
	"local/002-log-analyzer/pkg/scanner"
	"local/002-log-analyzer/pkg/result"
)

func main() {
	root := flag.String("root", "", "directory containing log files")
	fromStr := flag.String("from", "", "analyze logs from this date (YYYY-MM-DD)")
	toStr := flag.String("to", "", "analyze logs until this date (YYYY-MM-DD)")

	flag.Parse()

	rootPath := *root

	if rootPath == "" {
		log.Fatalf("error: root directory is required")
	}

	fileInfo, err := os.Stat(rootPath)

	if err != nil {
		log.Fatalf("error: cannot access %v : %v", rootPath, err)
	}

	if !fileInfo.IsDir() {
		log.Fatalf("error: %v is not a valid directory", rootPath)
	}

	entries, err := os.ReadDir(rootPath)

	if err != nil {
		log.Fatal(err)
	}

	if len(entries) == 0 {
		log.Fatalf("warn: %v is an empty directory", rootPath)
	}

	var fromDate, toDate *time.Time

	if *fromStr != "" {
		date, err := parseDate(*fromStr)

		if err != nil {
			log.Fatalf("error: invalid 'from' date format (%v), expected YYYY-MM-DD", *fromStr)
		}
		fromDate = &date
	}

	if *toStr != "" {
		date, err := parseDate(*toStr)

		if err != nil {
			log.Fatalf("error: invalid 'to' date format (%v), expected YYYY-MM-DD", *toStr)
		}
		toDate = &date
	}

	// Validate date range logic (from cannot be after to)
	if fromDate != nil && toDate != nil && fromDate.After(*toDate) {
		log.Fatalf("error: 'from' date (%v) cannot be after 'to' date (%v)", *fromStr, *toStr)
	}

	analytics, err := scanner.Scan(rootPath, fromDate, toDate)

	if err != nil {
		log.Fatal(err)
	}

	result.PrintResult(analytics, rootPath)
}

func parseDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
}