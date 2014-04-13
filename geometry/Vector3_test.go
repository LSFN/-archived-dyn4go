package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests the create methods.
 */

func TestVector3Create(t *testing.T) {
	v1 := new(Vector3)
	// should default to zero
	dyn4go.AssertEqual(t, 0.0, v1.X)
	dyn4go.AssertEqual(t, 0.0, v1.Y)
	dyn4go.AssertEqual(t, 0.0, v1.Z)

	v2 := NewVector3FromFloats(1.0, 2.0, 3.0)
	dyn4go.AssertEqual(t, 1.0, v2.X)
	dyn4go.AssertEqual(t, 2.0, v2.Y)
	dyn4go.AssertEqual(t, 3.0, v2.Z)

	v3 := NewVector3FromVector3(v2)
	dyn4go.AssertFalse(t, v3 == v2)
	dyn4go.AssertEqual(t, 1.0, v3.X)
	dyn4go.AssertEqual(t, 2.0, v3.Y)
	dyn4go.AssertEqual(t, 3.0, v3.Z)

	v4 := NewVector3FromFloatsDifference(0.0, 1.0, 1.0, 2.0, 3.0, 1.0)
	dyn4go.AssertEqual(t, 2.0, v4.X)
	dyn4go.AssertEqual(t, 2.0, v4.Y)
	dyn4go.AssertEqual(t, 0.0, v4.Z)

	v5 := NewVector3FromVector3Difference(v2, v1)
	dyn4go.AssertEqual(t, -1.0, v5.X)
	dyn4go.AssertEqual(t, -2.0, v5.Y)
	dyn4go.AssertEqual(t, -3.0, v5.Z)
}

/**
 * Tests the copy method.
 */

func TestVector3Copy(t *testing.T) {
	v := NewVector3FromFloats(1.0, 3.0, 2.0)
	vc := NewVector3FromVector3(v)

	dyn4go.AssertFalse(t, v == vc)
	dyn4go.AssertEqual(t, v.X, vc.X)
	dyn4go.AssertEqual(t, v.Y, vc.Y)
	dyn4go.AssertEqual(t, v.Z, vc.Z)
}

/**
 * Tests the distance methods.
 */

func TestVector3Distance(t *testing.T) {
	v := new(Vector3)

	dyn4go.AssertEqualWithinError(t, 4.000, v.DistanceSquaredFloats(2.0, 0.0, 0.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 9.000, v.DistanceSquaredFloats(2.0, -1.0, 2.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, v.DistanceFloats(2.0, 0.0, 0.0), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, v.DistanceFloats(2.0, -1.0, 2.0), 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 4.000, v.DistanceSquaredVector3(NewVector3FromFloats(2.0, 0.0, 0.0)), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 9.000, v.DistanceSquaredVector3(NewVector3FromFloats(2.0, -1.0, 2.0)), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, v.DistanceVector3(NewVector3FromFloats(2.0, 0.0, 0.0)), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, v.DistanceVector3(NewVector3FromFloats(2.0, -1.0, 2.0)), 1.0e-3)
}

/**
 * Tests the triple product method.
 */

func TestVector3TripleProduct(t *testing.T) {
	v1 := NewVector3FromFloats(1.0, 1.0, 0.0)
	v2 := NewVector3FromFloats(0.0, -1.0, 1.0)

	r := Vector3TripleProduct(v1, v2, v2)

	dyn4go.AssertEqualWithinError(t, -2.000, r.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, r.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, r.Z, 1.0e-3)
}

/**
 * Tests the equals method.
 */

func TestVector3Equals(t *testing.T) {
	v := NewVector3FromFloats(1.0, 2.0, -1.0)

	dyn4go.AssertTrue(t, *v == *v)
	dyn4go.AssertTrue(t, *v == *NewVector3FromVector3(v))
	dyn4go.AssertTrue(t, *v == *NewVector3FromFloats(1.0, 2.0, -1.0))

	dyn4go.AssertFalse(t, *v == *NewVector3FromVector3(v).SetFloats(2.0, 1.0, -1.0))
	dyn4go.AssertFalse(t, *v == *NewVector3FromFloats(2.0, 2.0, 3.0))
}

/**
 * Tests the set methods.
 */

func TestVector3Set(t *testing.T) {
	v := new(Vector3)

	v2 := NewVector3FromFloats(1.0, -3.0, 2.0)
	v.SetVector3(v2)

	dyn4go.AssertFalse(t, v == v2)
	dyn4go.AssertEqual(t, 1.0, v.X)
	dyn4go.AssertEqual(t, -3.0, v.Y)
	dyn4go.AssertEqual(t, 2.0, v.Z)

	v.SetFloats(-1.0, 0.0, 0.0)
	dyn4go.AssertEqual(t, -1.0, v.X)
	dyn4go.AssertEqual(t, 0.0, v.Y)
	dyn4go.AssertEqual(t, 0.0, v.Z)

	v.SetMagnitude(3.0)
	dyn4go.AssertEqualWithinError(t, -3.0, v.X, 1.0e-3)
	dyn4go.AssertEqual(t, 0.0, v.Y)
	dyn4go.AssertEqual(t, 0.0, v.Z)
}

/**
 * Tests the get methods.
 */

func TestVector3Get(t *testing.T) {
	v := NewVector3FromFloats(2.0, 1.0, -2.0)

	x := v.GetXComponent()
	y := v.GetYComponent()
	z := v.GetZComponent()

	dyn4go.AssertEqual(t, 2.0, x.X)
	dyn4go.AssertEqual(t, 0.0, x.Y)
	dyn4go.AssertEqual(t, 0.0, x.Z)

	dyn4go.AssertEqual(t, 0.0, y.X)
	dyn4go.AssertEqual(t, 1.0, y.Y)
	dyn4go.AssertEqual(t, 0.0, y.Z)

	dyn4go.AssertEqual(t, 0.0, z.X)
	dyn4go.AssertEqual(t, 0.0, z.Y)
	dyn4go.AssertEqual(t, -2.0, z.Z)

	dyn4go.AssertEqualWithinError(t, 3.000, v.GetMagnitude(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 9.000, v.GetMagnitudeSquared(), 1.0e-3)

	v2 := v.GetNegative()
	dyn4go.AssertEqual(t, -2.0, v2.X)
	dyn4go.AssertEqual(t, -1.0, v2.Y)
	dyn4go.AssertEqual(t, 2.0, v2.Z)

	v2 = v.GetNormalised()
	dyn4go.AssertEqualWithinError(t, 0.666, v2.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.333, v2.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.666, v2.Z, 1.0e-3)
}

/**
 * Tests the add and sum methods.
 */

func TestVector3Add(t *testing.T) {
	v1 := NewVector3FromFloats(1.0, 2.0, 3.0)
	v2 := NewVector3FromFloats(-2.0, 1.0, -1.0)

	v3 := v1.SumVector3(v2)
	dyn4go.AssertEqual(t, -1.0, v3.X)
	dyn4go.AssertEqual(t, 3.0, v3.Y)
	dyn4go.AssertEqual(t, 2.0, v3.Z)

	v3 = v1.SumFloats(3.0, -7.5, 2.0)
	dyn4go.AssertEqual(t, 4.0, v3.X)
	dyn4go.AssertEqual(t, -5.5, v3.Y)
	dyn4go.AssertEqual(t, 5.0, v3.Z)

	v1.AddVector3(v2)
	dyn4go.AssertEqual(t, -1.0, v1.X)
	dyn4go.AssertEqual(t, 3.0, v1.Y)
	dyn4go.AssertEqual(t, 2.0, v1.Z)

	v1.AddFloats(-2.0, 1.0, 0.0)
	dyn4go.AssertEqual(t, -3.0, v1.X)
	dyn4go.AssertEqual(t, 4.0, v1.Y)
	dyn4go.AssertEqual(t, 2.0, v1.Z)
}

/**
 * Tests the subtact and difference methods.
 */

func TestVector3Subtract(t *testing.T) {
	v1 := NewVector3FromFloats(1.0, 2.0, 3.0)
	v2 := NewVector3FromFloats(-2.0, 1.0, -1.0)

	v3 := v1.DifferenceVector3(v2)
	dyn4go.AssertEqual(t, 3.0, v3.X)
	dyn4go.AssertEqual(t, 1.0, v3.Y)
	dyn4go.AssertEqual(t, 4.0, v3.Z)

	v3 = v1.DifferenceFloats(3.0, -7.5, 2.0)
	dyn4go.AssertEqual(t, -2.0, v3.X)
	dyn4go.AssertEqual(t, 9.5, v3.Y)
	dyn4go.AssertEqual(t, 1.0, v3.Z)

	v1.SubtractVector3(v2)
	dyn4go.AssertEqual(t, 3.0, v1.X)
	dyn4go.AssertEqual(t, 1.0, v1.Y)
	dyn4go.AssertEqual(t, 4.0, v1.Z)

	v1.SubtractFloats(-2.0, 1.0, 0.0)
	dyn4go.AssertEqual(t, 5.0, v1.X)
	dyn4go.AssertEqual(t, 0.0, v1.Y)
	dyn4go.AssertEqual(t, 4.0, v1.Z)
}

/**
 * Tests the to method.
 */

func TestVector3To(t *testing.T) {
	p1 := NewVector3FromFloats(1.0, 1.0, 1.0)
	p2 := NewVector3FromFloats(0.0, 1.0, 0.0)

	r := p1.HereToVector3(p2)

	dyn4go.AssertEqual(t, -1.0, r.X)
	dyn4go.AssertEqual(t, 0.0, r.Y)
	dyn4go.AssertEqual(t, -1.0, r.Z)

	r = p1.HereToFloats(2.0, 0.0, -1.0)

	dyn4go.AssertEqual(t, 1.0, r.X)
	dyn4go.AssertEqual(t, -1.0, r.Y)
	dyn4go.AssertEqual(t, -2.0, r.Z)
}

/**
 * Tests the multiply and product methods.
 */

func TestVector3Multiply(t *testing.T) {
	v1 := NewVector3FromFloats(2.0, 1.0, -1.0)

	r := v1.Product(-1.5)
	dyn4go.AssertEqual(t, -3.0, r.X)
	dyn4go.AssertEqual(t, -1.5, r.Y)
	dyn4go.AssertEqual(t, 1.5, r.Z)

	v1.Multiply(-1.5)
	dyn4go.AssertEqual(t, -3.0, v1.X)
	dyn4go.AssertEqual(t, -1.5, v1.Y)
	dyn4go.AssertEqual(t, 1.5, v1.Z)
}

/**
 * Tests the dot method.
 */

func TestVector3Dot(t *testing.T) {
	v1 := NewVector3FromFloats(1.0, 1.0, -1.0)
	v2 := NewVector3FromFloats(0.0, 1.0, 0.0)

	dyn4go.AssertEqual(t, 1.0, v1.DotVector3(v2))

	dyn4go.AssertEqual(t, 1.0, v1.DotFloats(0.0, 1.0, 0.0))

	// test a perpendicular vector
	dyn4go.AssertEqual(t, 0.0, v1.DotFloats(-1.0, 1.0, 0.0))

	dyn4go.AssertEqual(t, 2.0, v1.DotFloats(1.0, 1.0, 0.0))
}

/**
 * Tests the cross product methods.
 */

func TestVector3Cross(t *testing.T) {
	v1 := NewVector3FromFloats(1.0, 1.0, 0.0)
	v2 := NewVector3FromFloats(0.0, 1.0, -1.0)

	r := v1.CrossVector3(v1)
	dyn4go.AssertEqual(t, 0.0, r.X)
	dyn4go.AssertEqual(t, 0.0, r.Y)
	dyn4go.AssertEqual(t, 0.0, r.Z)

	r = v1.CrossVector3(v2)
	dyn4go.AssertEqual(t, -1.0, r.X)
	dyn4go.AssertEqual(t, 1.0, r.Y)
	dyn4go.AssertEqual(t, 1.0, r.Z)

	r = v1.CrossFloats(1.0, 1.0, 0.0)
	dyn4go.AssertEqual(t, 0.0, r.X)
	dyn4go.AssertEqual(t, 0.0, r.Y)
	dyn4go.AssertEqual(t, 0.0, r.Z)

	r = v1.CrossFloats(0.0, 1.0, 1.0)
	dyn4go.AssertEqual(t, 1.0, r.X)
	dyn4go.AssertEqual(t, -1.0, r.Y)
	dyn4go.AssertEqual(t, 1.0, r.Z)

	r = v1.CrossFloats(-1.0, 1.0, -1.0)
	dyn4go.AssertEqual(t, -1.0, r.X)
	dyn4go.AssertEqual(t, 1.0, r.Y)
	dyn4go.AssertEqual(t, 2.0, r.Z)
}

/**
 * Tests the isOrthoganal method.
 */

func TestVector3IsOrthogonal(t *testing.T) {
	v1 := NewVector3FromFloats(1.0, 1.0, 0.0)
	v2 := NewVector3FromFloats(0.0, 1.0, 2.0)

	dyn4go.AssertFalse(t, v1.IsOrthogonalVector3(v2))
	dyn4go.AssertFalse(t, v1.IsOrthogonalVector3(v1))

	dyn4go.AssertFalse(t, v1.IsOrthogonalFloats(0.0, 1.0, 0.0))
	dyn4go.AssertTrue(t, v1.IsOrthogonalFloats(1.0, -1.0, 0.0))
	dyn4go.AssertTrue(t, v1.IsOrthogonalFloats(-1.0, 1.0, 0.0))
	dyn4go.AssertFalse(t, v1.IsOrthogonalFloats(1.0, 1.0, 0.0))
}

/**
 * Tests the isZero method.
 */

func TestVector3IsZero(t *testing.T) {
	v := new(Vector3)

	dyn4go.AssertTrue(t, v.IsZero())

	v.SetFloats(1.0, 0.0, 0.0)
	dyn4go.AssertFalse(t, v.IsZero())

	v.SetFloats(1.0, 1.0, 0.0)
	dyn4go.AssertFalse(t, v.IsZero())

	v.SetFloats(0.0, 1.0, 1.0)
	dyn4go.AssertFalse(t, v.IsZero())

	v.SetFloats(0.0, 0.0, 1.0)
	dyn4go.AssertFalse(t, v.IsZero())
}

/**
 * Tests the negate method.
 */

func TestVector3Negate(t *testing.T) {
	v := NewVector3FromFloats(1.0, -6.0, 2.0)

	v.Negate()
	dyn4go.AssertEqual(t, -1.0, v.X)
	dyn4go.AssertEqual(t, 6.0, v.Y)
	dyn4go.AssertEqual(t, -2.0, v.Z)
}

/**
 * Tests the zero method.
 */

func TestVector3Zero(t *testing.T) {
	v := NewVector3FromFloats(1.0, -2.0, 3.0)

	v.Zero()
	dyn4go.AssertEqual(t, 0.0, v.X)
	dyn4go.AssertEqual(t, 0.0, v.Y)
	dyn4go.AssertEqual(t, 0.0, v.Z)
}

/**
 * Tests the project method.
 */

func TestVector3Project(t *testing.T) {
	v1 := NewVector3FromFloats(1.0, 1.0, 0.0)
	v2 := NewVector3FromFloats(0.5, 1.0, 1.0)

	r := v1.Project(v2)

	dyn4go.AssertEqualWithinError(t, 0.333, r.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.666, r.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.666, r.Z, 1.0e-3)
}

/**
 * Tests the normalize method.
 */

func TestVector3Normalize(t *testing.T) {
	v := NewVector3FromFloats(2.0, 1.0, 2.0)
	v.Normalise()

	dyn4go.AssertEqualWithinError(t, 2.0/3.0, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.0/3.0, v.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.0/3.0, v.Z, 1.0e-3)
}
