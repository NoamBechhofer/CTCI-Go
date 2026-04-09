package lib

import (
	"fmt"
	"strings"
)

type GraphNode[T any] struct {
	Val      T
	Children []*GraphNode[T]
}

func bfsRunner[T any](queue Queue[*GraphNode[T]], visited map[*GraphNode[T]]struct{}, visitor func(*GraphNode[T])) {
	for !queue.IsEmpty() {
		node, _ := queue.Remove()
		visitor(node)
		for _, child := range node.Children {
			if _, seen := visited[child]; seen {
				continue
			}
			visited[child] = struct{}{}
			queue.Add(child)
		}
	}
}

func (root *GraphNode[T]) dfsDriver(visitor func(*GraphNode[T]), visited map[*GraphNode[T]]struct{}) {
	if _, seen := visited[root]; seen {
		return
	}
	visited[root] = struct{}{}
	visitor(root)
	for _, child := range root.Children {
		child.dfsDriver(visitor, visited)
	}
}

func (root *GraphNode[T]) BFS(visitor func(*GraphNode[T])) {
	queue := ArrayQueue[*GraphNode[T]]{}
	visited := map[*GraphNode[T]]struct{}{}

	visited[root] = struct{}{}
	queue.Add(root)

	bfsRunner(&queue, visited, visitor)
}

func (root *GraphNode[T]) DFS(visitor func(*GraphNode[T])) {
	visited := map[*GraphNode[T]]struct{}{}
	root.dfsDriver(visitor, visited)
}

func (root *GraphNode[T]) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "%p(%v): [", root, root.Val)
	for i, child := range root.Children {
		fmt.Fprintf(&builder, "%p(%v)", child, child.Val)
		if i < len(root.Children)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")

	return builder.String()
}

type Graph[T any] struct {
	Nodes []*GraphNode[T]
}

func (graph *Graph[T]) BFS(visitor func(*GraphNode[T])) {
	queue := ArrayQueue[*GraphNode[T]]{}
	visited := map[*GraphNode[T]]struct{}{}

	for _, node := range graph.Nodes {
		if _, seen := visited[node]; seen {
			continue
		}
		visited[node] = struct{}{}
		queue.Add(node)
		bfsRunner(&queue, visited, visitor)
	}
}

func (graph *Graph[T]) DFS(visitor func(*GraphNode[T])) {
	if len(graph.Nodes) == 0 {
		return
	}

	visited := map[*GraphNode[T]]struct{}{}
	for _, node := range graph.Nodes {
		node.dfsDriver(visitor, visited)
	}
}

func (graph *Graph[T]) String() string {
	if len(graph.Nodes) == 0 {
		return ""
	}

	var builder strings.Builder

	graph.DFS(func(node *GraphNode[T]) {
		fmt.Fprintf(&builder, "%s\n", node.String())
	})

	ret := builder.String()
	return ret[:len(ret)-1]
}
