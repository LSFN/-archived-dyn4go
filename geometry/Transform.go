package geometry2

import (
	"math"
)

type Transform struct {
	X, Y, m00, m01, m10, m11 float64
}

func NewTransform() *Transform {
	t := new(Transform)
	t.m00 = 1
	t.m11 = 1
	return t
}

func NewTransformFromTransform(t *Transform) *Transform {
	t2 := *t
	return &t2
}

func (t *Transform) RotateAboutOrigin(theta float64) {
	cos := math.Cos(theta)
	sin := math.Sin(theta)

	m00 := cos*t.m00 - sin*t.m10
	m01 := cos*t.m01 - sin*t.m11
	m10 := sin*t.m00 + cos*t.m10
	m11 := sin*t.m01 + cos*t.m11
	x := cos*t.X - sin*t.Y
	y := sin*t.X + cos*t.Y

	t.m00 = m00
	t.m01 = m01
	t.m10 = m10
	t.m11 = m11
	t.X = x
	t.Y = y
}

func (t *Transform) RotateAboutVector2(theta float64, v *Vector2) {
	t.RotateAboutXY(theta, v.X, v.Y)
}

func (t *Transform) RotateAboutXY(theta, x, y float64) {
	cm00 := t.m00
	cm01 := t.m01
	cx := t.X
	cm10 := t.m10
	cm11 := t.m11
	cy := t.Y

	cos := math.Cos(theta)
	sin := math.Sin(theta)
	rx := x - cos*x + sin*y
	ry := y - sin*x - cos*y

	t.m00 = cos*cm00 - sin*cm10
	t.m01 = cos*cm01 - sin*cm11
	t.X = cos*cx - sin*cy + rx
	t.m10 = sin*cm00 + cos*cm10
	t.m11 = sin*cm01 + cos*cm11
	t.Y = sin*cx + cos*cy + ry
}

func (t *Transform) TranslateXY(x, y float64) {
	t.X += x
	t.Y += y
}

func (t *Transform) TranslateVector2(v *Vector2) {
	t.X += v.X
	t.Y += v.Y
}

func (t *Transform) Set(t2 *Transform) {
	*t = *t2
}

func (t *Transform) Identity() {
	t.m00 = 1
	t.m01 = 0
	t.m10 = 0
	t.m11 = 1
	t.X = 0
	t.Y = 0
}

func (t *Transform) GetTransformedVector2(v *Vector2) *Vector2 {
	v2 := new(Vector2)
	v2.X = t.m00*v.X + t.m01*v.Y + t.X
	v2.Y = t.m10*v.X + t.m11*v.Y + t.Y
	return v2
}

func (t *Transform) GetTransformedVector2InDestination(v, dest *Vector2) {
	dest.X = t.m00*v.X + t.m01*v.Y + t.X
	dest.Y = t.m10*v.X + t.m11*v.Y + t.Y
}

func (t *Transform) Transform(v *Vector2) {
	x, y := v.X, v.Y
	v.X = t.m00*x + t.m01*y + t.X
	v.Y = t.m10*x + t.m11*y + t.Y
}

func (t *Transform) GetInverseTransformedVector2(v *Vector2) *Vector2 {
	v2 := new(Vector2)
	tx := v.X - t.X
	ty := v.Y - t.Y
	v2.X = t.m00*tx + t.m10*ty
	v2.Y = t.m01*tx + t.m11*ty
	return v2
}

func (t *Transform) GetInverseTransformedVector2InDestination(v, dest *Vector2) {
	tx := v.X - t.X
	ty := v.Y - t.Y
	dest.X = t.m00*tx + t.m10*ty
	dest.Y = t.m01*tx + t.m11*ty
}

func (t *Transform) InverseTransform(v *Vector2) {
	tx := v.X - t.X
	ty := v.Y - t.Y
	v.X = t.m00*tx + t.m10*ty
	v.Y = t.m01*tx + t.m11*ty
}

func (t *Transform) GetTransformedR(v *Vector2) *Vector2 {
	v2 := new(Vector2)
	v2.X = t.m00*v.X + t.m01*v.Y
	v2.Y = t.m10*v.X + t.m11*v.Y
	return v2
}

func (t *Transform) GetTransformedRInDestination(v, v2 *Vector2) {
	v2.X = t.m00*v.X + t.m01*v.Y
	v2.Y = t.m10*v.X + t.m11*v.Y
}

func (t *Transform) TransformR(v *Vector2) {
	x, y := v.X, v.Y
	v.X = t.m00*x + t.m01*y
	v.Y = t.m10*x + t.m11*y
}

func (t *Transform) GetInverseTransformedR(v *Vector2) *Vector2 {
	v2 := new(Vector2)
	v2.X = t.m00*v.X + t.m10*v.Y
	v2.Y = t.m01*v.X + t.m11*v.Y
	return v2
}

func (t *Transform) GetInverseTransformedRInDestination(v, v2 *Vector2) {
	v2.X = t.m00*v.X + t.m10*v.Y
	v2.Y = t.m01*v.X + t.m11*v.Y
}

func (t *Transform) InverseTransformR(v *Vector2) {
	x, y := v.X, v.Y
	v.X = t.m00*x + t.m10*y
	v.Y = t.m01*x + t.m11*y
}

func (t *Transform) GetTranslation() *Vector2 {
	return NewVector2FromXY(t.X, t.Y)
}

func (t *Transform) SetTranslationFromXY(x, y float64) {
	t.X = x
	t.Y = y
}

func (t *Transform) SetTranslationFromVector2(v *Vector2) {
	t.X = v.X
	t.Y = v.Y
}

func (t *Transform) GetTranslationTransform() *Transform {
	t2 := NewTransform()
	t2.TranslateXY(t.X, t.Y)
	return t2
}

func (t *Transform) GetRotation() float64 {
	return math.Atan2(t.m10, t.m00)
}

func (t *Transform) SetRotation(theta float64) float64 {
	r := t.GetRotation()
	t.RotateAboutXY(-r, t.X, t.Y)
	t.RotateAboutXY(theta, t.X, t.Y)
	return r
}

func (t *Transform) GetRotationTransform() *Transform {
	t2 := NewTransform()
	t2.RotateAboutOrigin(t.GetRotation())
	return t2
}

func (t *Transform) GetValues() []float64 {
	return []float64{t.m00, t.m01, t.X, t.m10, t.m11, t.Y}
}

func (t *Transform) Lerp(end *Transform, alpha float64) {
	x := (1-alpha)*t.X + alpha*end.X
	y := (1-alpha)*t.Y + alpha*end.Y
	rs := t.GetRotation()
	re := end.GetRotation()
	diff := re - rs
	if diff < -math.Pi {
		diff += TWO_PI
	} else if diff > math.Pi {
		diff -= TWO_PI
	}
	a := diff*alpha + rs
	t.Identity()
	t.RotateAboutOrigin(a)
	t.TranslateXY(x, y)
}

func (t *Transform) LerpInDestination(end *Transform, alpha float64, result *Transform) {
	x := (1-alpha)*t.X + alpha*end.X
	y := (1-alpha)*t.Y + alpha*end.Y
	rs := t.GetRotation()
	re := end.GetRotation()
	diff := re - rs
	if diff < -math.Pi {
		diff += TWO_PI
	} else if diff > math.Pi {
		diff -= TWO_PI
	}
	a := diff*alpha + rs
	result.Identity()
	result.RotateAboutOrigin(a)
	result.TranslateXY(x, y)
}

func (t *Transform) LerpDeltaInDestination(dp *Vector2, da, alpha float64, result *Transform) {
	result.Set(t)
	result.TranslateXY(dp.X*alpha, dp.Y*alpha)
	result.RotateAboutXY(da*alpha, result.X, result.Y)
}

func (t *Transform) LerpDelta(dp *Vector2, da, alpha float64) {
	t.TranslateXY(dp.X*alpha, dp.Y*alpha)
	t.RotateAboutXY(da*alpha, t.X, t.Y)
}

func (t *Transform) LerpedDelta(dp *Vector2, da, alpha float64) *Transform {
	result := NewTransform()
	result.Set(t)
	result.TranslateXY(dp.X*alpha, dp.Y*alpha)
	result.RotateAboutXY(da*alpha, result.X, result.Y)
	return result
}

func (t *Transform) Lerped(end *Transform, alpha float64) *Transform {
	x := (1-alpha)*t.X + alpha*end.X
	y := (1-alpha)*t.Y + alpha*end.Y
	rs := t.GetRotation()
	re := end.GetRotation()
	diff := re - rs
	if diff < -math.Pi {
		diff += TWO_PI
	} else if diff > math.Pi {
		diff -= TWO_PI
	}
	a := diff*alpha + rs
	tx := NewTransform()
	tx.RotateAboutOrigin(a)
	tx.TranslateXY(x, y)
	return tx
}
