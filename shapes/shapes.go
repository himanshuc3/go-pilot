package shapes

import (
	"math"
)

// NOTE:
// 1. For a language which is meant to be fast and efficient,
// interfaces are highly difficult to read and to be maintained.
// 2. On the other hand, they are easy to implement, use and
// provide decoupling between different parts of the code.
// 3. Interface resolution is implicit
type IShape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// NOTE:
// 1. Go and it's foolish minimalistic approach to object oriented programming
// with absurd this context name.
// 2. No function overloading in Go, so we are switching over to interfaces.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
