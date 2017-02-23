package grate

type Node interface {
	Load()
	Draw()
	Update() bool
}

type Engine interface {
	Run(Node)
}
