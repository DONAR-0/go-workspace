package main

import (
	"log"
	"net/http"

	"github.com/donar-0/go-workspace/l/di"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(di.MyGreeterHandler)))
}

