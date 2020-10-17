package main

import (
	"fmt"

	"design-patterns-in-go/adapter"
)

func main() {
	rc := adapter.NewRectangle(6, 4)
	a := adapter.VectorToRaster(rc)
	fmt.Print(adapter.DrawPoints(a))
}
