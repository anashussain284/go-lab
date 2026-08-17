package validator

import (
	"errors"
	"local/003-configuration-loader/pkg/config"
	"slices"
	"strconv"
	"strings"
)

func ValidateDebug(debugStr string) (bool, error) {
	debug, err := strconv.ParseBool(debugStr)

	if err != nil {
		return false, errors.New("DEBUG must be boolean value")
	}

	return debug, nil
}

func ValidatePort(portStr string) (int, error) {
	port, err := strconv.Atoi(portStr)

	if err != nil {
		return 0, errors.New("PORT must be integer value")
	}

	if port < 1 || port > 65535 {
		return 0, errors.New("PORT must be between 1 and 65535")
	}

	return port, nil
}

func RetrieveKeyValue(line string, configuration *config.Configuration) error {
	if line == "" {
		return errors.New("Empty line contained config file not accept")
	}

	equalIndex := strings.Index(line, "=")

	if equalIndex == -1 {
		return errors.New("Key:Value seperator (=) not found")
	}

	expectedKey := line[0:equalIndex]

	if expectedKey == "" {
		return errors.New("Key should not empty")
	}

	expectedValue := line[equalIndex+1:]

	whitelistedKeys := []string{"HOST", "PORT", "DEBUG"}
	hasKey := slices.Contains(whitelistedKeys, expectedKey)

	if !hasKey {
		return errors.New("Whitelisted key not found")
	}

	if expectedValue == "" {
		return errors.New("Value should not empty")
	}

	hasWhiteSpaceInValue := strings.Contains(expectedValue, " ")

	if hasWhiteSpaceInValue {
		expectedValue = strings.TrimSpace(expectedValue)
	}

	switch expectedKey {
	case "HOST":
		configuration.Host = expectedValue
	case "PORT":
		port, err := ValidatePort(expectedValue)
		if err != nil {
			return err
		}
		configuration.Port = port
	case "DEBUG":
		debug, err := ValidateDebug(expectedValue)
		if err != nil {
			return err
		}
		configuration.Debug = debug
	}

	return nil
}
