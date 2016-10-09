package model

import (
	"image"
	"image/color"
)

type TriangleFunc func()

func draw_line(v0, v1 Vec2i, img *image.RGBA, c color.Color) {
	Draw_line_solution(v0.X, v0.Y, v1.X, v1.Y, img, c)
}

func Triangle_first(t0, t1, t2 Vec2i, img *image.RGBA, c color.Color) {
	// sort the vertices, t0, t1, t2 lower−to−upper (bubblesort yay!)
	if t0.Y > t1.Y {
		t0, t1 = t1, t0
	}
	if t0.Y > t2.Y {
		t0, t2 = t2, t0
	}
	if t1.Y > t2.Y {
		t1, t2 = t2, t1
	}

	var red = color.RGBA{0xFF, 0, 0, 0xFF}
	var green = color.RGBA{0, 0xFF, 0, 0xFF}

	draw_line(t0, t1, img, green)
	draw_line(t1, t2, img, green)
	draw_line(t2, t0, img, red)
}

func Triangle_first_halve(t0, t1, t2 Vec2i, img *image.RGBA, c color.Color) {
	// sort the vertices, t0, t1, t2 lower−to−upper (bubblesort yay!)
	if t0.Y > t1.Y {
		t0, t1 = t1, t0
	}
	if t0.Y > t2.Y {
		t0, t2 = t2, t0
	}
	if t1.Y > t2.Y {
		t1, t2 = t2, t1
	}

	var red = color.RGBA{0xFF, 0, 0, 0xFF}
	var green = color.RGBA{0, 0xFF, 0, 0xFF}

	var total_height int = t2.Y - t0.Y
	for y := t0.Y; y <= t1.Y; y++ {
		var segment_height int = t1.Y - t0.Y + 1
		var alpha float64 = float64(y - (t0.Y)) / float64(total_height)
		var beta float64 = float64(y - (t0.Y)) / float64(segment_height) // be careful with divisions by zero

		var A Vec2i = t0.Add(t2.Subtract(t0).Scale(alpha))
		var B Vec2i = t0.Add(t1.Subtract(t0).Scale(beta))

		img.Set(A.X, y, red)
		img.Set(B.X, y, green)
	}
}

func Triangle_fill_first_halve(t0, t1, t2 Vec2i, img *image.RGBA, c color.Color) {
	// sort the vertices, t0, t1, t2 lower−to−upper (bubblesort yay!)
	if t0.Y > t1.Y {
		t0, t1 = t1, t0
	}
	if t0.Y > t2.Y {
		t0, t2 = t2, t0
	}
	if t1.Y > t2.Y {
		t1, t2 = t2, t1
	}

	var total_height int = t2.Y - t0.Y
	for y := t0.Y; y <= t1.Y; y++ {
		var segment_height int = t1.Y - t0.Y + 1
		var alpha float64 = float64(y - (t0.Y)) / float64(total_height)
		var beta float64 = float64(y - (t0.Y)) / float64(segment_height) // be careful with divisions by zero

		var A Vec2i = t0.Add(t2.Subtract(t0).Scale(alpha))
		var B Vec2i = t0.Add(t1.Subtract(t0).Scale(beta))

		if A.X > B.X {
			A, B = B, A
		}

		for x := A.X; x < B.X; x += 1 {
			img.Set(x, y, c)
		}
	}
}

func Triangle_fill_two_halves(t0, t1, t2 Vec2i, img *image.RGBA, c color.Color) {
	// sort the vertices, t0, t1, t2 lower−to−upper (bubblesort yay!)
	if t0.Y > t1.Y {
		t0, t1 = t1, t0
	}
	if t0.Y > t2.Y {
		t0, t2 = t2, t0
	}
	if t1.Y > t2.Y {
		t1, t2 = t2, t1
	}

	var total_height int = t2.Y - t0.Y
	for y := t0.Y; y <= t1.Y; y += 1 {
		var segment_height int = t1.Y - t0.Y + 1
		var alpha float64 = float64(y - (t0.Y)) / float64(total_height)
		var beta float64 = float64(y - (t0.Y)) / float64(segment_height) // be careful with divisions by zero

		var A Vec2i = t0.Add(t2.Subtract(t0).Scale(alpha))
		var B Vec2i = t0.Add(t1.Subtract(t0).Scale(beta))

		if A.X > B.X {
			A, B = B, A
		}

		for x := A.X; x < B.X; x += 1 {
			img.Set(x, y, c)
		}
	}

	for y := t1.Y; y <= t2.Y; y += 1 {
		var segment_height int = t2.Y - t1.Y + 1
		var alpha float64 = float64(y - (t0.Y)) / float64(total_height)
		var beta float64 = float64(y - (t1.Y)) / float64(segment_height) // be careful with divisions by zero

		var A Vec2i = t0.Add(t2.Subtract(t0).Scale(alpha))
		var B Vec2i = t1.Add(t2.Subtract(t1).Scale(beta))

		if A.X > B.X {
			A, B = B, A
		}

		for x := A.X; x < B.X; x += 1 {
			img.Set(x, y, c)
		}
	}
}
