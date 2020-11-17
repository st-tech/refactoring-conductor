package internal

import (
	"fmt"
	"regexp"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
)

// CountControlFlow is to count Control Flow statements.
func CountControlFlow(vbscript *domain.VBScript, str string) int {
	var count = 0
	isIf, err := regexp.MatchString(domain.VBScriptIfPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isElse, err := regexp.MatchString(domain.VBScriptElsePattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isEndIf, err := regexp.MatchString(domain.VBScriptEndIfPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	if isEndIf {
		vbscript.NestState--
	} else if isIf && !isElse {
		vbscript.CognitiveComplexity++
		vbscript.CognitiveComplexity += vbscript.NestState
		count += vbscript.NestState
		vbscript.NestState++
	} else if isElse {
		vbscript.CognitiveComplexity++
	}

	return count
}
