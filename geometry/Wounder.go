package geometry2

type Wounder interface {
	Shaper
	GetVertices() []*Vector2
	GetNormals() []*Vector2
}
