package main

import "./grate"

var ground grate.Image

type Game struct {}

func (Game) Load() {
	ground = graphics.Image("ground.png")
	ground.Load()
}

func (Game) Update() {

}

func (Game) Draw() {
	ground.Scale(graphics.Width(), 100)
	ground.Translate(0, graphics.Height()-100)
	ground.Draw()
}

//Grate initialisation.
var graphics grate.EbitenGraphics

func main() {
	grate.Ebiten{}.Run(Game{})
}
