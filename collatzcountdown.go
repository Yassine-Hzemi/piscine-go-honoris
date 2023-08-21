package piscine

// CollatzCountdown function returns the number of steps
// necessary to reach 1 using the Collatz conjecture.
// It returns -1 if start is equal to 0 or negative.
func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	n := start
	steps := 0

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			steps++
		} else {
			n = 3*n + 1
			steps++
		}
	}
	return steps
}
