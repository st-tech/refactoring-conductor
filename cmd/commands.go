package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/internal"
)

var rootCmd = &cobra.Command{
	Use:   "cog [file name]",
	Short: "Command line cognitive complexity",
	Args:  cobra.RangeArgs(1, 10),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}

		for _, arg := range args {
			vbscript := domain.VBScript{}

			internal.Read(arg, &vbscript)
			fmt.Printf("File Name: %s\n", arg)
			fmt.Printf("CognitiveComplexity: %d\n", vbscript.CognitiveComplexity)
			fmt.Printf("%+v\n", vbscript)
		}

		return nil
	},
}

// Execute is to execute main command line tool.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
