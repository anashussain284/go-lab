package main

import (
	"bufio"
	"fmt"
	"local/003-configuration-loader/pkg/validator"
	"log"
	"os"
	"strings"
)

func main() {
	HOST := "localhost"
	PORT := 8080
	DEBUG := false

	filePath := os.Args[1]

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "HOST") {
			host, err := validator.ValidateHost(line)

			if err != nil {
				log.Fatalf("Error: HOST validation: %v\n", err)
			}

			HOST = host
		} else if strings.Contains(line, "PORT") {

			port, err := validator.ValidatePort(line)

			if err != nil {
				log.Fatalf("Error: PORT validation: %v\n", err)
			}

			PORT = port

		} else if strings.Contains(line, "DEBUG") {

			// fmt.Printf("line: %v\n", line)

			debug, err := validator.ValidateDebug(line)

			if err != nil {
				log.Fatalf("Error: DEBUG validation: %v\n", err)
			}

			DEBUG = debug
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
