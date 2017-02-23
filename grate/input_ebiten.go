package grate

import "github.com/hajimehoshi/ebiten"

type EbitenInput struct {}

func (EbitenInput) KeyIsDown(key int) bool {
	return ebiten.IsKeyPressed(ebiten.Key(key)) 
}

func (EbitenInput) KeyW() int { return int(ebiten.KeyW) }
func (EbitenInput) KeyA() int { return int(ebiten.KeyA) }
func (EbitenInput) KeyS() int { return int(ebiten.KeyS) }
func (EbitenInput) KeyD() int { return int(ebiten.KeyD) }
func (EbitenInput) KeySpace() int { return int(ebiten.KeySpace) }
func (EbitenInput) KeyEnter() int { return int(ebiten.KeyEnter) }
