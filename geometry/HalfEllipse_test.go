package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

func TestHalfEllipseInterfaces(t *testing.T) {
	h := NewHalfEllipse(2, 1)
	var _ Convexer = h
	var _ Shaper = h
}

/**
 * Tests a zero width.
 */
func TestHalfEllipseCreateZeroWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewHalfEllipse(0.0, 1.0)
}

/**
 * Tests a negative width.
 */
func TestHalfEllipseCreateNegativeWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewHalfEllipse(-1.0, 1.0)
}

/**
 * Tests a zero height.
 */
func TestHalfEllipseCreateZeroHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewHalfEllipse(1.0, 0.0)
}

/**
 * Tests a negative height.
 */
func TestHalfEllipseCreateNegativeHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewHalfEllipse(1.0, -1.0)
}

/**
 * Tests the constructor.
 */
func TestHalfEllipseCreateSuccess(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	NewHalfEllipse(1.0, 2.0)
}

/**
 * Tests the contains method.
 */
func TestHalfEllipseContains(t *testing.T) {
	e := NewHalfEllipse(2.0, 0.5)
	transform := NewTransform()
	p := NewVector2FromXY(0.75, 0.35)

	// shouldn't be in
	dyn4go.AssertTrue(t, !e.ContainsVector2Transform(p, transform))

	p.SetToXY(0.75, -0.2)
	dyn4go.AssertTrue(t, !e.ContainsVector2Transform(p, transform))

	// move a bit
	p.SetToXY(0.75, 0.35)
	transform.TranslateXY(0.5, 0.0)

	// should be in
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))

	p.SetToXY(1.5, 0.0)

	// should be on the edge
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))

	// another test for failure case
	p.SetToXY(0.75, 0.35)
	e.TranslateVector2(e.GetCenter().GetNegative())
	dyn4go.AssertFalse(t, e.ContainsVector2Transform(p, transform))

	// try local rotation and translation
	e.RotateAboutOrigin(dyn4go.RadToDeg(90))
	e.TranslateXY(0.5, 1.0)

	p.SetToXY(0.3, 0.3)
	dyn4go.AssertFalse(t, e.ContainsVector2Transform(p, transform))

	p.SetToXY(0.7, 0.4)
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))
}

/**
 * Tests the project method.
 */
func TestHalfEllipseProject(t *testing.T) {
	e := NewHalfEllipse(2.0, 0.5)
	transform := NewTransform()
	x := NewVector2FromXY(1.0, 0.0)
	y := NewVector2FromXY(0.0, 1.0)

	// try some translation
	transform.TranslateXY(1.0, 0.5)

	i := e.ProjectVector2Transform(x, transform)
	dyn4go.AssertEqualWithinError(t, 0.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, i.max, 1.0e-3)

	i = e.ProjectVector2Transform(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.500, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, i.max, 1.0e-3)

	// try some rotation
	transform.RotateAboutVector2(dyn4go.DegToRad(30), transform.GetTransformedVector2(e.GetCenter()))

	i = e.ProjectVector2Transform(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.028, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.189, i.max, 1.0e-3)

	// try some local rotation
	e.TranslateXY(1.0, 0.5)
	e.RotateAboutCenter(dyn4go.DegToRad(30))

	i = e.ProjectVector2Transform(y, NewTransform())
	dyn4go.AssertEqualWithinError(t, 0.028, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.189, i.max, 1.0e-3)

	transform.Identity()
	transform.TranslateXY(0.0, 1.0)
	i = e.ProjectVector2Transform(y, transform)
	dyn4go.AssertEqualWithinError(t, 1.028, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.189, i.max, 1.0e-3)
}

/**
 * Tests the farthest methods.
 */
func TestHalfEllipseGetFarthest(t *testing.T) {
	e := NewHalfEllipse(2.0, 0.5)
	transform := NewTransform()
	x := NewVector2FromXY(1.0, 0.0)
	y := NewVector2FromXY(0.0, -1.0)

	// try some translation
	transform.TranslateXY(1.0, 0.5)

	p := e.GetFarthestPoint(x, transform)
	dyn4go.AssertEqualWithinError(t, 2.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.5, p.Y, 1.0e-3)

	// try some rotation
	transform.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	p = e.GetFarthestPoint(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.133, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.Y, 1.0e-3)

	// try some local rotation
	e.TranslateXY(1.0, 0.5)
	e.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	p = e.GetFarthestPoint(y, NewTransform())
	dyn4go.AssertEqualWithinError(t, 0.133, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.Y, 1.0e-3)

	transform.Identity()
	transform.TranslateXY(0.0, 1.0)
	p = e.GetFarthestPoint(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.133, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, p.Y, 1.0e-3)
}

/**
 * Tests the getAxes method.
 */
func TestHalfEllipseGetAxes(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	e := NewHalfEllipse(1.0, 0.5)
	e.GetAxes([]*Vector2{new(Vector2)}, NewTransform())
}

/**
 * Tests the getFoci method.
 */
func TestHalfEllipseGetFoci(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	e := NewHalfEllipse(1.0, 0.5)
	e.GetFoci(NewTransform())
}

/**
 * Tests the rotate methods.
 */
func TestHalfEllipseRotate(t *testing.T) {
	e := NewHalfEllipse(1.0, 0.25)

	// rotate about center
	e.TranslateXY(1.0, 1.0)
	e.RotateAboutCenter(dyn4go.DegToRad(30))
	dyn4go.AssertEqualWithinError(t, 1.000, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.106, e.center.Y, 1.0e-3)

	// rotate about the origin
	e.RotateAboutOrigin(dyn4go.DegToRad(90))
	dyn4go.AssertEqualWithinError(t, -1.106, e.center.X, 1.0e-3)
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
func TestHalfEllipseTranslate(t *testing.T) {
	e := NewHalfEllipse(1.0, 0.25)

	e.TranslateXY(1.0, -0.5)

	dyn4go.AssertEqualWithinError(t, 1.000, e.GetEllipseCenter().X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.500, e.GetEllipseCenter().Y, 1.0e-3)
}

/**
 * Tests the generated AABB.
 */
func TestHalfEllipseCreateAABB(t *testing.T) {
	e := NewHalfEllipse(1.0, 0.25)

	// using an identity transform
	aabb := e.CreateAABBTransform(NewTransform())
	dyn4go.AssertEqualWithinError(t, -0.500, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, aabb.GetMinY(), 1.0e-3)
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
	dyn4go.AssertEqualWithinError(t, 1.750, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.433, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.330, aabb.GetMaxY(), 1.0e-3)
}
