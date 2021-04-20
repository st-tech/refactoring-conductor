package test_test

import (
	"testing"

	"github.com/st-tech/refactoring-conductor/domain"
	"github.com/st-tech/refactoring-conductor/internal"
)

func TestSubStatements(t *testing.T) {
	tests := []struct {
		filename      string
		expectedValue int
	}{
		{"testdata/sub/sub.vbs", 2},
		{"testdata/sub/public_sub.vbs", 3},
		{"testdata/sub/public_default_sub.vbs", 4},
		{"testdata/sub/private_sub.vbs", 1},
	}
	for _, test := range tests {
		vbscript := domain.VBScript{}
		internal.Read(test.filename, &vbscript)

		if test.expectedValue != vbscript.CognitiveComplexity {
			t.Errorf("wrong cognitive complexity output: got %v, expected %v (%s)", vbscript.CognitiveComplexity, test.expectedValue, test.filename)
		}
	}
}
