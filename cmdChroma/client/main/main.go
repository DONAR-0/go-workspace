package main

import (
	"fmt"
	"log"

	"github.com/donar0/favecli/onnx"
)

func main() {
	// 1. Initialize the service once
	engine, err := onnx.NewEmbedder(
		"./models/minilm/model.onnx",
		"./models/minilm/tokenizer.json",
		"./models/onnx_runtime/onnxruntime-linux-x64-1.24.2/lib/libonnxruntime.so",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Close()

	// 2. Use it as many times as you want!
	vec1, _ := engine.Embed("What is ChromaDB?")
	vec2, _ := engine.Embed("How do I use ONNX with Go?")

	fmt.Printf("Vector 1 (first 3): %v\n", vec1[:3])
	fmt.Printf("Vector 2 (first 3): %v\n", vec2[:3])
}
