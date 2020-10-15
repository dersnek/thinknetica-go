package fibonacci

// Calculate returns Fibonacci number with a given index
func Calculate(n int) int {
	if n == 0 {
		return 0
	}

	a, b := 0, 1

	for i := 0; i < n-1; i++ {
		a, b = b, a+b
	}

	return b
}
