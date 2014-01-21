package geometry

type Edge struct {
	Feature
	vertex1, vertex2, max *Vertex
	edge                  *Vector2
	index                 int
}

func NewEdge(vertex1, vertex2, max *Vertex, edge *Vector2, index int) *Edge {
	e := new(Edge)
	e = EDGE_FEATURE
	e.vertex1 = vertex1
	e.vertex2 = vertex2
	e.edge = edge
	e.max = max
	e.index = index
	return e
}

func (e *Edge) GetVertex1() *Vertex {
	return e.vertex1
}

func (e *Edge) GetVertex2() *Vertex {
	return e.vertex2
}

func (e *Edge) GetEdge() *Vector2 {
	return e.edge
}

func (e *Edge) GetMaximum() *Vertex {
	return e.max
}

func (e *Edge) GetIndex() int {
	return e.index
}
