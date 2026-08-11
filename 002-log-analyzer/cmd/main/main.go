package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	// "path/filepath"
	"local/002-log-analyzer/pkg/scanner"
	"local/002-log-analyzer/pkg/result"
)

func main() {
	root := flag.String("root", "", "directory containing log files")

	flag.Parse()

	rootPath := *root

	if rootPath == "" {
		log.Fatalf("error: root directory is required")
	}

	/**
	 * check below
	 * 1. the directory exist
	 * 2. empty directory
	 * 3. else do the job
	 */
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

	fmt.Println(rootPath)

	analytics, err := scanner.Scan(rootPath)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(analytics)

	result.PrintResult(analytics)

}