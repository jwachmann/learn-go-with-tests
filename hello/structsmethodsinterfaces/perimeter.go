package structsmethodsinterfaces

import "math"

// Shape represents any shape my dude
type Shape interface {
	Area() float64
}

// Rectangle represents a rectangle my dude
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a rectangle with the given width and height
func Perimeter(rectangle Rectangle) float64 {
	return (2 * rectangle.Width) + (2 * rectangle.Height)
}

// Area calculates the area of a rectangle with the given width and height
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle represents a circle my dude
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle with the given width and height
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Triangle is a triangle my dude
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
