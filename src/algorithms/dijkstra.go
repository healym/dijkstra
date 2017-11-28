package algorithms

import (
	"fmt"
)

func (g *Graph) DijkstraMST(src) Graph, err {
	mst := make(Graph)
	h := make(Candidates, len(g.vertices))
	g.visit(src)
	v := g.vertices[src]
	for id, w := range v.arcs {
		g.vertices[id].dist = w
		h.Push(v)
		mst.vertices[v.id] = make(Vertex)
	}
	done := false
	for !done {
		if len(g.visited) == len(g.vertices) { // all vertices visited
			done = true
			continue
		}
		if h.IsEmpty() {
			return mst, error("Heap emptied!!")
		}
		v = h.Pop()
		if g.visited[v.id] {
			continue
		}
		g.visit(v.id)
		for w, d := range v.arcs {
			if g.visited[w] {
				continue
			}
			if d+v.dist < g.vertices[w].dist {
				g.vertices[w].dist = d + v.dist
				mst.vertices[w] = g.vertices[w]
			}
			h.Push(g.vertices[w])
		}
	}
	return mst, nil
}
