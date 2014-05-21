package concurrency_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// checkExpectedOutputs checks that all expected outputs are
// generated.
func checkExpectedOutputs(t *testing.T, outputs, expectedOutputs []string) {
	for _, expected := range expectedOutputs {
		found := false
		for _, output := range outputs {
			if output == expected {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Expected output %v never generated", expected)
		}
	}
}

// checkUnexpectedOutputs checks that no unexpected outputs are
// generated.
func checkUnexpectedOutputs(t *testing.T, outputs, expectedOutputs []string) {
	for _, output := range outputs {
		found := false
		for _, expected := range expectedOutputs {
			if output == expected {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Unexpected output %v generated", output)
		}
	}
}

// cleanupClosure returns a function that is used as the cleanup
// function of systematic tests.
func cleanupClosure(out *os.File) func() {
	return func() {
		fmt.Fprintf(out, "\n")
	}
}

// length computes the number of keys in the given set that hold the
// value 'true'.
func length(s map[int]bool) int {
	n := 0
	for _, ok := range s {
		if ok {
			n++
		}
	}
	return n
}

// processOutput processes the output file, returning a slice of all
// output lines generated by a test.
func processOutput(t *testing.T, f *os.File) []string {
	buffer, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Fatalf("ReadFile() failed: %v", err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(buffer))
	result := make([]string, 0)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("Scanning output file failed: %v", err)
	}
	return result
}

// setup is used as the setup function of systematic tests.
func setup() {}
