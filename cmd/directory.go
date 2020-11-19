package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/internal"
)

type Options struct {
	optionDirectory string
}

var (
	o = &Options{}
)

var directoryCommand = &cobra.Command{
	Use:   "directory",
	Short: "calculate files inside the directory",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(o.optionDirectory)
		files := internal.DirectoryInternalFiles(o.optionDirectory)

		for _, file := range files {
			vbscript := domain.VBScript{}

			isVBScript, err := regexp.MatchString(domain.VBScriptFileExtension, file)
			if err != nil {
				fmt.Printf("err: %v", err)
			}

			if isVBScript {
				internal.Read(file, &vbscript)
				fmt.Printf("CognitiveComplexity: %d\n", vbscript.CognitiveComplexity)
				fmt.Printf("%+v\n", vbscript)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(directoryCommand)
	directoryCommand.Flags().StringVarP(&o.optionDirectory, "dir", "d", "directory", "directory option")
}
