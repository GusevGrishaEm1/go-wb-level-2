package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func grep(data []byte, pattern string, flags map[string]bool) []string {
	lines := strings.Split(string(data), "n")
	matchedLines := make([]string, 0)

	regExp := regexp.MustCompile(pattern)

	for i, line := range lines {
		matched := regExp.MatchString(line)

		if flags["v"] {
			matched = !matched
		}

		if matched {
			matchedLines = append(matchedLines, line)
		} else {
			context := flags["A"] + flags["B"]
			if context > 0 && len(matchedLines) > 0 {
				if flags["A"] > 0 {
					for j := 1; j <= flags["A"]; j++ {
						if i+j < len(lines) {
							matchedLines = append(matchedLines, lines[i+j])
						}
					}
				}
				matchedLines = append(matchedLines, line)
				if flags["B"] > 0 {
					for j := 1; j <= flags["B"]; j++ {
						if i+j < len(lines) {
							matchedLines = append(matchedLines, lines[i+j])
						}
					}
				}
			}
		}
	}

	return matchedLines
}

func main() {
	after := flag.Int("A", 0, "Print N lines after match")
	before := flag.Int("B", 0, "Print N lines before match")
	context := flag.Bool("C", false, "Print N lines around match")
	count := flag.Bool("c", false, "Print count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert match")
	fixed := flag.Bool("F", false, "Exact string match")
	lineNum := flag.Bool("n", false, "Print line number")

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: grep [flags] pattern file")
		os.Exit(1)
	}

	data, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	pattern := args[0]
	if *ignoreCase {
		pattern = "(?i)" + pattern
	}
	if *fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	flags := map[string]bool{
		"A": *after,
		"B": *before,
		"C": *context,
		"c": *count,
		"i": *ignoreCase,
		"v": *invert,
		"F": *fixed,
		"n": *lineNum,
	}

	matchedLines := grep(data, pattern, flags)

	if *count {
		fmt.Println("Number of matching lines:", len(matchedLines))
	} else {
		for i, line := range matchedLines {
			if *lineNum {
				fmt.Printf("%d: %sn", i+1, line)
			} else {
				fmt.Println(line)
			}
		}
	}
}
