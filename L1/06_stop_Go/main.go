package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ch  chan bool
	wg *sync.WaitGroup
	ctx context.Context
}

func main() {
	thirdMethod()
}

func firstMethod() {
	ch := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	worker := Worker{
		ch: ch,
		wg: &wg,
		ctx: ctx,
	}

	wg.Add(1)
	go func() {
		select {
		case <-time.After(5 * time.Second):
			cancel()
		}
	}()
	go worker.work()
	wg.Wait()
}

func secondMethod() {
	ch := make(chan bool)
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	var wg sync.WaitGroup
	worker := Worker{
		ch: ch,
		wg: &wg,
		ctx: ctx,
	}
	wg.Add(1)
	go worker.work()
	wg.Wait()
}

func thirdMethod() {
	ch := make(chan bool)
	ctx := context.Background()
	var wg sync.WaitGroup
	worker := Worker{
		ch: ch,
		wg: &wg,
		ctx: ctx,
	}
	wg.Add(1)
	go worker.work()
	time.Sleep(5 * time.Second)
	ch <- true
	wg.Wait()


}

func (w *Worker) work() {
	start := time.Now()
	for{
		select{
		case <-w.ctx.Done():
			fmt.Printf("\nWork worked: %.0f seconds\n", time.Since(start).Seconds())
			defer w.wg.Done()
			return
		case <- w.ch:
			fmt.Printf("\nWork worked: %.0f seconds\n", time.Since(start).Seconds())
			defer w.wg.Done()
			return
		default:
			time.Sleep(time.Second)
			fmt.Print("work ")

		}
	}

}