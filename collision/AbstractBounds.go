package collision

import (
	"github.com/LSFN/dyn4go/geometry"
)

type AbstractBounds struct {
	transform *geometry.Transform
}

func InitAbstractBounds(a *AbstractBounds) {
	a.transform = geometry.NewTransform()
}

func InitAbstractBoundsTransform(a *AbstractBounds, transform *geometry.Transform) {
	if transform == nil {
		panic("Cannot create AbstractBounds from nil transform")
	}
	a.transform = transform
}

func (a *AbstractBounds) GetTransform() *geometry.Transform {
	return a.transform
}

func (a *AbstractBounds) RotateAboutOrigin(theta float64) {
	a.transform.RotateAboutOrigin(theta)
}

func (a *AbstractBounds) RotateAboutVector2(theta float64, point *geometry.Vector2) {
	a.transform.RotateAboutVector2(theta, point)
}

func (a *AbstractBounds) RotateAboutXY(theta, x, y float64) {
	a.transform.RotateAboutXY(theta, x, y)
}

func (a *AbstractBounds) TranslateXY(x, y float64) {
	a.transform.TranslateXY(x, y)
}

func (a *AbstractBounds) TranslateVector2(vector *geometry.Vector2) {
	a.transform.TranslateVector2(vector)
}

func (a *AbstractBounds) ShiftCoordinates(shift *geometry.Vector2) {
	a.transform.TranslateVector2(shift)
}
