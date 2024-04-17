package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		args        []string
		expectedOut string
	}{
		{
			name:        "No flags",
			input:       "Hello World\nGoodbye World",
			args:        []string{"World"},
			expectedOut: "Hello World\nGoodbye World",
		},
		{
			name:        "Ignore case",
			input:       "Hello World\nGoodbye World",
			args:        []string{"-i", "world"},
			expectedOut: "Hello World\nGoodbye World",
		},
		{
			name:        "Invert",
			input:       "Hello World\nGoodbye Wwow",
			args:        []string{"-v", "World"},
			expectedOut: "Goodbye Wwow",
		},
		{
			name:        "Count",
			input:       "Hello World\nGoodbye World\nHello ff",
			args:        []string{"-c", "World"},
			expectedOut: "2",
		},
		{
			name:        "Line number",
			input:       "Hello World\nGoodbye World\nHello World",
			args:        []string{"-n", "World"},
			expectedOut: "1:Hello World\n2:Goodbye World\n3:Hello World",
		},
		{
			name:        "After",
			input:       "1\n2\n3\n4\n5",
			args:        []string{"-A", "2", "3"},
			expectedOut: "3\n4\n5",
		},
		// {
		// 	name:        "Before",
		// 	input:       "1\n2\n3\n4\n5",
		// 	args:        []string{"-B", "2", "3"},
		// 	expectedOut: "1\n2\n3",
		// },
		// {
		// 	name:        "Context",
		// 	input:       "1\n2\n3\n4\n5",
		// 	args:        []string{"-C", "2", "3"},
		// 	expectedOut: "1\n2\n3\n4\n5",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "task.go")
			cmd.Args = append(cmd.Args, tt.args...)
			cmd.Stdin = bytes.NewBufferString(tt.input)
			output, err := cmd.Output()
			if err != nil {
				t.Errorf("Failed to execute command: %v", err)
			}
			if strings.TrimRight(string(output), "\n") != tt.expectedOut {
				t.Errorf("Expected output: %s, but got: %s", tt.expectedOut, string(output))
			}
		})
	}
}
