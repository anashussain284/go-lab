package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	HOST := "localhost"
	PORT := 8080
	DEBUG := false

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
				port, err := strconv.Atoi(searchValue)

				if err != nil {
					log.Fatalf("Error: string to int conversion: %v", err)
				}
				PORT = port
			}

		} else if strings.Contains(line, "DEBUG") {
			searchKey := "DEBUG"
			index := strings.Index(line, searchKey)
			searchValue := strings.TrimSpace(line[index+(len(searchKey))+1:])

			if searchValue != "" {
				debug, err := strconv.ParseBool(searchValue)

				if err != nil {
					log.Fatalf("Error: string to bool conversion: %v", err)
				}

				DEBUG = debug
			}
		}
	}

	err = scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Configuration Loaded:\n\n")
	fmt.Printf("Type: %T, HOST: %v\n", HOST, HOST)
	fmt.Printf("Type: %T, PORT: %v\n", PORT, PORT)
	fmt.Printf("Type: %T, DEBUG: %v\n", DEBUG, DEBUG)
}
