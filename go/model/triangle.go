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
	draw_line( t0, t1, img, c )
	draw_line( t1, t2, img, c )
	draw_line( t2, t0, img, c )
}
