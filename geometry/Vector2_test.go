package geometry2

import (
	"math"
	"testing"

	"github.com/LSFN/dyn4go"
)

func TestVector2Create(t *testing.T) {
	v1 := new(Vector2)
	// should default to zero
	dyn4go.AssertEqual(t, 0.0, v1.X)
	dyn4go.AssertEqual(t, 0.0, v1.Y)

	v2 := NewVector2FromXY(1.0, 2.0)
	dyn4go.AssertEqual(t, 1.0, v2.X)
	dyn4go.AssertEqual(t, 2.0, v2.Y)

	v3 := NewVector2FromVector2(v2)
	dyn4go.AssertFalse(t, v3 == v2)
	dyn4go.AssertEqual(t, 1.0, v3.X)
	dyn4go.AssertEqual(t, 2.0, v3.Y)

	v4 := NewVector2FromA2B_XY(0.0, 1.0, 2.0, 3.0)
	dyn4go.AssertEqual(t, 2.0, v4.X)
	dyn4go.AssertEqual(t, 2.0, v4.Y)

	v5 := NewVector2FromA2B(v2, v1)
	dyn4go.AssertEqual(t, -1.0, v5.X)
	dyn4go.AssertEqual(t, -2.0, v5.Y)

	v7 := NewVector2FromDirection(math.Pi / 6)
	dyn4go.AssertEqualWithinError(t, 1.000, v7.GetMagnitude(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 30.000, dyn4go.RadToDeg(v7.GetDirection()), 1.0E-4)

	v6 := NewVector2FromMagnitudeAndDirection(1.0, math.Pi/2)
	dyn4go.AssertEqualWithinError(t, 0.000, v6.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, v6.Y, 1.0e-3)
}

func TestVector2Copy(t *testing.T) {
	v := NewVector2FromXY(1.0, 3.0)
	vc := NewVector2FromVector2(v)

	dyn4go.AssertFalse(t, v == vc)
	dyn4go.AssertEqual(t, v.X, vc.X)
	dyn4go.AssertEqual(t, v.Y, vc.Y)
}

func TestVector2Distance(t *testing.T) {
	v := new(Vector2)

	dyn4go.AssertEqualWithinError(t, 4.000, v.DistanceSquaredFromXY(2.0, 0.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 5.000, v.DistanceSquaredFromXY(2.0, -1.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, v.DistanceFromXY(2.0, 0.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 5.000, v.DistanceFromXY(3.0, 4.0), 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 4.000, v.DistanceSquaredFromVector2(NewVector2FromXY(2.0, 0.0)), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 5.000, v.DistanceSquaredFromVector2(NewVector2FromXY(2.0, -1.0)), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, v.DistanceFromVector2(NewVector2FromXY(2.0, 0.0)), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 5.000, v.DistanceFromVector2(NewVector2FromXY(3.0, 4.0)), 1.0e-3)
}

func TestVector2DistanceBugInVersions_1_1_0_to_3_1_7(t *testing.T) {
	v := NewVector2FromXY(1.0, 2.0)
	dyn4go.AssertEqualWithinError(t, 2.236, v.GetMagnitude(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.236, v.DistanceFromXY(2.0, 0.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, v.DistanceFromXY(1.0, 2.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.414, v.DistanceFromXY(2.0, 1.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 4.242, v.DistanceFromXY(-2.0, -1.0), 1.0e-3)
}

func TestVector2TripleProduct(t *testing.T) {
	v1 := NewVector2FromXY(1.0, 1.0)
	v2 := NewVector2FromXY(1.0, -1.0)

	r := Vector2TripleProduct(v1, v2, v2)

	// the below would be -1.0 if the vectors were normalized
	dyn4go.AssertEqualWithinError(t, -2.000, r.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -2.000, r.Y, 1.0e-3)
}

func TestVector2Equals(t *testing.T) {
	v := NewVector2FromXY(1.0, 2.0)

	dyn4go.AssertTrue(t, v.EqualsVector2(v))
	dyn4go.AssertTrue(t, v.EqualsVector2(NewVector2FromVector2(v)))
	dyn4go.AssertTrue(t, v.EqualsVector2(NewVector2FromXY(1.0, 2.0)))
	dyn4go.AssertTrue(t, v.EqualsXY(1.0, 2.0))

	dyn4go.AssertFalse(t, v.EqualsVector2(NewVector2FromVector2(v).SetToXY(2.0, 1.0)))
	dyn4go.AssertFalse(t, v.EqualsXY(2.0, 2.0))
}

func TestVector2Set(t *testing.T) {
	v := new(Vector2)

	v2 := NewVector2FromXY(1.0, -3.0)
	v.SetToVector2(v2)

	dyn4go.AssertFalse(t, v == v2)
	dyn4go.AssertEqual(t, 1.0, v.X)
	dyn4go.AssertEqual(t, -3.0, v.Y)

	v.SetToXY(-1.0, 0.0)
	t.Log(v)
	dyn4go.AssertEqual(t, -1.0, v.X)
	dyn4go.AssertEqual(t, 0.0, v.Y)

	v.SetDirection(math.Pi / 2)
	dyn4go.AssertEqualWithinError(t, 0.0, v.X, 1E-10)
	dyn4go.AssertEqual(t, 1.0, v.Y)

	v.SetMagnitude(3.0)
	dyn4go.AssertEqualWithinError(t, 0.0, v.X, 1E-10)
	dyn4go.AssertEqual(t, 3.0, v.Y)
}

func TestVector2Get(t *testing.T) {
	v := NewVector2FromXY(3.0, 4.0)

	x := v.GetXComponent()
	y := v.GetYComponent()

	dyn4go.AssertEqual(t, 3.0, x.X)
	dyn4go.AssertEqual(t, 0.0, x.Y)
	dyn4go.AssertEqual(t, 0.0, y.X)
	dyn4go.AssertEqual(t, 4.0, y.Y)

	dyn4go.AssertEqualWithinError(t, 5.000, v.GetMagnitude(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 25.000, v.GetMagnitudeSquared(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 53.130, dyn4go.RadToDeg(v.GetDirection()), 1.0e-3)

	v2 := NewVector2FromXY(-4.0, 3.0)
	dyn4go.AssertEqualWithinError(t, 90.000, dyn4go.RadToDeg(v.GetAngleBetween(v2)), 1.0e-3)

	v2 = v.GetLeftHandOrthogonalVector()
	dyn4go.AssertEqual(t, 4.0, v2.X)
	dyn4go.AssertEqual(t, -3.0, v2.Y)

	v2 = v.GetRightHandOrthogonalVector()
	dyn4go.AssertEqual(t, -4.0, v2.X)
	dyn4go.AssertEqual(t, 3.0, v2.Y)

	v2 = v.GetNegative()
	dyn4go.AssertEqual(t, -3.0, v2.X)
	dyn4go.AssertEqual(t, -4.0, v2.Y)

	v2 = v.GetNormalized()
	dyn4go.AssertEqualWithinError(t, 0.600, v2.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.800, v2.Y, 1.0e-3)
}

func TestVector2Add(t *testing.T) {
	v1 := NewVector2FromXY(1.0, 2.0)
	v2 := NewVector2FromXY(-2.0, 1.0)

	v3 := v1.SumVector2(v2)
	dyn4go.AssertEqual(t, -1.0, v3.X)
	dyn4go.AssertEqual(t, 3.0, v3.Y)

	v3 = v1.SumXY(3.0, -7.5)
	dyn4go.AssertEqual(t, 4.0, v3.X)
	dyn4go.AssertEqual(t, -5.5, v3.Y)

	v1.AddVector2(v2)
	dyn4go.AssertEqual(t, -1.0, v1.X)
	dyn4go.AssertEqual(t, 3.0, v1.Y)

	v1.AddXY(-2.0, 1.0)
	dyn4go.AssertEqual(t, -3.0, v1.X)
	dyn4go.AssertEqual(t, 4.0, v1.Y)
}

func TestVector2Subtract(t *testing.T) {
	v1 := NewVector2FromXY(1.0, 2.0)
	v2 := NewVector2FromXY(-2.0, 1.0)

	v3 := v1.DifferenceVector2(v2)
	dyn4go.AssertEqual(t, 3.0, v3.X)
	dyn4go.AssertEqual(t, 1.0, v3.Y)

	v3 = v1.DifferenceXY(3.0, -7.5)
	dyn4go.AssertEqual(t, -2.0, v3.X)
	dyn4go.AssertEqual(t, 9.5, v3.Y)

	v1.SubtractVector2(v2)
	dyn4go.AssertEqual(t, 3.0, v1.X)
	dyn4go.AssertEqual(t, 1.0, v1.Y)

	v1.SubtractXY(-2.0, 1.0)
	dyn4go.AssertEqual(t, 5.0, v1.X)
	dyn4go.AssertEqual(t, 0.0, v1.Y)
}

func TestVector2To(t *testing.T) {
	p1 := NewVector2FromXY(1.0, 1.0)
	p2 := NewVector2FromXY(0.0, 1.0)

	r := p1.HereToVector2(p2)

	dyn4go.AssertEqual(t, -1.0, r.X)
	dyn4go.AssertEqual(t, 0.0, r.Y)

	r = p1.HereToXY(2.0, 0.0)

	dyn4go.AssertEqual(t, 1.0, r.X)
	dyn4go.AssertEqual(t, -1.0, r.Y)
}

func TestVector2Multiply(t *testing.T) {
	v1 := NewVector2FromXY(2.0, 1.0)

	r := v1.Product(-1.5)
	dyn4go.AssertEqual(t, -3.0, r.X)
	dyn4go.AssertEqual(t, -1.5, r.Y)

	v1.Multiply(-1.5)
	dyn4go.AssertEqual(t, -3.0, v1.X)
	dyn4go.AssertEqual(t, -1.5, v1.Y)
}

func TestVector2Dot(t *testing.T) {
	v1 := NewVector2FromXY(1.0, 1.0)
	v2 := NewVector2FromXY(0.0, 1.0)

	dyn4go.AssertEqual(t, 1.0, v1.DotVector2(v2))
	// test a perpendicular vector
	dyn4go.AssertEqual(t, 0.0, v1.DotVector2(v1.GetLeftHandOrthogonalVector()))
	dyn4go.AssertEqual(t, v1.GetMagnitudeSquared(), v1.DotVector2(v1))

	dyn4go.AssertEqual(t, 1.0, v1.DotXY(0.0, 1.0))
	// test a perpendicular vector
	dyn4go.AssertEqual(t, 0.0, v1.DotXY(-1.0, 1.0))
	dyn4go.AssertEqual(t, 2.0, v1.DotXY(1.0, 1.0))
}

func TestVector2Cross(t *testing.T) {
	v1 := NewVector2FromXY(1.0, 1.0)
	v2 := NewVector2FromXY(0.0, 1.0)

	dyn4go.AssertEqual(t, 0.0, v1.CrossVector2(v1))
	dyn4go.AssertEqual(t, 1.0, v1.CrossVector2(v2))
	dyn4go.AssertEqual(t, -2.0, v1.CrossVector2(v1.GetLeftHandOrthogonalVector()))

	dyn4go.AssertEqual(t, 0.0, v1.CrossXY(1.0, 1.0))
	dyn4go.AssertEqual(t, 1.0, v1.CrossXY(0.0, 1.0))
	dyn4go.AssertEqual(t, 2.0, v1.CrossXY(-1.0, 1.0))

	r := v1.CrossZ(3.0)

	dyn4go.AssertEqual(t, -3.0, r.X)
	dyn4go.AssertEqual(t, 3.0, r.Y)
}

func TestVector2IsOrthogonal(t *testing.T) {
	v1 := NewVector2FromXY(1.0, 1.0)
	v2 := NewVector2FromXY(0.0, 1.0)

	dyn4go.AssertFalse(t, v1.IsOrthogonalVector2(v2))
	dyn4go.AssertTrue(t, v1.IsOrthogonalVector2(v1.GetLeftHandOrthogonalVector()))
	dyn4go.AssertTrue(t, v1.IsOrthogonalVector2(v1.GetRightHandOrthogonalVector()))
	dyn4go.AssertFalse(t, v1.IsOrthogonalVector2(v1))

	dyn4go.AssertFalse(t, v1.IsOrthogonalXY(0.0, 1.0))
	dyn4go.AssertTrue(t, v1.IsOrthogonalXY(1.0, -1.0))
	dyn4go.AssertTrue(t, v1.IsOrthogonalXY(-1.0, 1.0))
	dyn4go.AssertFalse(t, v1.IsOrthogonalXY(1.0, 1.0))
}

func TestVector2IsZero(t *testing.T) {
	v := new(Vector2)

	dyn4go.AssertTrue(t, v.IsZero())

	v.SetToXY(1.0, 0.0)
	dyn4go.AssertFalse(t, v.IsZero())

	v.SetToXY(1.0, 1.0)
	dyn4go.AssertFalse(t, v.IsZero())

	v.SetToXY(0.0, 1.0)
	dyn4go.AssertFalse(t, v.IsZero())
}

func TestVector2Negate(t *testing.T) {
	v := NewVector2FromXY(1.0, -6.0)

	v.Negate()
	dyn4go.AssertEqual(t, -1.0, v.X)
	dyn4go.AssertEqual(t, 6.0, v.Y)
}

func TestVector2Zero(t *testing.T) {
	v := NewVector2FromXY(1.0, -2.0)

	v.Zero()
	dyn4go.AssertEqual(t, 0.0, v.X)
	dyn4go.AssertEqual(t, 0.0, v.Y)
}

func TestVector2Rotate(t *testing.T) {
	v := NewVector2FromXY(2.0, 1.0)

	v.RotateAboutOrigin(math.Pi / 2)
	dyn4go.AssertEqualWithinError(t, -1.000, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, v.Y, 1.0e-3)

	v.RotateAboutXY(math.Pi/3, 0.0, 1.0)
	dyn4go.AssertEqualWithinError(t, -1.366, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.634, v.Y, 1.0e-3)
}

func TestVector2Project(t *testing.T) {
	v1 := NewVector2FromXY(1.0, 1.0)
	v2 := NewVector2FromXY(0.5, 1.0)

	r := v1.Project(v2)

	dyn4go.AssertEqualWithinError(t, 0.600, r.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.200, r.Y, 1.0e-3)
}

func TestVector2Left(t *testing.T) {
	v := NewVector2FromXY(11.0, 2.5)
	v.Left()

	dyn4go.AssertEqual(t, 2.5, v.X)
	dyn4go.AssertEqual(t, -11.0, v.Y)
}

func TestVector2Right(t *testing.T) {
	v := NewVector2FromXY(11.0, 2.5)
	v.Right()

	dyn4go.AssertEqual(t, -2.5, v.X)
	dyn4go.AssertEqual(t, 11.0, v.Y)
}

func TestVector2Normalize(t *testing.T) {
	v := NewVector2FromXY(3.0, 4.0)
	v.Normalize()

	dyn4go.AssertEqualWithinError(t, 3.0/5.0, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 4.0/5.0, v.Y, 1.0e-3)
}

func TestVector2GetAngleBetweenRange(t *testing.T) {
	v1 := NewVector2FromXY(-1.0, 2.0)
	v2 := NewVector2FromXY(-2.0, -1.0)

	// this should return in the range of -pi,pi
	dyn4go.AssertTrue(t, math.Pi >= math.Abs(v1.GetAngleBetween(v2)))
}
