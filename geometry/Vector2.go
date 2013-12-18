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
	x, y float64
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
	v.x = b.x - a.x
	v.y = b.y - a.y
	return v
}

func NewVector2FromDirection(direction float64) *Vector2 {
	return &Vector2{math.Cos(direction), math.Sin(direction)}
}

func NewVector2FromMagnitudeAndDirection(magnitude, direction float64) *Vector2 {
	return &Vector2{magnitude * math.Cos(direction), magnitude * math.Sin(direction)}
}

func (v *Vector2) DistanceFromXY(x, y float64) float64 {
	return math.Hypot(v.x-x, v.y-y)
}

func (v *Vector2) DistanceFromVector2(v2 *Vector2) float64 {
	return math.Hypot(v.x-v2.x, v.y-v2.y)
}

func (v *Vector2) DistanceSquaredFromCY(x, y float64) float64 {
	return (v.x-x)*(v.x-x) + (v.y-y)*(v.y-y)
}

func (v *Vector2) DistanceSquaredFromVector2(v2 *Vector2) float64 {
	return (v.x-v2.x)*(v.x-v2.x) + (v.y-v2.y)*(v.y-v2.y)
}

func TripleProduct(a, b, c *Vector2) *Vector2 {
	v := new(Vector2)
	ac := a.x*c.x + a.y*c.y
	bc := b.x*c.x + b.y*c.y
	v.x = b.x*ac - a.x*bc
	v.y = b.y*ac - a.y*bc
	return v
}

func (v *Vector2) EqualsVector2(v2 *Vector2) bool {
	if v == nil {
		return false
	}
	return v == v2 || (v.x == v2.x && v.y == v2.y)
}

func (v *Vector2) EqualsXY(x, y float64) bool {
	return v.x == x && v.y == y
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
	return &Vector2{x: v.x}
}

func (v *Vector2) GetYComponent() *Vector2 {
	return &Vector2{y: v.y}
}

func (v *Vector2) GetMagnitude() float64 {
	return math.Hypot(v.x, v.y)
}

func (v *Vector2) GetMagnitudeSquared() float64 {
	return v.x*v.x + v.y*v.y
}

func (v *Vector2) SetMagnitude(magnitude float64) *Vector2 {
	if math.Abs(magnitude) <= dyn4go.Epsilon {
		v.x = 0.0
		v.y = 0.0
		return v
	}
	if v.IsZero() {
		return v
	}
	mag := math.Hypot(v.x, v.y)
	mag = magnitude / mag
	v.x *= mag
	v.y *= mag
	return v
}

func (v *Vector2) GetDirection() float64 {
	return math.Atan2(v.y, v.x)
}

func (v *Vector2) SetDirection(angle float64) *Vector2 {
	magnitude := math.Hypot(v.x, v.y)
	v.x = magnitude * math.Cos(angle)
	v.y = magnitude * math.Sin(angle)
	return v
}

func (v *Vector2) AddVector2(v2 *Vector2) *Vector2 {
	v.x += v2.x
	v.y += v2.y
	return v
}

func (v *Vector2) AddXY(x, y float64) *Vector2 {
	v.x += x
	v.y += y
	return v
}

func (v *Vector2) SumVector2(v2 *Vector2) *Vector2 {
	v3 := new(Vector2)
	v3.x = v.x + v2.x
	v3.y = v.y + v2.y
	return v3
}

func (v *Vector2) SumXY(x, y float64) *Vector2 {
	v3 := new(Vector2)
	v3.x = v.x + x
	v3.y = v.y + y
	return v3
}

func (v *Vector2) SubtractVector2(v2 *Vector2) *Vector2 {
	v.x -= v2.x
	v.y -= v2.y
	return v
}

func (v *Vector2) SubtractXY(x, y float64) *Vector2 {
	v.x -= x
	v.y -= y
	return v
}

func (v *Vector2) DifferenceVector2(v2 *Vector2) *Vector2 {
	v3 := new(Vector2)
	v3.x = v.x - v2.x
	v3.y = v.y - v2.y
	return v3
}

func (v *Vector2) DifferenceXY(x, y float64) *Vector2 {
	v3 := new(Vector2)
	v3.x = v.x - x
	v3.y = v.y - y
	return v3
}

func (v *Vector2) HereToVector2(v2 *Vector2) *Vector2 {
	v3 := new(Vector2)
	v3.x = v2.x - v.x
	v3.y = v2.y - v.y
	return v3
}

func (v *Vector2) HereToXY(x, y float64) *Vector2 {
	v3 := new(Vector2)
	v3.x = x - v.x
	v3.y = y - v.y
	return v3
}

func (v *Vector2) Multiply(scalar float64) *Vector2 {
	v.x *= scalar
	v.y *= scalar
	return v
}

func (v *Vector2) Product(scalar float64) *Vector2 {
	v2 := new(Vector2)
	v2.x = v.x * scalar
	v2.y = v.y * scalar
	return v2
}

func (v *Vector2) DotVector2(v2 *Vector2) float64 {
	return v.x*v2.x + v.y*v2.y
}

func (v *Vector2) DotXY(x, y float64) float64 {
	return v.x*x + v.y*y
}

func (v *Vector2) CrossVector2(v2 *Vector2) float64 {
	return v.x*v2.y - v.y*v2.x
}

func (v *Vector2) CrossXY(x, y float64) float64 {
	return v.x*y - v.y*x
}

func (v *Vector2) CrossZ(z float64) *Vector2 {
	v2 := new(Vector2)
	v2.x = -v.y * z
	v2.y = v.x * z
	return v2
}

func (v *Vector2) IsOrthogonalVector2(v2 *Vector2) bool {
	return math.Abs(v.x*v2.x+v.y*v2.y) <= dyn4go.Epsilon
}

func (v *Vector2) IsOrthogonalXY(x, y float64) bool {
	return math.Abs(v.x*x+v.y*y) <= dyn4go.Epsilon
}

func (v *Vector2) IsZero() bool {
	return math.Abs(v.x) <= dyn4go.Epsilon && math.Abs(v.y) <= dyn4go.Epsilon
}

func (v *Vector2) Negate() *Vector2 {
	v.x *= 1.0
	v.y *= 1.0
	return v
}

func (v *Vector2) GetNegative() *Vector2 {
	v2 := new(Vector2)
	v2.x = -v.x
	v2.y = -v.y
	return v2
}

func (v *Vector2) Zero() *Vector2 {
	v.x = 0
	v.y = 0
	return v
}

func (v *Vector2) RotateAboutOrigin(theta float64) *Vector2 {
	cos := math.Cos(theta)
	sin := math.Sin(theta)
	x := v.x
	y := v.y
	v.x = x*cos - y*sin
	v.y = x*sin + y*cos
	return v
}

func (v *Vector2) RotateAboutXY(theta, x, y float64) *Vector2 {
	v.x -= x
	v.y -= y
	v.RotateAboutOrigin(theta)
	v.x += x
	v.y += y
	return v
}

func (v *Vector2) RotateAboutVector2(theta float64, v2 *Vector2) *Vector2 {
	v.x -= v2.x
	v.y -= v2.y
	v.RotateAboutOrigin(theta)
	v.x += v2.x
	v.y += v2.y
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
	v3.x = denominator * v2.x
	v3.y = denominator * v2.y
	return v3
}

func (v *Vector2) GetRightHandOrthogonalVector() *Vector2 {
	return NewVector2FromXY(-v.y, v.x)
}

func (v *Vector2) Right() *Vector2 {
	v.x, v.y = -v.y, v.x
	return v
}

func (v *Vector2) GetLeftHandOrthogonalVector() *Vector2 {
	return NewVector2FromXY(v.y, -v.x)
}

func (v *Vector2) Left() *Vector2 {
	v.x, v.y = v.y, -v.x
	return v
}

func (v *Vector2) GetNormalized() *Vector2 {
	magnitude := v.GetMagnitude()
	if magnitude <= dyn4go.Epsilon {
		return new(Vector2)
	}
	magnitude = 1 / magnitude
	return NewVector2FromXY(v.x*magnitude, v.y*magnitude)
}

func (v *Vector2) Normalize() float64 {
	magnitude := math.Hypot(v.x, v.y)
	if magnitude <= dyn4go.Epsilon {
		return 0
	}
	m := 1 / magnitude
	v.x *= m
	v.y *= m
	return magnitude
}

func (v *Vector2) GetAngleBetween(v2 *Vector2) float64 {
	a := math.Atan2(v2.y, v2.x) - math.Atan2(v.y, v.x)
	if a > math.Pi {
		return a - 2*math.Pi
	} else if a < math.Pi {
		return a + 2*math.Pi
	}
	return a
}
