package grate

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"

var screen *ebiten.Image
var options ebiten.DrawImageOptions
var matrix ebiten.DrawImageOptions
var err error

type EbitenImage struct {
	*ebiten.Image
	options ebiten.DrawImageOptions
	path string
}

func (img *EbitenImage) Load() {
	println("loading!")
	img.Image, _, err = ebitenutil.NewImageFromFile(img.path, ebiten.FilterNearest)
	if err != nil {
		println(err.Error())
	}
}

func (img *EbitenImage) Draw() {
	screen.DrawImage(img.Image, &(img.options))
	img.options = matrix
}

func (img *EbitenImage) Scale(x, y float64) {
	img.options.GeoM.Scale(x, y)
}

func (img *EbitenImage) Translate(x, y float64) {
	img.options.GeoM.Translate(x, y)
}

func (img *EbitenImage) Rotate(angle float64) {
	img.options.GeoM.Rotate(angle)
}

type EbitenGraphics struct {}

func (EbitenGraphics) Image(path string) Image {
	return &EbitenImage{path:path}
}

func (EbitenGraphics) Width() float64 {
	w, _ := screen.Size()
	return float64(w)
}

func (EbitenGraphics) Height() float64 {
	_, h := screen.Size()
	return float64(h)
}

func (EbitenGraphics) Translate(x, y float64) {
	matrix.GeoM.Translate(x, y)
}
