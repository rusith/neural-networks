package singleinput

import "testing"

func TestNewSingleInputNeuron(t *testing.T) {
	n := NewSingleInputNeuron(1, "hardlin")
	if n == nil {
		t.Error("NewSingleInputNeuron Doesn't Work")
	}
}

func TestUnmatchedTransfer(t *testing.T) {
	n := NewSingleInputNeuron(1, "test")
	if n.Do(1, 1) != 0 {
		t.Error("Unmatched transfer function names should return 0")
	}
}
