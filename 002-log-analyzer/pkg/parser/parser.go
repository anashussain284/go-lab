package parser

import (
	"time"
	"strings"
)

func LogLevelParser(level string) []string {
	defaultLevel := []string{"[notice]","[warn]","[error]"}
	switch level {
		case "INFO":
			return defaultLevel
		case "WARN":
			return []string{"[warn]","[error]"}
		case "ERROR":
			return []string{"[error]"}
		default:
			return defaultLevel
	}
}

func ParseDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
}

// Helper to extract date timestamp from a standard log line
func ExtractDateFromLine(line string) (time.Time, bool) {
	// Find opening '[' and closing ']' for timestamp bracket
	startIdx := strings.Index(line, "[")
	endIdx := strings.Index(line, "]")

	if startIdx == -1 || endIdx == -1 || startIdx >= endIdx {
		return time.Time{}, false
	}

	// Extract content inside brackets: "Thu Jun 01 06:07:04 2005"
	rawTimestamp := line[startIdx+1 : endIdx]

	// Go reference layout for "Day Mon DD HH:MM:SS YYYY"
	layout := "Mon Jan 02 15:04:05 2006"

	t, err := time.Parse(layout, rawTimestamp)
	if err != nil {
		return time.Time{}, false
	}

	return t, true
}