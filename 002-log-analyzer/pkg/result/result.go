package result

import (
	"fmt"
	"local/002-log-analyzer/pkg/scanner"
)

func PrintResult(analytics scanner.Analytics) {
	fmt.Println("Log Analytics Report")
	fmt.Println("--------------------")
	fmt.Printf("Files analyzed: %v\n", analytics.Files)
	fmt.Printf("Lines analyzed: %v\n\n", analytics.Lines)
	fmt.Printf("INFO: \t%v\n", analytics.Notice)
	fmt.Printf("WARN: \t%v\n", analytics.Warn)
	fmt.Printf("ERROR: \t%v\n", analytics.Error)
}