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
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}

		vbscript := domain.VBScript{}

		internal.Read(args[0], &vbscript)

		fmt.Printf("CognitiveComplexity: %d\n", vbscript.CognitiveComplexity)

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
