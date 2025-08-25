package inheri

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Eat() {
	fmt.Println(a.Name, "is eating...")
}

type Dog struct {
	Animal // Embedding Animal struct
}

func (d Dog) Bark() {
	fmt.Println(d.Name, "is Barking")
}
