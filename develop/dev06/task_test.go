package main

import (
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	config := config{
		Fields:    "2,3",
		Delimiter: ",",
		Separated: false,
	}

	input := "1,2,3\n4,5,6\n7,8,9"
	expectedOutput := "2\t3\n5\t6\n8\t9"

	output, err := cut(strings.Split(input, "\n"), config)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if output != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, output)
	}
}
