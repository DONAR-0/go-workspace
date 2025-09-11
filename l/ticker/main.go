package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			case <-done:
				fmt.Println("Ticker Stopped")
				return
			}
		}
	}()

	time.Sleep(7 * time.Second)

	done <- true
}
