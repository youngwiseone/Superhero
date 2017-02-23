package main

import "./grate"

type Enemy struct {
	X, Y float64
	img grate.Image
	gravity float64
	jumpaccel float64
	jumping bool
	jumpstrength float64
	speed float64
	runRight bool
	runLeft bool
}

func (enemy *Enemy) Load() {
	enemy.img = graphics.Image("Enemy.png")
	enemy.img.Load()
	enemy.gravity = 5
	enemy.speed = 1
	enemy.runLeft = true
	enemy.runRight = false
}

func (enemy *Enemy) Update() bool {

	//Enemy Movement
	if enemy.runLeft == true {
		if enemy.speed < 5 && enemy.speed > 0{
			enemy.speed += 1
		}
		if enemy.speed == 5 {
			enemy.runRight = true
			enemy.speed = -1
			enemy.runLeft = false
		}
	}
	if enemy.runRight == true {
		if enemy.speed > -5 && enemy.speed < 0{
			enemy.speed -= 1
		}
		if enemy.speed == -5 {
			enemy.runLeft = true
			enemy.speed = 1
			enemy.runRight = false
		}
	}
	
	enemy.X += enemy.speed
	
	for {
		var message = mailbox.GetMessage(enemy.X, enemy.Y, 50)
		if message.Data == "" {
			break
		}
		
		if message.Data == "kill" {
			return true
		}
	}
	
	if enemy.jumping {
		enemy.Y -= enemy.jumpstrength
		enemy.jumpaccel+=1
		enemy.Y += enemy.jumpaccel/4
	}
	
	//Enemy Gravity
	enemy.Y+= enemy.gravity
	
	//This code stops the Enemy from falling through the bottom of the screen
	if enemy.Y > graphics.Height()-100 {
		enemy.Y = graphics.Height()-100
		enemy.jumpaccel = 0
		enemy.jumping = false
		enemy.jumpstrength = 20
	}
	
	return false
}

func (enemy *Enemy) Draw() {
	var w, h = enemy.img.Width(), enemy.img.Height()
	enemy.img.Translate(-w/2, -h/2)
	enemy.img.Translate(enemy.X, enemy.Y)
	enemy.img.Translate(-hero.X+graphics.Width()/2, -hero.Y+graphics.Height()/2)
	enemy.img.Draw()
}
