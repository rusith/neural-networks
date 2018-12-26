package singleinput

import "testing"
func TestSHardlim(t *testing.T) {
	n := NewSingleInputNeuron(1, "s-hardlim")
	a := n.Do(-5, 1)
	if a != -1 { t.Error("a should be -1 for -5, 1")}
	a = n.Do(-1, 1)
	if a != 1 { t.Error("a should be 1 for -1, 1")}
	a = n.Do(5, 2)
	if a != 1 { t.Error("a should be 1 for 5, 2")}
}