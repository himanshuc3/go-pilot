package sum

func Sum(numbers [5]int) int {
	sum := 0
	// NOTE
	// 1. range is the declarative way of looping over an array.
	// 2. Opposite parameter convention to js, first is the index and second is the value.
	for _, number := range numbers {
		sum += number
	}
	return sum
}
