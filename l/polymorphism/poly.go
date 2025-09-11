package polymorphism

import "fmt"

type Animal interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "barks")
}
