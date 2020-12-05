package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/st-tech/refactoring-conductor/domain"
	"github.com/st-tech/refactoring-conductor/internal"
)

var directoryCommand = &cobra.Command{
	Use:   "directory [path to directory]",
	Short: "calculate files inside the directory",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}

		files := internal.DirectoryInternalFiles(args[0])
		VBScriptJSON := domain.VBScriptJSON{}
		VBScriptFileExtensionPattern := regexp.MustCompile(domain.VBScriptFileExtension)

		for _, file := range files {
			vbscript := domain.VBScript{}
			vbscript.FileName = file

			isVBScript := VBScriptFileExtensionPattern.MatchString(file)

			if isVBScript {
				internal.Read(file, &vbscript)
				VBScriptJSON.VBScript = append(VBScriptJSON.VBScript, vbscript)
			}
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
	rootCmd.AddCommand(directoryCommand)
}
