package geometry

type Vertex struct {
	point *Vector2
	index int
}

const (
	NOT_INDEXED = -1
)

func NewVertexFromVector2(v *Vector2) *Vertex {
	return NewVertexFromVector2Int(v, NOT_INDEXED)
}

func NewVertexFromVector2Int(v *Vector2, i int) *Vertex {
	vertex := new(Vertex)
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

func (v *Vertex) IsEdge() bool {
	return false
}

func (v *Vertex) IsVertex() bool {
	return true
}
