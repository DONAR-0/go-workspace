package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/DONAR-0/go-workspace/assertions/pkg/utils"
)

const sockPath = "/tmp/echo.sock"

func main() {
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		panic(err)
	}

	utils.DeferCheck(conn.Close)

	_, _ = fmt.Fprintln(conn, "ping")
	_, _ = fmt.Fprintln(conn, "hanuman ji ki jai")
	_, _ = fmt.Fprintln(conn, "quit")

	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
