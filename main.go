package main

import "./grate"

var ground grate.Image
var scene grate.Scene
var hero = new(Hero)

type Game struct {}

func (Game) Load() {
	ground = graphics.Image("ground.png")
	ground.Load()
	
	//Add an enemy to the scene
	var enemy = new(Enemy)
	enemy.X = 400
	enemy.Y = 50
	scene.Add(enemy)
	
	//Add a hero to the scene
	scene.Add(hero)
	
	scene.Load()
}

func (Game) Update() {
	scene.Update()
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
