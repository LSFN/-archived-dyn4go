package geometry

type Convexer interface {
	GetAxes(foci []*Vector2, t *Transform) []*Vector2
	GetFoci(t *Transform) []*Vector2
	GetFarthestFeature(v *Vector2, t *Transform) Featurer
	GetFarthestPoint(v *Vector2, t *Transform) *Vector2
}
