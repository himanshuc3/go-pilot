package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		// NOTE
		// 1. Defining array using 2 ways:
		// 2. [N]type{values} & [...]type{values}
		// 3. `go test -cover` shows the coverage of the tests.
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// NOTE
	// 1. Cannot use equality operator to compare slices.
	// 2. `reflect.DeepEqual` is not type safe.
	// 3. slices.Equal is used for shallow comparison.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	// NOTE
	// 1. This is a closure, a function that can access variables defined outside of it.
	// 2. Another advantage is code isolation/encapsulation, not allowing other functions to use it outside of the test function.
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 10})
		want := []int{5, 19}

		checkSums(t, got, want)
	})

	// NOTE
	// 1. Compile time errors are our friend because they help
	// us write software that works, runtime errors are our
	// enemies because they affect our users.
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)

	})

}
