package internal

import (
	"os"
	"reflect"
	"testing"
)

func TestSortByRadix(t *testing.T) {
	sorter := &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	expected := []string{"apple", "banana", "banana", "mango", "orange"}
	result := sorter.SortByRadix(sorter.Lines)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortByRadix failed: got %v, want %v", result, expected)
	}
}

func TestSortByMerge(t *testing.T) {
	sorter := &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	expected := []string{"apple", "banana", "banana", "mango", "orange"}
	result := sorter.SortByMerge(sorter.Lines)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortByMerge failed: got %v, want %v", result, expected)
	}
}

func TestSortFileByLines(t *testing.T) {
	sorter := &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	expected := []string{"apple", "banana", "banana", "mango", "orange"}
	result, err := sorter.SortFileByLines(MergeSort)
	if err != nil {
		t.Errorf("SortFileByLines failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortFileByLines failed: got %v, want %v", result, expected)
	}

	// Test with Radix sort
	sorter = &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	expected = []string{"apple", "banana", "banana", "mango", "orange"}
	result, err = sorter.SortFileByLines(RadixSort)
	if err != nil {
		t.Errorf("SortFileByLines failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortFileByLines failed: got %v, want %v", result, expected)
	}

	// Test with Quick sort
	sorter = &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	expected = []string{"apple", "banana", "banana", "mango", "orange"}
	result, err = sorter.SortFileByLines(QuickSort)
	if err != nil {
		t.Errorf("SortFileByLines failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortFileByLines failed: got %v, want %v", result, expected)
	}

	// Test with Heap sort
	sorter = &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	expected = []string{"apple", "banana", "banana", "mango", "orange"}
	result, err = sorter.SortFileByLines(HeapSort)
	if err != nil {
		t.Errorf("SortFileByLines failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortFileByLines failed: got %v, want %v", result, expected)
	}
}

func TestSortFileByUniqueLines(t *testing.T) {
	sorter := &FileSorter{
		FileData: map[string]uint64{"banana": 2, "apple": 1, "orange": 1, "mango": 1},
	}
	expected := []string{"apple", "banana", "mango", "orange"}
	result, err := sorter.SortFileByUniqueLines(MergeSort)
	if err != nil {
		t.Errorf("SortFileByUniqueLines failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortFileByUniqueLines failed: got %v, want %v", result, expected)
	}

	// Test with Radix sort
	sorter = &FileSorter{
		FileData: map[string]uint64{"banana": 2, "apple": 1, "orange": 1, "mango": 1},
	}
	expected = []string{"apple", "banana", "mango", "orange"}
	result, err = sorter.SortFileByUniqueLines(RadixSort)
	if err != nil {
		t.Errorf("SortFileByUniqueLines failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortFileByUniqueLines failed: got %v, want %v", result, expected)
	}

	// Test with Quick sort
	sorter = &FileSorter{
		FileData: map[string]uint64{"banana": 2, "apple": 1, "orange": 1, "mango": 1},
	}
	expected = []string{"apple", "banana", "mango", "orange"}
	result, err = sorter.SortFileByUniqueLines(QuickSort)
	if err != nil {
		t.Errorf("SortFileByUniqueLines failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortFileByUniqueLines failed: got %v, want %v", result, expected)
	}

	// Test with Heap sort
	sorter = &FileSorter{
		FileData: map[string]uint64{"banana": 2, "apple": 1, "orange": 1, "mango": 1},
	}
	expected = []string{"apple", "banana", "mango", "orange"}
	result, err = sorter.SortFileByUniqueLines(HeapSort)
	if err != nil {
		t.Errorf("SortFileByUniqueLines failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortFileByUniqueLines failed: got %v, want %v", result, expected)
	}

	// Test with invalid sort algorithm
	sorter = &FileSorter{
		FileData: map[string]uint64{"banana": 2, "apple": 1, "orange": 1, "mango": 1},
	}
	_, err = sorter.SortFileByUniqueLines("invalid")
	if err == nil {
		t.Errorf("SortFileByUniqueLines failed with error: %v", err)
	}
}

func TestSortByQuick(t *testing.T) {
	sorter := &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	expected := []string{"apple", "banana", "banana", "mango", "orange"}
	sorter.SortByQuick(sorter.Lines)
	if !reflect.DeepEqual(sorter.Lines, expected) {
		t.Errorf("SortByQuick failed: got %v, want %v", sorter.Lines, expected)
	}
}

func TestSortByHeap(t *testing.T) {
	sorter := &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	expected := []string{"apple", "banana", "banana", "mango", "orange"}
	sorter.SortByHeap(sorter.Lines)
	if !reflect.DeepEqual(sorter.Lines, expected) {
		t.Errorf("SortByHeap failed: got %v, want %v", sorter.Lines, expected)
	}
}

func TestSortByRandom(t *testing.T) {
	sorter := &FileSorter{
		Lines: []string{"banana", "apple", "orange", "mango", "banana"},
	}
	original := make([]string, len(sorter.Lines))
	copy(original, sorter.Lines)
	sorter.SortByRandom(sorter.Lines)
	if reflect.DeepEqual(sorter.Lines, original) {
		t.Errorf("SortByRandom failed: results are the same as original, expected a change")
	}
	if len(sorter.Lines) != len(original) {
		t.Errorf("SortByRandom failed: resulting length %d does not match original %d", len(sorter.Lines), len(original))
	}
	tempMap := make(map[string]int)
	for _, item := range original {
		tempMap[item]++
	}
	for _, item := range sorter.Lines {
		tempMap[item]--
	}
	for item, count := range tempMap {
		if count != 0 {
			t.Errorf("SortByRandom failed: imbalance in item %s, count %d", item, count)
		}
	}
}

func TestProcessFileToSort(t *testing.T) {
	content := "banana\napple\norange\nmango\nbanana"
	file, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatal("Failed to create temp file:", err)
	}
	defer os.Remove(file.Name())
	_, err = file.WriteString(content)
	if err != nil {
		t.Fatal("Failed to write to temp file:", err)
	}
	_, err = file.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal("Failed to seek to start of file:", err)
	}

	sorter, err := ProcessFileToSort(file)
	if err != nil {
		t.Errorf("ProcessFileToSort failed: %v", err)
	}
	if len(sorter.Lines) != 5 {
		t.Errorf("Expected 5 lines, got %d", len(sorter.Lines))
	}
	if sorter.FileData["banana"] != 2 {
		t.Errorf("Expected 'banana' to appear 2 times, got %d", sorter.FileData["banana"])
	}
}
