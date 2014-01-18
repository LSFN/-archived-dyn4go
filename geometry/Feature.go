package geometry

type Feature int

func (f Feature) IsEdge() bool {
	return f == EDGE_FEATURE
}

func (f Feature) IsVertex() bool {
	return f == VERTEX_FEATURE
}

const (
	NOT_INDEXED = iota
	EDGE_FEATURE
	VERTEX_FEATURE
)
