package main

import "testing"

// NOTE
// 1. Testing is a first class citizen in Go.
// 2. Testing files rules: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go - mostly everything is suffixed/prefixed with "test".
func TestHello(t *testing.T) {
	got := hello("Sarmooh")
	want := "Hello, Sarmooh!"

	if got != want {
		// NOTE
		// 1. `t.Errorf` has string formatting
		// 2. Run `go test` to run the tests - need to initiate a go project first using `go mod init <project-name>`
		t.Errorf("got %q want %q", got, want)
	}
}
