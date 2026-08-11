package main

import (
	"fmt"
	"log"
	"os"
	"local/006-file-statistics/pkg/scanner"
	"local/006-file-statistics/pkg/output"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: directory path is required, eg: /var/www/html")
		os.Exit(1)
	}

	path := os.Args[1]

	fileInfo, err := os.Stat(path)

	if err != nil {
		log.Fatalf("error: cannot access %q: %v", path, err)
	}

	if !fileInfo.IsDir() {
		log.Fatalf("error: path '%v' is not a directory", path)
	}

	statistics, err := scanner.Scan(path)

	if err != nil {
		log.Fatalf("error: directory scanning failed: %v", err)
	}

	output.PrintStatistics(statistics)
}