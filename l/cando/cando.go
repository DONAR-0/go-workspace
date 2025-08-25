package cando

import "fmt"

type Car struct {
	Color, Make, Model string
	Year               int
}

func (c Car) DisplayInfo() {
	fmt.Printf("%v", c)
}
