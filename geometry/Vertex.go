package geometry2

type Vertex struct {
	Feature
	point *Vector2
	index int
}

func NewVertexVector2(v *Vector2) *Vertex {
	return NewVertexVector2Int(v, FEATURE_NOT_INDEXED)
}

func NewVertexVector2Int(v *Vector2, i int) *Vertex {
	vertex := new(Vertex)
	vertex.featureType = FEATURE_VERTEX
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
