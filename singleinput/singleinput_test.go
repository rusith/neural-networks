package singleinput

import "testing"

func TestNewSingleInputNeuron(t *testing.T) {
	n := NewSingleInputNeuron(1, "hardlin")
	if n == nil {
		t.Error("NewSingleInputNeuron Doesn't Work")
	}
}
