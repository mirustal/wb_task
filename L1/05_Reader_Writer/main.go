package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Worker struct {
	ch  chan int
	ctx context.Context
}

func (w *Worker) Reader() {
	for {
		select {
		case num := <-w.ch:
			fmt.Printf("Reader receive: %d\n", num)
		case <-w.ctx.Done():
			return
		}
	}
}

func (w *Worker) Writer() {
	i := 0
	for {
		select {
		case <-w.ctx.Done():
			return
		case w.ch <- i:
			i++
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go run main.go <time works>")
		return
	}
	N, _ := strconv.Atoi(os.Args[1])

	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	worker := Worker{
		ch: ch,
		ctx: ctx,
	}

	go worker.Reader()
	go worker.Writer()

	duration := time.Duration(N) * time.Second
	time.Sleep(duration)
	cancel()

}