package model

import (
	"image"
	"image/color"
)

type TriangleFunc func()

func draw_line( v0, v1 Vec2i, img *image.RGBA , c color.Color ) {
	Draw_line_solution( v0.X, v0.Y, v1.X, v1.Y, img, c)
}

func Triangle_first( t0, t1, t2 Vec2i, img *image.RGBA, c color.Color ) {
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

	var red = color.RGBA{0xFF,0,0,0xFF}
	var green = color.RGBA{0,0xFF,0,0xFF}

	draw_line( t0, t1, img, green )
	draw_line( t1, t2, img, green )
	draw_line( t2, t0, img, red )
}
