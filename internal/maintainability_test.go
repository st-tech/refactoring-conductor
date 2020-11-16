package internal_test

import (
	"testing"

	"github.com/st-tech/search-tools/2020internship-yoshikawa/src/domain"
)

func TestCalculateMaintainability(t *testing.T) {
	// MAX(0,(171 - 5.2 * ln(Halstead Volume) - 0.23 * (Cyclomatic Complexity) - 16.2 * ln(Lines of Code))*100 / 171)
	// 1. n1 - ユニークなオペレーションコード数（命令、関数、式など）
	// 2. n2 - ユニークなオペランド数（変数、パラメータ、ファイル名など）
	// 3. N1 - オペレーションコードの総数
	// 4. N2 - サブルーチンのすべての変数インスタンスの総数
	// 5. A = n1 + n2 (プログラム内で使用されている命令、変数などの種類数、最小単位）
	// 6. B = N1 + N2 (プログラム内で使用されている命令、変数などの総数）
	// 7. Halstead Volume = B * log(2)A (プログラムソースコードの情報量）
	expectedMaintainability := 0
	vbscript := domain.VBScript{}

	if expectedMaintainability != vbscript.Maintainability {
		t.Errorf("wrong output: got %v, expected %v", vbscript.Maintainability, expectedMaintainability)
	}
}
