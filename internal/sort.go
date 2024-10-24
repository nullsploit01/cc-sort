package internal

import (
	"bufio"
	"errors"
	"os"
)

type SortAlgorithm string

const (
	RadixSort SortAlgorithm = "radix"
	MergeSort SortAlgorithm = "merge"
	QuickSort SortAlgorithm = "quick"
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

func (f *FileSorter) SortFileByLines(algorithm SortAlgorithm) ([]string, error) {
	var sortedLines []string
	switch algorithm {
	case RadixSort:
		sortedLines = f.SortByRadix()
	case MergeSort:
		sortedLines = f.SortByMerge()
	case QuickSort:
		sortedLines = f.SortByQuick()
	default:
		return nil, errors.New("unsupported sort algorithm")
	}
	return sortedLines, nil
}

func (f *FileSorter) SortFileByUniqueLines(algorithm SortAlgorithm) ([]string, error) {
	var uniqueLines []string
	for line := range f.FileData {
		uniqueLines = append(uniqueLines, line)
	}

	f.Lines = uniqueLines
	var sortedLines []string

	switch algorithm {
	case RadixSort:
		sortedLines = f.SortByRadix()
	case MergeSort:
		sortedLines = f.SortByMerge()
	case QuickSort:
		sortedLines = f.SortByQuick()
	default:
		return nil, errors.New("unsupported sorting algorithm")
	}

	return sortedLines, nil
}

func (f *FileSorter) SortByRadix() []string {
	maxLength := getMaxLineLength(f.Lines)

	for i := maxLength - 1; i >= 0; i-- {
		f.Lines = countingSortByPosition(f.Lines, i)
	}

	return f.Lines
}

func (f *FileSorter) SortByMerge() []string {
	return mergeSort(f.Lines)
}

func (f *FileSorter) SortByQuick() []string {
	quickSort(f.Lines, 0, len(f.Lines)-1)

	return f.Lines
}

func countingSortByPosition(lines []string, position int) []string {
	count := make([]int, 256)
	output := make([]string, len(lines))

	for i := 0; i < len(lines); i++ {
		charIndex := getCharIndex(lines[i], position)
		count[charIndex]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i := len(lines) - 1; i >= 0; i-- {
		charIndex := getCharIndex(lines[i], position)
		output[count[charIndex]-1] = lines[i]
		count[charIndex] -= 1
	}

	return output
}

func getMaxLineLength(lines []string) int {
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	return maxLength
}

func getCharIndex(line string, position int) int {
	if position >= len(line) {
		return 0
	}

	return int(line[position]) + 1
}

func mergeSort(lines []string) []string {
	if len(lines) <= 1 {
		return lines
	}

	mid := len(lines) / 2
	left := mergeSort(lines[:mid])
	right := mergeSort(lines[mid:])

	return merge(left, right)
}

func merge(left, right []string) []string {
	var merged []string

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

func quickSort(lines []string, low, high int) {
	if low < high {
		pi := partition(lines, low, high)
		quickSort(lines, low, pi-1)
		quickSort(lines, pi+1, high)
	}
}

func partition(lines []string, low, high int) int {
	pivot := lines[high]
	i := low - 1
	for j := low; j < high; j++ {
		if lines[j] < pivot {
			i++
			lines[i], lines[j] = lines[j], lines[i]
		}
	}
	lines[i+1], lines[high] = lines[high], lines[i+1]
	return i + 1
}
