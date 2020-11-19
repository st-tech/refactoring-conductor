package cmd

import (
	"bytes"
	"encoding/json"
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

		VBScriptJSON := domain.VBScriptJson{}

		for _, arg := range args {
			vbscript := domain.VBScript{}

			internal.Read(arg, &vbscript)
			VBScriptJSON.VBScript = append(VBScriptJSON.VBScript, vbscript)
		}

		jsonBytes, err := json.Marshal(VBScriptJSON)
		if err != nil {
			fmt.Println("JSON Marshal error:", err)
		}

		out := new(bytes.Buffer)
		json.Indent(out, jsonBytes, "", " ")
		fmt.Println(out.String())

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
