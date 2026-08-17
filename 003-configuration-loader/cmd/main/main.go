package main

import (
	"bufio"
	"fmt"
	"local/003-configuration-loader/pkg/config"
	"local/003-configuration-loader/pkg/validator"
	"log"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Error: application require a .env file path")
	}

	filePath := os.Args[1]

	fileInfo, err := os.Stat(filePath)

	if err != nil {
		log.Fatal(err)
	}

	if fileInfo.IsDir() {
		log.Fatalf("Error: application not accept directory")
	}

	if filepath.Ext(filePath) != ".env" {
		log.Fatalf("Error: application only accept .env file")
	}

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cfg := config.Configuration{
		Host:  "localhost",
		Port:  8080,
		Debug: true,
	}

	var configuration config.Configuration

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		configuration, err = validator.RetrieveKeyValue(line, &cfg)

		if err != nil {
			log.Fatal(err)
		}

	}

	err = scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Configuration Loaded:\n\n")
	fmt.Printf("HOST:\t%v\n", configuration.Host)
	fmt.Printf("PORT:\t%v\n", configuration.Port)
	fmt.Printf("DEBUG:\t%v\n", configuration.Debug)
}
