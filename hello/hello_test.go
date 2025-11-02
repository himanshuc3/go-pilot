package main

import (
	"bytes"
	"testing"
)

// NOTE
// 1. Testing is a first class citizen in Go.
// 2. Testing files rules: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go - mostly everything is suffixed/prefixed with "test".
// 3. Test functions are like describe blocks and subtests are individual case assertions.
func TestHello(t *testing.T) {
	t.Run("saying hello to Sarmooh, which I haven't been able to do after being blocked.", func(t *testing.T) {
		got := hello("Sarmooh", "")
		want := "Hello, Sarmooh!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, stranger!' when an empty string is passed", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, stranger!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := hello("Sarmooh", "Spanish")
		want := "Hola, Sarmooh!"
		assertCorrectMessage(t, got, want)
	})

}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sarmooh")

	got := buffer.String()
	expected := "Hello, Sarmooh!"

	assertCorrectMessage(t, got, expected)
}

// NOTE
// 1. `testing.TB` is an interface that is implemented by `*testing.T` and `*testing.B`.
// 2. Shorten parameter types
func assertCorrectMessage(t testing.TB, got, want string) {
	// NOTE
	// 1. `t.Helper()` is used to mark the function as a helper.
	// 2. Helpers help report the line number in the error message from original test suite instead of the
	// helper.
	t.Helper()
	if got != want {
		// NOTE
		// 1. `t.Errorf` has string formatting
		// 2. Run `go test` to run the tests - need to initiate a go project first using `go mod init <project-name>`
		t.Errorf("got %q want %q", got, want)
	}
}
