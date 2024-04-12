package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
        - "a4bc2d5e" => "aaaabccddddde"
        - "abcd" => "abcd"
        - "45" => "" (некорректная строка)
        - "" => ""
Дополнительное задание: поддержка escape - последовательностей
        - qwe\4\5 => qwe45 (*)
        - qwe\45 => qwe44444 (*)
        - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func uncompressString(input string) (string, error) {
	var output string
	escape := false

	for _, char := range input {
		if escape || unicode.IsLetter(char) {
			output += string(char)
			escape = false
		} else {
			if char == '\\' {
				escape = true
			} else {
				if unicode.IsDigit(char) {
					if len(output) == 0 {
						return "", errors.New("некорректная строка")
					}
					times, _ := strconv.Atoi(string(char))
					output += strings.Repeat(string(output[len(output)-1]), times-1)
				}
			}
		}
	}

	return output, nil
}

func main() {
	output, err := uncompressString(`qwe\45`)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	fmt.Print(output)
}
