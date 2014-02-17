package geometry

const (
	NOT_INDEXED = iota
	EDGE_FEATURE
	VERTEX_FEATURE
)

type Feature interface {
	IsEdge() bool
	IsVertex() bool
}
