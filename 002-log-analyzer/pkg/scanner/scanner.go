package scanner

import (
	"bufio"
	// "fmt"
	"os"
	// "log"
	// "fmt"
	"io/fs"
	"path/filepath"
)

type Analytics struct {
	Files int
	Lines int
}

func Scan(rootPath string) (Analytics, error) {
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

		return analyzeFile(path, &analytics)
		// return nil
	})

	if err != nil {
		return analytics, err
	}

	return analytics, nil
}

func analyzeFile(path string, analytics *Analytics) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// line := scanner.Text()

		analytics.Lines ++
	}

	return scanner.Err()
}