# 002-log-analyzer

## Project overview

This is a GO project for learning and practicing go functionalities for build project. It is used to get the analysiz of .log files conteined in a directory

## Project Phases

### Phase - 1
- CLI only accept valid directory
- CLI not accept empty directory
- Scan full directory only consider `.log` files and do the following
	- Count number files analyzed
	- Count number of directories analyzed

### Phase - 2
- Level wise log counting

### Phase - 3
- Filter data between fromDate and toDate via CLI, this is optional value
- if date found it will validate following
	- date should be follow YYYY-MM-DD format
	- fromDate should be before toDate
- Make decision via function is analyzable or not
	- Extract date from line
	- If date found in the line time.Parse it into following format "Day Mon DD HH:MM:SS YYYY" and return (time, true) else return (nil, false)
	- If above function return true
		- Check the fromDate and toDate is withinDate range
			- parsed time format "Day Mon DD HH:MM:SS YYYY" convert into following format "YYYY-MM-DD" as logDay for easy validation
			- if from != nil && logDay.Before(*from) return false
			- if to != nil
				- make endOfDay use to.Add(24 * time.Hour)
				- if !logDay.Before(endOfDay) return false
			- return true
		- If above function return false skip the line else analyze it
- Analyze the line and take count

## How to run it

```
run go ./cmd/main/main.go "root path"
```
