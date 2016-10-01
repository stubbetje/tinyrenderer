package main

import (
	"image"
	"image/png"
	"os"
	"image/color"
	"image/draw"
)

func newImage( width int, height int, c color.Color ) image.Image {
	dim := image.Rect(0,0,width,height)
	img := image.NewNRGBA(dim)

	draw.Draw(img, dim, &image.Uniform{c}, image.ZP, draw.Src )

	return img
}

func main() {

	img := newImage( 100, 100, color.Black );

	file, err := os.Create( "output/image.png")
	if err != nil {
		panic( err )
	}
	png.Encode( file, img )
}
