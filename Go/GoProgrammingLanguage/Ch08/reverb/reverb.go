package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() { // 处理同一个客户端连接中的输入
		go echo(conn, input.Text(), 1*time.Second) // 对于同一个连接中陆续收到的内容，启动并发 goroutine 来进行处理
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn) // 启动并发 goroutine 来处理客户端链接
	}
}
