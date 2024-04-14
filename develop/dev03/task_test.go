package main

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	input := []string{"d a", "c b", "b c", "a d"}
	config := Config{
		Key:             "1",
		SortByNumerical: false,
		ReverseSort:     true,
		UniqueValues:    false,
	}

	expected := []string{"d a", "c b", "b c", "a d"}
	result, _ := Sort(input, config)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sort function did not sort as expected. Expected: %v, Got: %v", expected, result)
	}
}

func TestGetCompareFunc(t *testing.T) {
	config1 := Config{SortByNumerical: true, ReverseSort: true}
	compareFunc1 := getCompareFunc(config1)

	if !compareFunc1("10", "5") {
		t.Errorf("Numerical reverse sort failed")
	}

	config2 := Config{SortByNumerical: false, ReverseSort: false}
	compareFunc2 := getCompareFunc(config2)

	if !compareFunc2("apple", "banana") {
		t.Errorf("String sort failed")
	}
}
