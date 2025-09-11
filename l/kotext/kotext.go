package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Worker: stopping, reason:", ctx.Err())
				return
			default:
				log.Printf("Worker: working...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	log.Printf("Main: Cancelling the worker")
	cancel()
	time.Sleep(1 * time.Second)
}
