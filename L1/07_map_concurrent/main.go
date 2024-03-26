package main

import (
	"fmt"
	"sync"
)


type Counters struct {
    mx sync.Mutex
    m map[string]int
}

func (c *Counters) Load(key string) (val int) {
    c.mx.Lock()
    defer c.mx.Unlock()
    val, _ = c.m[key]
    return val
}

func (c *Counters) Save(key string, value int) {
    c.mx.Lock()
    defer c.mx.Unlock()
    c.m[key] = value
}

func NewCounters() *Counters {
    return &Counters{
        m: make(map[string]int),
    }
}

func main() {
	m := NewCounters()
	m.Save("wildberries", 10)
	fmt.Println(m.Load("wildberries"))
}