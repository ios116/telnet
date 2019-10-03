package main

import (
	"bufio"
	"context"
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

// ReadingAndWriter Reading from server and writing to server
func ReadingAndWriter(ctx context.Context, r io.Reader, w io.Writer) {
	ch := make(chan string)
	scanner := bufio.NewScanner(r)
	go func() {
		for {
			if scanner.Scan() {
				ch <- scanner.Text()
			}
		}
	}()
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		case text := <-ch:
			_, err := w.Write([]byte(fmt.Sprintf("The message: %s \n", text)))
			if err != nil {
				log.Println(err)
				break OUTER
			}
		}
	}
}

func main() {
	var timeout = flag.IntP("timeout", "t", 60, "connection timeout")
	flag.Parse()
	fmt.Println("time out is ", *timeout)
	dialer := &net.Dialer{}
	ctx := context.Background()
	dur := time.Duration(*timeout) * time.Second
	ctx, _ = context.WithTimeout(ctx, dur)
	conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:3302")
	if err != nil {
		log.Fatalf("Cannot conection: %v", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ReadingAndWriter(ctx, conn, os.Stdout)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		ReadingAndWriter(ctx, os.Stdin, conn)
		wg.Done()
	}()
	wg.Wait()
	conn.Close()
	log.Println("Finished reading routine")
}
