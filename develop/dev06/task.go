package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var fields string
var delimiter string
var separated bool

type Cuter struct {
	sl        []string
	Fields    string
	Delimiter string
	Separated bool
	result    string
}

//Возвращает массив строки разделенной по делителю
func (c *Cuter) split(text string) []string {
	return strings.Split(text, c.Delimiter)
}

func (c *Cuter) Cut(text string) string {

	c.sl = c.split(text)
	// Если не нашлись разделители
	if len(c.sl) <= 1 {
		// Не выводим строки если нет разделителя
		if c.Separated {
			return ""
		}
		c.result = c.sl[0]
		return c.result
	}

	d := strings.Split(c.Fields, ",")

	for _, v := range d {
		k, _ := strconv.Atoi(string(v))

		if len(c.sl)-1 > k {
			c.result += c.sl[k] + " "
		}
	}

	return c.result
}

func main() {
	flag.StringVar(&fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "только строки с разделителем")
	flag.Parse()
	text := flag.Arg(0)

	c := Cuter{
		Fields:    fields,
		Delimiter: delimiter,
		Separated: separated,
	}
	fmt.Println(fields, delimiter, separated)

	res := c.Cut(text)
	fmt.Println(res)
}
