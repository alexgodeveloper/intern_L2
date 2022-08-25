package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

//Шаблон "Посетитель" - это поведенческий паттерн проектирования, который позволяет добавлять поведение в структуру без фактического её изменения.

type Shape interface {
	GetType() string
	Accept(Visitor)
}

type square struct {
	side int
}

func NewSquare(side int) *square {
	return &square{
		side: side,
	}
}

func (s *square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *square) GetType() string {
	return "Square"
}

type Visitor interface {
	VisitForSquare(*square)
}

type areaCalculator struct {
	area int
}

func NewAreaCalculator() *areaCalculator {
	return &areaCalculator{}
}

func (a *areaCalculator) VisitForSquare(s *square) {
	// Вычисляем площадь для квадрата. После вычисления площади присваиваем её в
	// переменную area экземпляра
	a.area = s.side * 2
	fmt.Printf("Calculating area for square:%d\n", a.area)
}

func Task3() {
	s := NewSquare(2)
	ac := NewAreaCalculator()
	s.Accept(ac)
}
