package internal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/st-tech/refactoring-conductor/domain"
)

// Read is to read file and scan it line by line.
func Read(filename string, vbscript *domain.VBScript) int {
	fp, err := os.Open(filename)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		CountControlFlow(vbscript, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("err: %v", err)
	}

	return vbscript.CognitiveComplexity
}
