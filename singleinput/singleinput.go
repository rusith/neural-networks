package singleinput

import ac "github.com/rusith/neural-networks/activators"

// Neuron is a neuron which has only one input and one output
type Neuron struct {
	Bias float32
	Activation string
}

// NewSingleInputNeuron creates a new SingleInputNeuron 
func NewSingleInputNeuron(b float32, t string) *Neuron {
	return &Neuron{ Bias: b, Activation: t }
}

// Do will calculate result
func (ne *Neuron) Do(p float32, w float32) float32 {
	n := (p * w) + ne.Bias
	switch ne.Activation {
	case "hardlim":
		return ac.Hardlim(n)
	case "linear":
		return ac.Linear(n)
	case "s-linear": 
		return ac.SaturatedLinear(n)
	case "s-hardlim":
		return ac.SHardlim(n)
	}

	return 0
}