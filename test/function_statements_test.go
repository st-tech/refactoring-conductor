package test_test

import (
	"testing"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/internal"
)

func TestFunctionStatements(t *testing.T) {
	var expectedCognitiveComplexity = 3
	var expectedFirstCognitiveComplexity = 2
	var expectedSecondCognitiveComplexity = 1
	filename := "testdata/function.vbs"
	vbscript := domain.VBScript{}

	vbscript.CognitiveComplexity = internal.Read(filename, &vbscript)

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}

	if expectedFirstCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.Functions[0].CognitiveComplexity, expectedFirstCognitiveComplexity)
	}

	if expectedSecondCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.Functions[1].CognitiveComplexity, expectedSecondCognitiveComplexity)
	}
}
