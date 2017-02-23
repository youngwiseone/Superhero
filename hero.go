package main

import "./grate"

type Hero struct {
	X, Y float64
	img grate.Image
}

func (hero *Hero) Load() {
	hero.img = graphics.Image("Hero.png")
	hero.img.Load()
}

func (hero *Hero) Update() {
	if input.KeyIsDown(input.KeyA()) {
		hero.X--
	}
	if input.KeyIsDown(input.KeyD()) {
		hero.X++
	}
	hero.Y+=5
	if hero.Y > graphics.Height()-100 {
		hero.Y = graphics.Height()-100
	}
}

func (hero *Hero) Draw() {
	var w, h = hero.img.Width(), hero.img.Height()
	hero.img.Translate(-w/2, -h/2)
	hero.img.Translate(hero.X, hero.Y)
	hero.img.Draw()
}
