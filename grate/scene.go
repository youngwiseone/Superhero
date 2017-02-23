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
func (scene *Scene) Update() bool {
	var del []int
	for i, node := range scene.Nodes {
		if node.Update() {
			del = append(del, i)
		}
	}
	for _, id := range del {
		scene.Nodes = scene.Nodes[:id+copy(scene.Nodes[id:], scene.Nodes[id+1:])]
	}
	return false
}
func (scene *Scene) Draw() {
	for _, node := range scene.Nodes {
		node.Draw()
	}
}
