package main

import "./grate"

type Tree struct {
	X, Y float64
	img grate.Image
}

func (tree *Tree) Load() {
	tree.img = graphics.Image("images/Tree.png")
	tree.img.Load()
}

func (tree *Tree) Update() bool {
	return false
}

func (tree *Tree) Draw() {
	var w, h = tree.img.Width(), tree.img.Height()
	tree.img.Translate(-w/2, -h/2)
	tree.img.Translate(tree.X, tree.Y)
	tree.img.Translate(-hero.X+graphics.Width()/2, -hero.Y+graphics.Height()/2)
	tree.img.Draw()
}
