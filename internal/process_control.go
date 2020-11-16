package internal

import (
	"fmt"
	"regexp"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
)

// CountControlFlow is to count Control Flow statements.
func CountControlFlow(vbscript domain.VBScript, str string) int {
	var cognitiveComplexityCount = 2

	isIf, err := regexp.MatchString(domain.VBScriptControlFlowPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	fmt.Printf("%v\n", isIf)

	return cognitiveComplexityCount
}
