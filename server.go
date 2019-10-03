package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handlerConn(conn net.Conn) {

	log.Println("new connection")
	defer conn.Close()
	conn.Write([]byte(fmt.Sprintf("Welcom to %s, from %s \n\n", conn.LocalAddr(), conn.RemoteAddr())))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			return
		}
		fmt.Println("Text:", text)
		conn.Write([]byte(fmt.Sprintf("I have recived text: %s \n\n", text)))
	}
	log.Println("Connection close", conn.RemoteAddr())
}

func server() {
	l, err := net.Listen("tcp", "0.0.0.0:3302")
	log.Println("Hello I am server!")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlerConn(conn)
	}
}

func main() {
	server()
}
