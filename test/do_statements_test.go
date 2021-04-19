package test_test

import (
	"testing"

	"github.com/st-tech/refactoring-conductor/domain"
	"github.com/st-tech/refactoring-conductor/internal"
)

func TestDoStatements(t *testing.T) {
	tests := []struct {
		filename      string
		expectedValue int
	}{
		{"testdata/do/do.vbs", 1},
		{"testdata/do/do_while.vbs", 1},
		{"testdata/do/do_until.vbs", 1},
		{"testdata/do/do_nest.vbs", 3},
		{"testdata/do/do_multi.vbs", 2},
		{"testdata/do/do_whitespace.vbs", 2},
		{"testdata/do/do_triple_nest.vbs", 6},
	}
	for _, test := range tests {
		vbscript := domain.VBScript{}
		internal.Read(test.filename, &vbscript)

		if test.expectedValue != vbscript.CognitiveComplexity {
			t.Errorf("wrong cognitive complexity output: got %v, expected %v (%s)", vbscript.CognitiveComplexity, test.expectedValue, test.filename)
		}
	}
}
