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

### Phase - 4
- Aceept search keyword from CLI
- Skip the line if search keyword is not found the line

### Phase - 5
- Aceept searchable level from CLI, this is optional and only accept following levels `INFO, WARN, ERROR` and default level is INFO

### Phase - 6 : print top errors
- Store errors and count into a map
	- create a empty slice with 0 length and len(map) capacity
	- append all keys of map intot the newly created slice
	- sort the slice use `sort.SliceStable(keys, func(i, j int) bool)`, use this function sort the map values
	- print only 5 key:values of map, which mean top errors

### Phase - 7 : save the report into a file
- use `func Create(name string) (*File, error)` to create a file and defer Close file
- use `func NewWriter(w io.Writer) *Writer` to setup bufferedFile for save the item into memory
- use `func MultiWriter(writers ...Writer) Writer` to create duplicate writer to handle write the report into terminal and file simultaneously
- use `func Fprintln(w io.Writer, a ...any) (n int, err error)` use to write the report

### Phase - 8 : refactoring
- Count every scanned line
- Loglevel parser name change to bettern understanding
- Don't calculate allotted level for each line, change it only for one time
- Change log level name INFO to NOTICE

## How to run it

```
go run ./cmd/main/main.go --root /var/www/html --from 2026-08-01 --to 2026-08-01 --search "database connection" --level INFO
```