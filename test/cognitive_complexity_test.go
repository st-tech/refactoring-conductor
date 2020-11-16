package test_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestCognitiveComplexity(t *testing.T) {
	filename := "testdata/if.vbs"

	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("err: %v", err)
	}
}
