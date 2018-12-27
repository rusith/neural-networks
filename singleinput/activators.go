package singleinput

func hardlim(n float32) float32 {
	if n < 0 { return 0 }
	return 1
}

func linear(n float32) float32 {
	return n
}

func sHardlim(n float32) float32  {
	if n < 0  { return -1 }
	return 1
}

func saturatedLinear(n float32) float32  {
	if n >= 0 && n <= 1 { return n }
	if n < 0 { return 0 }
	return 1
}