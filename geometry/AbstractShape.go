package geometry

type AbstractShape interface {
	GetID() string
	GetCenter() *Vector2
	GetUserData() interface{}
	SetUserData(data interface{})
	RotateAboutOrigin(theta float64)
	RotateAboutCenter(theta float64)
	RotateAboutVector2(theta float64, v *Vector2)
	RotateAboutXY(theta, x, y float64)
	TranslateXY(x, y float64)
	TranslateVector2(v *Vector2)
	Project(v *Vector2) *Interval
	Contains(v *Vector2) bool
	CreateAABB() *AABB
}
