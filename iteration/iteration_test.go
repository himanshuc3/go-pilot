package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// NOTE
// 1. Benchmarks are also first-class citizens in Go.
// 2. `go test -bench=.` runs benchmarks.
// 3. Strings in Go are immutable, meaning every concatenation, such as in our Repeat function,
// involves copying memory to accommodate the new string
func BenchmarkRepeat(b *testing.B) {
	// setup code runs before each benchmark
	for b.Loop() {
		Repeat("a") // running the code being benchmarked
	}
	// teardown code runs after each benchmark
}
