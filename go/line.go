package main

import (
	"image"
	"image/color"
	"math"
)

type LineFunc func(x0 int, y0 int, x1 int, y1 int, img *image.RGBA, c color.Color);

func draw_line_example(x0 int, y0 int, x1 int, y1 int, img *image.RGBA, c color.Color) {

	var t float64;

	for t = 0; t < 1; t += 0.01 {
		x := float64(x0) * ( 1.0 - t ) + float64(x1) * t
		y := float64(y0) * ( 1.0 - t ) + float64(y1) * t

		img.Set(int(x), int(y), c)
	}
}

//func draw_line_steps(x0 int, y0 int, x1 int, y1 int, img *image.RGBA, c color.Color) {
//	var stepX, stepY float64;
//	var dx, dy int;
//
//	if ( x0 > x1 ) {
//		dx = x0 - x1
//	} else {
//		dx = - ( x1 - x0 )
//	}
//
//	if ( y0 > y1 ) {
//		dy = y0 - y1
//	} else {
//		dy = y1 - y0
//	}
//
//	stepX = float64(dx) / float64(dy);
//	stepY = float64(dy) / float64(dx);
//
//	var x, y float64;
//
//	for i := math.Max(dx, dy); i > 0; i -= 1 {
//
//	}
//
//	fmt.Printf("stepx = %f, stepy = %f\n", stepX, stepY)
//
//}

func draw_line_solution(x0 int, y0 int, x1 int, y1 int, img *image.RGBA, c color.Color) {

	var steep bool = false;
	if math.Abs(float64(x0 - x1)) < math.Abs(float64(y0 - y1)) {
		// if the line is steep, we transpose the image
		x0, y0 = y0, x0
		x1, y1 = y1, x1
		steep = true
	}
	if (x0 > x1) {
		// make it left−to−right
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}
	for x := x0; x < x1; x += 1 {
		t := float64(x - x0) / float64(x1 - x0)
		y := float64(y0) * (1. - t) + float64(y1) * t
		//int y = y0*(1.-t) + y1*t;
		if (steep) {
			img.Set(int(y), x, c) // if transposed, de−transpose
		} else {
			img.Set(x, int(y), c)
		}
	}
}
