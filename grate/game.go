package grate

type Game interface {
	Load()
	Draw()
	Update()
}

type Engine interface {
	Run(Game)
}
