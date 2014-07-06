package test

import (
	"github.com/LSFN/dyn4go/collision/broadphase"
	"github.com/LSFN/dyn4go/collision/manifold"
	"github.com/LSFN/dyn4go/collision/narrowphase"
)

type AbstractTestAABBDetector struct {
	sat   *narrowphase.SAT
	gjk   *narrowphase.GJK
	sapI  *broadphase.SapIncremental
	sapBF *broadphase.SapBruteForce
	sapT  *broadphase.SapTree
	dynT  *broadphase.DynamicAABBTree
	cmfs  *manifold.ClippingManifoldSolver
}

func InitAbastractTestAABBDetector(a *AbstractTestAABBDetector) {
	a.sat = new(narrowphase.SAT)
	a.gjk = narrowphase.NewGJK()
	a.sapI = broadphase.NewSapIncremental()
	a.sapBF = broadphase.NewSapBruteForce()
	a.sapT = broadphase.NewSapTree()
	a.dynT = broadphase.NewDynamicAABBTree()
	a.cmfs = new(manifold.ClippingManifoldSolver)
}
