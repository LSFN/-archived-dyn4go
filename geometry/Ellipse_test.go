package geometry2

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

func TestEllipseInterfaces(t *testing.T) {
	e := NewEllipse(2, 1)
	var _ Convexer = e
	var _ Shaper = e
}

/**
 * Tests a zero width.
 */
func TestEllipseCreateZeroWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewEllipse(0.0, 1.0)
}

/**
 * Tests a negative width.
 */
func TestEllipseCreateNegativeWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewEllipse(-1.0, 1.0)
}

/**
 * Tests a zero height.
 */
func TestEllipseCreateZeroHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewEllipse(1.0, 0.0)
}

/**
 * Tests a negative height.
 */
func TestEllipseCreateNegativeHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewEllipse(1.0, -1.0)
}

/**
 * Tests the constructor.
 */
func TestEllipseCreateSuccess(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	e := NewEllipse(1.0, 2.0)
	dyn4go.AssertEqual(t, 1.0, e.GetHalfHeight())
	dyn4go.AssertEqual(t, 0.5, e.GetHalfWidth())
	dyn4go.AssertEqual(t, 1.0, e.GetWidth())
	dyn4go.AssertEqual(t, 2.0, e.GetHeight())
}

/**
 * Tests the contains method.
 */
func TestEllipseContains(t *testing.T) {
	e := NewEllipse(2.0, 1.0)
	transform := NewTransform()
	p := NewVector2FromXY(0.75, 0.35)

	// shouldn't be in the circle
	dyn4go.AssertTrue(t, !e.ContainsVector2Transform(p, transform))

	// move the circle a bit
	transform.TranslateXY(0.5, 0.0)

	// should be in the circle
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))

	p.SetToXY(1.5, 0.0)

	// should be on the edge
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))

	// test with local translation
	e.RotateAboutOrigin(dyn4go.DegToRad(90))
	e.TranslateXY(0.5, 1.0)

	dyn4go.AssertFalse(t, e.ContainsVector2Transform(p, transform))
	p.SetToXY(1.0, 2.1)
	dyn4go.AssertFalse(t, e.ContainsVector2Transform(p, transform))
	p.SetToXY(1.0, 2.0)
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))
}

/**
 * Tests the project method.
 */
func TestEllipseProject(t *testing.T) {
	e := NewEllipse(2.0, 1.0)
	transform := NewTransform()
	x := NewVector2FromXY(1.0, 0.0)
	y := NewVector2FromXY(0.0, -1.0)

	// try some translation
	transform.TranslateXY(1.0, 0.5)
	i := e.ProjectVector2Transform(x, transform)
	dyn4go.AssertEqualWithinError(t, 0.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, i.max, 1.0e-3)

	// try some rotation
	transform.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)
	i = e.ProjectVector2Transform(y, transform)
	dyn4go.AssertEqualWithinError(t, -1.161, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.161, i.max, 1.0e-3)

	// try some local rotation
	e.TranslateXY(1.0, 0.5)
	e.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)
	i = e.ProjectVector2Transform(y, NewTransform())
	dyn4go.AssertEqualWithinError(t, -1.161, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.161, i.max, 1.0e-3)

	transform.Identity()
	transform.TranslateXY(0.0, 1.0)
	i = e.ProjectVector2Transform(y, transform)
	dyn4go.AssertEqualWithinError(t, -2.161, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.839, i.max, 1.0e-3)
}

/**
 * Tests the farthest methods.
 */
func TestEllipseGetFarthest(t *testing.T) {
	e := NewEllipse(2.0, 1.0)
	transform := NewTransform()
	x := NewVector2FromXY(1.0, 0.0)
	y := NewVector2FromXY(0.0, -1.0)

	// try some translation
	transform.TranslateXY(1.0, 0.5)

	p := e.GetFarthestPoint(x, transform)
	dyn4go.AssertEqualWithinError(t, 2.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, p.Y, 1.0e-3)

	// try some rotation
	transform.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	p = e.GetFarthestPoint(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.509, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.161, p.Y, 1.0e-3)

	// try some local rotation
	e.TranslateXY(1.0, 0.5)
	e.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	p = e.GetFarthestPoint(y, NewTransform())
	dyn4go.AssertEqualWithinError(t, 0.509, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.161, p.Y, 1.0e-3)

	transform.Identity()
	transform.TranslateXY(0.0, 1.0)
	p = e.GetFarthestPoint(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.509, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.838, p.Y, 1.0e-3)
}

/**
 * Tests the getAxes method.
 */
func TestEllipseGetAxes(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	e := NewEllipse(1.0, 0.5)
	e.GetAxes([]*Vector2{new(Vector2)}, NewTransform())
}

/**
 * Tests the getFoci method.
 */
func TestEllipseGetFoci(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	e := NewEllipse(1.0, 0.5)
	e.GetFoci(NewTransform())
}

/**
 * Tests the rotate methods.
 */
func TestEllipseRotate(t *testing.T) {
	e := NewEllipse(1.0, 0.5)

	// rotate about center
	e.TranslateXY(1.0, 1.0)
	e.RotateAboutCenter(dyn4go.DegToRad(30))
	dyn4go.AssertEqualWithinError(t, 1.000, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, e.center.Y, 1.0e-3)

	// rotate about the origin
	e.RotateAboutOrigin(dyn4go.DegToRad(90))
	dyn4go.AssertEqualWithinError(t, -1.000, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, e.center.Y, 1.0e-3)
	e.TranslateVector2(e.GetCenter().GetNegative())

	// should move the center
	e.RotateAboutXY(dyn4go.DegToRad(90), 1.0, -1.0)
	dyn4go.AssertEqualWithinError(t, 0.000, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -2.000, e.center.Y, 1.0e-3)
}

/**
 * Tests the translate methods.
 */
func TestEllipseTranslate(t *testing.T) {
	e := NewEllipse(1.0, 0.5)

	e.TranslateXY(1.0, -0.5)

	dyn4go.AssertEqualWithinError(t, 1.000, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.500, e.center.Y, 1.0e-3)
}

/**
 * Tests the generated AABB.
 */
func TestEllipseCreateAABB(t *testing.T) {
	e := NewEllipse(1.0, 0.5)

	// using an identity transform
	aabb := e.CreateAABBTransform(NewTransform())
	dyn4go.AssertEqualWithinError(t, -0.500, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.250, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, aabb.GetMaxY(), 1.0e-3)

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
	dyn4go.AssertEqualWithinError(t, 0.549, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.669, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.450, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.330, aabb.GetMaxY(), 1.0e-3)
}
