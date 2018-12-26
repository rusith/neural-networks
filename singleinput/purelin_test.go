package singleinput

import "testing"
func TestPurelin(t *testing. T) {
	n := NewSingleInputNeuron(1, "purelin")
	a := n.Do(-5, 1)
	if a != -4 { t.Error("a should be -4 for -5, 1")}
	a = n.Do(-1, 1)
	if a != 0 { t.Error("a should be 0 for -1, 1")}
	a = n.Do(5, 2)
	if a != 11 { t.Error("a should be 11 for 5, 2")}
}