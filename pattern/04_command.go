package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

//Команда" — это поведенческий шаблон проектирования. Он предлагает инкапсулировать запрос как отдельный объект.
//Созданный объект имеет всю информацию о запросе и поэтому может выполнять его независимо.

type Command interface {
	Execute()
}

type Device interface {
	On()
	Off()
}

//Стуктура кнопка
type button struct {
	command Command
}

func NewButton(command Command) *button {
	return &button{
		command: command,
	}
}

func (b *button) Press() {
	b.command.Execute()
}

//Включение
type onCommand struct {
	device Device
}

func NewOnCommand(device Device) *onCommand {
	return &onCommand{
		device: device,
	}
}

func (c *onCommand) Execute() {
	c.device.On()
}

//Выключение
type offCommand struct {
	device Device
}

func NewOffCommand(device Device) *offCommand {
	return &offCommand{
		device: device,
	}
}

func (c *offCommand) Execute() {
	c.device.Off()
}

//Структура ...
type tv struct {
	isRunning bool
}

func NewTv() *tv {
	return &tv{}
}

func (t *tv) On() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) Off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func Task4() {
	tv1 := NewTv()
	onCommand := NewOnCommand(tv1)
	offCommand := NewOffCommand(tv1)
	onButton := NewButton(onCommand)
	onButton.Press()
	offButton := NewButton(offCommand)
	offButton.Press()
}
