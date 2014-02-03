package geometry

import (
	"math"

	"github.com/LSFN/dyn4go"
)

type Vector3 struct {
	X, Y, Z float64
}

func NewVector3FromVector3(v *Vector3) *Vector3 {
	v2 := new(Vector3)
	*v2 = *v
	return v2
}

func NewVector3FromFloats(x, y, z float64) *Vector3 {
	v := new(Vector3)
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

func NewVector3FromFloatsDifference(x1, y1, z1, x2, y2, z2 float64) *Vector3 {
	v := new(Vector3)
	v.X = x2 - x1
	v.y = y2 - y1
	v.Z = z2 - z1
	return v
}

func NewVector3FromVector3Difference(p1, p2 *Vector3) *Vector3 {
	v := new(Vector3)
	v.X = p2.X - p1.X
	v.Y = p2.Y - p1.Y
	v.Z = p2.Z - p1.Z
	return v
}

func (v *Vector3) DistanceFloats(x, y, z float64) float64 {
	xd := v.X - x
	yd := v.Y - y
	zd := v.Z - z
	return math.Sqrt(xd*xd + yd*yd + zd*zd)
}

func (v *Vector3) DistanceVector3(v2 *Vector3) float64 {
	xd := v.X - v2.X
	yd := v.Y - v2.Y
	zd := v.Z - v2.Z
	return math.Sqrt(xd*xd + yd*yd + zd*zd)
}

func (v *Vector3) DistanceSquaredFloats(x, y, z float64) float64 {
	xd := v.X - x
	yd := v.Y - y
	zd := v.Z - z
	return xd*xd + yd*yd + zd*zd
}

func (v *Vector3) DistanceSquaredVector3(v2 *Vector3) float64 {
	xd := v.X - v2.X
	yd := v.Y - v2.Y
	zd := v.Z - v2.Z
	return xd*xd + yd*yd + zd*zd
}

func TripleProduct(a, b, c *Vector3) *Vector3 {
	v := new(Vector3)
	ac := a.X*c.X + a.Y*c.Y + a.Z*c.Z
	bc := b.X*c.X + b.Y*c.Y + b.Z*c.Z
	v.X = b.X*ac - a.X*bc
	v.Y = b.Y*ac - a.Y*bc
	v.Z = b.Z*ac - a.Z*bc
	return v
}

func (v *Vector3) Set(v2 *Vector3) *Vector3 {
	*v = *v2
	return v
}

func (v *Vector3) GetXComponent() *Vector3 {
	return NewVector3FromFloats(v.X, 0, 0)
}

func (v *Vector3) GetYComponent() *Vector3 {
	return NewVector3FromFloats(0, v.Y, 0)
}

func (v *Vector3) GetXComponent() *Vector3 {
	return NewVector3FromFloats(v.X, 0, 0)
}

func (v *Vector3) GetMagnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3) GetMagnitudeSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vector3) SetMagnitude(magnitude float64) *Vector3 {
	if math.Abs(magnitude) <= dyn4go.Epsilon {
		v.X = 0
		v.Y = 0
		v.Z = 0
		return v
	}
	mag := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	mag = magnitude / mag
	v.X *= mag
	v.Y *= mag
	v.Z *= mag
	return v
}

func (v *Vector3) AddVector3(v2 *Vector3) *Vector3 {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
	return v
}

func (v *Vector3) AddFloats(x, y, z float64) *Vector3 {
	v.X += x
	v.Y += y
	v.Z += z
	return v
}

func (v *Vector3) SumVector3(v2 *Vector3) *Vector3 {
	v3 := new(Vector3)
	v3.X = v.X + v2.X
	v3.Y = v.Y + v2.Y
	v3.Z = v.Z + v2.Z
	return v3
}

func (v *Vector3) SumFloats(x, y, z float64) *Vector3 {
	v2 := new(Vector3)
	v2.X = v.X + x
	v2.Y = v.Y + y
	v2.Z = v.Z + z
	return v2
}

func (v *Vector3) SubtractVector3(v2 *Vector3) *Vector3 {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
	return v
}

func (v *Vector3) SubtractFloats(x, y, z float64) *Vector3 {
	v.X -= x
	v.Y -= y
	v.Z -= z
	return v
}

func (v *Vector3) DifferenceVector3(v2 *Vector3) *Vector3 {
	v3 := new(Vector3)
	v3.X = v.X - v2.X
	v3.Y = v.Y - v2.Y
	v3.Z = v.Z - v2.Z
	return v3
}

func (v *Vector3) DifferenceFloats(x, y, z float64) *Vector3 {
	v2 := new(Vector3)
	v2.X = v.X - x
	v2.Y = v.Y - y
	v2.Z = v.Z - z
	return v2
}

func (v *Vector3) HereToVector3(v2 *Vector3) *Vector3 {
	return NewVector3FromFloats(v2.X-v.X, v2.Y-v.Y, v2.Z-v.Z)
}

func (v *Vector3) HereToFloats(x, y, z float64) *Vector3 {
	return NewVector3FromFloats(x-v.X, y-v.Y, z-v.Z)
}

func (v *Vector3) Multiply(s float64) *Vector3 {
	v.X * s
	v.Y * s
	v.Z * s
	return v
}

func (v *Vector3) Product(s float64) *Vector3 {
	return NewVector3FromFloats(v.X*s, v.Y*s, v.Z*s)
}

func (v *Vector3) DotVector3(v2 *Vector3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v *Vector3) DotFloats(x, y, z float64) float64 {
	return v.X*x + v.Y*y + v.Z*z
}

func (v *Vector3) CrossVector3(v2 *Vector3) *Vector3 {
	return NewVector3FromFloats(v.Y*v2.Z-v.Z*v2.Y, v.Z*v2.X-v.X*v2.Z, v.X*v2.Y-v.Y*v2.X)
}

func (v *Vector3) CrossFloats(x, y, z float64) *Vector3 {
	return NewVector3FromFloats(v.Y*z-v.Z*y, v.Z*x-v.X*z, v.X*y-v.Y*x)
}

func (v *Vector3) IsOrthogonal(v2 *Vector3) bool {
	return math.Abs(v.X*v2.X+v.Y*v2.Y+v.Z*v2.Z) <= dyn4go.Epsilon
}

func (v *Vector3) IsZero() bool {
	return math.Abs(v.X) <= dyn4go.Epsilon && math.Abs(v.Y) <= dyn4go.Epsilon && math.Abs(v.Z) <= dyn4go.Epsilon
}

func (v *Vector3) Negate() *Vector3 {
	v.X *= -1
	v.Y *= -1
	v.Z *= -1
	return v
}

func (v *Vector3) GetNegative() *Vector3 {
	return NewVector3FromFloats(-v.X, -v.Y, -v.Z)
}

func (v *Vector3) Zero() *Vector3 {
	v.X = 0
	v.Y = 0
	v.Z = 0
	return v
}

func (v *Vector3) Project(v2 *Vector3) *Vector3 {
	dotProd := v.DotVector3(v2)
	denom := v.DotVector3(v)
	if denom <= dyn4go.Epsilon {
		return new(Vector3)
	}
	denom = dotProd / denom
	return NewVector3FromFloats(denom*v2.X, denom*v2.Y, denom*v2.Z)
}

func (v *Vector3) GetNormailised() *Vector3 {
	mag := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if mag <= dyn4go.Epsilon {
		return new(Vector3)
	}
	mag = 1 / mag
	return NewVector3FromFloats(v.X*mag, v.Y*mag, v.Z*mag)
}

func (v *Vector3) Normalise() float64 {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	if magnitude <= dyn4go.Epsilon {
		return new(Vector3)
	}
	m = 1 / magnitude
	v.X *= m
	v.Y *= m
	v.Z *= m
	return magnitude
}
