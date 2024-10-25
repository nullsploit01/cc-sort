package internal

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"hash/fnv"
	"math/big"
	"os"
	"sort"
)

type SortAlgorithm string

const (
	RadixSort  SortAlgorithm = "radix"
	MergeSort  SortAlgorithm = "merge"
	QuickSort  SortAlgorithm = "quick"
	HeapSort   SortAlgorithm = "heap"
	RandomSort SortAlgorithm = "random"
)

type FileSorter struct {
	FileData map[string]uint64
	Lines    []string
}

func ProcessFileToSort(file *os.File) (*FileSorter, error) {
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

func (f *FileSorter) applySort(lines []string, algorithm SortAlgorithm) ([]string, error) {
	switch algorithm {
	case RadixSort:
		return f.SortByRadix(lines), nil
	case MergeSort:
		return f.SortByMerge(lines), nil
	case QuickSort:
		return f.SortByQuick(lines), nil
	case HeapSort:
		return f.SortByHeap(lines), nil
	case RandomSort:
		return f.SortByRandom(lines), nil
	default:
		return nil, errors.New("unsupported sort algorithm")
	}
}

func (f *FileSorter) SortFileByLines(algorithm SortAlgorithm) ([]string, error) {
	return f.applySort(f.Lines, algorithm)
}

func (f *FileSorter) SortFileByUniqueLines(algorithm SortAlgorithm) ([]string, error) {
	uniqueLines := make([]string, 0, len(f.FileData))
	for line := range f.FileData {
		uniqueLines = append(uniqueLines, line)
	}
	return f.applySort(uniqueLines, algorithm)
}

func (f *FileSorter) SortByRadix(lines []string) []string {
	maxLength := getMaxLineLength(lines)

	for i := maxLength - 1; i >= 0; i-- {
		lines = countingSortByPosition(lines, i)
	}

	return lines
}

func (f *FileSorter) SortByMerge(lines []string) []string {
	return mergeSort(lines)
}

func (f *FileSorter) SortByQuick(lines []string) []string {
	quickSort(lines, 0, len(lines)-1)

	return lines
}

func (f *FileSorter) SortByHeap(lines []string) []string {
	heapSort(lines)

	return lines
}

func (f *FileSorter) SortByRandom(lines []string) []string {
	seed, _ := getRandomSeed()

	hasher := fnv.New64a()
	hashValues := make(map[string]uint64)

	for _, lines := range lines {
		hasher.Reset()
		seedBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(seedBytes, seed.Uint64())

		_, _ = hasher.Write(seedBytes)
		_, _ = hasher.Write([]byte(lines))
		hashValues[lines] = hasher.Sum64()
	}

	sort.Slice(lines, func(i, j int) bool {
		return hashValues[lines[i]] < hashValues[lines[j]]
	})

	return lines
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

func heapSort(lines []string) {
	n := len(lines)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(lines, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		lines[0], lines[i] = lines[i], lines[0]
		heapify(lines, i, 0)
	}
}

func heapify(lines []string, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && lines[left] > lines[largest] {
		largest = left
	}

	if right < n && lines[right] > lines[largest] {
		largest = right
	}

	if largest != i {
		lines[i], lines[largest] = lines[largest], lines[i]
		heapify(lines, n, largest)
	}
}

func getRandomSeed() (*big.Int, error) {
	var seed big.Int
	seedBytes := make([]byte, 8)

	_, err := rand.Read(seedBytes)
	if err != nil {
		return nil, err
	}

	seed.SetBytes(seedBytes)
	return &seed, nil
}
