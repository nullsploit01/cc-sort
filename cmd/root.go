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

var rootCmd = &cobra.Command{
	Use:   "cc-sort",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		fs, err := internal.ProcessFileToSorter(file)
		if err != nil {
			cmd.PrintErrln(err)
		}

		if isUnique {
			lines, err = fs.SortFileByUniqueLines()
		} else {
			lines, err = fs.SortFileByLines()
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
		// Unwrap the error
		return opErr.Err == syscall.EPIPE
	}
	return false
}

func init() {
	rootCmd.Flags().BoolVarP(&isUnique, "unique", "u", false, "Unique Keys")
}
