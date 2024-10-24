package internal

import (
	"bufio"
	"os"
	"sort"
)

func SortFileByLine(file *os.File) ([]string, error) {
	if _, err := file.Seek(0, 0); err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	sort.Strings(lines)

	return lines, nil
}
