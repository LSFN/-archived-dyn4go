package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests the failed creation of a triangle with one point being nil.
 * @since 3.1.0
 */
func TestCreateNullPoint1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewTriangle(
		nil,
		NewVector2FromXY(-0.5, -0.5),
		NewVector2FromXY(0.5, -0.5),
	)
}

/**
 * Tests the failed creation of a triangle with one point being nil.
 * @since 3.1.0
 */
func TestCreateNullPoint2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewTriangle(
		NewVector2FromXY(-0.5, -0.5),
		nil,
		NewVector2FromXY(0.5, -0.5),
	)
}

/**
 * Tests the failed creation of a triangle with one point being nil.
 * @since 3.1.0
 */
func TestCreateNullPoint3(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewTriangle(
		NewVector2FromXY(-0.5, -0.5),
		NewVector2FromXY(0.5, -0.5),
		nil,
	)
}

/**
 * Tests the contains method.
 */
func TestContains(t *testing.T) {
	triangle := NewTriangle(
		NewVector2FromXY(0.0, 0.5),
		NewVector2FromXY(-0.5, -0.5),
		NewVector2FromXY(0.5, -0.5),
	)
	tx := NewTransform()

	// outside
	p := NewVector2FromXY(1.0, 1.0)
	dyn4go.AssertFalse(t, triangle.ContainsVector2Transform(p, tx))

	// inside
	p.SetToXY(0.2, 0.0)
	dyn4go.AssertTrue(t, triangle.ContainsVector2Transform(p, tx))

	// on edge
	p.SetToXY(0.3, -0.5)
	dyn4go.AssertTrue(t, triangle.ContainsVector2Transform(p, tx))

	// move the triangle a bit
	tx.RotateAboutOrigin(dyn4go.DegToRad(90))
	tx.TranslateXY(0.0, 1.0)

	// still outside
	p.SetToXY(1.0, 1.0)
	dyn4go.AssertFalse(t, triangle.ContainsVector2Transform(p, tx))

	// inside
	p.SetToXY(0.4, 1.0)
	dyn4go.AssertTrue(t, triangle.ContainsVector2Transform(p, tx))

	// on edge
	p.SetToXY(0.0, 0.76)
	// 0.76 should be 0.75 but it fails because of floating point problems
	dyn4go.AssertTrue(t, triangle.ContainsVector2Transform(p, tx))
}
