package shapes

import "testing"

// NOTE:
// 1. go test -run TestPerimeter => to provide more specificity to testing
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {

	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	expected := 72.0

	if got != expected {
		t.Errorf("got %.2f :: expected %.2f", got, expected)
	}
}
