package util

// A simple utility function that sums an array.
// Given the nature of the project, I'll be doing this
// a lot.
func SumArray(seq []int) int {
	sum := 0
	for _, num := range seq {
		sum += num
	}

	return sum
}
