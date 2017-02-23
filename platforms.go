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
	if hero.X+25 > platform.X && hero.X-25 < platform.X+platform.W &&
		hero.Y+50 > platform.Y && hero.Y+50 < platform.Y+platform.H {
			
			
			/*if hero.X-platform.X < platform.W/2 {
				hero.X = platform.X
			}
			if hero.X-platform.X >= platform.W/2 {
				hero.X = platform.X+platform.W
			}*/
			
			if hero.Y-platform.Y < platform.Y/2 && (hero.jumpaccel > 20 || hero.jumpstrength < 20 || !hero.jumping) {
				hero.Y = platform.Y-50
				hero.jumpaccel = 0
				hero.jumping = false
				hero.jumpstrength = 20
			}
	}
	return false
}

func (platform *Platform) Draw() {
	platform.img.Translate(platform.X, platform.Y)
	platform.img.Translate(-hero.X+graphics.Width()/2, -hero.Y+graphics.Height()/2)
	platform.img.Draw()
}
