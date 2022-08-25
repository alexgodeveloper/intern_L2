package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	c := Cuter{
		sl:        nil,
		Fields:    "1,3",
		Delimiter: "\t",
		Separated: false,
		result:    "",
	}

	want := "два три пять "

	text := "раз	два три	четыре	пять	шесть"
	r := c.Cut(text)

	if !reflect.DeepEqual(r, want) {
		t.Errorf("Строка из ввода: %s не равна ожидаемой: %s", c.result, want)
	}
}
