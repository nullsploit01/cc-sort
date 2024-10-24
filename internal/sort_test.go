package internal_test

import (
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/nullsploit01/cc-sort/internal"
)

func createTempFile(t *testing.T, content string) *os.File {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Unable to create temporary file: %v", err)
	}

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatalf("Failed to seek to start of file: %v", err)
	}

	return tmpFile
}

func TestProcessFileToSorter(t *testing.T) {
	content := "hello\nworld\nhello\n"
	file := createTempFile(t, content)
	defer file.Close()

	sorter, err := internal.ProcessFileToSorter(file)
	if err != nil {
		t.Fatalf("ProcessFileToSorter failed: %v", err)
	}

	expected := map[string]uint64{"hello": 2, "world": 1}
	if !reflect.DeepEqual(sorter.FileData, expected) {
		t.Errorf("Expected %v, got %v", expected, sorter.FileData)
	}
}

func TestSortFileByLines(t *testing.T) {
	fileData := map[string]uint64{"hello": 2, "world": 1}
	sorter := &internal.FileSorter{FileData: fileData}

	sortedLines, _ := sorter.SortFileByLines()
	expected := []string{"hello", "hello", "world"}

	if !reflect.DeepEqual(sortedLines, expected) {
		t.Errorf("Expected %v, got %v", expected, sortedLines)
	}
}

func TestSortFileByUniqueLines(t *testing.T) {
	fileData := map[string]uint64{"banana": 1, "apple": 2, "cherry": 1}
	sorter := &internal.FileSorter{FileData: fileData}

	sortedLines, _ := sorter.SortFileByUniqueLines()
	expected := []string{"apple", "banana", "cherry"}
	sort.Strings(expected)

	if !reflect.DeepEqual(sortedLines, expected) {
		t.Errorf("Expected %v, got %v", expected, sortedLines)
	}
}
