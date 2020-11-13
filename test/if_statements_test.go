package test_test

import (
	"testing"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/internal"
)

func TestIfStatements(t *testing.T) {
	var expectedCognitiveComplexity = 2
	var testCode = `If nOne < nTwo Then
						Min = nOne
					Else
						Min = nTwo
					End If`

	if v := internal.CountControlFlow(testCode); expectedCognitiveComplexity != v {
		t.Errorf("wrong output: got %v, expected %v", v, expectedCognitiveComplexity)
	}
}

func TestIfStatementsSingleState(t *testing.T) {
	var expectedCognitiveComplexity = 1
	var testCode = `If i=10 Then alert("Hello")`
	vbscript := domain.VBScript{}
	vbscript.NestState = 0

	if v := internal.CountControlFlow(testCode); expectedCognitiveComplexity != v {
		t.Errorf("wrong output: got %v, expected %v", v, expectedCognitiveComplexity)
	}
}

func TestIfStatementsNest(t *testing.T) {
	var expectedCognitiveComplexity = 5
	var testCode = `If nOne < nTwo Then // +1
						If nOne <= nTwo Then // +2
							Min = nTwo
						Else // +1
							Min = nOne
						End if
					Else // +1
						Min = nTwo
					End If`

	if v := internal.CountControlFlow(testCode); expectedCognitiveComplexity != v {
		t.Errorf("wrong output: got %v, expected %v", v, expectedCognitiveComplexity)
	}
}

func TestIfStatementsNestElseIf(t *testing.T) {
	var expectedCognitiveComplexity = 5
	var testCode = `If nOne < nTwo Then // +1
						If nOne <= nTwo Then // +2
							Min = nTwo
						ElseIf // +2
							Min = nOne
						End if
					Else // +1
						Min = nTwo
					End If`

	if v := internal.CountControlFlow(testCode); expectedCognitiveComplexity != v {
		t.Errorf("wrong output: got %v, expected %v", v, expectedCognitiveComplexity)
	}
}

func TestIfStatementsTripleNest(t *testing.T) {
	var expectedCognitiveComplexity = 8
	var testCode = `If nOne < nTwo Then // +1
						If nOne <= nTwo Then // +2
							If nOne <= nTwo Then // +3
								Min = nTwo
							Else // +1
								Min = nOne
							End if
						End if
					Else // +1
						Min = nTwo
					End If`

	if v := internal.CountControlFlow(testCode); expectedCognitiveComplexity != v {
		t.Errorf("wrong output: got %v, expected %v", v, expectedCognitiveComplexity)
	}
}
