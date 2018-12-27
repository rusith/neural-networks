package singleinput

func hardlim(n int) int {
	if n < 0 { return 0 }
	return 1
}

func linear(n int) int {
	return n
}

func sHardlim(n int) int  {
	if n < 0  { return -1 }
	return 1
}