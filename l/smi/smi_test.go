package smi

import (
	"testing"

	"github.com/DONAR-0/go-workspace/assertions/pkg/tablewriter"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}
	got := rectangle.Perimeter()
	want := 40.0

	tablewriter.AssertFloatGotWant(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		tablewriter.AssertFloatGotWant(t, got, want)
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}

		checkArea(t, rectangle, 100.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{
			Radius: 10.0,
		}

		checkArea(t, circle, 314.1592653589793)
	})

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 12, Height: 6}, 72.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
