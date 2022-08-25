package main

import (
	"fmt"
	"sort"
	"strings"
)

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

func anagram(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, w := range words {
		word := strings.ToLower(w)
		wordSorted := sortString(word)
		anagrams[wordSorted] = append(anagrams[wordSorted], word)
	}

	res := make(map[string][]string)
	for _, v := range anagrams {
		if len(v) > 1 {
			res[v[0]] = v
			sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
		}
	}
	return res
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Лунь", "нуль", "горечь"}
	fmt.Println("Словарь:\n\t", words)

	anagrams := anagram(words)

	fmt.Println("Анаграммы:")
	for k, v := range anagrams {
		fmt.Println("\t", k, v)
	}

}
