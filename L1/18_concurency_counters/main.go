package main

import (
	"fmt"
	"sync"
)

type Counters struct {
	mx  sync.Mutex
	wg  sync.WaitGroup
	num int
}

func (c *Counters) worker(workerId int) {
	defer c.wg.Done()
	c.mx.Lock()
	c.num++
	c.mx.Unlock()
	fmt.Printf("Worker %d, incremented\n", workerId)
}

func main() {
	var counters Counters
	for i := 1; i <= 10; i++ {
		counters.wg.Add(1)
		go counters.worker(i)
	}

	counters.wg.Wait()

	fmt.Printf("Increment value:%d\n", counters.num)
}
