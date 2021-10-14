package slices

// Sum returns sum of numbers
func Sum(numbers []int) int {
	r := 0
	for _, n := range numbers {
		r += n
	}
	return r
}
