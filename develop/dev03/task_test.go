package main

import (
	"reflect"
	"testing"
)

func TestSortWithColumn(t *testing.T) {
	config := Config{Column: "2"}

	expected := []string{"d a", "c b", "b c", "a d"}

	output, err := Sort("test1.txt", config)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected: %v, got: %v", expected, output)
	}
}

// func TestSortNumerical(t *testing.T) {
// 	input := []string{"3", "1", "10", "2"}
// 	config := Config{SortNumerical: true}
// 	expected := "1n2n3n10n"
// 	result, err := Sort(input, config)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if result != expected {
// 		t.Errorf("Sort didn't produce the expected result. Got: %s, Expected: %s", result, expected)
// 	}
// }
