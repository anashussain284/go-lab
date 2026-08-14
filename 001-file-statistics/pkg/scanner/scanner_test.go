package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan(t *testing.T) {
	root := t.TempDir()

	files := map[string]string{
		"a.txt": "hello",
		"b.txt": "hello world",
		"c.go":  "hello world!",
	}

	for name, content := range files {
		path := filepath.Join(root, name)

		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	stats, err := Scan(root)
	if err != nil {
		t.Fatal(err)
	}

	if stats.Files != 3 {
		t.Errorf("expected 3 files, got %d", stats.Files)
	}

	if stats.Directories != 0 {
		t.Errorf("expected 0 directories, got %d", stats.Directories)
	}

	if stats.TotalSize != 28 {
		t.Errorf("expected total size 28, got %d", stats.TotalSize)
	}

	if stats.LargestFile != "c.go" {
		t.Errorf("expected largest file c.go, got %s", stats.LargestFile)
	}

	if stats.LargestSize != 12 {
		t.Errorf("expected largest size 12, got %d", stats.LargestSize)
	}

	if stats.Extensions[".txt"] != 2 {
		t.Errorf("expected .txt count 2, got %d", stats.Extensions[".txt"])
	}

	if stats.Extensions[".go"] != 1 {
		t.Errorf("expected .go count 1, got %d", stats.Extensions[".go"])
	}
}