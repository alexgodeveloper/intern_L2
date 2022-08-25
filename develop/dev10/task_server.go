package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// Создаем сервер телнет
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(conn)
	}
}

// Обработчик обращения к серверу
func handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}
		fmt.Print("Message incoming: " + message)
		conn.Write([]byte("Message received:" + message))
	}
}
