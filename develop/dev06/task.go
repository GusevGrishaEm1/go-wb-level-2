package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Fields    string
	Delimiter string
	Separated bool
}

func main() {
	config := parseFlags()
	lines := scanInput()
	output, err := cut(lines, config)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(output)
}

func parseFlags() Config {
	config := Config{}

	config.Fields = *flag.String("f", "", "Specify the columns to output")
	config.Delimiter = *flag.String("d", "\t", "Specify the delimiter")
	config.Separated = *flag.Bool("s", false, "Output only the lines that contain the delimiter")

	flag.Parse()

	return config
}

func scanInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func cut(lines []string, config Config) (string, error) {
	var output strings.Builder
	fieldsMap := make(map[int]bool)
	if config.Fields != "" {
		fields := strings.Split(config.Fields, ",")
		for _, field := range fields {
			index, err := strconv.Atoi(field)
			if err != nil {
				return "", err
			}
			fieldsMap[index] = true
		}
	}
	for i := 0; i < len(lines); i++ {
		columns := strings.Split(lines[i], config.Delimiter)
		if config.Separated && !strings.Contains(lines[i], config.Delimiter) {
			continue
		}
		for i, col := range columns {
			if config.Fields == "" || fieldsMap[i+1] {
				output.WriteString(col)
				if i != len(columns)-1 {
					output.WriteString("\t")
				}
			}
		}
		if i != len(lines)-1 {
			output.WriteString("\n")
		}
	}
	return output.String(), nil
}
