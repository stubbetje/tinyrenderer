package main

import (
	"image"
	"image/png"
	"os"
	"image/color"
	"image/draw"
	"github.com/disintegration/imaging"
	"./model"
)

func newImage(width int, height int, c color.Color) *image.RGBA {
	dim := image.Rect(0, 0, width, height)
	img := image.NewRGBA(dim)

	draw.Draw(img, dim, &image.Uniform{c}, image.ZP, draw.Src)

	return img
}

func saveImage(filename string, img image.Image) {

	im := imaging.FlipV( img )

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	png.Encode(file, im)
}

func main() {

	var width, height int = 800, 800
	img := newImage(width, height, color.Black);

	var line LineFunc = draw_line_solution;

	m, err := model.LoadFromFile("../obj/african_head/african_head.obj")
	if err != nil {
		panic(err)
	}

	//line(13, 20, 80, 40, img, color.White);

	for i := 0; i < m.Nfaces(); i += 1 {
		face := m.Face(i)

		for j := 0; j < 3; j += 1 {
			v0 := m.Vertice(face[j])
			v1 := m.Vertice(face[(j + 1) % 3])

			x0 := (v0.X + 1.0) * float64(width) / 2.0
			y0 := (v0.Y + 1.0) * float64(height) / 2.0

			x1 := (v1.X + 1.0) * float64(width) / 2.0
			y1 := (v1.Y + 1.0) * float64(height) / 2.0

			line( int(x0), int(y0), int(x1), int(y1), img, color.White)
		}
	}

	saveImage("output/image.png", img)
}
