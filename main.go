package main

import (
	"embed"
	"embed-test/server"
	"fmt"
)

var (
	//go:embed hello.txt
	hello embed.FS

	//go:embed assets/nested.txt
	nested embed.FS
)

func main() {
	fmt.Printf("file: %v\n", hello)
	data, _ := hello.ReadFile("hello.txt")
	fmt.Println(string(data))

	fmt.Printf("file: %v\n", nested)
	data2, _ := nested.ReadFile("assets/nested.txt")
	fmt.Println(string(data2))

	server.Run()
}
