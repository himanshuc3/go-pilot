package integers

import (
	"fmt"
	"testing"
)

// NOTE
// 1. `go test` takes the context of the current package and tests for all the functions in the package.
func TestAdder(t *testing.T) {
	sum := Add(3, 3)
	expected := 6

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// NOTE
// 1. Executable examples are a great way to document your code
// and shows up as executable examples in godoc.
// 2. Names always start with Example
// 3. `go test` by default runs the tests and examples.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

// NOTE
// 1. The 'Ouput comment' above ensures Example code is executed
// rather than just being compiled when running `go test`.
