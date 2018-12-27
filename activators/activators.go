package activators

func Hardlim(n float32) float32 {
	if n < 0 { return 0 }
	return 1
}

func Linear(n float32) float32 {
	return n
}

func SHardlim(n float32) float32  {
	if n < 0  { return -1 }
	return 1
}

func SaturatedLinear(n float32) float32  {
	if n >= 0 && n <= 1 { return n }
	if n < 0 { return 0 }
	return 1
}