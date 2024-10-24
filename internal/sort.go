package internal

import (
	"bufio"
	"os"
	"sort"
)

type FileSorter struct {
	FileData map[string]uint64
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

	return &FileSorter{
		FileData: fileData,
	}, nil
}

func (f *FileSorter) SortFileByLines() ([]string, error) {
	var lines []string
	for line := range f.FileData {
		lines = append(lines, line)
	}

	sort.Slice(lines, func(i, j int) bool {
		return f.FileData[lines[i]] > f.FileData[lines[j]]
	})

	return lines, nil
}

func (f *FileSorter) SortFileByUniqueLines() ([]string, error) {
	var lines []string
	for line := range f.FileData {
		lines = append(lines, line)
	}

	sort.Strings(lines)

	return lines, nil
}
