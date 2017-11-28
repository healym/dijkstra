package algorithms

import (
	"fmt"
)

func KruskalMST(g *Graph) Graph, err {
	AdjMat := make([][]int, len(g.vertices))
	for i := range AdjMat { // start at head
		AdjMat[i] = make([]int, len(g.vertices))
	}
	mst := make(Graph)
	h = make(Candidates, len(g.vertices))
	for v := range g.vertices {
		for i := range v.arcs {
			AdjMat[v][i] = v.arcs[i]
			e := make(Edge)
			e.head = v
			e.tail = i
			e.weight = v.arcs[i]
			h.Push(e)
		}
	}
	done := false
	for !done {
		if len(g.visited) == len(g.vertices) { // all vertices visited
			done = true
			continue
		}
		if h.IsEmpty {
			log.fatal("Heap emptied!!")
		}
		e := h.Pop()
		if g.visited[e.head] {
			continue
		}
		g.visit(e.head)
		if !mst.vertices[e.head] {
			mst.vertices[e.head] = Vertex{id: e.head}
		}
		mst.vertices[e.head].arcs[e.tail] = e.weight
	}
	return mst, nil

}
