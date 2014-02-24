package geometry

type Convex interface {
	GetAxes(foci []*Vector2, t *Transform) []*Vector2
	GetFoci(t *Transform) []*Vector2
	GetFarthestFeature(v *Vector2, t *Transform) Feature
	GetFarthestPoint(v *Vector2, t *Transform) *Vector2
}
