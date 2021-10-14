package iteration

// Repeat repeats the sequence 5 times
func Repeat(sequence string, times int) string {
	chars := len(sequence)
	r := make([]rune, times*chars)
	idx := 0
	for i := 0; i < times; i++ {
		for c := 0; c < chars; c++ {
			r[idx] = rune(sequence[c])
			idx++
		}
	}
	return string(r)
}
