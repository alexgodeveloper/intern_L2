package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
//Используйте шаблон проектирования "Состояние", когда объект может быть в различных состояниях. В зависимости от текущего запроса объекту необходимо изменять свое текущее состояние.
//Используйте шаблон, когда объект может по-разному отвечать на один и тот же запрос в зависимости от текущего состояния. В этом случае использование шаблона проектирования "Состояние" заменит множество условных операторов.

// MobileAlertStater обеспечивает общий интерфейс для различных состояний.
type MobileAlertStater interface {
	Alert() string
}

// MobileAlert реализует оповещение в зависимости от своего состояния.
type MobileAlert struct {
	state MobileAlertStater
}

// Alert возвращает строку
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState изменяет состояние
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert это конструктор MobileAlert.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration реализует вибрации
type MobileAlertVibration struct {
}

// Alert возвращает вибрации предупреждения
func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertSong реализует звуковой сигнал
type MobileAlertSong struct {
}

// Alert возвращает звуковое предупреждения
func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}

func Task8() {
	s := &MobileAlertSong{}
	ma := NewMobileAlert()

	// Вызываем оповещение телефона (по умолчанию вибрация)
	fmt.Println(ma.Alert())

	// Меняем состояние на звук
	ma.SetState(s)
	fmt.Println(ma.Alert())
}
