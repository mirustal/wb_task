package main

import (
	"context"
	"fmt"
	"sync"
)

type Worker struct {
	ch  chan int
	wg *sync.WaitGroup
	ctx context.Context
}

func (w *Worker) Reader() {
	for {
		select {
		case num := <-w.ch:
			fmt.Printf("Reader receive and square:  %d\n", num*num)
			w.wg.Done() 
		case <-w.ctx.Done():
			return
		}
	}
}

func (w *Worker) Writer(arr []int) {
	for i := 0; i < len(arr); i++ { 
		select {
		case <-w.ctx.Done():
			return
		case w.ch <- arr[i]:
		}
	}
}

func main() {
	

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	worker := Worker{
		ch: ch,
		wg: wg,
		ctx: ctx,
	}
	
	arr := []int{2, 4, 6, 8, 10}
	wg.Add(len(arr))
	go worker.Reader()
	go worker.Writer(arr)

	wg.Wait()
	cancel()

}