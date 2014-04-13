package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

func TestCapsuleInterfaces(t *testing.T) {
	c := NewCapsule(2, 1)
	var _ Convexer = c
	var _ Shaper = c
}

/**
 * Tests a zero width.
 */
func TestCapsuleCreateZeroWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewCapsule(0.0, 1.0)
}

/**
 * Tests a negative width.
 */
func TestCapsuleCreateNegativeWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewCapsule(-1.0, 1.0)
}

/**
 * Tests a zero height.
 */
func TestCapsuleCreateZeroHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewCapsule(1.0, 0.0)
}

/**
 * Tests a negative height.
 */
func TestCapsuleCreateNegativeHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewCapsule(1.0, -1.0)
}

/**
 * Tests the constructor.
 */
func TestCapsuleCreateSuccessHorizontal(t *testing.T) {
	cap := NewCapsule(2.0, 1.0)
	x := cap.localXAxis
	dyn4go.AssertEqualWithinError(t, 1.000, x.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, x.Y, 1.0e-3)
}

/**
 * Tests the constructor.
 */
func TestCapsuleCreateSuccessVertical(t *testing.T) {
	cap := NewCapsule(1.0, 2.0)
	x := cap.localXAxis
	dyn4go.AssertEqualWithinError(t, 0.000, x.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, x.Y, 1.0e-3)
}

/**
 * Tests the contains method.
 */
func TestCapsuleContains(t *testing.T) {
	e := NewCapsule(2.0, 1.0)
	transform := NewTransform()
	p := NewVector2FromXY(0.8, -0.45)

	// shouldn't be inside
	dyn4go.AssertTrue(t, !e.ContainsVector2Transform(p, transform))

	// move it a bit
	transform.TranslateXY(0.5, 0.0)

	// should be inside
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))

	p.SetToXY(1.5, 0.0)
	// should be on the edge
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))
	p.SetToXY(0.75, 0.5)
	// should be on the edge
	dyn4go.AssertTrue(t, e.ContainsVector2Transform(p, transform))
}

/**
 * Tests the project method.
 */
func TestCapsuleProject(t *testing.T) {
	e := NewCapsule(2.0, 1.0)
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
	dyn4go.AssertEqualWithinError(t, -1.25, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.25, i.max, 1.0e-3)

	// try some local rotation
	e.TranslateXY(1.0, 0.5)
	e.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	i = e.ProjectVector2Transform(y, NewTransform())
	dyn4go.AssertEqualWithinError(t, -1.25, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.25, i.max, 1.0e-3)

	transform.Identity()
	transform.TranslateXY(0.0, 1.0)
	i = e.ProjectVector2Transform(y, transform)
	dyn4go.AssertEqualWithinError(t, -2.25, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.75, i.max, 1.0e-3)
}

/**
 * Tests the farthest methods.
 */
func TestCapsuleGetFarthest(t *testing.T) {
	e := NewCapsule(2.0, 1.0)
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
	dyn4go.AssertEqualWithinError(t, 0.566, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.25, p.Y, 1.0e-3)

	// try some local rotation
	e.TranslateXY(1.0, 0.5)
	e.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	p = e.GetFarthestPoint(y, NewTransform())
	dyn4go.AssertEqualWithinError(t, 0.566, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.25, p.Y, 1.0e-3)

	transform.Identity()
	transform.TranslateXY(0.0, 1.0)
	p = e.GetFarthestPoint(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.566, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.75, p.Y, 1.0e-3)
}

/**
 * Tests the getAxes method.
 */
func TestCapsuleGetAxes(t *testing.T) {
	e := NewCapsule(1.0, 0.5)
	// should be two axes + number of foci
	foci := []*Vector2{
		NewVector2FromXY(2.0, -0.5),
		NewVector2FromXY(1.0, 3.0),
	}
	axes := e.GetAxes(foci, NewTransform())
	dyn4go.AssertEqual(t, 4, len(axes))

	// make sure we get back the right axes
	axes = e.GetAxes(nil, NewTransform())
	dyn4go.AssertEqualWithinError(t, 1.000, axes[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, axes[0].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, axes[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, axes[1].Y, 1.0e-3)
}

/**
 * Tests the getFoci method.
 */
func TestCapsuleGetFoci(t *testing.T) {
	e := NewCapsule(1.0, 0.5)
	foci := e.GetFoci(NewTransform())
	// should be two foci
	dyn4go.AssertEqual(t, 2, len(foci))
	// make sure the foci are correct
	dyn4go.AssertEqualWithinError(t, -0.250, foci[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, foci[0].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, foci[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, foci[1].Y, 1.0e-3)
}

/**
 * Tests the rotate methods.
 */
func TestCapsuleRotate(t *testing.T) {
	e := NewCapsule(1.0, 0.5)

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
func TestCapsuleTranslate(t *testing.T) {
	e := NewCapsule(1.0, 0.5)

	e.TranslateXY(1.0, -0.5)

	dyn4go.AssertEqualWithinError(t, 1.000, e.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.500, e.center.Y, 1.0e-3)
}

/**
 * Tests the generated AABB.
 */
func TestCapsuleCreateAABB(t *testing.T) {
	e := NewCapsule(1.0, 0.5)

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
	dyn4go.AssertEqualWithinError(t, 0.533, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.625, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.466, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.375, aabb.GetMaxY(), 1.0e-3)
}
