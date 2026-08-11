package result

import (
	"fmt"
	"local/002-log-analyzer/pkg/scanner"
)

func PrintResult(analytics scanner.Analytics) {
	fmt.Println("Log Analytics Report")
	fmt.Println("--------------------")
	fmt.Printf("Files analyzed: %v\n", analytics.Files)
	fmt.Printf("Lines analyzed: %v\n", analytics.Lines)
}