package grate

type Input interface {
	IsKeyDown(int) bool
	
	KeyW() int
	KeyA() int
	KeyS() int
	KeyD() int
}


