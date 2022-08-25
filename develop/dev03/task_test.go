package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {

	fs := FlagsSort{
		column:  1,
		reverse: true,
		unique:  true,
		byNum:   true,
	}

	want := []string{
		"4 aaaa",
		"3 afsa",
		"2 dsa",
		"1 asde",
	}

	f, err := os.Open("text.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fscan := bufio.NewScanner(f)
	sl = readScan(fscan)

	sl = strings.Split(string(Sort(sl, &fs)), "\n")

	if !reflect.DeepEqual(sl, want) {
		t.Errorf("Массив из файла %s не равен ожидаемому %s", sl, want)
	}

}
