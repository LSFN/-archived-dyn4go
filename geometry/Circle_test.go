package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests a zero radius.
 */
func TestCreateZeroRadius(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewCircle(0.0)
}

/**
 * Tests a negative radius.
 */
func TestCreateNegativeRadius(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewCircle(-1.0)
}

/**
 * Tests the constructor.
 */
func TestCreateSuccess(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	NewCircle(1.0)
}

/**
 * Tests the contains method.
 */
func TestContains(t *testing.T) {
	c := NewCircle(2.0)
	transform := NewTransform()
	p := NewVector2FromXY(2.0, 4.0)

	// shouldn't be in the circle
	dyn4go.AssertTrue(t, !c.ContainsTransform(p, transform))

	// move the circle a bit
	transform.TranslateXY(2.0, 2.5)

	// should be in the circle
	dyn4go.AssertTrue(t, c.ContainsTransform(p, transform))

	transform.TranslateXY(0.0, -0.5)

	// should be on the edge
	dyn4go.AssertTrue(t, c.ContainsTransform(p, transform))
}

/**
 * Tests the project method.
 */
func TestProject(t *testing.T) {
	c := NewCircle(1.5)
	transform := NewTransform()
	x := NewVector2FromXY(1.0, 0.0)
	y := NewVector2FromXY(0.0, 1.0)

	transform.TranslateXY(1.0, 0.5)

	i := c.ProjectTransform(x, transform)

	dyn4go.AssertEqualWithinError(t, -0.500, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.500, i.max, 1.0e-3)

	// rotating about the center shouldn't effect anything
	transform.RotateAboutXY(dyn4go.DegToRad(30), 1.0, 0.5)

	i = c.ProjectTransform(y, transform)

	dyn4go.AssertEqualWithinError(t, -1.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, i.max, 1.0e-3)
}

/**
 * Tests the farthest methods.
 */
func TestGetFarthest(t *testing.T) {
	c := NewCircle(1.5)
	transform := NewTransform()
	y := NewVector2FromXY(0.0, -1.0)

	f := c.GetFarthestFeature(y, transform)
	dyn4go.AssertTrue(t, f.IsVertex())
	dyn4go.AssertEqualWithinError(t, 0.000, f.point.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.500, f.point.Y, 1.0e-3)

	p := c.GetFarthestPoint(y, transform)
	dyn4go.AssertEqualWithinError(t, 0.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.500, p.Y, 1.0e-3)

	// move the circle a bit
	transform.TranslateXY(0.0, -0.5)

	f = c.GetFarthestFeature(y.GetNegative(), transform)
	dyn4go.AssertTrue(t, f.IsVertex())
	dyn4go.AssertEqualWithinError(t, 0.000, f.point.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, f.point.Y, 1.0e-3)

	p = c.GetFarthestPoint(y.GetNegative(), transform)
	dyn4go.AssertEqualWithinError(t, 0.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, p.Y, 1.0e-3)
}

/**
 * Tests the getAxes method.
 */
func TestGetAxes(t *testing.T) {
	c := NewCircle(1.5)
	transform := NewTransform()
	// a cicle has infinite axes so it should be nil
	axes := c.GetAxes(nil, transform)
	dyn4go.AssertTrue(t, axes == nil)
}

/**
 * Tests the getFoci method.
 */
func TestGetFoci(t *testing.T) {
	c := NewCircle(1.5)
	transform := NewTransform()
	// should only return one
	foci := c.GetFoci(transform)
	dyn4go.AssertEqual(t, 1, len(foci))
}

/**
 * Tests the rotate methods.
 */
func TestRotate(t *testing.T) {
	// center is at 0,0
	c := NewCircle(1.0)

	// rotate about center
	c.TranslateXY(1.0, 1.0)
	c.RotateAboutCenter(dyn4go.DegToRad(30))
	dyn4go.AssertEqualWithinError(t, 1.000, c.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, c.center.Y, 1.0e-3)

	// rotate about the origin
	c.RotateAboutOrigin(dyn4go.DegToRad(90))
	dyn4go.AssertEqualWithinError(t, -1.000, c.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, c.center.Y, 1.0e-3)
	c.TranslateVector2(c.GetCenter().GetNegative())

	// should move the center
	c.RotateAboutXY(dyn4go.DegToRad(90), 1.0, -1.0)
	dyn4go.AssertEqualWithinError(t, 0.000, c.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -2.000, c.center.Y, 1.0e-3)
}

/**
 * Tests the translate methods.
 */
func TestTranslate(t *testing.T) {
	// center is at 0,0
	c := NewCircle(1.0)

	c.TranslateXY(1.0, -0.5)

	dyn4go.AssertEqualWithinError(t, 1.000, c.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.500, c.center.Y, 1.0e-3)
}

/**
 * Tests the generated AABB.
 * @since 3.1.0
 */
func TestCreateAABB(t *testing.T) {
	c := NewCircle(1.2)

	// using an identity transform
	aabb := c.CreateAABBTransform(NewTransform())
	dyn4go.AssertEqualWithinError(t, -1.2, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.2, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.2, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.2, aabb.GetMaxY(), 1.0e-3)

	// try using the default method
	aabb2 := c.CreateAABB()
	dyn4go.AssertEqual(t, aabb.GetMinX(), aabb2.GetMinX())
	dyn4go.AssertEqual(t, aabb.GetMinY(), aabb2.GetMinY())
	dyn4go.AssertEqual(t, aabb.GetMaxX(), aabb2.GetMaxX())
	dyn4go.AssertEqual(t, aabb.GetMaxY(), aabb2.GetMaxY())

	// test using a rotation and translation matrix
	tx := NewTransform()
	tx.RotateAboutOrigin(dyn4go.DegToRad(30.0))
	tx.TranslateXY(1.0, 2.0)

	aabb = c.CreateAABBTransform(tx)
	dyn4go.AssertEqualWithinError(t, -0.2, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.8, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.2, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.2, aabb.GetMaxY(), 1.0e-3)
}
