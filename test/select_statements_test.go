package test_test

import (
	"testing"

	"github.com/st-tech/refactoring-conductor/domain"
	"github.com/st-tech/refactoring-conductor/internal"
)

func TestSelectStatements(t *testing.T) {
	tests := []struct {
		filename      string
		expectedValue int
	}{
		{"testdata/select/select.vbs", 1},
		{"testdata/select/select_multi.vbs", 2},
		{"testdata/select/select_nest.vbs", 3},
		{"testdata/select/select_triple_nest.vbs", 6},
		{"testdata/select/select_whitespace.vbs", 2},
	}
	for _, test := range tests {
		vbscript := domain.VBScript{}
		internal.Read(test.filename, &vbscript)

		if test.expectedValue != vbscript.CognitiveComplexity {
			t.Errorf("wrong cognitive complexity output: got %v, expected %v (%s)", vbscript.CognitiveComplexity, test.expectedValue, test.filename)
		}
	}
}
