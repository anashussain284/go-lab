package scanner

import (
	"bufio"
	// "fmt"
	"os"
	// "log"
	// "fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

type Analytics struct {
	Files int
	Lines int
	Notice int
	Warn int
	Error int
	Total int
}

// Scan now accepts optional from and to date pointers
func Scan(rootPath string, fromDate *time.Time, toDate *time.Time, searchStr string) (Analytics, error) {
	analytics := Analytics {}

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".log" {
			return nil
		}

		analytics.Files++

		return analyzeFile(path, &analytics, fromDate, toDate, searchStr)
		// return nil
	})

	if err != nil {
		return analytics, err
	}

	return analytics, nil
}

func analyzeFile(path string, analytics *Analytics, from *time.Time, to *time.Time, search string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Extract date from log line
		logDate, ok := extractDateFromLine(line)

		// fmt.Printf("logDate: %v and ok: %v\n", logDate, ok)

		// Skip date checks if the line does not contain a valid date timestamp
		if ok {
			// If a line is outside the requested date range, skip processing it
			if !withinDateRange(logDate, from, to) {
				continue
			}
		}

		if !withinString(line, search) {
			continue
		}

		// Count line & log level if it passed the date filter
		analytics.Lines ++
		countByLevel(line, analytics)
	}

	return scanner.Err()
}


func withinString(line string, search string) bool {
	if strings.Contains(line, search) {
		return true
	}

	return false
}

// Helper to check if a log date falls inside [from, to] inclusive
func withinDateRange(logDate time.Time, from *time.Time, to *time.Time) bool {
	// Truncate time portion from logDate to compare purely by calendar date (YYYY-MM-DD)
	logDay := time.Date(logDate.Year(), logDate.Month(), logDate.Day(), 0, 0, 0, 0, logDate.Location())

	// fmt.Printf("logDay: %v\n", logDay)

	if from != nil && logDay.Before(*from) {
		return false
	}

	if to != nil {
		endOfDay := to.Add(24 * time.Hour)
		if !logDay.Before(endOfDay) {
			return false
		}
	}

	return true
}

// Helper to extract date timestamp from a standard log line
func extractDateFromLine(line string) (time.Time, bool) {
	// Find opening '[' and closing ']' for timestamp bracket
	startIdx := strings.Index(line, "[")
	endIdx := strings.Index(line, "]")

	// fmt.Printf("startIdx: %v and endIdx: %v\n", startIdx, endIdx)

	if startIdx == -1 || endIdx == -1 || startIdx >= endIdx {
		return time.Time{}, false
	}

	// Extract content inside brackets: "Thu Jun 01 06:07:04 2005"
	rawTimestamp := line[startIdx+1 : endIdx]
	// fmt.Printf("rawTimestamp: %v\n", rawTimestamp)

	// Go reference layout for "Day Mon DD HH:MM:SS YYYY"
	layout := "Mon Jan 02 15:04:05 2006"

	t, err := time.Parse(layout, rawTimestamp)
	if err != nil {
		return time.Time{}, false
	}
	// fmt.Printf("t from log file: %v\n", t)

	return t, true
}

func countByLevel(line string, analytics *Analytics) {
	if strings.Contains(line, "[notice]") {
		analytics.Notice++
	}

	if strings.Contains(line, "[warn]") {
		analytics.Warn++
	}

	if strings.Contains(line, "[error]") {
		analytics.Error++
	}

	analytics.Total++
}