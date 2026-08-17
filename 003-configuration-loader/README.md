# 003-configuration-loader

## Project overview

The project is Configuration Loader, It is for learning and practicing go functionalities purpose.

## Project Phases

### Phase - 1
- Print default configuration in CLI

### Phase - 2
- Accept config file from CLI
- Read and assign config-value to the config-key if config-value not empty

### Phase - 3
- Convert string to int use `strconv.Atoi()` package
- Convert string to bool use `strconv.ParseBool()` package
### Phase - 4
- Validate HOST
- Validate PORT (must be integer between 1 and 65535)
- Validate DEBUG (must be bool)

### Phase - 5 (Refactoring)
- Remove redundant file validations and checking
- Change data type to store config values for easy maintain

## How to run it

```
go run ./cmd/main/main.go