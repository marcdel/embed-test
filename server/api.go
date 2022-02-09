package server

import (
	"embed-test/assets"
	"fmt"
)

func Run() {
	fmt.Println("Running the server!")

	assets.GetNested()

	assets.GetFallbackImage()
}
