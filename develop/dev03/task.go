package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	config := parseFlags()
	input := []string{"d a", "c b", "b c", "a d"}
	result, err := sortByConfig(input, config)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	fmt.Print(result)
}

type config struct {
	Key              string
	SortByNumerical  bool
	ReverseSort      bool
	UniqueValues     bool
	SortByMonth      bool
	IgnoreTrailing   bool
	CheckSorted      bool
	SortSuffixValues bool
}

func parseFlags() config {
	config := config{}

	flag.StringVar(&config.Key, "k", "", "Specify the column to sort by")
	flag.BoolVar(&config.SortByNumerical, "n", false, "Sort by numerical value")
	flag.BoolVar(&config.ReverseSort, "r", false, "Sort in reverse order")
	flag.BoolVar(&config.UniqueValues, "u", false, "Do not display duplicate lines")
	flag.BoolVar(&config.SortByMonth, "M", false, "Sort by month name")
	flag.BoolVar(&config.IgnoreTrailing, "b", false, "Ignore trailing whitespace")
	flag.BoolVar(&config.CheckSorted, "c", false, "Check if the data is sorted")
	flag.BoolVar(&config.SortSuffixValues, "h", false, "Sort by numerical value with suffixes")

	flag.Parse()

	return config
}

type element struct {
	str string
	key string
}

func sortByConfig(input []string, c config) ([]string, error) {
	elementsToSort := make([]element, 0, len(input))

	// Определяем элементы и их ключи для сортировки
	for _, v := range input {
		if c.Key != "" {
			arrKeys := strings.Split(v, " ")
			numKey, err := strconv.Atoi(c.Key)
			if err != nil {
				return nil, err
			}
			elementsToSort = append(elementsToSort, element{v, arrKeys[numKey-1]})
		} else {
			elementsToSort = append(elementsToSort, element{v, v})
		}
	}

	// Определяем функцию для сортировки
	compare := getCompareFunc(c)

	// Сортируем
	sort.Slice(elementsToSort, func(i, j int) bool {
		return compare(elementsToSort[i].key, elementsToSort[j].key)
	})

	// Убираем уникальные элементы
	if c.UniqueValues {
		mp := make(map[string]struct{}, len(elementsToSort))
		output := make([]string, 0, len(elementsToSort))
		for _, vl := range elementsToSort {
			_, ok := mp[vl.str]
			if ok {
				continue
			}
			mp[vl.str] = struct{}{}
			output = append(output, vl.str)
		}
		return output, nil
	}

	output := make([]string, 0, len(elementsToSort))
	for _, vl := range elementsToSort {
		output = append(output, vl.str)
	}
	return output, nil
}

// получаем функцию для сравнения элементов
func getCompareFunc(config config) func(v1 string, v2 string) bool {
	if config.SortByNumerical {
		return func(v1 string, v2 string) bool {
			num1, err1 := strconv.Atoi(v1)
			num2, err2 := strconv.Atoi(v2)
			if err1 == nil && err2 == nil {
				if config.ReverseSort {
					return num1 > num2
				}
				return num1 < num2
			}
			return false
		}
	}
	return func(v1 string, v2 string) bool {
		if config.ReverseSort {
			return v1 > v2
		}
		return v1 < v2
	}
}
