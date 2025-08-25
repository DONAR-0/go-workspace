package main

import "github.com/donar-0/go-workspace/l/inheri"

func main() {
	myDog := inheri.Dog{
		Animal: inheri.Animal{
			Name: "Buddy",
		},
	}
	myDog.Eat()
	myDog.Bark()
}
