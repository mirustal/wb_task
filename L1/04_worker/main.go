package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
)

type Worker struct {
	ID  int
	ctx context.Context
	ch  <-chan int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go run main.go <num_workers>")
		return
	}

	N, _ := strconv.Atoi(os.Args[1])
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

for i := 1; i <= N; i++ {
		worker := Worker{
			ID:  i,
			ctx: ctx,
			ch:  ch,
		}
		go worker.work()
	}
	
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		<-sigCh
		fmt.Println("Signal ctrl+c")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case ch <- rand.Intn(100):
		}
	}
}


func (w *Worker) work() {
	for {
		select {
		case <-w.ctx.Done():
			fmt.Printf("Worker %d stop\n", w.ID)
			return
		case num := <-w.ch:
			fmt.Printf("Worker %d received: %d\n", w.ID, num)
		}
	}
}