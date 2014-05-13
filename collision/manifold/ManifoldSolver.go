package manifold

import (
	"github.com/LSFN/dyn4go/collision/narrowphase"
	"github.com/LSFN/dyn4go/geometry"
)

type ManifoldSolver interface {
	GetManifold(penetration *narrowphase.Penetration, convex1 *geometry.Convexer, transform1 *geometry.Transform, convex2 *geometry.Convexer, transform2 *geometry.Transform, manifold *Manifold) bool
}
