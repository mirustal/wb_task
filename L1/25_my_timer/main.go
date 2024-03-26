package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	second := 2
	fmt.Printf("Запущен сон на %v секунды\n", second)
	Sleep(2)
	fmt.Println("Прошло времени:", time.Since(start))	
}


func Sleep(sec int) {
    <-time.After(time.Duration(sec) * time.Second)
}

