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
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	config := parseFlags()

	if config.FilePath != "" {
		file, err := os.Open(config.FilePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		io.Copy(os.Stdout, file)
	}

	filter := func(line string) bool {
		if config.IgnoreCase {
			line = strings.ToLower(line)
		}

		pattern := os.Args[len(os.Args)-1]

		if config.Fixed {
			return line == pattern
		}
		return strings.Contains(line, pattern)
	}

	pattern := flag.Arg(0)
	if pattern == "" {
		fmt.Fprintln(os.Stderr, "Error: No pattern provided")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	lineNumsToPrint := make(map[int]struct{}, 0)
	count := 0
	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		if filter(line) {
			if !config.Invert {
				lineNumsToPrint[idx] = struct{}{}
				if config.Count {
					count++
				}
				if config.Context != 0 {
					size := idx - config.Context
					for i := idx; i >= size; i-- {
						lineNumsToPrint[i] = struct{}{}
					}
					size = idx + config.Context
					for i := idx; i <= size; i++ {
						lineNumsToPrint[i] = struct{}{}
					}
				}
				if config.After != 0 {
					size := idx + config.After
					for i := idx; i <= size; i++ {
						lineNumsToPrint[i] = struct{}{}
					}
				}
				if config.Before != 0 {
					size := idx - config.Before
					for i := idx; i >= size; i-- {
						lineNumsToPrint[i] = struct{}{}
					}
				}
			}
		} else {
			if config.Invert {
				lineNumsToPrint[idx] = struct{}{}
				if config.Count {
					count++
				}
			}
		}
		lines = append(lines, line)
		idx++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	if config.Count {
		fmt.Print(count)
		return
	}

	for i, v := range lines {
		if _, ok := lineNumsToPrint[i]; ok {
			if config.LineNumber {
				fmt.Printf("%d:", i+1)
			}
			fmt.Print(v)
			if i != len(lines)-1 {
				fmt.Print("\n")
			}
		}
	}
}

type config struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNumber bool
	FilePath   string
}

func parseFlags() config {
	config := config{}

	flag.IntVar(&config.After, "A", 0, "print +N lines after a match")
	flag.IntVar(&config.Before, "B", 0, "print +N lines before a match")
	flag.IntVar(&config.Context, "C", 0, "print +N lines around a match")
	flag.BoolVar(&config.Count, "c", false, "count the number of lines")
	flag.BoolVar(&config.IgnoreCase, "i", false, "ignore case distinctions")
	flag.BoolVar(&config.Invert, "v", false, "select non-matching lines")
	flag.BoolVar(&config.Fixed, "F", false, "match exact string, not regex")
	flag.BoolVar(&config.LineNumber, "n", false, "print line number with output lines")
	flag.StringVar(&config.FilePath, "file", "", "path to the file to filter")

	flag.Parse()

	return config
}
