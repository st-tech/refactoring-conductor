package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/st-tech/refactoring-conductor/domain"
	"github.com/st-tech/refactoring-conductor/internal"
)

var singleFileCmd = &cobra.Command{
	Use:   "single [file name]",
	Short: "Calculate cognitive complexity of single file.",
	Args:  cobra.RangeArgs(1, 10),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}

		VBScriptJSON := domain.VBScriptJSON{}

		for _, arg := range args {
			vbscript := domain.VBScript{}
			vbscript.FileName = arg

			internal.Read(arg, &vbscript)
			VBScriptJSON.VBScript = append(VBScriptJSON.VBScript, vbscript)
		}

		jsonBytes, err := json.Marshal(VBScriptJSON)
		if err != nil {
			fmt.Println("JSON Marshal error:", err)
		}

		out := new(bytes.Buffer)
		err = json.Indent(out, jsonBytes, "", " ")
		if err != nil {
			fmt.Println("JSON Indent error:", err)
		}
		fmt.Println(out.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(singleFileCmd)
}
