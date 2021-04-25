package internal

import (
	"fmt"
	"regexp"

	"github.com/st-tech/refactoring-conductor/domain"
)

// CountControlFlow is to count Control Flow statements.
func CountControlFlow(vbscript *domain.VBScript, str string) {
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
		if isEndNestStatement(str) {
			vbscript.EndNest()
			getLastFunction(vbscript).EndNest()
		} else if isBeginNestStatement(str) {
			vbscript.BeginNest()
			getLastFunction(vbscript).BeginNest()
		} else if isIncrementStatement(str) {
			vbscript.Increment()
			getLastFunction(vbscript).Increment()
		}
	} else { // Function外の計算を行う
		if isEndNestStatement(str) {
			vbscript.EndNest()
		} else if isBeginNestStatement(str) {
			vbscript.BeginNest()
		} else if isIncrementStatement(str) {
			vbscript.Increment()
		}
	}
}

func getLastFunction(vbscript *domain.VBScript) *domain.Function {
	return &vbscript.Functions[len(vbscript.Functions)-1]
}

func isBeginNestStatement(str string) bool {
	isIf, err := regexp.MatchString(domain.VBScriptIfPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isFor, err := regexp.MatchString(domain.VBScriptForPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isDo, err := regexp.MatchString(domain.VBScriptDoPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isWhile, err := regexp.MatchString(domain.VBScriptWhilePattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	return isIf || isFor || isDo || isWhile
}

func isEndNestStatement(str string) bool {
	isEndIf, err := regexp.MatchString(domain.VBScriptEndIfPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isNext, err := regexp.MatchString(domain.VBScriptNextPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isLoop, err := regexp.MatchString(domain.VBScriptLoopPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	isWhileEnd, err := regexp.MatchString(domain.VBScriptWhileEndPattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	return isEndIf || isNext || isLoop || isWhileEnd
}

func isIncrementStatement(str string) bool {
	isElse, err := regexp.MatchString(domain.VBScriptElsePattern, str)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	return isElse
}
