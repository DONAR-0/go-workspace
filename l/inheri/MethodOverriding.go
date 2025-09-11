package inheri

import "fmt"

func (a Animal) Speak() {
	fmt.Println(a.Name, "makes a sound")
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "Barks")
}
