package main

import "./grate"
import "math"

type Hero struct {
	X, Y float64
	img grate.Image
	gravity float64
	jumpaccel float64
	jumping bool
	jumpstrength float64
	
	attacking bool
	attackstep int
}

func (hero *Hero) Load() {
	hero.img = graphics.Image("images/Hero.png")
	hero.img.Load()
	hero.gravity = 5
}


func (hero *Hero) Update() bool {

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
	
	if input.KeyIsDown(input.KeyEnter()) {
		hero.attacking = true
	}
	
	if hero.jumping {
		hero.Y -= hero.jumpstrength
		hero.jumpaccel+=1
		hero.Y += hero.jumpaccel/2
	}
	
	if hero.attacking {
		hero.attackstep += 20
		if hero.attackstep > 360 {
			hero.attacking = false
			hero.attackstep = 0
			mailbox.SendMessage("kill", hero.X, hero.Y, 100)
		}
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
	
	return false
}

func (hero *Hero) Draw() {
	var w, h = hero.img.Width(), hero.img.Height()
	hero.img.Translate(-w/2, -h/2)
	if hero.attacking {
		hero.img.Rotate((float64(hero.attackstep)*math.Pi)/180)
	}
	hero.img.Translate(graphics.Width()/2, graphics.Height()/2)
	hero.img.Draw()
}
