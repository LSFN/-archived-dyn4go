package test

import (
	"github.com/LSFN/dyn4go"
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/collision/broadphase"
	"github.com/LSFN/dyn4go/geometry"
	"math"
	"testing"
)

type BroadphaseTest struct {
	sapI  *broadphase.SapIncremental
	sapBF *broadphase.SapBruteForce
	sapT  *broadphase.SapTree
	dynT  *broadphase.DynamicAABBTree
}

func NewBroadphaseTest() *BroadphaseTest {
	b := new(BroadphaseTest)
	b.sapI = broadphase.NewSapIncremental()
	b.sapBF = broadphase.NewSapBruteForce()
	b.sapT = broadphase.NewSapTree()
	b.dynT = broadphase.NewDynamicAABBTree()
	return b
}

func ColliderSliceContains(container []collision.Collider, containee collision.Collider) bool {
	for _, v := range container {
		if v == containee {
			return true
		}
	}
	return false
}

/**
 * Tests the add method.
 */

func TestAdd(t *testing.T) {
	this := NewBroadphaseTest()
	ct := NewCollidableTestShape(geometry.CreateCircle(1.0))

	// make sure its not there first
	dyn4go.AssertTrue(t, this.sapI.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.sapBF.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.sapT.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.dynT.GetAABB(ct) == nil)

	// add the item to the broadphases
	this.sapI.Add(ct)
	this.sapBF.Add(ct)
	this.sapT.Add(ct)
	this.dynT.Add(ct)

	// make sure they are there
	dyn4go.AssertFalse(t, this.sapI.GetAABB(ct) == nil)
	dyn4go.AssertFalse(t, this.sapBF.GetAABB(ct) == nil)
	dyn4go.AssertFalse(t, this.sapT.GetAABB(ct) == nil)
	dyn4go.AssertFalse(t, this.dynT.GetAABB(ct) == nil)
}

/**
 * Tests the remove method.
 */

func TestRemove(t *testing.T) {
	this := NewBroadphaseTest()
	ct := NewCollidableTestShape(geometry.CreateCircle(1.0))

	// add the item to the broadphases
	this.sapI.Add(ct)
	this.sapBF.Add(ct)
	this.sapT.Add(ct)
	this.dynT.Add(ct)

	// make sure they are there
	dyn4go.AssertFalse(t, this.sapI.GetAABB(ct) == nil)
	dyn4go.AssertFalse(t, this.sapBF.GetAABB(ct) == nil)
	dyn4go.AssertFalse(t, this.sapT.GetAABB(ct) == nil)
	dyn4go.AssertFalse(t, this.dynT.GetAABB(ct) == nil)

	// then remove them from the broadphases
	this.sapI.Remove(ct)
	this.sapBF.Remove(ct)
	this.sapT.Remove(ct)
	this.dynT.Remove(ct)

	// make sure they aren't there any more
	dyn4go.AssertTrue(t, this.sapI.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.sapBF.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.sapT.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.dynT.GetAABB(ct) == nil)
}

/**
 * Tests the update method where the collidable moves very little.
 */

func TestUpdateSmall(t *testing.T) {
	this := NewBroadphaseTest()
	ct := NewCollidableTestShape(geometry.CreateCircle(1.0))

	// add the item to the broadphases
	this.sapI.Add(ct)
	this.sapBF.Add(ct)
	this.sapT.Add(ct)
	this.dynT.Add(ct)

	// make sure they are there
	aabbSapI := this.sapI.GetAABB(ct)
	aabbSapBF := this.sapBF.GetAABB(ct)
	aabbSapT := this.sapT.GetAABB(ct)
	aabbDynT := this.dynT.GetAABB(ct)
	dyn4go.AssertFalse(t, aabbSapI == nil)
	dyn4go.AssertFalse(t, aabbSapBF == nil)
	dyn4go.AssertFalse(t, aabbSapT == nil)
	dyn4go.AssertFalse(t, aabbDynT == nil)

	// move the collidable a bit
	ct.TranslateXY(0.05, 0.0)

	// update the broadphases
	this.sapI.Update(ct)
	this.sapBF.Update(ct)
	this.sapT.Update(ct)
	this.dynT.Update(ct)

	// the aabbs should not have been updated because of the expansion code
	dyn4go.AssertEqual(t, aabbSapI, this.sapI.GetAABB(ct))
	dyn4go.AssertEqual(t, aabbSapBF, this.sapBF.GetAABB(ct))
	dyn4go.AssertEqual(t, aabbSapT, this.sapT.GetAABB(ct))
	dyn4go.AssertEqual(t, aabbDynT, this.dynT.GetAABB(ct))
}

/**
 * Tests the update method where the collidable moves enough to update the AABB.
 */

func TestUpdateLarge(t *testing.T) {
	this := NewBroadphaseTest()
	ct := NewCollidableTestShape(geometry.CreateCircle(1.0))

	// add the item to the broadphases
	this.sapI.Add(ct)
	this.sapBF.Add(ct)
	this.sapT.Add(ct)
	this.dynT.Add(ct)

	// make sure they are there
	aabbSapI := this.sapI.GetAABB(ct)
	aabbSapBF := this.sapBF.GetAABB(ct)
	aabbSapT := this.sapT.GetAABB(ct)
	aabbDynT := this.dynT.GetAABB(ct)
	dyn4go.AssertFalse(t, aabbSapI == nil)
	dyn4go.AssertFalse(t, aabbSapBF == nil)
	dyn4go.AssertFalse(t, aabbSapT == nil)
	dyn4go.AssertFalse(t, aabbDynT == nil)

	// move the collidable a bit
	ct.TranslateXY(0.5, 0.0)

	// update the broadphases
	this.sapI.Update(ct)
	this.sapBF.Update(ct)
	this.sapT.Update(ct)
	this.dynT.Update(ct)

	// the aabbs should not have been updated because of the expansion code
	dyn4go.AssertNotEqual(t, aabbSapI, this.sapI.GetAABB(ct))
	dyn4go.AssertNotEqual(t, aabbSapBF, this.sapBF.GetAABB(ct))
	dyn4go.AssertNotEqual(t, aabbSapT, this.sapT.GetAABB(ct))
	dyn4go.AssertNotEqual(t, aabbDynT, this.dynT.GetAABB(ct))
}

/**
 * Tests the clear method.
 */

func TestClear(t *testing.T) {
	this := NewBroadphaseTest()
	ct := NewCollidableTestShape(geometry.CreateCircle(1.0))

	// add the item to the broadphases
	this.sapI.Add(ct)
	this.sapBF.Add(ct)
	this.sapT.Add(ct)
	this.dynT.Add(ct)

	// clear all the broadphases
	this.sapI.Clear()
	this.sapBF.Clear()
	this.sapT.Clear()
	this.dynT.Clear()

	// check for the aabb
	dyn4go.AssertTrue(t, this.sapI.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.sapBF.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.sapT.GetAABB(ct) == nil)
	dyn4go.AssertTrue(t, this.dynT.GetAABB(ct) == nil)
}

/**
 * Tests the getAABB method.
 */

func TestGetAABB(t *testing.T) {
	this := NewBroadphaseTest()
	ct := NewCollidableTestShape(geometry.CreateCircle(1.0))

	// add the item to the broadphases
	this.sapI.Add(ct)
	this.sapBF.Add(ct)
	this.sapT.Add(ct)
	this.dynT.Add(ct)

	// make sure they are there
	aabbSapI := this.sapI.GetAABB(ct)
	aabbSapBF := this.sapBF.GetAABB(ct)
	aabbSapT := this.sapT.GetAABB(ct)
	aabbDynT := this.dynT.GetAABB(ct)

	aabb := ct.CreateAABB()
	// don't forget that the aabb is expanded
	aabb.Expand(broadphase.DEFAULT_AABB_EXPANSION)
	dyn4go.AssertTrue(t, isEqualAABBAABB(aabbSapI, aabb))
	dyn4go.AssertTrue(t, isEqualAABBAABB(aabbSapBF, aabb))
	dyn4go.AssertTrue(t, isEqualAABBAABB(aabbSapT, aabb))
	dyn4go.AssertTrue(t, isEqualAABBAABB(aabbDynT, aabb))
}

/**
 * Helper method for the getAABB test method.
 * @param aabb1 the first aabb
 * @param aabb2 the second aabb
 * @return boolean true if they are basically the same
 */
func isEqualAABBAABB(aabb1, aabb2 *geometry.AABB) bool {
	return !((math.Abs(aabb1.GetMinX()-aabb2.GetMinX()) >= 1.0E-8) ||
		(math.Abs(aabb1.GetMinY()-aabb2.GetMinY()) >= 1.0E-8) ||
		(math.Abs(aabb1.GetMaxX()-aabb2.GetMaxX()) >= 1.0E-8) ||
		(math.Abs(aabb1.GetMaxY()-aabb2.GetMaxY()) >= 1.0E-8))
}

/**
 * Tests the {@link AbstractAABBDetector} detect methods.
 * @since 3.1.0
 */

func TestDetectAbstract(t *testing.T) {
	this := NewBroadphaseTest()
	ct1 := NewCollidableTestShape(geometry.CreateCircle(1.0))
	ct2 := NewCollidableTestShape(geometry.CreateUnitCirclePolygon(5, 0.5))
	ct1.TranslateXY(-2.0, 0.0)
	ct2.TranslateXY(-1.0, 1.0)

	dyn4go.AssertTrue(t, this.dynT.DetectColliders(ct1, ct2))
	dyn4go.AssertTrue(t, this.dynT.DetectConvexTransform(ct1.GetFixture(0).GetShape(), ct1.transform, ct2.GetFixture(0).GetShape(), ct2.transform))

	ct1.TranslateXY(-1.0, 0.0)
	dyn4go.AssertFalse(t, this.dynT.DetectColliders(ct1, ct2))
	dyn4go.AssertFalse(t, this.dynT.DetectConvexTransform(ct1.GetFixture(0).GetShape(), ct1.transform, ct2.GetFixture(0).GetShape(), ct2.transform))
}

/**
 * Tests the detect method.
 * @since 3.1.0
 */

func TestDetect(t *testing.T) {
	this := NewBroadphaseTest()
	ct1 := NewCollidableTestShape(geometry.CreateCircle(1.0))
	ct2 := NewCollidableTestShape(geometry.CreateUnitCirclePolygon(5, 0.5))
	ct3 := NewCollidableTestShape(geometry.CreateRectangle(1.0, 0.5))
	ct4 := NewCollidableTestShape(geometry.CreateVerticalSegment(2.0))

	ct1.TranslateXY(-2.0, 0.0)
	ct2.TranslateXY(-1.0, 1.0)
	ct3.TranslateXY(0.5, -2.0)
	ct4.TranslateXY(1.0, 1.0)

	// add the items to the broadphases
	this.sapI.Add(ct1)
	this.sapI.Add(ct2)
	this.sapI.Add(ct3)
	this.sapI.Add(ct4)

	this.sapBF.Add(ct1)
	this.sapBF.Add(ct2)
	this.sapBF.Add(ct3)
	this.sapBF.Add(ct4)

	this.sapT.Add(ct1)
	this.sapT.Add(ct2)
	this.sapT.Add(ct3)
	this.sapT.Add(ct4)

	this.dynT.Add(ct1)
	this.dynT.Add(ct2)
	this.dynT.Add(ct3)
	this.dynT.Add(ct4)

	pairs := this.sapI.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapBF.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.dynT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
}

/**
 * Tests the detect method using an AABB.
 */

func TestDetectAABB(t *testing.T) {
	this := NewBroadphaseTest()
	ct1 := NewCollidableTestShape(geometry.CreateCircle(1.0))
	ct2 := NewCollidableTestShape(geometry.CreateUnitCirclePolygon(5, 0.5))
	ct3 := NewCollidableTestShape(geometry.CreateRectangle(1.0, 0.5))
	ct4 := NewCollidableTestShape(geometry.CreateVerticalSegment(2.0))

	ct1.TranslateXY(-2.0, 0.0)
	ct2.TranslateXY(-1.0, 1.0)
	ct3.TranslateXY(0.5, -2.0)
	ct4.TranslateXY(1.0, 1.0)

	// add the items to the broadphases
	this.sapI.Add(ct1)
	this.sapI.Add(ct2)
	this.sapI.Add(ct3)
	this.sapI.Add(ct4)

	this.sapBF.Add(ct1)
	this.sapBF.Add(ct2)
	this.sapBF.Add(ct3)
	this.sapBF.Add(ct4)

	this.sapT.Add(ct1)
	this.sapT.Add(ct2)
	this.sapT.Add(ct3)
	this.sapT.Add(ct4)

	this.dynT.Add(ct1)
	this.dynT.Add(ct2)
	this.dynT.Add(ct3)
	this.dynT.Add(ct4)

	// this aabb should include:
	// ct3 and ct4
	aabb := geometry.NewAABBFromFloats(0.0, -2.0, 1.0, 1.0)

	list := this.sapI.DetectAABB(aabb)
	dyn4go.AssertEqual(t, 2, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct3))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.sapBF.DetectAABB(aabb)
	dyn4go.AssertEqual(t, 2, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct3))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.sapT.DetectAABB(aabb)
	dyn4go.AssertEqual(t, 2, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct3))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.dynT.DetectAABB(aabb)
	dyn4go.AssertEqual(t, 2, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct3))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))

	// should include:
	// ct2, ct3, and ct4
	aabb = geometry.NewAABBFromFloats(-0.75, -3.0, 2.0, 1.0)

	list = this.sapI.DetectAABB(aabb)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct3))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.sapBF.DetectAABB(aabb)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct3))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.sapT.DetectAABB(aabb)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct3))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.dynT.DetectAABB(aabb)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct3))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
}

/**
 * Tests the raycast method.
 */

func TestRaycast(t *testing.T) {
	this := NewBroadphaseTest()
	ct1 := NewCollidableTestShape(geometry.CreateCircle(1.0))
	ct2 := NewCollidableTestShape(geometry.CreateUnitCirclePolygon(5, 0.5))
	ct3 := NewCollidableTestShape(geometry.CreateRectangle(1.0, 0.5))
	ct4 := NewCollidableTestShape(geometry.CreateVerticalSegment(2.0))

	ct1.TranslateXY(-2.0, 0.0)
	ct2.TranslateXY(-1.0, 1.0)
	ct3.TranslateXY(0.5, -2.0)
	ct4.TranslateXY(1.0, 1.2)

	// add the items to the broadphases
	this.sapI.Add(ct1)
	this.sapI.Add(ct2)
	this.sapI.Add(ct3)
	this.sapI.Add(ct4)

	this.sapBF.Add(ct1)
	this.sapBF.Add(ct2)
	this.sapBF.Add(ct3)
	this.sapBF.Add(ct4)

	this.sapT.Add(ct1)
	this.sapT.Add(ct2)
	this.sapT.Add(ct3)
	this.sapT.Add(ct4)

	this.dynT.Add(ct1)
	this.dynT.Add(ct2)
	this.dynT.Add(ct3)
	this.dynT.Add(ct4)

	// ray that points in the positive x direction and starts at the origin
	r := geometry.NewRayFromVector2(geometry.NewVector2FromXY(1.0, 0.0))
	// infinite length
	l := 0.0

	list := this.sapI.Raycast(r, l)
	dyn4go.AssertEqual(t, 0, len(list))
	list = this.sapBF.Raycast(r, l)
	dyn4go.AssertEqual(t, 0, len(list))
	list = this.sapT.Raycast(r, l)
	dyn4go.AssertEqual(t, 0, len(list))
	list = this.dynT.Raycast(r, l)
	dyn4go.AssertEqual(t, 0, len(list))

	// try a different ray
	r = geometry.NewRayFromVector2Vector2(geometry.NewVector2FromXY(-3.0, 0.75), geometry.NewVector2FromXY(1.0, 0.0))
	list = this.sapI.Raycast(r, l)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct1))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.sapBF.Raycast(r, l)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct1))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.sapT.Raycast(r, l)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct1))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.dynT.Raycast(r, l)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct1))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))

	// try one more ray
	r = geometry.NewRayFromVector2Vector2(geometry.NewVector2FromXY(-1.0, -1.0), geometry.NewVector2FromXY(0.85, 0.35))
	list = this.sapI.Raycast(r, l)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct1))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.sapBF.Raycast(r, l)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct1))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.sapT.Raycast(r, l)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct1))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
	list = this.dynT.Raycast(r, l)
	dyn4go.AssertEqual(t, 3, len(list))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct1))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct2))
	dyn4go.AssertTrue(t, ColliderSliceContains(list, ct4))
}

/**
 * Tests the get/set expansion methods.
 */

func TestExpansion(t *testing.T) {
	this := NewBroadphaseTest()
	// test the default
	dyn4go.AssertEqual(t, broadphase.DEFAULT_AABB_EXPANSION, this.sapI.GetAABBExpansion())
	dyn4go.AssertEqual(t, broadphase.DEFAULT_AABB_EXPANSION, this.sapBF.GetAABBExpansion())
	dyn4go.AssertEqual(t, broadphase.DEFAULT_AABB_EXPANSION, this.sapT.GetAABBExpansion())
	dyn4go.AssertEqual(t, broadphase.DEFAULT_AABB_EXPANSION, this.dynT.GetAABBExpansion())

	// test changing the expansion
	this.sapI.SetAABBExpansion(0.3)
	this.sapBF.SetAABBExpansion(0.3)
	this.sapT.SetAABBExpansion(0.3)
	this.dynT.SetAABBExpansion(0.3)
	dyn4go.AssertEqual(t, 0.3, this.sapI.GetAABBExpansion())
	dyn4go.AssertEqual(t, 0.3, this.sapBF.GetAABBExpansion())
	dyn4go.AssertEqual(t, 0.3, this.sapT.GetAABBExpansion())
	dyn4go.AssertEqual(t, 0.3, this.dynT.GetAABBExpansion())

	// test the new expansion value
	ct := NewCollidableTestShape(geometry.CreateCircle(1.0))

	// add the item to the broadphases
	this.sapI.Add(ct)
	this.sapBF.Add(ct)
	this.sapT.Add(ct)
	this.dynT.Add(ct)

	// make sure they are there
	aabbSapI := this.sapI.GetAABB(ct)
	aabbSapBF := this.sapBF.GetAABB(ct)
	aabbSapT := this.sapT.GetAABB(ct)
	aabbDynT := this.dynT.GetAABB(ct)

	aabb := ct.CreateAABB()
	// don't forget that the aabb is expanded
	aabb.Expand(0.3)
	dyn4go.AssertTrue(t, isEqualAABBAABB(aabbSapI, aabb))
	dyn4go.AssertTrue(t, isEqualAABBAABB(aabbSapBF, aabb))
	dyn4go.AssertTrue(t, isEqualAABBAABB(aabbSapT, aabb))
	dyn4go.AssertTrue(t, isEqualAABBAABB(aabbDynT, aabb))
}

/**
 * Tests the shiftCoordinates method.
 */

func TestShiftCoordinates2(t *testing.T) {
	this := NewBroadphaseTest()
	ct1 := NewCollidableTestShape(geometry.CreateCircle(1.0))
	ct2 := NewCollidableTestShape(geometry.CreateUnitCirclePolygon(5, 0.5))
	ct3 := NewCollidableTestShape(geometry.CreateRectangle(1.0, 0.5))
	ct4 := NewCollidableTestShape(geometry.CreateVerticalSegment(2.0))

	ct1.TranslateXY(-2.0, 0.0)
	ct2.TranslateXY(-1.0, 1.0)
	ct3.TranslateXY(0.5, -2.0)
	ct4.TranslateXY(1.0, 1.0)

	// add the items to the broadphases
	this.sapI.Add(ct1)
	this.sapI.Add(ct2)
	this.sapI.Add(ct3)
	this.sapI.Add(ct4)

	this.sapBF.Add(ct1)
	this.sapBF.Add(ct2)
	this.sapBF.Add(ct3)
	this.sapBF.Add(ct4)

	this.sapT.Add(ct1)
	this.sapT.Add(ct2)
	this.sapT.Add(ct3)
	this.sapT.Add(ct4)

	this.dynT.Add(ct1)
	this.dynT.Add(ct2)
	this.dynT.Add(ct3)
	this.dynT.Add(ct4)

	// perform a detect on the whole broadphase
	pairs := this.sapI.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapBF.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.dynT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))

	// shift the broadphases
	shift := geometry.NewVector2FromXY(1.0, -2.0)
	this.sapI.ShiftCoordinates(shift)
	this.sapBF.ShiftCoordinates(shift)
	this.sapT.ShiftCoordinates(shift)
	this.dynT.ShiftCoordinates(shift)

	// the number of pairs detected should be identical
	pairs = this.sapI.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapBF.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.dynT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
}

/**
 * Tests creating a SapBruteForce detector using a negative capacity.
 */
func TestSapBruteForceNegativeInitialCapacity(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	broadphase.NewSapBruteForceInt(-10)
}

/**
 * Tests creating a SapIncremental detector using a negative capacity.
 */
func TestSapIncrementalNegativeInitialCapacity(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	broadphase.NewSapIncrementalInt(-10)
}

/**
 * Tests creating a SapTree detector using a negative capacity.
 */
func TestSapTreeNegativeInitialCapacity(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	broadphase.NewSapTreeInt(-10)
}

/**
 * Tests creating a DynamicAABBTree detector using a negative capacity.
 */
func TestDynamicAABBTreeNegativeInitialCapacity(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	broadphase.NewDynamicAABBTreeInt(-10)
}
