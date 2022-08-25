package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

var after int
var before int
var context int
var count bool
var ignoreCase bool
var invert bool
var fixed bool
var lineNum bool

var fileName string
var sl []string
var str string

type Grep struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	lines      []string
	grepStr    string
	index      int
	amount     int
	result     []string
}

// Сканируем файл
func readScan(scan *bufio.Scanner) []string {
	s := make([]string, 0)

	for scan.Scan() {
		s = append(s, scan.Text())
	}

	return s
}

//Ищем индекс текущей строки
func FindIndex(g *Grep) {
	var k = 0 //для подсчета совпадений
	for i, r := range g.lines {
		if strings.Contains(r, g.grepStr) {
			g.index = i
			k++
		}
	}
	if k == 0 {
		g.result = append(g.result, "Нет совпадений по строке")
		os.Exit(0)
	}

}

//Печатает текущую строку
func CurrentString(g *Grep) {
	g.result = append(g.result, "Текущая строка:", g.lines[g.index])

}

// Печать N строк после совпадения
func After(g *Grep) {
	g.result = append(g.result, "Строки после совпадения:")
	j := g.index + g.after
	k := g.index + 1
	if j < len(g.lines) {
		for k < len(g.lines) {
			g.result = append(g.result, g.lines[k])
			k++
		}
	} else {
		g.result = append(g.result, "В файле недостаточно строк")

	}
}

// Печать N строк до совпадения
func Before(g *Grep) {
	g.result = append(g.result, "Строки до совпадения:")
	j := g.index - g.before
	k := g.index - 1
	if j >= 0 {
		for k >= 0 {
			g.result = append(g.result, g.lines[k])
			k--
		}

	} else {
		g.result = append(g.result, "В файле недостаточно строк")

	}
}

// Печать N строк вокруг совпадения
func Context(g *Grep) {
	Before(g)
	After(g)

}

// Печать кол-ва строк в файле
func PrintCount(g *Grep) {
	r := strconv.Itoa(len(g.lines))
	g.result = append(g.result, "Кол-во строк в файле:", r)

}

// Текущий номер строки
func CurrentNumStr(g *Grep) {
	r := strconv.Itoa(g.index + 1)
	g.result = append(g.result, "Текущий номер строки:", r)

}

// Обработчик условий
func GrepF(g *Grep) []string {
	if g.ignoreCase && !g.fixed {
		var lsl []string
		for _, r := range g.lines {
			lsl = append(lsl, strings.ToLower(r))
		}
		g.lines = lsl
		g.grepStr = strings.ToLower(g.grepStr)
	}
	FindIndex(g)

	if !g.invert {
		CurrentString(g)
	}
	if g.after > 0 {
		After(g)

	}

	if g.before > 0 {
		Before(g)
	}

	if g.context > 0 {
		Context(g)

	}

	if g.count {
		PrintCount(g)
	}

	if g.lineNum {
		CurrentNumStr(g)
	}

	return g.result

}

func main() {
	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&count, "c", false, "количество строк")
	flag.BoolVar(&ignoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&lineNum, "n", false, "напечатать номер строки")
	flag.Parse()
	fileName = flag.Arg(0)
	str = flag.Arg(1)

	// Открываем файл
	r, err := os.Open(fileName)
	defer func(r *os.File) {
		err := r.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Создаем сканера, который считает данные по строчно и добавит их в слайс
	sc := bufio.NewScanner(r)
	sl = readScan(sc)
	// создаем структуру из переданных параметров и флагов пользователя
	g := &Grep{
		lines:      sl,
		ignoreCase: ignoreCase,
		after:      after,
		before:     before,
		context:    context,
		count:      count,
		invert:     invert,
		fixed:      fixed,
		lineNum:    lineNum,
	}

	GrepF(g)
	fmt.Println(g.result)
}
