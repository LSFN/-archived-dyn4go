package geometry

type Shape interface {
	Transformable
	GetID() string
	GetCenter() *Vector2
	GetRadius() float64
	GetRadiusVector2(center *Vector2) float64
	GetUserData() *interface{}
	SetUserData(data *interface{})
	RotateAboutCenter(theta float64)
	Project(v *Vector2) *Interval
	ProjectTransform(v *Vector2, t *Transform) *Interval
	Contains(v *Vector2) bool
	ContainsTransform(v *Vector2, t *Transform) bool
	CreateMass(density float64)
	CreateAABB() *AABB
	CreateAABBTransform(t *Transform)
}
