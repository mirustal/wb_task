package inheritance

/*
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)
*/

import "fmt"

type Action struct {
	Human
}

type Human struct {
	name   string
	gender string
	age    int8
}

func (human Human) Say() string {
	return fmt.Sprintf("%s say: hello interviewer!", human.name)
}

func Inheritance(name string) {
	action := Action{
		Human{
			name:   name,
			gender: "male",
			age:    21}}

	fmt.Println(action.Say())
}
