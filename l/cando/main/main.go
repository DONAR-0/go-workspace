package main

import (
	"fmt"

	"github.com/donar-0/go-workspace/l/cando"
)

func main() {
	car1 := cando.Car{
		Color: "Red",
		Make:  "Toyota",
		Model: "Corolla ",
		Year:  2020,
	}
	car2 := cando.Car{
		Color: "Blue",
		Make:  "Ford",
		Model: "Mustang",
		Year:  2021,
	}

	fmt.Printf("\n-------------Car 1 --------------------\n")

	car1.DisplayInfo()
	fmt.Printf("\n-------------Car 2 --------------------\n")
	car2.DisplayInfo()
	fmt.Println()
}
