package bridge

import (
	"fmt"
)

// decoupling abstraction from implementation
type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {
	//
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle with Radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (v *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle with Radius", radius)
}

type Circle struct {
	Renderer Renderer
	Radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{
		Renderer: renderer,
		Radius:   radius,
	}
}
func (c *Circle) Draw() {
	c.Renderer.RenderCircle(c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}
