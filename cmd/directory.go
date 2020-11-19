package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/internal"
)

var directoryCommand = &cobra.Command{
	Use:   "directory",
	Short: "calculate files inside the directory",
	Args:  cobra.RangeArgs(1, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}

		files := internal.DirectoryInternalFiles(args[0])
		VBScriptJson := domain.VBScriptJson{}

		for _, file := range files {
			vbscript := domain.VBScript{}

			isVBScript, err := regexp.MatchString(domain.VBScriptFileExtension, file)
			if err != nil {
				fmt.Printf("err: %v", err)
			}

			if isVBScript {
				internal.Read(file, &vbscript)
				fmt.Printf("%+v\n", vbscript)
			}
			VBScriptJson.VBScript = append(VBScriptJson.VBScript, vbscript)
		}

		jsonBytes, err := json.Marshal(VBScriptJson)
		if err != nil {
			fmt.Println("JSON Marshal error:", err)
		}

		out := new(bytes.Buffer)
		// プリフィックスなし、スペース4つでインデント
		json.Indent(out, jsonBytes, "", " ")
		fmt.Println(out.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(directoryCommand)
}
