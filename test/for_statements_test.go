package test_test

import (
	"testing"

	"github.com/st-tech/refactoring-conductor/domain"
	"github.com/st-tech/refactoring-conductor/internal"
)

func TestForStatements(t *testing.T) {
	tests := []struct {
		filename      string
		expectedValue int
	}{
		{"testdata/for/for.vbs", 1},
		{"testdata/for/for_multi.vbs", 2},
		{"testdata/for/for_nest.vbs", 3},
		{"testdata/for/for_triple_nest.vbs", 6},
		{"testdata/for/for_whitespace.vbs", 2},
	}
	for _, test := range tests {
		vbscript := domain.VBScript{}
		internal.Read(test.filename, &vbscript)

		if test.expectedValue != vbscript.CognitiveComplexity {
			t.Errorf("wrong cognitive complexity output: got %v, expected %v (%s)", vbscript.CognitiveComplexity, test.expectedValue, test.filename)
		}
	}
}
