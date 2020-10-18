package main

import (
	"fmt"

	"design-patterns-in-go/composite"
)

func main() {
	drawing := composite.GraphicObject{Name: "My Drawing"}
	drawing.Children = append(drawing.Children, *composite.NewCircle("Red"))
	drawing.Children = append(drawing.Children, *composite.NewSquare("Yellow"))

	group := composite.GraphicObject{Name: "Group 1"}
	group.Children = append(drawing.Children, *composite.NewCircle("Blue"))
	group.Children = append(drawing.Children, *composite.NewSquare("Green"))
	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())

	// neuron layers
	neuron1, neuron2 := &composite.Neuron{}, &composite.Neuron{}
	layer1, layer2 := composite.NewNeuronLayer(3), composite.NewNeuronLayer(4)

	composite.Connect(neuron1, neuron2)
	composite.Connect(neuron1, layer1)
	composite.Connect(layer2, neuron1)
	composite.Connect(layer1, layer2)
}
