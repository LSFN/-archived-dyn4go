package geometry

type Shaper interface {
	Transformer
	GetID() string
	GetCenter() *Vector2
	GetRadius() float64
	GetRadiusVector2(center *Vector2) float64
	GetUserData() interface{}
	SetUserData(data interface{})
	RotateAboutCenter(theta float64)
	ProjectVector2(v *Vector2) *Interval
	ProjectVector2Transform(v *Vector2, t *Transform) *Interval
	ContainsVector2(v *Vector2) bool
	ContainsVector2Transform(v *Vector2, t *Transform) bool
	CreateMass(density float64) *Mass
	CreateAABB() *AABB
	CreateAABBTransform(t *Transform) *AABB
}
