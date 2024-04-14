package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слон", "нос"}
	expected := map[string][]string{
		"акптя":  {"пятак", "пятка", "тяпка"},
		"иклост": {"листок", "слиток", "столик"},
	}

	result := findAnagrams(words)

	for key, anagrams := range expected {
		sort.Strings(anagrams)
		expected[key] = anagrams
	}
	for _, value := range result {
		sort.Strings(value)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFindAnagramsEmpty(t *testing.T) {
	words := []string{}
	expected := map[string][]string{}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
