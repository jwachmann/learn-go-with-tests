package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Height: 10.0, Width: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 2.0}, want: 20.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 2, Height: 3}, want: 3},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("Got '%g' want '%g'", got, tt.want)
			}
		})
	}
}
