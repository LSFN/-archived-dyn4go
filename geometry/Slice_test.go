package geometry

import (
	"math"
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests a zero radius.
 */
func TestCreateZeroRadius(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewSlice(0.0, dyn4go.DegToRad(50))

}

/**
 * Tests a negative radius.
 */
func TestCreateNegativeRadius(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewSlice(-1.0, dyn4go.DegToRad(50))
}

/**
 * Tests a zero theta.
 */
func TestCreateZeroTheta(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewSlice(1.0, 0)
}

/**
 * Tests a negative theta.
 */
func TestCreateNegativeTheta(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewSlice(1.0, -dyn4go.DegToRad(50))
}

/**
 * Tests the constructor.
 */
func TestCreateSuccess(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	slice := NewSlice(1.0, dyn4go.DegToRad(50))

	// the circle center should be the origin
	dyn4go.AssertEqualWithinError(t, 0.000, slice.GetCircleCenter().X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, slice.GetCircleCenter().Y, 1.0e-3)
}

/**
 * Tests the contains method.
 */
func TestContains(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))
	transform := NewTransform()
	p := NewVector2FromXY(0.5, -0.3)

	// shouldn't be inside
	dyn4go.AssertTrue(t, !e.ContainsVector2Transform(p, transform))

	// move it a bit
	transform.TranslateXY(-0.25, 0.0)

	// should be inside
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))

	p.SetToXY(0.75, 0.0)
	// should be on the edge
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))
}

/**
 * Tests the project method.
 */
func TestProject(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))
	transform := NewTransform()
	x := NewVector2FromXY(1.0, 0.0)
	y := NewVector2FromXY(0.0, 1.0)

	// try some translation
	transform.TranslateXY(1.0, 0.5)

	i := e.ProjectVector2Transform(x, transform)
	dyn4go.AssertEqualWithinError(t, 1.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, i.max, 1.0e-3)

	// try some rotation
	transform.RotateAboutVector2(dyn4go.DegToRad(30), transform.GetTransformedVector2(e.GetCenter()))

	i = e.ProjectVector2Transform(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.177, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.996, i.max, 1.0e-3)

	// try some local rotation
	e.TranslateXY(1.0, 0.5)
	e.RotateAboutCenter(dyn4go.DegToRad(30))

	i = e.ProjectVector2Transform(y, NewTransform())
	dyn4go.AssertEqualWithinError(t, 0.177, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.996, i.max, 1.0e-3)

	transform.Identity()
	transform.TranslateXY(0.0, 1.0)
	i = e.ProjectVector2Transform(y, transform)
	dyn4go.AssertEqualWithinError(t, 1.177, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.996, i.max, 1.0e-3)
}

/**
 * Tests the farthest methods.
 */
func TestGetFarthest(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))
	transform := NewTransform()
	x := NewVector2FromXY(1.0, 0.0)
	y := NewVector2FromXY(0.0, 1.0)

	// try some translation
	transform.TranslateXY(1.0, 0.5)

	p := e.GetFarthestPoint(x, transform)
	dyn4go.AssertEqualWithinError(t, 2.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, p.Y, 1.0e-3)

	// try some rotation
	transform.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	p = e.GetFarthestPoint(y, transform)
	dyn4go.AssertEqualWithinError(t, 1.573, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.319, p.Y, 1.0e-3)

	// try some local rotation
	e.TranslateXY(1.0, 0.5)
	e.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	p = e.GetFarthestPoint(y, NewTransform())
	dyn4go.AssertEqualWithinError(t, 1.573, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.319, p.Y, 1.0e-3)

	transform.Identity()
	transform.TranslateXY(0.0, 1.0)
	p = e.GetFarthestPoint(y, transform)
	dyn4go.AssertEqualWithinError(t, 1.573, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.319, p.Y, 1.0e-3)
}

/**
 * Tests the getAxes method.
 */
func TestGetAxes(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))
	// should be two axes + number of foci
	foci := []*Vector2{
		NewVector2FromXY(2.0, -0.5),
		NewVector2FromXY(1.0, 3.0),
	}
	axes := e.GetAxes(foci, NewTransform())
	dyn4go.AssertEqual(t, 4, len(axes))

	// make sure we get back the right axes
	axes = e.GetAxes(nil, NewTransform())
	dyn4go.AssertEqualWithinError(t, -0.422, axes[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.906, axes[0].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.422, axes[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.906, axes[1].Y, 1.0e-3)
}

/**
 * Tests the getFoci method.
 */
func TestGetFoci(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))
	foci := e.GetFoci(NewTransform())
	// should be two foci
	dyn4go.AssertEqual(t, 1, len(foci))
	// make sure the foci are correct
	dyn4go.AssertEqualWithinError(t, 0.000, foci[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, foci[0].Y, 1.0e-3)
}

/**
 * Tests the rotate methods.
 */
func TestRotate(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))
	// note: the center is not at the origin

	// rotate about center
	e.TranslateXY(1.0, 1.0)
	e.RotateAboutCenter(dyn4go.DegToRad(30))
	dyn4go.AssertEqualWithinError(t, 1.645, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, e.center.Y, 1.0e-3)

	// rotate about the origin
	e.RotateAboutOrigin(dyn4go.DegToRad(90))
	dyn4go.AssertEqualWithinError(t, -1.000, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.645, e.center.Y, 1.0e-3)
	e.TranslateVector2(e.GetCenter().GetNegative())

	// should move the center
	e.RotateAboutXY(dyn4go.DegToRad(90), 1.0, -1.0)
	dyn4go.AssertEqualWithinError(t, 0.000, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -2.000, e.center.Y, 1.0e-3)
}

/**
 * Tests the translate methods.
 */
func TestTranslate(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))

	e.TranslateXY(1.0, -0.5)

	dyn4go.AssertEqualWithinError(t, 1.645, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.500, e.center.Y, 1.0e-3)
}

/**
 * Tests the generated AABB.
 */
func TestCreateAABB(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))

	// using an identity transform
	aabb := e.CreateAABBTransform(NewTransform())
	dyn4go.AssertEqualWithinError(t, 0.000, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.422, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.422, aabb.GetMaxY(), 1.0e-3)

	// try using the default method
	aabb2 := e.CreateAABB()
	dyn4go.AssertEqual(t, aabb.GetMinX(), aabb2.GetMinX())
	dyn4go.AssertEqual(t, aabb.GetMinY(), aabb2.GetMinY())
	dyn4go.AssertEqual(t, aabb.GetMaxX(), aabb2.GetMaxX())
	dyn4go.AssertEqual(t, aabb.GetMaxY(), aabb2.GetMaxY())

	// test using a rotation and translation matrix
	tx := NewTransform()
	tx.RotateAboutOrigin(dyn4go.DegToRad(30.0))
	tx.TranslateXY(1.0, 2.0)

	aabb = e.CreateAABBTransform(tx)
	dyn4go.AssertEqualWithinError(t, 1.000, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.996, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.819, aabb.GetMaxY(), 1.0e-3)
}

/**
 * Verifies the output of the getRadius and getSliceRadius methods.
 */
func TestSliceRadius(t *testing.T) {
	e := NewSlice(1.0, dyn4go.DegToRad(50))
	dyn4go.AssertEqualWithinError(t, 1.000, e.GetSliceRadius(), 1.0e-3)
	dyn4go.AssertFalse(t, math.Abs(1.0-e.GetRadius()) < dyn4go.Epsilon)
	dyn4go.AssertFalse(t, math.Abs(e.GetSliceRadius()-e.GetRadius()) < dyn4go.Epsilon)
}
