package broadphase

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
)

type AbstractAABBDetector struct {
	expansion float64
}

func InitAbstractAABBDetector(abstractAABBDetector *AbstractAABBDetector) {
	abstractAABBDetector.expansion = 0.2
}

func (a *AbstractAABBDetector) DetectColliders(b, c collision.Collider) bool {
	bAABB := b.CreateAABB()
	cAABB := c.CreateAABB()
	return bAABB.Overlaps(cAABB)
}

func (a *AbstractAABBDetector) DetectConvexTransform(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform) bool {
	b := convex1.CreateAABBTransform(transform1)
	c := convex2.CreateAABBTransform(transform2)
	return b.Overlaps(c)
}

func (a *AbstractAABBDetector) GetAABBExpansion() float64 {
	return a.expansion
}

func (a *AbstractAABBDetector) SetAABBExpansion(expansion float64) {
	a.expansion = expansion
}
