package geometry2

const (
	FEATURE_NOT_INDEXED = -1
	FEATURE_EDGE        = 0
	FEATURE_VERTEX      = 1
)

type Featurer interface {
	IsEdge() bool
	IsVertex() bool
}
