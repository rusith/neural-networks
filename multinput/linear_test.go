package multiinput

import "testing"

func TestMultiInputLinear(t *testing.T) {
	ne := NewMultiInputNeuron(1, "linear")
	res:= ne.Do([]float32{2,3}, []float32{1,1})
	if res != 5 { t.Error("Linear should be 5")}
	res = ne.Do([]float32{0,1}, []float32{1,1})
	if res != 1 { t.Error("Linear should be 1")}
}