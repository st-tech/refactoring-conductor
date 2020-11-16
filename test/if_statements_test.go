package test_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/internal"
)

func TestIfStatements(t *testing.T) {
	var expectedCognitiveComplexity = 2
	filename := "testdata/if.vbs"
	vbscript := domain.VBScript{}

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		internal.CountControlFlow(vbscript, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("err: %v", err)
	}

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}

func TestIfStatementsSingleState(t *testing.T) {
	var expectedCognitiveComplexity = 1
	filename := "testdata/if_single_line.vbs"
	vbscript := domain.VBScript{}
	vbscript.NestState = 0

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		internal.CountControlFlow(vbscript, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("err: %v", err)
	}

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}

func TestIfStatementsNest(t *testing.T) {
	var expectedCognitiveComplexity = 5
	filename := "testdata/if_nest.vbs"
	vbscript := domain.VBScript{}
	vbscript.NestState = 0

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		internal.CountControlFlow(vbscript, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("err: %v", err)
	}

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}

func TestIfStatementsTripleNest(t *testing.T) {
	var expectedCognitiveComplexity = 8
	filename := "testdata/if_triple_nest.vbs"
	vbscript := domain.VBScript{}
	vbscript.NestState = 0

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		internal.CountControlFlow(vbscript, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("err: %v", err)
	}

	if expectedCognitiveComplexity != vbscript.CognitiveComplexity {
		t.Errorf("wrong output: got %v, expected %v", vbscript.CognitiveComplexity, expectedCognitiveComplexity)
	}
}
