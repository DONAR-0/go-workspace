package main

import (
	"bufio"
	"fmt"
	"net"
)

const sockPath = "/tmp/echo.sock"

func main() {
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "ping")
	fmt.Fprintln(conn, "hanuman ji ki jai")
	fmt.Fprintln(conn, "quit")

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
