package main

import "fmt"

type NeuronInterface interface {
	Iter() []*Neuron
}

type Neuron struct {
	In, Out []*Neuron
}

func NewNeuron() *Neuron {
	return &Neuron{}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

type NeuronLayer struct {
	Neurons []Neuron
}

func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}
	return result
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func neuralNetworkExample() {
	n1, n2 := NewNeuron(), NewNeuron()
	l1, l2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(n1, n2)
	Connect(n1, l1)
	Connect(l2, n1)
	Connect(l1, l2)

	for _, el := range n1.Out {
		fmt.Println(el)
	}
}
