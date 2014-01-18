package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

func TestCreateSuccess(t *testing.T) {
	NewAABBFromFloats(0.0, 0.0, 1.0, 1.0)
	NewAABBFromFloats(-2.0, 2.0, -1.0, 5.0)
	NewAABBFromVector2(NewVector2FromXY(-3.0, 0.0), NewVector2FromXY(-2.0, 2.0))
}

func TestCreateRadius(t *testing.T) {
	aabb := NewAABBFromRadius(0.5)
	dyn4go.AssertEqualWithinError(t, -0.500, aabb.getMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.500, aabb.getMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, aabb.getMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, aabb.getMaxY(), 1.0e-3)

	aabb = NewAABBFromCenterRadius(NewVector2FromXY(-1.0, 1.0), 0.5)
	dyn4go.AssertEqualWithinError(t, -1.500, aabb.getMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, aabb.getMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.500, aabb.getMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.500, aabb.getMaxY(), 1.0e-3)
}

func TestCreateCopy(t *testing.T) {
	aabb1 := NewAABBFromVector2(NewVector2FromXY(-3.0, 0.0), NewVector2FromXY(-2.0, 2.0))
	aabb2 := NewAABBFromAABB(aabb1)
	dyn4go.AssertNotEqual(t, aabb1, aabb2)
	dyn4go.AssertEqualWithinError(t, aabb1.getMinX(), aabb2.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, aabb1.getMinY(), aabb2.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, aabb1.getMaxX(), aabb2.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, aabb1.getMaxY(), aabb2.getMaxY(), 1.0E-4)
}

func TestCreateFailure1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewAABBFromFloats(0.0, 0.0, -1.0, 2.0)
}

func TestCreateFailure2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewAABBFromVector2(NewVector2FromXY(0.0, 0.0), NewVector2FromXY(-1.0, 2.0))
}

func TestPerimeter(t *testing.T) {
	aabb := NewAABBFromFloats(-2.0, 0.0, 2.0, 1.0)
	// 4 + 1 = 5
	dyn4go.AssertEqualWithinError(t, 10.0, aabb.GetPerimeter(), 1.0E-4)
}

func TestArea(t *testing.T) {
	aabb := NewAABBFromFloats(-2.0, 0.0, 2.0, 1.0)
	// 4
	dyn4go.AssertEqualWithinError(t, 4.0, aabb.GetArea(), 1.0E-4)
}

func TestUnion(t *testing.T) {
	// overlapping AABBs
	aabb1 := NewAABBFromFloats(-2.0, 0.0, 2.0, 1.0)
	aabb2 := NewAABBFromFloats(-1.0, -2.0, 5.0, 0.5)

	// test the GetUnion method
	aabbr := aabb1.GetUnion(aabb2)
	dyn4go.AssertEqualWithinError(t, -2.0, aabbr.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, -2.0, aabbr.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 5.0, aabbr.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 1.0, aabbr.getMaxY(), 1.0E-4)

	// test the GetUnion method using separated aabbs
	aabb3 := NewAABBFromFloats(-4.0, 2.0, -3.0, 4.0)
	aabbr = aabb1.GetUnion(aabb3)
	dyn4go.AssertEqualWithinError(t, -4.0, aabbr.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabbr.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabbr.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 4.0, aabbr.getMaxY(), 1.0E-4)

	aabb1.Union(aabb2)
	dyn4go.AssertEqualWithinError(t, -2.0, aabb1.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, -2.0, aabb1.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 5.0, aabb1.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb1.getMaxY(), 1.0E-4)
}

func TestExpand(t *testing.T) {
	aabb := NewAABBFromFloats(-2.0, 0.0, 4.0, 4.0)
	aabb2 := aabb.GetExpanded(2.0)

	dyn4go.AssertNotEqual(t, aabb, aabb2)

	aabb.Expand(1.0)
	dyn4go.AssertEqualWithinError(t, -2.5, aabb.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, -0.5, aabb.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 4.5, aabb.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 4.5, aabb.getMaxY(), 1.0E-4)

	// the second aabb will have different values
	dyn4go.AssertEqualWithinError(t, -3.0, aabb2.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, -1.0, aabb2.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 5.0, aabb2.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 5.0, aabb2.getMaxY(), 1.0E-4)

	// test negative expansion
	aabb2 = aabb.GetExpanded(-1.0)
	dyn4go.AssertEqualWithinError(t, -2.0, aabb2.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabb2.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 4.0, aabb2.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 4.0, aabb2.getMaxY(), 1.0E-4)
	aabb.Expand(-1.0)
	dyn4go.AssertEqualWithinError(t, -2.0, aabb.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabb.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 4.0, aabb.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 4.0, aabb.getMaxY(), 1.0E-4)

	// test an overly negative expansion
	aabb2 = aabb.GetExpanded(-8.0)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb2.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabb2.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb2.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabb2.getMaxY(), 1.0E-4)
	aabb.Expand(-8.0)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabb.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabb.getMaxY(), 1.0E-4)
}

func TestOverlaps(t *testing.T) {
	// test overlap
	aabb1 := NewAABBFromFloats(-2.0, 0.0, 2.0, 1.0)
	aabb2 := NewAABBFromFloats(-1.0, -2.0, 5.0, 2.0)
	dyn4go.AssertTrue(t, aabb1.Overlaps(aabb2))
	dyn4go.AssertTrue(t, aabb2.Overlaps(aabb1))

	// test no overlap
	aabb3 := NewAABBFromFloats(3.0, 2.0, 4.0, 3.0)
	dyn4go.AssertFalse(t, aabb1.Overlaps(aabb3))
	dyn4go.AssertFalse(t, aabb3.Overlaps(aabb1))

	// test containment
	aabb4 := NewAABBFromFloats(-1.0, 0.25, 1.0, 0.75)
	dyn4go.AssertTrue(t, aabb1.Overlaps(aabb4))
	dyn4go.AssertTrue(t, aabb4.Overlaps(aabb1))
}

func TestContains(t *testing.T) {
	// test overlap
	aabb1 := NewAABBFromFloats(-2.0, 0.0, 2.0, 1.0)
	aabb2 := NewAABBFromFloats(-1.0, -2.0, 5.0, 2.0)
	dyn4go.AssertFalse(t, aabb1.ContainsAABB(aabb2))
	dyn4go.AssertFalse(t, aabb2.ContainsAABB(aabb1))

	// test no overlap
	aabb3 := NewAABBFromFloats(3.0, 2.0, 4.0, 3.0)
	dyn4go.AssertFalse(t, aabb1.ContainsAABB(aabb3))
	dyn4go.AssertFalse(t, aabb3.ContainsAABB(aabb1))

	// test containment
	aabb4 := NewAABBFromFloats(-1.0, 0.25, 1.0, 0.75)
	dyn4go.AssertTrue(t, aabb1.ContainsAABB(aabb4))
	dyn4go.AssertFalse(t, aabb4.ContainsAABB(aabb1))
}

func TestGetWidth(t *testing.T) {
	aabb := NewAABBFromFloats(-2.0, 0.0, 1.0, 1.0)

	dyn4go.AssertEqual(t, 3.0, aabb.GetWidth())
}

func TestGetHeight(t *testing.T) {
	aabb := NewAABBFromFloats(-2.0, 0.0, 1.0, 1.0)

	dyn4go.AssertEqual(t, 1.0, aabb.GetHeight())
}

func TestTranslate(t *testing.T) {
	aabb := NewAABBFromFloats(-2.0, 0.0, 1.0, 1.0)

	aabb2 := aabb.GetTranslated(NewVector2FromXY(-1.0, 2.0))
	dyn4go.AssertNotEqual(t, aabb, aabb2)

	dyn4go.AssertEqualWithinError(t, -2.0, aabb.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabb.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.getMaxY(), 1.0E-4)

	dyn4go.AssertEqualWithinError(t, -3.0, aabb2.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabb2.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabb2.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 3.0, aabb2.getMaxY(), 1.0E-4)

	aabb.Translate(NewVector2FromXY(-1.0, 2.0))

	dyn4go.AssertEqualWithinError(t, -3.0, aabb.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabb.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabb.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 3.0, aabb.getMaxY(), 1.0E-4)
}

func TestContainsPoint(t *testing.T) {
	aabb := NewAABBFromFloats(-2.0, 0.0, 2.0, 1.0)

	// test containment
	dyn4go.AssertTrue(t, aabb.ContainsXY(0.0, 0.5))

	// test no containment
	dyn4go.AssertFalse(t, aabb.ContainsXY(0.0, 2.0))

	// test on edge
	dyn4go.AssertTrue(t, aabb.ContainsXY(0.0, 1.0))
}

func TestIntersection(t *testing.T) {
	aabb1 := NewAABBFromFloats(-2.0, 0.0, 2.0, 1.0)
	aabb2 := NewAABBFromFloats(-1.0, -2.0, 5.0, 0.5)

	aabbr := aabb1.GetIntersection(aabb2)
	dyn4go.AssertEqualWithinError(t, -1.0, aabbr.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabbr.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabbr.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.5, aabbr.getMaxY(), 1.0E-4)

	// test using separated aabbs (should give a zero AABB)
	aabb3 := NewAABBFromFloats(-4.0, 2.0, -3.0, 4.0)
	aabbr = aabb1.GetIntersection(aabb3)
	dyn4go.AssertEqualWithinError(t, 0.0, aabbr.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabbr.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabbr.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabbr.getMaxY(), 1.0E-4)

	aabb1.Intersection(aabb2)
	dyn4go.AssertEqualWithinError(t, -1.0, aabb1.getMinX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.0, aabb1.getMinY(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 2.0, aabb1.getMaxX(), 1.0E-4)
	dyn4go.AssertEqualWithinError(t, 0.5, aabb1.getMaxY(), 1.0E-4)
}

func TestDegenerate(t *testing.T) {
	aabb := NewAABBFromFloats(0.0, 0.0, 0.0, 0.0)
	dyn4go.AssertTrue(t, aabb.IsDegenerate())

	aabb = NewAABBFromFloats(1.0, 2.0, 1.0, 3.0)
	dyn4go.AssertTrue(t, aabb.IsDegenerate())

	aabb = NewAABBFromFloats(1.0, 0.0, 2.0, 1.0)
	dyn4go.AssertFalse(t, aabb.IsDegenerate())

	aabb = NewAABBFromFloats(1.0, 0.0, 1.000001, 2.0)
	dyn4go.AssertFalse(t, aabb.IsDegenerate())
	dyn4go.AssertFalse(t, aabb.IsDegenerateWithError(dyn4go.Epsilon))
	dyn4go.AssertTrue(t, aabb.IsDegenerateWithError(0.000001))
}
