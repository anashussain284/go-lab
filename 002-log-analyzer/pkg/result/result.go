package result

import (
	"bufio"
	"fmt"
	"io"
	"local/002-log-analyzer/pkg/scanner"
	"log"
	"os"
	"sort"
)

func PrintResult(analytics scanner.Analytics, rootPath string) {
	file, err := os.Create("report.txt")

	if err != nil {
		log.Fatalf("error: create file, %v", err)
	}
	defer file.Close()

	bufferedFile := bufio.NewWriter(file)
	defer bufferedFile.Flush()

	w := io.MultiWriter(os.Stdout, bufferedFile)

	fmt.Fprintln(w, "Log Analytics Report")
	fmt.Fprintln(w, "--------------------")
	fmt.Fprintf(w, "Dir:\t%v\n", rootPath)
	fmt.Fprintf(w, "Files analyzed: %v\n", analytics.Files)
	fmt.Fprintf(w, "Lines analyzed: %v\n\n", analytics.Lines)
	fmt.Fprintf(w, "INFO: \t%v\n", analytics.Notice)
	fmt.Fprintf(w, "WARN: \t%v\n", analytics.Warn)
	fmt.Fprintf(w, "ERROR: \t%v\n\n", analytics.Error)
	fmt.Fprintf(w, "Total: \t%v\n\n", analytics.Total)

	topErrors(&analytics, w)
}

func topErrors(analytics *scanner.Analytics, w io.Writer) {
	allErrors := analytics.TopErrors

	keys := make([]string, 0, len(allErrors))

	for key := range allErrors {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return allErrors[keys[i]] > allErrors[keys[j]]
	})

	fmt.Fprintf(w, "Top Errors: \n")

	limit := min(5, len(keys))

	topKeys := keys[:limit]

	for _, k := range topKeys {
		fmt.Fprintf(w, "(%d):\t%v\n", allErrors[k], k)
	}
}
