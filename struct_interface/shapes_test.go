package struct_interface

import "testing"

//type Rectangle struct {
//	Width float64
//	Height float64
//}

//func TestPerimeter(t *testing.T) {
//	rectangle := Rectangle{10.0, 10.0}
//	got := Perimeter(rectangle)
//	want := 40.0
//
//	if got != want {
//		t.Errorf("got %.2f want %.2f", got, want)
//	}
//}
//
//func TestArea(t *testing.T) {
//	rectangle := Rectangle{12.0, 6.0}
//	got := Area(rectangle)
//	want := 72.0
//
//	if got != want {
//		t.Errorf("got %.2f want %.2f", got, want)
//	}
//}
//func TestArea(t *testing.T) {
//
//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{12, 6}
//		got := rectangle.Area()
//		want := 72.0
//
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10}
//		got := circle.Area()
//		want := 314.1592653589793
//
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	})
//
//}
//
//func TestArea(t *testing.T) {
//
//	checkArea := func(t testing.TB, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//	t.Run("rectangles", func(t *testing.T) {
//		rectangle := Rectangle{12, 6}
//		checkArea(t, rectangle, 72.0)
//	})
//
//	t.Run("circles", func(t *testing.T) {
//		circle := Circle{10}
//		checkArea(t, circle, 314.1592653589793)
//	})
//
//}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base:   12, Height: 6}, want:  36},

	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}