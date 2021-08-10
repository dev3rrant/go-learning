package main

import "testing"

func TestPerimeter(t *testing.T) {
	assertEqualFloats := func(t *testing.T, expected, got float64) {
		t.Helper()
		if expected != got {
			t.Errorf("expected %.2f, received %.2f", expected, got)
		}
	}
	t.Run("get perimeter of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		expected := 40.0
		assertEqualFloats(t, expected, got)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{3, 5}, want: 7.5},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			hasArea := tt.shape.Area()
			if hasArea != tt.want {
				t.Errorf("got %g want %g", hasArea, tt.want)
			}
		})
	}
}
