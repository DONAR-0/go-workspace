package main

import "github.com/donar-0/go-workspace/l/polymorphism"

func main() {
	var myAnimal polymorphism.Animal = polymorphism.Dog{Name: "Buddy"}
	myAnimal.Speak()
}
