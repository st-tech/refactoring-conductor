package test_test

import (
	"testing"

	"github.com/st-tech/refactoring-conductor/domain"
	"github.com/st-tech/refactoring-conductor/internal"
)

func TestWhileStatements(t *testing.T) {
	tests := []struct {
		filename      string
		expectedValue int
	}{
		{"testdata/while/while.vbs", 1},
		{"testdata/while/while_nest.vbs", 3},
		{"testdata/while/while_triple_nest.vbs", 6},
		{"testdata/while/while_multi.vbs", 2},
		{"testdata/while/while_wend.vbs", 2},
	}
	for _, test := range tests {
		vbscript := domain.VBScript{}
		internal.Read(test.filename, &vbscript)

		if test.expectedValue != vbscript.CognitiveComplexity {
			t.Errorf("wrong cognitive complexity output: got %v, expected %v (%s)", vbscript.CognitiveComplexity, test.expectedValue, test.filename)
		}
	}
}
