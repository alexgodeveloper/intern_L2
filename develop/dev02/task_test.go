package main

import "testing"

func TestStringUnpacking(t *testing.T) {
	m := make(map[string]string)
	m["a4bc2d5e"] = "aaaabccddddde"
	m["abcd"] = "abcd"
	m["45"] = ""
	m["qwe\\4\\5"] = "qwe45"
	m["qwe\\45"] = "qwe44444"
	m["qwe\\\\5"] = "qwe\\\\\\\\\\"

	for k, v := range m {
		if s, _ := StringUnpacking(k); s != v {
			t.Error("result != value")
		}
	}
}
