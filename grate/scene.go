package grate

type Scene struct {
	Nodes []Node
}


func (scene *Scene) Add(node Node) {
	scene.Nodes = append(scene.Nodes, node)
}

func (scene *Scene) Load() {
	for _, node := range scene.Nodes {
		node.Load()
	}
}
func (scene *Scene) Update() {
	for _, node := range scene.Nodes {
		node.Update()
	}
}
func (scene *Scene) Draw() {
	for _, node := range scene.Nodes {
		node.Draw()
	}
}
