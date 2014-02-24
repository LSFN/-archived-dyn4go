package geometry

const (
	NOT_INDEXED    = -1
	EDGE_FEATURE   = 0
	VERTEX_FEATURE = 1
)

type Feature interface {
	IsEdge() bool
	IsVertex() bool
}
