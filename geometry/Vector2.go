package geometry

import (
	"math"

	"github.com/LSFN/dyn4go"
)

var (
	X_AXIS = Vector2{1.0, 0.0}
	Y_AXIS = Vector2{0.0, 1.0}
)

type Vector2 struct {
	X, Y float64
}

func NewVector2FromXY(x, y float64) *Vector2 {
	return &Vector2{x, y}
}

func NewVector2FromVector2(vOrig *Vector2) *Vector2 {
	v := *vOrig
	return &v
}

func NewVector2FromA2B_XY(xa, ya, xb, yb float64) *Vector2 {
	return &Vector2{xb - xa, yb - ya}
}

func NewVector2FromA2B(a, b *Vector2) *Vector2 {
	v := new(Vector2)
	v.X = b.X - a.X
	v.Y = b.Y - a.Y
	return v
}

func NewVector2FromDirection(direction float64) *Vector2 {
	return &Vector2{math.Cos(direction), math.Sin(direction)}
}

func NewVector2FromMagnitudeAndDirection(magnitude, direction float64) *Vector2 {
	return &Vector2{magnitude * math.Cos(direction), magnitude * math.Sin(direction)}
}

func (v *Vector2) DistanceFromXY(x, y float64) float64 {
	return math.Hypot(v.X-x, v.Y-y)
}

func (v *Vector2) DistanceFromVector2(v2 *Vector2) float64 {
	return math.Hypot(v.X-v2.X, v.Y-v2.Y)
}

func (v *Vector2) DistanceSquaredFromCY(x, y float64) float64 {
	return (v.X-x)*(v.X-x) + (v.Y-y)*(v.Y-y)
}

func (v *Vector2) DistanceSquaredFromVector2(v2 *Vector2) float64 {
	return (v.X-v2.X)*(v.X-v2.X) + (v.Y-v2.Y)*(v.Y-v2.Y)
}

func TripleProduct(a, b, c *Vector2) *Vector2 {
	v := new(Vector2)
	ac := a.X*c.X + a.Y*c.Y
	bc := b.X*c.X + b.Y*c.Y
	v.X = b.X*ac - a.X*bc
	v.Y = b.Y*ac - a.Y*bc
	return v
}

func (v *Vector2) EqualsVector2(v2 *Vector2) bool {
	if v == nil {
		return false
	}
	return v == v2 || (v.X == v2.X && v.Y == v2.Y)
}

func (v *Vector2) EqualsXY(x, y float64) bool {
	return v.X == x && v.Y == y
}

func (v *Vector2) SetToVector2(v2 *Vector2) *Vector2 {
	*v = *v2
	return v
}

func (v *Vector2) SetToXY(x, y float64) *Vector2 {
	v = &Vector2{x, y}
	return v
}

func (v *Vector2) GetXComponent() *Vector2 {
	return &Vector2{X: v.X}
}

func (v *Vector2) GetYComponent() *Vector2 {
	return &Vector2{Y: v.Y}
}

func (v *Vector2) GetMagnitude() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v *Vector2) GetMagnitudeSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector2) SetMagnitude(magnitude float64) *Vector2 {
	if math.Abs(magnitude) <= dyn4go.Epsilon {
		v.X = 0.0
		v.Y = 0.0
		return v
	}
	if v.IsZero() {
		return v
	}
	mag := math.Hypot(v.X, v.Y)
	mag = magnitude / mag
	v.X *= mag
	v.Y *= mag
	return v
}

func (v *Vector2) GetDirection() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v *Vector2) SetDirection(angle float64) *Vector2 {
	magnitude := math.Hypot(v.X, v.Y)
	v.X = magnitude * math.Cos(angle)
	v.Y = magnitude * math.Sin(angle)
	return v
}

func (v *Vector2) AddVector2(v2 *Vector2) *Vector2 {
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v *Vector2) AddXY(x, y float64) *Vector2 {
	v.X += x
	v.Y += y
	return v
}

func (v *Vector2) SumVector2(v2 *Vector2) *Vector2 {
	v3 := new(Vector2)
	v3.X = v.X + v2.X
	v3.Y = v.Y + v2.Y
	return v3
}

func (v *Vector2) SumXY(x, y float64) *Vector2 {
	v3 := new(Vector2)
	v3.X = v.X + x
	v3.Y = v.Y + y
	return v3
}

func (v *Vector2) SubtractVector2(v2 *Vector2) *Vector2 {
	v.X -= v2.X
	v.Y -= v2.Y
	return v
}

func (v *Vector2) SubtractXY(x, y float64) *Vector2 {
	v.X -= x
	v.Y -= y
	return v
}

func (v *Vector2) DifferenceVector2(v2 *Vector2) *Vector2 {
	v3 := new(Vector2)
	v3.X = v.X - v2.X
	v3.Y = v.Y - v2.Y
	return v3
}

func (v *Vector2) DifferenceXY(x, y float64) *Vector2 {
	v3 := new(Vector2)
	v3.X = v.X - x
	v3.Y = v.Y - y
	return v3
}

func (v *Vector2) HereToVector2(v2 *Vector2) *Vector2 {
	v3 := new(Vector2)
	v3.X = v2.X - v.X
	v3.Y = v2.Y - v.Y
	return v3
}

func (v *Vector2) HereToXY(x, y float64) *Vector2 {
	v3 := new(Vector2)
	v3.X = x - v.X
	v3.Y = y - v.Y
	return v3
}

func (v *Vector2) Multiply(scalar float64) *Vector2 {
	v.X *= scalar
	v.Y *= scalar
	return v
}

func (v *Vector2) Product(scalar float64) *Vector2 {
	v2 := new(Vector2)
	v2.X = v.X * scalar
	v2.Y = v.Y * scalar
	return v2
}

func (v *Vector2) DotVector2(v2 *Vector2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vector2) DotXY(x, y float64) float64 {
	return v.X*x + v.Y*y
}

func (v *Vector2) CrossVector2(v2 *Vector2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v *Vector2) CrossXY(x, y float64) float64 {
	return v.X*y - v.Y*x
}

func (v *Vector2) CrossZ(z float64) *Vector2 {
	v2 := new(Vector2)
	v2.X = -v.Y * z
	v2.Y = v.X * z
	return v2
}

func (v *Vector2) IsOrthogonalVector2(v2 *Vector2) bool {
	return math.Abs(v.X*v2.X+v.Y*v2.Y) <= dyn4go.Epsilon
}

func (v *Vector2) IsOrthogonalXY(x, y float64) bool {
	return math.Abs(v.X*x+v.Y*y) <= dyn4go.Epsilon
}

func (v *Vector2) IsZero() bool {
	return math.Abs(v.X) <= dyn4go.Epsilon && math.Abs(v.Y) <= dyn4go.Epsilon
}

func (v *Vector2) Negate() *Vector2 {
	v.X *= 1.0
	v.Y *= 1.0
	return v
}

func (v *Vector2) GetNegative() *Vector2 {
	v2 := new(Vector2)
	v2.X = -v.X
	v2.Y = -v.Y
	return v2
}

func (v *Vector2) Zero() *Vector2 {
	v.X = 0
	v.Y = 0
	return v
}

func (v *Vector2) RotateAboutOrigin(theta float64) *Vector2 {
	cos := math.Cos(theta)
	sin := math.Sin(theta)
	x := v.X
	y := v.Y
	v.X = x*cos - y*sin
	v.Y = x*sin + y*cos
	return v
}

func (v *Vector2) RotateAboutXY(theta, x, y float64) *Vector2 {
	v.X -= x
	v.Y -= y
	v.RotateAboutOrigin(theta)
	v.X += x
	v.Y += y
	return v
}

func (v *Vector2) RotateAboutVector2(theta float64, v2 *Vector2) *Vector2 {
	v.X -= v2.X
	v.Y -= v2.Y
	v.RotateAboutOrigin(theta)
	v.X += v2.X
	v.Y += v2.Y
	return v
}

func (v *Vector2) Project(v2 *Vector2) *Vector2 {
	dotProd := v.DotVector2(v2)
	denominator := v2.DotVector2(v2)
	v3 := new(Vector2)
	if denominator <= dyn4go.Epsilon {
		return v3
	}
	denominator = dotProd / denominator
	v3.X = denominator * v2.X
	v3.Y = denominator * v2.Y
	return v3
}

func (v *Vector2) GetRightHandOrthogonalVector() *Vector2 {
	return NewVector2FromXY(-v.Y, v.X)
}

func (v *Vector2) Right() *Vector2 {
	v.X, v.Y = -v.Y, v.X
	return v
}

func (v *Vector2) GetLeftHandOrthogonalVector() *Vector2 {
	return NewVector2FromXY(v.Y, -v.X)
}

func (v *Vector2) Left() *Vector2 {
	v.X, v.Y = v.Y, -v.X
	return v
}

func (v *Vector2) GetNormalized() *Vector2 {
	magnitude := v.GetMagnitude()
	if magnitude <= dyn4go.Epsilon {
		return new(Vector2)
	}
	magnitude = 1 / magnitude
	return NewVector2FromXY(v.X*magnitude, v.Y*magnitude)
}

func (v *Vector2) Normalize() float64 {
	magnitude := math.Hypot(v.X, v.Y)
	if magnitude <= dyn4go.Epsilon {
		return 0
	}
	m := 1 / magnitude
	v.X *= m
	v.Y *= m
	return magnitude
}

func (v *Vector2) GetAngleBetween(v2 *Vector2) float64 {
	a := math.Atan2(v2.Y, v2.X) - math.Atan2(v.Y, v.X)
	if a > math.Pi {
		return a - math.Pi * 2
	} else if a < math.Pi {
		return a + math.Pi * 2
	}
	return a
}
