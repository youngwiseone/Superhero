package grate

type Node interface {
	Load()
	Draw()
	Update()
}

type Engine interface {
	Run(Node)
}
