package scanner

import (
	"path/filepath"
	"io/fs"
)

type Statistics struct {
	Files int
	Directories int
	TotalSize int64
	LargestFile string
	LargestSize int64
	Extensions map[string]int
}

func Scan(rootPath string) (Statistics, error)  {
	stats := Statistics {
		Extensions : make(map[string]int),
	}

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if path == rootPath {
				return nil
			}

			stats.Directories++
			return nil
		}
		stats.Files++

		fileInfo, err := d.Info()

		if err != nil {
			return err
		}

		fileSize := fileInfo.Size()

		stats.TotalSize += fileSize

		if stats.LargestSize < fileSize {
			stats.LargestFile = filepath.Base(path)
			stats.LargestSize = fileSize
		}

		ext := filepath.Ext(path)

		if ext == "" {
			ext = "no-extension"
		}

		stats.Extensions[ext]++

		return nil
	})

	if err != nil {
		return stats, err
	}

	return stats, nil
}