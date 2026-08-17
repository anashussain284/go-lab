package validator

import (
	"errors"
	"strconv"
	"strings"
)

func ValidateDebug(line string) (bool, error) {
	searchKey := "DEBUG"

	index := strings.Index(line, searchKey)

	if index == -1 {
		return false, errors.New("DEBUG not found in config file")
	}

	valueStart := index + (len(searchKey)) + 1

	if valueStart > len(line) {
		return false, errors.New("DEBUG config is malformed or missing a value separator")
	}

	searchValue := strings.TrimSpace(line[valueStart:])

	if searchValue == "" {
		return false, errors.New("DEBUG cannot be empty in config file")
	}

	debug, err := strconv.ParseBool(searchValue)

	if err != nil {
		return false, errors.New("DEBUG must be boolean value")
	}

	return debug, nil
}

func ValidatePort(line string) (int, error) {
	searchKey := "PORT"
	index := strings.Index(line, searchKey)

	if index == -1 {
		return 0, errors.New("PORT not found in config file")
	}

	valueStart := index + (len(searchKey)) + 1

	if valueStart > len(line) {
		return 0, errors.New("PORT config is malformed or missing a value separator")
	}

	searchValue := strings.TrimSpace(line[valueStart:])

	if searchValue == "" {
		return 0, errors.New("PORT cannot be empty in config-file")
	}

	port, err := strconv.Atoi(searchValue)

	if err != nil {
		return 0, err
	}

	if port < 1 || port > 65535 {
		return 0, errors.New("PORT must be between 1 and 65535")
	}

	return port, nil
}

func ValidateHost(line string) (string, error) {
	searchKey := "HOST"
	index := strings.Index(line, searchKey)

	if index == -1 {
		return "", errors.New("HOST not found in config file")
	}

	valueStart := index + index + (len(searchKey)) + 1

	if valueStart > len(line) {
		return "", errors.New("HOST config is malformed or missing a value separator")
	}

	searchValue := strings.TrimSpace(line[valueStart:])

	if searchValue == "" {
		return "", errors.New("HOST cannot be empty in config-file")
	}

	return searchValue, nil
}

func ValidateLine(line string) (bool, error) {
	if !strings.Contains(line, "=") {
		return false, errors.New("Invalid value assigning")
	}

	equalIndex := strings.Index(line, "=")
	beforeEqual := line[:equalIndex]
	lenBeforeEqual := len(beforeEqual)

	if equalIndex < 0 || lenBeforeEqual < 1 {
		return false, errors.New("Invalid equal position")
	}

	return true, nil
}
