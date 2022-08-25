package main

import (
	"errors"
	"fmt"
	"os"
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

func StringUnpacking(s string) (string, error) {
	// обработка случая пустой строки
	if len(s) == 0 {
		return s, nil
	}

	var result []rune
	rstr := []rune(s)

	// строка, начинающаяся с числа некорректная
	if unicode.IsDigit(rstr[0]) {
		return "", errors.New("Стока начинается с цифры.  Неправильный формат.")
	}

	for i := 0; i < len(rstr); i++ {
		switch {
		case unicode.IsDigit(rstr[i]):

			// n - количество повторений символа
			n := int(rstr[i] - '0')

			// число, начинающееся с 0, некорректно
			if n == 0 {
				return "", errors.New("Число начинается с 0.  ")
			}

			// если число повторений 1, то добавлять ничего не нужно
			if n == 1 {
				break
			}

			// слайс повторяющихся значений добавляется одним append'ом
			pack := make([]rune, n-1)
			for i := range pack {
				pack[i] = result[len(result)-1]
			}
			result = append(result, pack...)
		case rstr[i] == '\\':

			// одиночный '\' в конце строки некорректен
			if i == len(rstr)-1 {
				return "", errors.New("одиночный '\\' в конце строки некорректен")
			}

			// руна после '\' обрабатывается как символ
			result = append(result, rstr[i+1])
			i++
		default:
			result = append(result, rstr[i])
		}
	}

	return string(result), nil
}

func main() {

	s, err := StringUnpacking("qwe\\\\5")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		return
	}
	fmt.Println(s)

}

