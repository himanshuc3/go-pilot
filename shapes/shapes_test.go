package shapes

import (
	"math"
	"testing"
)

// NOTE:
// 1. go test -run TestPerimeter => to provide more specificity to testing
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {
	// NOTE:
	// 1. I would probably prefer to prefix my interfaces with I,
	// to atleast have it be very obvious and hence easier to read.
	checkArea := func(t testing.TB, shape IShape, expected float64) {
		t.Helper()
		got := shape.Area()
		if math.Abs(got-expected) > 0.01 {
			// NOTE:
			// 1. %g has more precision than %f.
			// 2. Ensure test messages are meaningful and helpful, because
			// god blesseth, you were to maintain the code one
			// year from now?
			t.Errorf("got %g :: expected area to be %g", got, expected)
		}
	}

	// NOTE:
	// 1. Creating slice of anonymous struct
	areaTests := []struct {
		shape    IShape
		expected float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.16},
	}

	// t.Run("rectangles", func(t *testing.T) {

	// 	rectangle := Rectangle{12.0, 6.0}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.16)
	// })

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.expected)
	}
}
