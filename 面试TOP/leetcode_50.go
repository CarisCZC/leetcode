package main

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n < 0 {
		return 1 / powHelp(x, -n)
	}
	return powHelp(x, n)
}

func powHelp(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	half := powHelp(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
