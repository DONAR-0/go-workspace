package main

import (
	"fmt"
	"os"

	"github.com/donar-0/go-workspace/l/mocking"
)

func main() {
	fmt.Printf("Something I am doing %v", "gello")
	// mocking.CountDown(os.Stdout)
	sleeper := &mocking.DefaultSleeper{}
	mocking.CountDown(os.Stdout, sleeper)
}
