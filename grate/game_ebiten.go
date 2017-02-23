package grate

import "github.com/hajimehoshi/ebiten"

type Ebiten struct {}

func (Ebiten) Run(game Node) {
	//Run game.
	game.Load()
	var loaded bool
	ebiten.Run(func(s *ebiten.Image) error {
		if !loaded {
			ebiten.SetScreenScale(1)
			loaded = true
		}
		screen = s
		game.Update()
		game.Draw()
		return nil
	}, 800, 600, 2, "Grate")
}
