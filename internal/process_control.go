package internal

import (
	"fmt"
	"regexp"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
)

// CountControlFlow is to count Control Flow statements.
func CountControlFlow(vbscript *domain.VBScript, str string) {

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

	isFunction, err := regexp.MatchString(domain.VBScriptFunctionPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isEndFunction, err := regexp.MatchString(domain.VBScriptEndFunctionPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	if isEndFunction {
		vbscript.IsBeginFunction = false
	} else if isFunction && !isEndFunction {
		vbscript.IsBeginFunction = true
		slice := domain.Function{
			NestState:           0,
			CognitiveComplexity: 0,
			FunctionName:        str,
		}
		vbscript.Functions = append(vbscript.Functions, slice)
	}

	if isEndIf {
		vbscript.NestState--
		if vbscript.IsBeginFunction {
			vbscript.Functions[len(vbscript.Functions)-1].NestState--
		}
	} else if isIf && !isElse {
		vbscript.CognitiveComplexity++
		vbscript.CognitiveComplexity += vbscript.NestState
		vbscript.NestState++

		if vbscript.IsBeginFunction {
			vbscript.Functions[len(vbscript.Functions)-1].CognitiveComplexity++
			// Add (NestState -1), Because added vbscript.NestState++
			vbscript.Functions[len(vbscript.Functions)-1].CognitiveComplexity += (vbscript.NestState - 1)
			vbscript.Functions[len(vbscript.Functions)-1].NestState++
		}
	} else if isElse {
		vbscript.CognitiveComplexity++
		if vbscript.IsBeginFunction {
			vbscript.Functions[len(vbscript.Functions)-1].CognitiveComplexity++
		}
	}
}
