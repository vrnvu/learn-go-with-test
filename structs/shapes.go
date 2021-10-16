package structs

import "math"

// Shape trait Area
type Shape interface {
	Area() float64
}

// Rectangle holds width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle holds radius
type Circle struct {
	Radius float64
}

// Triangle holds base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter of width and height
func Perimeter(r Rectangle) float64 {
	return 2.0 * (r.Width + r.Height)
}

// Area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area of triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2.0
}
