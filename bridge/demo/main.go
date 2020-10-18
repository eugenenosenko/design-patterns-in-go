package main

import (
	"fmt"

	"design-patterns-in-go/bridge"
)

func main() {
	raster := bridge.RasterRenderer{}
	vector := bridge.VectorRenderer{}
	circle := bridge.NewCircle(
		&vector, 32,
	)
	circle1 := bridge.NewCircle(
		&raster, 32,
	)

	fmt.Println("Rendering via vector", circle)
	fmt.Println("Rendering via raster", circle1)
}
