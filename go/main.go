package main

import (
	"image"
	"image/png"
	"os"
	"image/color"
	"image/draw"
)

func newImage(width int, height int, c color.Color) *image.RGBA {
	dim := image.Rect(0, 0, width, height)
	img := image.NewRGBA(dim)

	draw.Draw(img, dim, &image.Uniform{c}, image.ZP, draw.Src)

	return img
}

func saveImage(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	png.Encode(file, img)
}

func main() {

	img := newImage(100, 100, color.Black);

	var line LineFunc = draw_line_solution;

	line(13, 20, 80, 40, img, color.White);

	saveImage("output/image.png", img)
}
