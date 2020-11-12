package test_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestCognitiveComplexity(t *testing.T) {
	filename := "testdata/test.vbs"

	// ファイルオープン
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		// ここで一行ずつ処理
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		// エラー処理
	}
}
