package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func or(channels ...<-chan interface{}) <-chan interface{} {
	// слайс select-case
	var s []reflect.SelectCase
	c := make(chan interface{})

	for _, ch := range channels {
		// добавляем неизвестное количество каналов в слайс с кейсами
		s = append(s, reflect.SelectCase{
			Dir:  reflect.SelectRecv,  // Если Dir — SelectRecv, случай представляет операцию получения: case <-Chan
			Chan: reflect.ValueOf(ch), // Конкретный канал на чтение
		})
	}

	go func() {
		// Как и в обычном select, здесь будет блокировка до тех пор, пока не будет выполнено хотя бы одно
		// из чтений done-каналов
		reflect.Select(s)
		// После чтения из канала закроется канал, который возобновит работу go-main
		close(c)
	}()

	return c
}

func main() {

	var sig = func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	// Ожидаем запись в канал, что основанная горутина не завершилась всех остальных
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
