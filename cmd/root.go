package cmd

import (
	"os"

	"github.com/merschformann/gobenchtransform/benchconvert"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobenchtransform",
	Short: "Transform Go benchmark results into a format that can be used by other tools.",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the input and output files.
		inputFile, err := cmd.Flags().GetString("input")
		if err != nil {
			panic(err)
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			panic(err)
		}
		quiet, err := cmd.Flags().GetBool("quiet")
		if err != nil {
			panic(err)
		}

		// Get input and output streams.
		var input *os.File
		if inputFile == "" {
			input = os.Stdin
		} else {
			input, err = os.Open(inputFile)
			if err != nil {
				panic(err)
			}
		}
		var output *os.File
		if outputFile == "" {
			output = os.Stdout
		} else {
			output, err = os.Create(outputFile)
			if err != nil {
				panic(err)
			}
		}

		// Convert the input to CSV.
		err = benchconvert.ConvertToCSV(input, output, quiet)
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("input", "i", "", "input file (default is stdin)")
	rootCmd.Flags().StringP("output", "o", "", "output file (default is stdout)")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress repeating output to stdout")
}
