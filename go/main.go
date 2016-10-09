package main

import (
	"image"
	"image/png"
	"os"
	"image/color"
	"image/draw"
	"github.com/disintegration/imaging"
	"./model"
	"math/rand"
)

func newImage(width int, height int, c color.Color) *image.RGBA {
	dim := image.Rect(0, 0, width, height)
	img := image.NewRGBA(dim)

	draw.Draw(img, dim, &image.Uniform{c}, image.ZP, draw.Src)

	return img
}

func saveImage(filename string, img image.Image) {

	im := imaging.FlipV(img)

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	png.Encode(file, im)
}

func do_triangles(img *image.RGBA) {

	var t0 []model.Vec2i = []model.Vec2i{{10,70},{50,160},{70,80}}
	var t1 []model.Vec2i = []model.Vec2i{{180,50},{150,1},{70,180}}
	var t2 []model.Vec2i = []model.Vec2i{{180,150},{120,160},{130,180}}

	var red = color.RGBA{0xFF,0,0,0XFF}
	var white = color.White
	var green = color.RGBA{0,0XFF,0,0XFF}

	model.Triangle_fill_two_halves(t0[0], t0[1], t0[2], img, red);
	model.Triangle_fill_two_halves(t1[0], t1[1], t1[2], img, white);
	model.Triangle_fill_two_halves(t2[0], t2[1], t2[2], img, green);
}

func do_render_head(img *image.RGBA) {

	m, err := model.LoadFromFile("../obj/african_head/african_head.obj")
	if err != nil {
		panic(err)
	}

	//line(13, 20, 80, 40, img, color.White);

	dim := img.Bounds().Size()

	// float width factor  = fw
	// float height factor = fh
	var fw, fh float64
	fw = float64(dim.X) / 2.0
	fh = float64(dim.Y) / 2.0

	for i := 0; i < m.Nfaces(); i += 1 {
		face := m.Face(i)

		var points []model.Vec2i = make([]model.Vec2i, 3)

		for j := 0; j < 3; j += 1 {
			world_coords := m.Vertice(face[j])
			points[j] = model.Vec2i{
				X: int((world_coords.X + 1.0 ) * fw),
				Y: int((world_coords.Y + 1.0 ) * fh),
			}

		}

		c := color.RGBA{ R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A:0xFF}

		model.Triangle_fill_two_halves(points[0], points[1], points[2], img, c)
	}
}

func main() {

	var width, height int = 800, 800
	img := newImage(width, height, color.Black);

	do_render_head(img)

	//do_triangles( img )

	saveImage("output/image.png", img)
}
