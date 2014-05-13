package geometry

type WounderConvexer interface {
	GetVertices() []*Vector2
	GetNormals() []*Vector2
	Convexer
}
