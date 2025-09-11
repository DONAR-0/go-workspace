package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var sockPath = "/tmp/echo.sock"

func main() {
	_ = os.Remove(sockPath)

	ln, err := net.Listen("unix", sockPath)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	_ = os.Chmod(sockPath, 0660)
	log.Println("UDS echo server listening on", sockPath)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				if ctx.Err() != nil {
					errCh <- nil
					return
				}

				errCh <- err

				return
			}

			go handleConn(conn)
		}
	}()

	select {
	case <-ctx.Done():
		log.Printf("\nShutting Down")
	case err := <-errCh:
		if err != nil {
			fmt.Println("listen/accept error:", err)
		}
	}

	_ = ln.Close()
	_ = os.Remove(sockPath)

	fmt.Println("clean exit")
}

func handleConn(c net.Conn) {
	defer c.Close()

	addr := c.RemoteAddr()
	fmt.Fprintf(c, "Hello from %s\n", filepath.Base(sockPath))

	sc := bufio.NewScanner(c)
	for sc.Scan() {
		line := sc.Text()
		// Echo the line back
		fmt.Fprintf(c, "echo: %s\n", line)
	}
	// If needed, check sc.Err() for read errors
	fmt.Println("client disconnected:", addr)
}
