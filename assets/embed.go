package assets

import (
	"embed"
	"fmt"
)

var (
	//go:embed nested.txt
	nested embed.FS
)

func GetNested() {
	fmt.Println("from embed.go")
	fmt.Printf("file: %v\n", nested)
	data, _ := nested.ReadFile("nested.txt")
	fmt.Println(string(data))
}
