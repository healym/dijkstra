package tree-algorithms

type Edge struct {
    tail int
    head int
    weight int
}

type Vertex struct {
    id int
    dist int
    arcs map[int]int        // arcs[vertex id] = weight
}

type Graph struct {
    visited map[int]bool
    vertices map[int]Vertex
}
