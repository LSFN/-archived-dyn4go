package geometry

type Vertex struct {
	Feature
	point *Vector2
	index int
}

func NewVertexFromPoint(v *Vector2) *Vertex {
	return NewVertexFromPointIndex(v, NOT_INDEXED)
}

func NewVertexFromPointIndex(v *Vector2, i int) *Vertex {
	vertex := new(Vertex)
	vertex = VERTEX_FEATURE
	vertex.point = v
	vertex.index = i
	return vertex
}

func (v *Vertex) GetPoint() *Vector2 {
	return v.point
}

func (v *Vertex) GetIndex() int {
	return v.index
}
