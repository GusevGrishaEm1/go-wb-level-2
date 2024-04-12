package main

import (
	"os/exec"
	"testing"
	"time"
)

// Test #1: Check output datetime format.
func TestPrintNTP(t *testing.T) {
	cmd := exec.Command("go", "run", "task.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
	 	t.Errorf("Error running command: %v", err)
	}
	if err := isDateValid(string(out)); err != nil {
		t.Errorf("Date is not valid: %v", err)
	}
}

func isDateValid(dateString string) error {
	_, err := time.Parse(time.RFC3339, dateString)
	return err
}