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
