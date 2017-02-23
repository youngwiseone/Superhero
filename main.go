package main

import "./grate"

var ground grate.Image
var hero Hero

type Game struct {}

func (Game) Load() {
	ground = graphics.Image("ground.png")
	ground.Load()
	hero.Load()
}

func (Game) Update() {
	hero.Update()
}

func (Game) Draw() {
	ground.Scale(graphics.Width(), 100)
	ground.Translate(0, graphics.Height()-100)
	ground.Draw()
	hero.Draw()
}

//Grate initialisation.
var graphics grate.EbitenGraphics
var input grate.EbitenInput

func main() {
	grate.Ebiten{}.Run(Game{})
}
