package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

//Паттерн «Фасад» предоставляет унифицированный интерфейс вместо набора интерфейсов некоторой подсистемы.
//Фасад определяет интерфейс более высокого уровня, который упрощает использование подсистемы.

// NewMan создает новго человека
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Man реализует человека и фасад
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Что можно сделать с человеком
func (m *Man) Facade() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

// House реализует подсистему дом
type House struct {
}

// Build реализация.
func (h *House) Build() string {
	return "Build house"
}

type Tree struct {
}

func (t *Tree) Grow() string {
	return "Tree grow"
}

type Child struct {
}

func (c *Child) Born() string {
	return "Child born"
}

func Pat1() {
	m := NewMan()
	fmt.Println(m.Facade())

}
