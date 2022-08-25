package pattern

import (
	"errors"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
//Шаблон "Фабрика" - это порождающий шаблон проектирования, а также один из наиболее часто используемых шаблонов. Этот шаблон позволяет скрыть логику создания генерируемых экземпляров.
//Клиент взаимодействует только с фабричной структурой и сообщает, какие экземпляры необходимо создать. Класс фабрики взаимодействует с соответствующими конкретными структурами и возвращает правильный экземпляр.

type Gun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) SetName(name string) {
	g.name = name
}

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) SetPower(power int) {
	g.power = power
}

func (g *gun) GetPower() int {
	return g.power
}

type ak47 struct {
	gun
}

func NewAk47() Gun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type maverick struct {
	gun
}

func NewMaverick() Gun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}

func GetGun(gunType string) (Gun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "maverick" {
		return NewMaverick(), nil
	}
	return nil, errors.New("wrong gun type passed")
}

func Task6() {
	ak47, err := GetGun("ak47")
	if err != nil {
		log.Fatalf("Cannot create ak47 gun. Error %v", err)
	}
	maverick, err := GetGun("maverick")
	if err != nil {
		log.Fatalf("Cannot create maverick gun. Error %v", err)
	}
	fmt.Println(ak47)
	fmt.Println(maverick)
}
