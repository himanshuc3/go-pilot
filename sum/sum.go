package sum

func Sum(numbers []int) int {
	sum := 0
	// NOTE
	// 1. range is the declarative way of looping over an array.
	// 2. Opposite parameter convention to js, first is the index and second is the value.
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// NOTE
// 1. `...` is used to pass a variable number of arguments to a function and collect 'em all.
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	// NOTE
	// 1. `make` is used to create a slice of a given length.
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
