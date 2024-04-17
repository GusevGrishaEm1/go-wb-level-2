package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		w := strings.ToLower(word)
		runes := []rune(w)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		sortedWord := string(runes)

		if _, ok := anagramMap[sortedWord]; !ok {
			anagramMap[sortedWord] = []string{w}
		} else {
			anagramMap[sortedWord] = append(anagramMap[sortedWord], w)
		}
	}

	for key, value := range anagramMap {
		if len(value) < 2 {
			delete(anagramMap, key)
		} else {
			sort.Strings(value)
			anagramMap[key] = value
		}
	}

	return anagramMap
}

func main() {
	// Пример использования
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)

	for key, value := range anagrams {
		fmt.Printf("Множество анаграмм с ключом %s: %vn", key, value)
	}
}
