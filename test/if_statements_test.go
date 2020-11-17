package test_test

import (
	"testing"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/internal"
)

func TestIfStatements(t *testing.T) {
	var expectedCognitiveComplexity = 2
	filename := "testdata/if.vbs"
	vbscript := domain.VBScript{}

	vbscript.CognitiveComplexity = internal.Read(filename, &vbscript)

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}

func TestElseIfStatements(t *testing.T) {
	var expectedCognitiveComplexity = 3
	filename := "testdata/elseif.vbs"
	vbscript := domain.VBScript{}

	vbscript.CognitiveComplexity = internal.Read(filename, &vbscript)

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}

func TestIfStatementsSingleState(t *testing.T) {
	var expectedCognitiveComplexity = 1
	filename := "testdata/if_single_line.vbs"
	vbscript := domain.VBScript{}
	vbscript.NestState = 0

	vbscript.CognitiveComplexity = internal.Read(filename, &vbscript)

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}

func TestIfStatementsNest(t *testing.T) {
	var expectedCognitiveComplexity = 5
	filename := "testdata/if_nest.vbs"
	vbscript := domain.VBScript{}
	vbscript.NestState = 0

	vbscript.CognitiveComplexity = internal.Read(filename, &vbscript)

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}

func TestIfStatementsTripleNest(t *testing.T) {
	var expectedCognitiveComplexity = 8
	filename := "testdata/if_triple_nest.vbs"
	vbscript := domain.VBScript{}
	vbscript.NestState = 0

	vbscript.CognitiveComplexity = internal.Read(filename, &vbscript)

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}
