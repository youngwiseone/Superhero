package main

import "./grate"

type Hero struct {
	X, Y float64
	img grate.Image
	gravity float64
	jumpaccel float64
	jumping bool
	jumpstrength float64
}

func (hero *Hero) Load() {
	hero.img = graphics.Image("Hero.png")
	hero.img.Load()
	hero.gravity = 5
}

func (hero *Hero) Update() {

	//Hero Controls
	if input.KeyIsDown(input.KeyA()) {
		hero.X-=5
	}
	if input.KeyIsDown(input.KeyD()) {
		hero.X+=5
	}
	if !hero.jumping && input.KeyIsDown(input.KeySpace()) {
		hero.jumping = true
	} else if !input.KeyIsDown(input.KeySpace()) && hero.jumping {
		hero.jumpstrength--
	}
	
	if hero.jumping {
		hero.Y -= hero.jumpstrength
		hero.jumpaccel+=1
		hero.Y += hero.jumpaccel/2
	}
	
	//Hero Gravity
	hero.Y+= hero.gravity
	
	//This code stops the Hero from falling through the bottom of the screen
	if hero.Y > graphics.Height()-100 {
		hero.Y = graphics.Height()-100
		hero.jumpaccel = 0
		hero.jumping = false
		hero.jumpstrength = 20
	}
}

func (hero *Hero) Draw() {
	var w, h = hero.img.Width(), hero.img.Height()
	hero.img.Translate(-w/2, -h/2)
	hero.img.Translate(graphics.Width()/2, graphics.Height()/2)
	hero.img.Draw()
}
