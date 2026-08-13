package main

import (
	"flag"
	"local/002-log-analyzer/pkg/result"
	"local/002-log-analyzer/pkg/scanner"
	"local/002-log-analyzer/pkg/parser"
	"log"
	"os"
	"slices"
	"time"
	"fmt"
	// "sort"
)

func main2() {
	data := []byte("Hello, Go!\nThis is a quick way to write files.")
	err := os.WriteFile("output.txt", data, 644)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File writed successfull")
}

func main() {
	root := flag.String("root", "", "directory containing log files")
	fromStr := flag.String("from", "", "analyze logs from this date (YYYY-MM-DD)")
	toStr := flag.String("to", "", "analyze logs until this date (YYYY-MM-DD)")
	searchStr := flag.String("search", "", "search text inside log entries")
	levelStr := flag.String("level", "", "filter by level: INFO, WARN, ERROR")

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
		date, err := parser.ParseDate(*fromStr)

		if err != nil {
			log.Fatalf("error: invalid 'from' date format (%v), expected YYYY-MM-DD", *fromStr)
		}
		fromDate = &date
	}

	if *toStr != "" {
		date, err := parser.ParseDate(*toStr)

		if err != nil {
			log.Fatalf("error: invalid 'to' date format (%v), expected YYYY-MM-DD", *toStr)
		}
		toDate = &date
	}

	if fromDate != nil && toDate != nil && fromDate.After(*toDate) {
		log.Fatalf("error: 'from' date (%v) cannot be after 'to' date (%v)", *fromStr, *toStr)
	}

	level := *levelStr

	if level != "" {
		logLevelSlice := []string{"INFO","WARN","ERROR"}
		hasLevel := slices.Contains(logLevelSlice, level)

		if !hasLevel {
			log.Fatalf("error: 'invalid level' (%v), expected 'INFO' or 'WARN' or 'ERROR'", level)
		}
	}

	analytics, err := scanner.Scan(rootPath, fromDate, toDate, *searchStr, level)

	if err != nil {
		log.Fatal(err)
	}

	result.PrintResult(analytics, rootPath)
}