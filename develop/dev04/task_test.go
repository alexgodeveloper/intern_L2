package main

import (
	"reflect"
	"testing"
)

func TestAnagrams(t *testing.T) {
	anagrams := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Лунь", "нуль", "горечь"}

	testM := make(map[string][]string)
	testM["листок"] = append(testM["листок"], "листок", "слиток", "столик")
	testM["пятак"] = append(testM["пятак"], "пятак", "пятка", "тяпка")
	testM["лунь"] = append(testM["лунь"], "лунь", "нуль")

	m := anagram(anagrams)
	for k := range m {
		if ok := reflect.DeepEqual(testM[k], m[k]); !ok {
			t.Error("result != value")
		}
	}
}
