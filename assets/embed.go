package assets

import (
	"embed"
	"fmt"
)

var (
	//go:embed nested.txt
	nested embed.FS

	//go:embed accelerate_fallback.svg
	accelerate_fallback embed.FS
)

func GetFallbackImage() {
	fmt.Println("from embed.go")
	fmt.Printf("file: %v\n", accelerate_fallback)
	data, _ := accelerate_fallback.ReadFile("accelerate_fallback.svg")
	fmt.Println(string(data))
}

func GetNested() {
	fmt.Println("from embed.go")
	fmt.Printf("file: %v\n", nested)
	data, _ := nested.ReadFile("nested.txt")
	fmt.Println(string(data))
}
