package internal

import (
	"fmt"
	"regexp"

	"github.com/st-tech/refactoring-conductor/domain"
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
			Cognitive: domain.Cognitive{
				NestState:           0,
				CognitiveComplexity: 0,
			},
			FunctionName: str,
		}
		vbscript.Functions = append(vbscript.Functions, slice)
	}

	if vbscript.IsBeginFunction {
		if isEndIf {
			vbscript.EndNest()
			vbscript.Functions[len(vbscript.Functions)-1].EndNest()
		} else if isIf && !isElse {
			vbscript.BeginNest()
			vbscript.Functions[len(vbscript.Functions)-1].BeginNest()
		} else if isElse {
			vbscript.Increment()
			vbscript.Functions[len(vbscript.Functions)-1].Increment()
		}
	} else { // Function外の計算を行う
		if isEndIf {
			vbscript.EndNest()
		} else if isIf && !isElse {
			vbscript.BeginNest()
		} else if isElse {
			vbscript.Increment()
		}
	}
}
