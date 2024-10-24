package internal

import (
	"bufio"
	"os"
	"sort"
)

func SortFileByLine(file *os.File) ([]string, error) {
	defer file.Seek(0, 0)
	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	sortLines := make([]string, len(lines))
	copy(sortLines, lines)
	sort.Sort(sort.StringSlice(sortLines))
	return sortLines, nil
}
