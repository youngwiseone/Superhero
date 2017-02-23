package main

import "./grate"

var ground grate.Image
var scene grate.Scene
var hero = new(Hero)
var mailbox grate.Mailbox

type Game struct {}

func (Game) Load() {
	ground = graphics.Image("ground.png")
	ground.Load()
	
	scene.Add(&Tree{X:150, Y:430})
	scene.Add(&Tree{X:350, Y:430})
	scene.Add(&Tree{X:50, Y:430})

	scene.Add(&Enemy{X:400, Y:50})
	scene.Add(&Enemy{X:500, Y:50})
	scene.Add(&Enemy{X:300, Y:50})
	scene.Add(&Enemy{X:200, Y:50})
	scene.Add(&Enemy{X:100, Y:50})
	
	//Add a hero to the scene
	scene.Add(hero)
	scene.Add(&mailbox)
	scene.Load()
}

func (Game) Update() bool {
	return scene.Update()
}

func (Game) Draw() {
	ground.Scale(graphics.Width(), 100)
	ground.Translate(0, graphics.Height()-100)
	ground.Translate(0, -hero.Y+graphics.Height()/2)
	ground.Draw()
	scene.Draw()
}

//Grate initialisation.
var graphics grate.EbitenGraphics
var input grate.EbitenInput

func main() {
	grate.Ebiten{}.Run(Game{})
}
