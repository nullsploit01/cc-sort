package cmd

import (
	"io"
	"os"
	"strings"
	"syscall"

	"github.com/nullsploit01/cc-sort/internal"
	"github.com/spf13/cobra"
)

var isUnique bool
var sortAlgorithm string

var rootCmd = &cobra.Command{
	Use:   "ccsort",
	Short: "Sorts files using various algorithms.",
	Long: `cc-sort is a powerful CLI tool for sorting contents of a file using various algorithms including Radix, Merge, Quick, Heap, and Random sorts. 
	This tool supports unique sorting where duplicate lines are ignored.

	Usage:
	ccsort [options] [file]

	Examples:
	cc-sort --algorithm merge --unique file.txt
	cc-sort --algorithm quick file.txt

	Flags:
	-u, --unique       Sort unique lines only.
	-a, --algorithm    Specify the sorting algorithm to use (default "radix").`,

	Run: func(cmd *cobra.Command, args []string) {
		var lines []string
		if len(args) < 1 {
			cmd.PrintErr("Error: A file name is required as an argument.\n")
			cmd.Usage()
			return
		}

		file, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fs, err := internal.ProcessFileToSort(file)
		if err != nil {
			cmd.PrintErrln(err)
		}

		if isUnique {
			lines, err = fs.SortFileByUniqueLines(internal.SortAlgorithm(sortAlgorithm))
		} else {
			lines, err = fs.SortFileByLines(internal.SortAlgorithm(sortAlgorithm))
		}
		if err != nil {
			cmd.PrintErrln(err)
		}

		output := strings.Join(lines, "\n")
		_, err = cmd.OutOrStdout().Write([]byte(output))
		if err != nil && !isBrokenPipeError(err) {
			cmd.Println(err)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func isBrokenPipeError(err error) bool {
	if err == syscall.EPIPE {
		return true
	}
	if err == io.ErrClosedPipe {
		return true
	}
	if opErr, ok := err.(*os.PathError); ok {
		return opErr.Err == syscall.EPIPE
	}
	return false
}

func init() {
	rootCmd.Flags().BoolVarP(&isUnique, "unique", "u", false, "--unique")
	rootCmd.Flags().StringVarP(&sortAlgorithm, "algorithm", "a", string(internal.RadixSort), "--algorithm <algorithm>")
}
