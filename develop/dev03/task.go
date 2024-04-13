package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-m — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	config := ParseFlags()
	result, err := Sort("test1.txt", config)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	fmt.Print(result)
}

type Config struct {
	Column           string
	SortNumerical    bool
	ReverseSort      bool
	UniqueValues     bool
	SortByMonth      bool
	IgnoreTrailing   bool
	CheckSorted      bool
	SortSuffixValues bool
}

func ParseFlags() Config {
	config := Config{}

	flag.StringVar(&config.Column, "k", "", "Specify the column to sort by")
	flag.BoolVar(&config.SortNumerical, "n", false, "Sort by numerical value")
	flag.BoolVar(&config.ReverseSort, "r", false, "Sort in reverse order")
	flag.BoolVar(&config.UniqueValues, "u", false, "Do not display duplicate lines")
	flag.BoolVar(&config.SortByMonth, "M", false, "Sort by month name")
	flag.BoolVar(&config.IgnoreTrailing, "b", false, "Ignore trailing whitespace")
	flag.BoolVar(&config.CheckSorted, "c", false, "Check if the data is sorted")
	flag.BoolVar(&config.SortSuffixValues, "h", false, "Sort by numerical value with suffixes")

	flag.Parse()

	return config
}

func Sort(input string, c Config) ([]string, error) {
	cmd := exec.Command("sort")

	if c.Column != "" {
		cmd.Args = append(cmd.Args, "-k"+c.Column)
	}

	if c.SortNumerical {
		cmd.Args = append(cmd.Args, "-n")
	}

	if c.ReverseSort {
		cmd.Args = append(cmd.Args, "-r")
	}

	if c.UniqueValues {
		cmd.Args = append(cmd.Args, "-u")
	}

	if c.SortByMonth {
		cmd.Args = append(cmd.Args, "-M")
	}

	if c.IgnoreTrailing {
		cmd.Args = append(cmd.Args, "-b")
	}

	if c.CheckSorted {
		cmd.Args = append(cmd.Args, "-c")
	}

	if c.SortSuffixValues {
		cmd.Args = append(cmd.Args, "-h")
	}

	cmd.Args = append(cmd.Args, input)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(string(out), "\r\n"), nil
}
