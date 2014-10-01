package test

import (
	"github.com/LSFN/dyn4go/collision/broadphase"
	"github.com/LSFN/dyn4go/collision/manifold"
	"github.com/LSFN/dyn4go/collision/narrowphase"
	"github.com/LSFN/dyn4go/geometry"
)

type AbstractTest struct {
	aabb  *AbstractTestAABBDetector
	sat   *narrowphase.SAT
	gjk   *narrowphase.GJK
	sapI  *broadphase.SapIncremental
	sapBF *broadphase.SapBruteForce
	sapT  *broadphase.SapTree
	dynT  *broadphase.DynamicAABBTree
	cmfs  *manifold.ClippingManifoldSolver
}

func InitAbastractTest(a *AbstractTest) {
	a.aabb = NewAbstractTestAABBDetector()
	a.sat = new(narrowphase.SAT)
	a.gjk = narrowphase.NewGJK()
	a.sapI = broadphase.NewSapIncremental()
	a.sapBF = broadphase.NewSapBruteForce()
	a.sapT = broadphase.NewSapTree()
	a.dynT = broadphase.NewDynamicAABBTree()
	a.cmfs = new(manifold.ClippingManifoldSolver)
}

type AbstractTestAABBDetector struct {
	broadphase.AbstractAABBDetector
}

func NewAbstractTestAABBDetector() *AbstractTestAABBDetector {
	a := new(AbstractTestAABBDetector)
	broadphase.InitAbstractAABBDetector(&a.AbstractAABBDetector)
	return a
}

func (a *AbstractTestAABBDetector) Add(collidable *CollidableTest) {

}

func (a *AbstractTestAABBDetector) Detect() []*broadphase.BroadphasePair {
	return nil
}

func (a *AbstractTestAABBDetector) Remove(collidable *CollidableTest) {
}

func (a *AbstractTestAABBDetector) Update(collidable *CollidableTest) {
}

func (a *AbstractTestAABBDetector) DetectAABB(aabb *geometry.AABB) []*CollidableTest {
	return nil
}

func (a *AbstractTestAABBDetector) Clear() {
}

func (a *AbstractTestAABBDetector) GetAABB(collidable *CollidableTest) *geometry.AABB {
	return nil
}

func (a *AbstractTestAABBDetector) Raycast(ray *geometry.Ray, length float64) []*CollidableTest {
	return nil
}

func (a *AbstractTestAABBDetector) ShiftCoordinates(shift *geometry.Vector2) {

}
