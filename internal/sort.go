package internal

import "os"

func SortFileByLine(file *os.File) error {
	defer file.Seek(0, 0)
	return nil
}
