package composite

type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(n.Out, other)
}

type NeuronLayer struct {
	Neurons []Neuron
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}
