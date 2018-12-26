package singleinput

import "testing"

func TestHardlim(t *testing. T) {
	n := NewSingleInputNeuron(1, "hardlim")
	a := n.Do(-5, 2)
	if a != 0 { t.Error("a should be 0 for -5, 2")}
	a = n.Do(-1, 1)
	if a != 1 { t.Error("a should be 1 for -1, 1")}
	a = n.Do(5, 2)
	if a != 1 { t.Error("a should be 1 for 5, 2")}
}