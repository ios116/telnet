package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

func sendingEmail(ctx context.Context, id int) {
	randValue := rand.Intn(5000)
	duration := time.Duration(randValue) * time.Millisecond
	select {
	case <-time.After(duration):
		log.Printf("task is complited %d duration %d", id, randValue)
	case <-ctx.Done():
		log.Printf("task %d is cansaled", id)
	}
}

func handlerConn(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte(fmt.Sprintf("Hello! local host %s, remote host %s ", conn.LocalAddr(), conn.RemoteAddr())))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			return
		}
		fmt.Println("Text:", text)
		conn.Write([]byte(fmt.Sprintf("I have recived text: %s", text)))
	}
	log.Println("Connection close", conn.RemoteAddr())
}

func server() {
	l, err := net.Listen("tcp", "0.0.0.0:3302")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handlerConn(conn)
	}
}

func someCont(){
	w := sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func() {
			sendingEmail(ctx, i)
			cancel()
			w.Done()
		}()
	}
	w.Wait()
}

func main() {
   server()
}
