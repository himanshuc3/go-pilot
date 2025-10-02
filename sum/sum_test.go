package sum

import "testing"

func TestSum(t *testing.T) {
	// NOTE
	// 1. Defining array using 2 ways:
	// 2. [N]type{values} & [...]type{values}
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
