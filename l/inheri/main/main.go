package main

import (
	"fmt"

	"github.com/donar-0/go-workspace/l/inheri"
)

func main() {
	myDog := inheri.Dog{
		Animal: inheri.Animal{
			Name: "Buddy",
		},
	}
	myDog.Eat()
	myDog.Bark()

	myCar := inheri.Car{Engine: inheri.Engine{HorsePower: 200}, Wheels: inheri.Wheels{Count: 4}}
	fmt.Println("HorsePower: ", myCar.HorsePower)
	fmt.Println("Wheels: ", myCar.Count)

	// ---  Method Overiding
	myDog.Speak()
	// Animal is parent
	myDog.Animal.Speak()
}
