package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius int
}
type Triangle struct {
	Base   int
	Height int
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius*c.Radius)
}

func (t Triangle) Area() float64 {
	return 0.5 * float64(t.Base*t.Height)
}
