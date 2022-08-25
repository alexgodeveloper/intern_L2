package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	g := Grep{
		after:      2,
		before:     1,
		context:    1,
		count:      true,
		ignoreCase: true,
		invert:     false,
		fixed:      false,
		lineNum:    true,
		lines:      nil,
		grepStr:    "душой",
		index:      0,
		amount:     0,
	}

	want := []string{
		"Текущая строка:",
		"их мало, с опытной душой,",
		"Строки после совпадения:",
		"В файле недостаточно строк",
		"Строки до совпадения:",
		"не падал, не блевал и не ругался?",
		"ну кто ж из нас на палубе большой",
		"Строки до совпадения:",
		"не падал, не блевал и не ругался?",
		"ну кто ж из нас на палубе большой",
		"Строки после совпадения:",
		"В файле недостаточно строк",
		"Кол-во строк в файле:",
		"4",
		"Текущий номер строки:",
		"3",
	}

	f, err := os.Open("text.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fscan := bufio.NewScanner(f)
	sl = readScan(fscan)
	g.lines = sl
	r := GrepF(&g)

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Массив из файла %s не равен ожидаемому %s", g.result, want)
	}
}
