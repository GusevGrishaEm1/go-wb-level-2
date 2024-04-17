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
		} else {
			return strings.Contains(line, pattern)
		}
	}

	pattern := flag.Arg(0)
	if pattern == "" {
		fmt.Fprintln(os.Stderr, "Error: No pattern provided")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	ha
	for scanner.Scan() {
		line := scanner.Text()
		if filter(line) {

		}
		lines = append(lines, line)
	}

	for i, line := range lines {
		if filter(line) {
			if config.Invert {
				continue
			}
			if config.Count {
				count++
				continue
			}
			if config.LineNumber {
				fmt.Print(i + 1, ":")
			}
			fmt.Println(line)
		} else {
			if config.Invert {
				if config.Count {
					count++
					continue
				}
				if config.LineNumber {
					fmt.Print(i + 1, ":")
				}
				fmt.Println(line)
			}
		}

	}
		line := scanner.Text()


		if filter(line) {

			if config.Invert {
				continue
			}

			if config.Count {
				count++
				continue
			}

			if config.Context != 0 || config.After != 0 || config.Before != 0 {
				if lastMatchLine == 0 {
					lastMatchLine = count
				}

				if config.Before != 0 && count-lastMatchLine <= config.Before {
					if config.LineNumber {
						fmt.Print(count, ":")
					}
					fmt.Println(line)
				} else if config.Context != 0 && count-lastMatchLine <= config.Context {
					if config.LineNumber {
						fmt.Print(count, ":")
					}
					fmt.Println(line)
				} else if config.After != 0 && count-lastMatchLine > config.After {
					if config.LineNumber {
						fmt.Print(count, ":")
					}
					fmt.Println(line)
				}
			} else {
				if config.LineNumber {
					fmt.Print(count, ":")
				}
				fmt.Println(line)
			}
		} else {
			if config.Invert {
				if config.Count {
					count++
				}
				if config.LineNumber {
					fmt.Print(count, ":")
				}
				fmt.Println(line)
			}
		}
	}

	if config.Count {
		fmt.Print(count)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}

type Config struct {
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

func parseFlags() Config {
	config := Config{}

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
