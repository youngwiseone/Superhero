package main

import "./grate"

type Platform struct {
	X, Y, W, H float64
	img grate.Image
}


func (platform *Platform) Load() {
	platform.img = graphics.Image("platform.png")
	platform.img.Load()
	platform.W = 100
	platform.H = 27
}

func (platform *Platform) Update() bool {
	if hero.X > platform.X && hero.X < platform.X+platform.W &&
		hero.Y > platform.Y && hero.Y < platform.Y+platform.H {
			
			
			if hero.X-platform.X < platform.W/2 {
				hero.X = platform.X
			}
			if hero.X-platform.X >= platform.W/2 {
				hero.X = platform.X+platform.W
			}
			
			if hero.Y-platform.Y < platform.Y/2 {
				hero.Y = platform.Y
			}
			if hero.Y-platform.Y >= platform.Y/2 {
				hero.Y = platform.Y+platform.Y
			}
	}
	return false
}

func (platform *Platform) Draw() {
	platform.img.Translate(platform.X, platform.Y)
	platform.img.Draw()
}
