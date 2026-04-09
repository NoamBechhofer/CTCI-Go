package routebetweennodes

import "github.com/NoamBechhofer/CTCI-Go/lib"

func RouteBetweenNodes[T any](graph *lib.Graph[T], s, e *lib.GraphNode[T]) bool {
	queue := lib.ArrayQueue[*lib.GraphNode[T]]{}
	visited := map[*lib.GraphNode[T]]struct{}{}

	queue.Add(s)
	visited[s] = struct{}{}

	for !queue.IsEmpty() {
		node, _ := queue.Remove()
		if node == e {
			return true
		}
		for _, child := range node.Children {
			if _, seen := visited[child]; seen {
				continue
			}
			visited[child] = struct{}{}
			queue.Add(child)
		}
	}

	return false
}
