package scanner

import (
	"bufio"
	"slices"
	"os"
	"io/fs"
	"local/002-log-analyzer/pkg/parser"
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

func Scan(rootPath string, fromDate *time.Time, toDate *time.Time, searchStr string, levelStr string) (Analytics, error) {
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

		return analyzeFile(path, &analytics, fromDate, toDate, searchStr, levelStr)
		// return nil
	})

	if err != nil {
		return analytics, err
	}

	return analytics, nil
}

func analyzeFile(path string, analytics *Analytics, from *time.Time, to *time.Time, search string, level string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Extract date from log line
		logDate, ok := parser.ExtractDateFromLine(line)

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
		selectedLogLevel := parser.LogLevelParser(level)
		countByLevel(line, analytics, selectedLogLevel)
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

func countByLevel(line string, analytics *Analytics, allowedLevels []string) {
	// 1. Identify which level this line actually contains
	lineLevel := ""
	if strings.Contains(line, "[notice]") {
		lineLevel = "[notice]"
	} else if strings.Contains(line, "[warn]") {
		lineLevel = "[warn]"
	} else if strings.Contains(line, "[error]") {
		lineLevel = "[error]"
	}

	// If the line has no recognized log level tag, return early
	if lineLevel == "" {
		return
	}

	// 2. Filter: Check if the line's level is in the allowed levels list
	if !slices.Contains(allowedLevels, lineLevel) {
		return
	}

	// 3. Increment specific analytics counters based on the line's level
	switch lineLevel {
		case "[notice]":
			analytics.Notice++
		case "[warn]":
			analytics.Warn++
		case "[error]":
			analytics.Error++
	}

	analytics.Total++
}