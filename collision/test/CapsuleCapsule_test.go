package test

import (
	"github.com/LSFN/dyn4go"
	"github.com/LSFN/dyn4go/collision/broadphase"
	"github.com/LSFN/dyn4go/collision/manifold"
	"github.com/LSFN/dyn4go/collision/narrowphase"
	"github.com/LSFN/dyn4go/geometry"
	"testing"
)

/**
 * Sets up the test.
 */
type CapsuleCapsuleTest struct {
	AbstractTest
	capsule1, capsule2 *geometry.Capsule
}

func NewCapsuleCapsuleTest() *CapsuleCapsuleTest {
	this := new(CapsuleCapsuleTest)
	InitAbastractTest(&this.AbstractTest)
	this.capsule1 = geometry.NewCapsule(0.5, 1.0)
	this.capsule2 = geometry.NewCapsule(1.0, 0.5)
	this.sapI.Clear()
	this.sapBF.Clear()
	this.sapT.Clear()
	this.dynT.Clear()
	return this
}

/**
 * Tests {@link Shape} AABB.
 */

func TestDetectShapeAABB(t *testing.T) {
	this := NewCapsuleCapsuleTest()
	t1 := geometry.NewTransform()
	t2 := geometry.NewTransform()

	// test containment
	dyn4go.AssertTrue(t, this.aabb.DetectConvexTransform(this.capsule1, t1, this.capsule2, t2))
	dyn4go.AssertTrue(t, this.aabb.DetectConvexTransform(this.capsule2, t2, this.capsule1, t1))

	// test overlap
	t1.TranslateXY(-0.5, 0.0)
	dyn4go.AssertTrue(t, this.aabb.DetectConvexTransform(this.capsule1, t1, this.capsule2, t2))
	dyn4go.AssertTrue(t, this.aabb.DetectConvexTransform(this.capsule2, t2, this.capsule1, t1))

	// test only AABB overlap
	t2.TranslateXY(0.0, 0.7)
	dyn4go.AssertTrue(t, this.aabb.DetectConvexTransform(this.capsule1, t1, this.capsule2, t2))
	dyn4go.AssertTrue(t, this.aabb.DetectConvexTransform(this.capsule2, t2, this.capsule1, t1))

	// test no overlap
	t2.TranslateXY(1.0, 0.0)
	dyn4go.AssertFalse(t, this.aabb.DetectConvexTransform(this.capsule1, t1, this.capsule2, t2))
	dyn4go.AssertFalse(t, this.aabb.DetectConvexTransform(this.capsule2, t2, this.capsule1, t1))
}

/**
 * Tests {@link Collidable} AABB.
 */

func TestDetectCollidableAABB(t *testing.T) {
	this := NewCapsuleCapsuleTest()
	// create some collidables
	ct1 := NewCollidableTestShape(this.capsule1)
	ct2 := NewCollidableTestShape(this.capsule2)

	// test containment
	dyn4go.AssertTrue(t, this.aabb.DetectColliders(ct1, ct2))
	dyn4go.AssertTrue(t, this.aabb.DetectColliders(ct2, ct1))

	// test overlap
	ct1.TranslateXY(-0.5, 0.0)
	dyn4go.AssertTrue(t, this.aabb.DetectColliders(ct1, ct2))
	dyn4go.AssertTrue(t, this.aabb.DetectColliders(ct2, ct1))

	// test only AABB overlap
	ct2.TranslateXY(0.0, 0.7)
	dyn4go.AssertTrue(t, this.aabb.DetectColliders(ct1, ct2))
	dyn4go.AssertTrue(t, this.aabb.DetectColliders(ct2, ct1))

	// test no overlap
	ct2.TranslateXY(1.0, 0.0)
	dyn4go.AssertFalse(t, this.aabb.DetectColliders(ct1, ct2))
	dyn4go.AssertFalse(t, this.aabb.DetectColliders(ct2, ct1))
}

/**
 * Tests {@link SapIncremental}.
 */

func TestDetectBroadphase(t *testing.T) {
	this := NewCapsuleCapsuleTest()
	pairs := make([]*broadphase.BroadphasePair, 0)

	// create some collidables
	ct1 := NewCollidableTestShape(this.capsule1)
	ct2 := NewCollidableTestShape(this.capsule2)

	this.sapI.Add(ct1)
	this.sapI.Add(ct2)
	this.sapBF.Add(ct1)
	this.sapBF.Add(ct2)
	this.sapT.Add(ct1)
	this.sapT.Add(ct2)
	this.dynT.Add(ct1)
	this.dynT.Add(ct2)

	// test containment
	pairs = this.sapI.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapBF.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.dynT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))

	// test overlap
	ct1.TranslateXY(-0.5, 0.0)
	this.sapI.Update(ct1)
	this.sapBF.Update(ct1)
	this.sapT.Update(ct1)
	this.dynT.Update(ct1)
	pairs = this.sapI.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapBF.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.dynT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))

	// test only AABB overlap
	ct2.TranslateXY(0.0, 0.7)
	this.sapI.Update(ct2)
	this.sapBF.Update(ct2)
	this.sapT.Update(ct2)
	this.dynT.Update(ct2)
	pairs = this.sapI.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapBF.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.sapT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))
	pairs = this.dynT.Detect()
	dyn4go.AssertEqual(t, 1, len(pairs))

	// test no overlap
	ct2.TranslateXY(1.0, 0.0)
	this.sapI.Update(ct2)
	this.sapBF.Update(ct2)
	this.sapT.Update(ct2)
	this.dynT.Update(ct2)
	pairs = this.sapI.Detect()
	dyn4go.AssertEqual(t, 0, len(pairs))
	pairs = this.sapBF.Detect()
	dyn4go.AssertEqual(t, 0, len(pairs))
	pairs = this.sapT.Detect()
	dyn4go.AssertEqual(t, 0, len(pairs))
	pairs = this.dynT.Detect()
	dyn4go.AssertEqual(t, 0, len(pairs))
}

/**
 * Tests that sat is unsupported.
 */

func TestDetectSat(t *testing.T) {
	this := NewCapsuleCapsuleTest()
	p := narrowphase.NewPenetration()
	t1 := geometry.NewTransform()
	t2 := geometry.NewTransform()

	// test containment
	dyn4go.AssertTrue(t, this.sat.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p))
	dyn4go.AssertTrue(t, this.sat.Detect(this.capsule1, t1, this.capsule2, t2))
	n := p.GetNormal()
	dyn4go.AssertEqualWithinError(t, 0.000, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.750, p.GetDepth(), 1.0e-3)
	// try reversing the shapes
	dyn4go.AssertTrue(t, this.sat.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p))
	dyn4go.AssertTrue(t, this.sat.Detect(this.capsule2, t2, this.capsule1, t1))
	n = p.GetNormal()
	dyn4go.AssertEqualWithinError(t, 1.000, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.750, p.GetDepth(), 1.0e-3)

	// test overlap
	t1.TranslateXY(-0.5, 0.0)
	dyn4go.AssertTrue(t, this.sat.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p))
	dyn4go.AssertTrue(t, this.sat.Detect(this.capsule1, t1, this.capsule2, t2))
	n = p.GetNormal()
	dyn4go.AssertEqualWithinError(t, 1.000, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, p.GetDepth(), 1.0e-3)
	// try reversing the shapes
	dyn4go.AssertTrue(t, this.sat.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p))
	dyn4go.AssertTrue(t, this.sat.Detect(this.capsule2, t2, this.capsule1, t1))
	n = p.GetNormal()
	dyn4go.AssertEqualWithinError(t, -1.000, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, p.GetDepth(), 1.0e-3)

	// test AABB overlap
	t2.TranslateXY(0.0, 0.7)
	dyn4go.AssertFalse(t, this.sat.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p))
	dyn4go.AssertFalse(t, this.sat.Detect(this.capsule1, t1, this.capsule2, t2))
	// try reversing the shapes
	dyn4go.AssertFalse(t, this.sat.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p))
	dyn4go.AssertFalse(t, this.sat.Detect(this.capsule2, t2, this.capsule1, t1))

	// test no overlap
	t2.TranslateXY(1.0, 0.0)
	dyn4go.AssertFalse(t, this.sat.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p))
	dyn4go.AssertFalse(t, this.sat.Detect(this.capsule1, t1, this.capsule2, t2))
	// try reversing the shapes
	dyn4go.AssertFalse(t, this.sat.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p))
	dyn4go.AssertFalse(t, this.sat.Detect(this.capsule2, t2, this.capsule1, t1))
}

/**
 * Tests {@link Gjk}.
 */

func TestDetectGjk(t *testing.T) {
	this := NewCapsuleCapsuleTest()
	p := narrowphase.NewPenetration()
	t1 := geometry.NewTransform()
	t2 := geometry.NewTransform()

	// test containment
	dyn4go.AssertTrue(t, this.gjk.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p))
	dyn4go.AssertTrue(t, this.gjk.Detect(this.capsule1, t1, this.capsule2, t2))
	n := p.GetNormal()
	dyn4go.AssertEqualWithinError(t, -1.000, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.750, p.GetDepth(), 1.0e-3)
	// try reversing the shapes
	dyn4go.AssertTrue(t, this.gjk.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p))
	dyn4go.AssertTrue(t, this.gjk.Detect(this.capsule2, t2, this.capsule1, t1))
	n = p.GetNormal()
	dyn4go.AssertEqualWithinError(t, 1.000, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.750, p.GetDepth(), 1.0e-3)

	// test overlap
	t1.TranslateXY(-0.5, 0.0)
	dyn4go.AssertTrue(t, this.gjk.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p))
	dyn4go.AssertTrue(t, this.gjk.Detect(this.capsule1, t1, this.capsule2, t2))
	n = p.GetNormal()
	dyn4go.AssertEqualWithinError(t, 1.000, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, p.GetDepth(), 1.0e-3)
	// try reversing the shapes
	dyn4go.AssertTrue(t, this.gjk.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p))
	dyn4go.AssertTrue(t, this.gjk.Detect(this.capsule2, t2, this.capsule1, t1))
	n = p.GetNormal()
	dyn4go.AssertEqualWithinError(t, -1.000, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, p.GetDepth(), 1.0e-3)

	// test AABB overlap
	t2.TranslateXY(0.0, 0.7)
	dyn4go.AssertFalse(t, this.gjk.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p))
	dyn4go.AssertFalse(t, this.gjk.Detect(this.capsule1, t1, this.capsule2, t2))
	// try reversing the shapes
	dyn4go.AssertFalse(t, this.gjk.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p))
	dyn4go.AssertFalse(t, this.gjk.Detect(this.capsule2, t2, this.capsule1, t1))

	// test no overlap
	t2.TranslateXY(1.0, 0.0)
	dyn4go.AssertFalse(t, this.gjk.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p))
	dyn4go.AssertFalse(t, this.gjk.Detect(this.capsule1, t1, this.capsule2, t2))
	// try reversing the shapes
	dyn4go.AssertFalse(t, this.gjk.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p))
	dyn4go.AssertFalse(t, this.gjk.Detect(this.capsule2, t2, this.capsule1, t1))
}

/**
 * Tests the {@link Gjk} distance method.
 */

func TestGjkDistance(t *testing.T) {
	this := NewCapsuleCapsuleTest()
	s := narrowphase.NewSeparation()

	t1 := geometry.NewTransform()
	t2 := geometry.NewTransform()

	// test containment
	dyn4go.AssertFalse(t, this.gjk.Distance(this.capsule1, t1, this.capsule2, t2, s))
	// try reversing the shapes
	dyn4go.AssertFalse(t, this.gjk.Distance(this.capsule2, t2, this.capsule1, t1, s))

	// test overlap
	t1.TranslateXY(-0.5, 0.0)
	dyn4go.AssertFalse(t, this.gjk.Distance(this.capsule1, t1, this.capsule2, t2, s))
	// try reversing the shapes
	dyn4go.AssertFalse(t, this.gjk.Distance(this.capsule2, t2, this.capsule1, t1, s))

	// test AABB overlap
	t2.TranslateXY(0.0, 0.7)
	dyn4go.AssertTrue(t, this.gjk.Distance(this.capsule1, t1, this.capsule2, t2, s))
	n := s.GetNormal()
	p1 := s.GetPoint1()
	p2 := s.GetPoint2()
	dyn4go.AssertEqualWithinError(t, 0.014, s.GetDistance(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.485, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.874, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.378, p1.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.468, p1.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.371, p2.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.481, p2.Y, 1.0e-3)
	// try reversing the shapes
	dyn4go.AssertTrue(t, this.gjk.Distance(this.capsule2, t2, this.capsule1, t1, s))
	n = s.GetNormal()
	p1 = s.GetPoint1()
	p2 = s.GetPoint2()
	dyn4go.AssertEqualWithinError(t, 0.014, s.GetDistance(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.485, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.874, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.371, p1.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.481, p1.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.378, p2.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.468, p2.Y, 1.0e-3)

	// test no overlap
	t2.TranslateXY(1.0, 0.0)
	dyn4go.AssertTrue(t, this.gjk.Distance(this.capsule1, t1, this.capsule2, t2, s))
	n = s.GetNormal()
	p1 = s.GetPoint1()
	p2 = s.GetPoint2()
	dyn4go.AssertEqualWithinError(t, 0.828, s.GetDistance(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.940, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.338, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.264, p1.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.334, p1.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.514, p2.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.615, p2.Y, 1.0e-3)
	// try reversing the shapes
	dyn4go.AssertTrue(t, this.gjk.Distance(this.capsule2, t2, this.capsule1, t1, s))
	n = s.GetNormal()
	p1 = s.GetPoint1()
	p2 = s.GetPoint2()
	dyn4go.AssertEqualWithinError(t, 0.828, s.GetDistance(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.940, n.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.338, n.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.514, p1.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.615, p1.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.264, p2.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.334, p2.Y, 1.0e-3)
}

/**
 * Test the {@link ClippingManifoldSolver}.
 */

func TestGetClipManifold(t *testing.T) {
	this := NewCapsuleCapsuleTest()
	m := manifold.NewManifold()
	p := narrowphase.NewPenetration()

	t1 := geometry.NewTransform()
	t2 := geometry.NewTransform()

	// test containment gjk
	this.gjk.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p)
	dyn4go.AssertTrue(t, this.cmfs.GetManifold(p, this.capsule1, t1, this.capsule2, t2, m))
	dyn4go.AssertEqual(t, 1, len(m.GetPoints()))
	// try reversing the shapes
	this.gjk.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p)
	dyn4go.AssertTrue(t, this.cmfs.GetManifold(p, this.capsule2, t2, this.capsule1, t1, m))
	dyn4go.AssertEqual(t, 1, len(m.GetPoints()))

	// test containment sat
	this.sat.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p)
	dyn4go.AssertTrue(t, this.cmfs.GetManifold(p, this.capsule1, t1, this.capsule2, t2, m))
	dyn4go.AssertEqual(t, 1, len(m.GetPoints()))
	// try reversing the shapes
	this.sat.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p)
	dyn4go.AssertTrue(t, this.cmfs.GetManifold(p, this.capsule2, t2, this.capsule1, t1, m))
	dyn4go.AssertEqual(t, 1, len(m.GetPoints()))

	t1.TranslateXY(-0.5, 0.0)

	// test overlap gjk
	this.gjk.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p)
	dyn4go.AssertTrue(t, this.cmfs.GetManifold(p, this.capsule1, t1, this.capsule2, t2, m))
	dyn4go.AssertEqual(t, 1, len(m.GetPoints()))
	mp := m.GetPoints()[0]
	p1 := mp.GetPoint()
	dyn4go.AssertEqualWithinError(t, -0.500, p1.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p1.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, mp.GetDepth(), 1.0e-3)
	// try reversing the shapes
	this.gjk.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p)
	dyn4go.AssertTrue(t, this.cmfs.GetManifold(p, this.capsule2, t2, this.capsule1, t1, m))
	dyn4go.AssertEqual(t, 1, len(m.GetPoints()))
	mp = m.GetPoints()[0]
	p1 = mp.GetPoint()
	dyn4go.AssertEqualWithinError(t, -0.500, p1.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p1.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, mp.GetDepth(), 1.0e-3)

	// test overlap sat
	this.sat.DetectPenetration(this.capsule1, t1, this.capsule2, t2, p)
	dyn4go.AssertTrue(t, this.cmfs.GetManifold(p, this.capsule1, t1, this.capsule2, t2, m))
	dyn4go.AssertEqual(t, 1, len(m.GetPoints()))
	mp = m.GetPoints()[0]
	p1 = mp.GetPoint()
	dyn4go.AssertEqualWithinError(t, -0.500, p1.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p1.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, mp.GetDepth(), 1.0e-3)
	// try reversing the shapes
	this.sat.DetectPenetration(this.capsule2, t2, this.capsule1, t1, p)
	dyn4go.AssertTrue(t, this.cmfs.GetManifold(p, this.capsule2, t2, this.capsule1, t1, m))
	dyn4go.AssertEqual(t, 1, len(m.GetPoints()))
	mp = m.GetPoints()[0]
	p1 = mp.GetPoint()
	dyn4go.AssertEqualWithinError(t, -0.500, p1.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p1.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, mp.GetDepth(), 1.0e-3)
}
