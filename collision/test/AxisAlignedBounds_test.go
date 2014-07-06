package test

import (
	"github.com/LSFN/dyn4go"
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
	"testing"
)

/**
 * Tests the width and height getters.
 */

func TestGetWidthAndHeight(t *testing.T) {
	ab := collision.NewAxisAlignedBounds(10.0, 7.0)
	dyn4go.AssertEqual(t, 10.0, ab.GetWidth())
	dyn4go.AssertEqual(t, 7.0, ab.GetHeight())
}

/**
 * Tests the getTranslation method.
 */

func TestGetTranslation(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)
	bounds.TranslateXY(1.0, -2.0)
	tx := bounds.GetTranslation()
	dyn4go.AssertEqual(t, 1.0, tx.X)
	dyn4go.AssertEqual(t, -2.0, tx.Y)
}

/**
 * Tests the getBounds method.
 */

func TestGetBounds(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)
	aabb := bounds.GetBounds()
	// should be centered about the origin
	dyn4go.AssertEqual(t, -10.0, aabb.GetMinX())
	dyn4go.AssertEqual(t, -10.0, aabb.GetMinY())
	dyn4go.AssertEqual(t, 10.0, aabb.GetMaxX())
	dyn4go.AssertEqual(t, 10.0, aabb.GetMaxY())

	// move it a bit
	bounds.TranslateXY(1.0, -2.0)
	aabb = bounds.GetBounds()
	dyn4go.AssertEqual(t, -9.0, aabb.GetMinX())
	dyn4go.AssertEqual(t, -12.0, aabb.GetMinY())
	dyn4go.AssertEqual(t, 11.0, aabb.GetMaxX())
	dyn4go.AssertEqual(t, 8.0, aabb.GetMaxY())
}

/**
 * Verifies the rotate methods do not modify the internal
 * structure of the bounds.
 */

func TestRotateNoOp(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)
	// perform some rotations
	bounds.RotateAboutOrigin(dyn4go.DegToRad(30.0))
	bounds.RotateAboutXY(dyn4go.DegToRad(-15.0), 3.0, -4.0)
	bounds.RotateAboutVector2(dyn4go.DegToRad(7.5), geometry.NewVector2FromXY(1.0, 0.0))
	// verify that the bounds are left unchanged
	aabb := bounds.GetBounds()
	// should be centered about the origin
	dyn4go.AssertEqual(t, -10.0, aabb.GetMinX())
	dyn4go.AssertEqual(t, -10.0, aabb.GetMinY())
	dyn4go.AssertEqual(t, 10.0, aabb.GetMaxX())
	dyn4go.AssertEqual(t, 10.0, aabb.GetMaxY())
}

/**
 * Tests creating a {@link AxisAlignedBounds} with invalid bounds.
 */
func TestCreateInvalidBounds1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	collision.NewAxisAlignedBounds(0, 1)
}

/**
 * Tests creating a {@link AxisAlignedBounds} with invalid bounds.
 */
func TestCreateInvalidBounds2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	collision.NewAxisAlignedBounds(1, 0)
}

/**
 * Tests creating a {@link AxisAlignedBounds} with invalid bounds.
 */
func TestCreateInvalidBounds3(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	collision.NewAxisAlignedBounds(1, -1)
}

/**
 * Tests creating a {@link AxisAlignedBounds} with invalid bounds.
 */
func TestCreateInvalidBounds4(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	collision.NewAxisAlignedBounds(-1, 1)
}

/**
 * Tests the isOutside method on a {@link Circle}.
 */

func TestIsOutsideCircle(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)
	// create some shapes
	c := geometry.NewCircle(1.0)
	ct := NewCollidableTestShape(c)

	// should be in
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test half way in and out
	ct.transform.TranslateXY(9.5, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(1.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-1.5, 9.5)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test moving the bounds
	bounds.TranslateXY(2.0, 1.0)

	// test half way in and out
	ct.transform.TranslateXY(2.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(1.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-1.5, 1.5)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))
}

/**
 * Tests the isOutside method on a {@link Ellipse}.
 */

func TestIsOutsideEllipse(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)
	// create some shapes
	c := geometry.NewEllipse(1.0, 0.5)
	ct := NewCollidableTestShape(c)

	// should be in
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test half way in and out
	ct.transform.TranslateXY(9.5, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(1.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-1.5, 9.5)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test moving the bounds
	bounds.TranslateXY(2.0, 1.0)

	// test half way in and out
	ct.transform.TranslateXY(2.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(1.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-1.5, 1.5)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))
}

/**
 * Tests the isOutside method on a {@link Rectangle}.
 */

func TestIsOutsideRectangle(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)

	// create some shapes
	r := geometry.NewRectangle(1.0, 1.0)
	ct := NewCollidableTestShape(r)

	// should be in
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test half way in and out
	ct.transform.TranslateXY(10.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(0.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-0.6, 10.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test moving the bounds
	bounds.TranslateXY(2.0, 1.0)

	// test half way in and out
	ct.transform.TranslateXY(2.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(1.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-1.5, 1.5)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))
}

/**
 * Tests the isOutside method on a {@link Polygon}.
 */

func TestIsOutsidePolygon(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)

	// create some shapes
	p := geometry.CreateUnitCirclePolygon(6, 0.5)
	ct := NewCollidableTestShape(p)

	// should be in
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test half way in and out
	ct.transform.TranslateXY(10.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(0.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-0.6, 10.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test moving the bounds
	bounds.TranslateXY(2.0, 1.0)

	// test half way in and out
	ct.transform.TranslateXY(2.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(1.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-1.5, 1.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))
}

/**
 * Tests the isOutside method on a {@link Triangle}.
 */

func TestIsOutsideTriangle(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)

	// create some shapes
	tr := geometry.NewTriangle(
		geometry.NewVector2FromXY(0.0, 0.5),
		geometry.NewVector2FromXY(-0.5, -0.5),
		geometry.NewVector2FromXY(0.5, -0.5),
	)
	ct := NewCollidableTestShape(tr)

	// should be in
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test half way in and out
	ct.transform.TranslateXY(10.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(0.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-0.6, 10.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test moving the bounds
	bounds.TranslateXY(2.0, 1.0)

	// test half way in and out
	ct.transform.TranslateXY(2.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(1.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-1.5, 1.5)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))
}

/**
 * Tests the isOutside method on a {@link Segment}.
 */

func TestIsOutsideSegment(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)

	// create some shapes
	s := geometry.NewSegment(geometry.NewVector2FromXY(0.5, -0.5), geometry.NewVector2FromXY(-0.5, 0.5))
	ct := NewCollidableTestShape(s)

	// should be in
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test half way in and out
	ct.transform.TranslateXY(10.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(0.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-0.6, 10.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test moving the bounds
	bounds.TranslateXY(2.0, 1.0)

	// test half way in and out
	ct.transform.TranslateXY(2.0, 0.0)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))

	// test all the way out
	ct.transform.TranslateXY(1.6, 0.0)
	dyn4go.AssertTrue(t, bounds.IsOutside(ct))

	// test half way out a corner
	ct.transform.TranslateXY(-1.5, 1.5)
	dyn4go.AssertFalse(t, bounds.IsOutside(ct))
}

/**
 * Tests shifting the coordinates of the bounds.
 */

func TestShiftCoordinates(t *testing.T) {
	bounds := collision.NewAxisAlignedBounds(20, 20)

	tx := bounds.GetTransform().GetTranslation()
	dyn4go.AssertEqualWithinError(t, 0.000, tx.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, tx.Y, 1.0e-3)

	// test the shifting which is really just a translation
	bounds.ShiftCoordinates(geometry.NewVector2FromXY(1.0, 1.0))
	tx = bounds.GetTransform().GetTranslation()
	dyn4go.AssertEqualWithinError(t, 1.000, tx.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, tx.Y, 1.0e-3)
}
