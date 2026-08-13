package result

import (
	"os"
	"fmt"
	"sort"
	"local/002-log-analyzer/pkg/scanner"
)

func PrintResult(analytics scanner.Analytics, rootPath string) {
	fmt.Println("Log Analytics Report")
	fmt.Println("--------------------")
	fmt.Printf("Dir:\t%v\n", rootPath)
	fmt.Printf("Files analyzed: %v\n", analytics.Files)
	fmt.Printf("Lines analyzed: %v\n\n", analytics.Lines)
	fmt.Printf("INFO: \t%v\n", analytics.Notice)
	fmt.Printf("WARN: \t%v\n", analytics.Warn)
	fmt.Printf("ERROR: \t%v\n\n", analytics.Error)
	fmt.Printf("Total: \t%v\n", analytics.Total)
	fmt.Println()

	topErrors(&analytics)
}

func topErrors(analytics *scanner.Analytics) {
	allErrors := analytics.TopErrors

	keys := make([]string, 0, len(allErrors))

	for key := range allErrors {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return allErrors[keys[i]] > allErrors[keys[j]]
	})

	fmt.Printf("Top Errors: \n")

	limit := min(5, len(keys))

	topKeys := keys[:limit]

	for _, k := range topKeys {
		fmt.Printf("(%d):\t%v\n",allErrors[k], k)
	}
}