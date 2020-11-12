package internal_test

import (
	"testing"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/internal"
)

// 制御フローに関する処理の算出のテストを行う
func TestCountControlFlow(t *testing.T) {
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
