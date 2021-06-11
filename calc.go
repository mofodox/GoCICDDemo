package calc

// Add two numbers and return the result.
func Add(val1, val2 int) int {
	return val1 + val2
}

// Subtract two numbers and return the result.
func Subtract(val1, val2 int) int {
	return val1 - val2
}

// Multiply two numbers and return the result.
func Multiply(val1, val2 int) int {
	return val1 * val2
}

// Divide two numbers and return the result.
func Divide(val1, val2 int) float64 {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	return float64(val1) / float64(val2)
}
