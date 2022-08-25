package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeout int
var host string
var port string

func main() {
	// Сигнал для завершение работы телнет
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGQUIT)

	conn := tcpConnect()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	go func() {
		for {
			str, err := reader.ReadString('\n')
			if err != nil {
				break
			}

			if len(str) > 0 {
				fmt.Print(str)
			}
		}
	}()
	// Читаем данные от сервера телнет
	go func() {
		defer close(sig)
		nr := bufio.NewScanner(os.Stdin)
		for nr.Scan() {
			fmt.Fprintf(conn, nr.Text()+"\n")
		}
	}()

	select {
	case <-sig:
		fmt.Println("end")
	}

}

// Создаем подключение к серверу телнет
func tcpConnect() net.Conn {
	flag.IntVar(&timeout, "timeout", 10, "таймаут")
	flag.Parse()

	to := time.Duration(timeout) * time.Second

	host = flag.Arg(0)
	port = flag.Arg(1)

	conn, err := net.DialTimeout("tcp", host+":"+port, to)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Connected to " + host + ":" + port + ". For exit press Ctrl+D\n")

	return conn
}
