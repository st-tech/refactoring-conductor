package internal

import (
	"fmt"
	"regexp"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
)

// CountControlFlow is to count Control Flow statements.
func CountControlFlow(vbscript domain.VBScript, str string) int {
	var count = 0
	isIf, err := regexp.MatchString(domain.VBScriptControlFlowPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	if isIf {
		vbscript.CognitiveComplexity++
		count++
	}

	return count
}
