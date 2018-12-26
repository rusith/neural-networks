package singleinput

// Neuron is a neuron which has only one input and one output
type Neuron struct {
	Bias int
	Activation string
}

// NewSingleInputNeuron creates a new SingleInputNeuron 
func NewSingleInputNeuron(b int, t string) *Neuron {
	return &Neuron{ Bias: b, Activation: t }
}

// Do will calculate result
func (ne *Neuron) Do(p int, w int) int {
	n := (p * w) + ne.Bias
	switch ne.Activation {
	case "hardlim":
		return hardlim(n)
	}

	return 0
}