package main

import (
	"context"
	"log"
	"math/rand"
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

func main()  {
    someCont()
}