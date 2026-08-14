package output

import (
	"fmt"
	"sort"
	"local/006-file-statistics/pkg/scanner"
)

func PrintStatistics(statistics scanner.Statistics) {
	fmt.Printf("Files:\t\t%v\n", statistics.Files)
	fmt.Printf("Directories:\t%v\n", statistics.Directories)
	fmt.Printf("Total Size:\t%v bytes\n", statistics.TotalSize)
	fmt.Printf("Largest File:\t%v (%v bytes)\n", statistics.LargestFile, statistics.LargestSize)
	fmt.Println()

	keys := make([]string, 0, len(statistics.Extensions))

	for ext := range statistics.Extensions {
		keys = append(keys, ext)
	}

	sort.Strings(keys)

	for _, ext := range keys {
		fmt.Printf("%v\t%v\n", ext, statistics.Extensions[ext])
	}
}