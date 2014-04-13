package geometry2

type Transformer interface {
	RotateAboutOrigin(theta float64)
	RotateAboutVector2(theta float64, v *Vector2)
	RotateAboutXY(theta, x, y float64)
	TranslateXY(x, y float64)
	TranslateVector2(v *Vector2)
}
