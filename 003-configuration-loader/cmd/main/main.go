package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	HOST := "localhost"
	PORT := "8080"
	DEBUG := "false"

	filePath := os.Args[1]

	// fmt.Println(filePath)

	// fileContent, err := os.ReadFile(filePath)

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "HOST") {
			searchKey := "HOST"
			index := strings.Index(line, searchKey)
			searchValue := strings.TrimSpace(line[index+(len(searchKey))+1:])

			if searchValue != "" {
				HOST = searchValue
			}

		} else if strings.Contains(line, "PORT") {
			searchKey := "PORT"
			index := strings.Index(line, searchKey)
			searchValue := strings.TrimSpace(line[index+(len(searchKey))+1:])

			if searchValue != "" {
				PORT = searchValue
			}

		} else if strings.Contains(line, "DEBUG") {
			searchKey := "DEBUG"
			index := strings.Index(line, searchKey)
			searchValue := strings.TrimSpace(line[index+(len(searchKey))+1:])

			if searchValue != "" {
				DEBUG = searchValue
			}
		}
	}

	err = scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Configuration Loaded:\n\n")
	fmt.Printf("HOST: %v\n", HOST)
	fmt.Printf("PORT: %v\n", PORT)
	fmt.Printf("DEBUG: %v\n", DEBUG)
}
