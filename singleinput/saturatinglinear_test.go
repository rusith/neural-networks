package singleinput

import "testing"


func TestSaturatedLinear(t *testing.T) {
	n := NewSingleInputNeuron(1, "s-linear")
	a := n.Do(-5, 1)
	if a != 0 { t.Error("a should be 0 for -5, 1")}
	a = n.Do(-1, 1)
	if a != 0 { t.Error("a should be 0 for -1, 1")}
	a = n.Do(-1, 0)
	if a != 1 { t.Error("a should be 0 for -1, 0")}
	a = n.Do(-0.5, 1)
	if a != 0.5 { t.Error("a should be 0.5 for -0.5, 1")}
	a = n.Do(5, 2)
	if a != 1 { t.Error("a should be 1 for 5, 2")}
}