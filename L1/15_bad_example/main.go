package main

import "fmt"



/*
	происхоидт утечка памяти, 1 << 10 создает строку из 1024 символа, а затем мы запоминаем ее в глобальную переменную и сохраняем первые 100, но строка хранит в себе выделенную память
	var justString string
	func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	}

	func main() {
	someFunc()
	}
*/



func createHugeString(len int) string {
	str := make([]rune, len, len)
    return string(str)
}

func someFunc() {
    v := createHugeString(1 << 10)
    justString := string(v[:100])

	justString = v[:100]
    fmt.Println(len(justString))
}

func main() {
    someFunc()
}