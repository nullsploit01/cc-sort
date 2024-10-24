package internal

import (
	"bufio"
	"os"
	"sort"
)

type FileSorter struct {
	FileData map[string]uint64
	Lines    []string
}

func ProcessFileToSorter(file *os.File) (*FileSorter, error) {
	if _, err := file.Seek(0, 0); err != nil {
		return nil, err
	}

	fileData := make(map[string]uint64)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fileData[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var lines []string
	for line, freq := range fileData {
		for i := uint64(0); i < freq; i++ {
			lines = append(lines, line)
		}
	}

	return &FileSorter{
		FileData: fileData,
		Lines:    lines,
	}, nil
}

func (f *FileSorter) SortFileByLines() ([]string, error) {
	sort.Strings(f.Lines)
	return f.Lines, nil
}

func (f *FileSorter) SortFileByUniqueLines() ([]string, error) {
	var uniqueLines []string
	for line := range f.FileData {
		uniqueLines = append(uniqueLines, line)
	}

	sort.Strings(uniqueLines)

	return uniqueLines, nil
}
