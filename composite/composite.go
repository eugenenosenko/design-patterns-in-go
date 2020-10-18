package composite

import (
	"strings"
)

type NeuronInterface interface {
	Iter() []*Neuron
}

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)

	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}

	return result
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{
		Name: "Circle", Color: color}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{
		Name:  "Square",
		Color: color,
	}
}
