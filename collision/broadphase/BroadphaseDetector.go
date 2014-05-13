package broadphase

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
)

type BroadphaseDetector interface {
	Add(collidable Collider)
	Remove(collidable Collider)
	Update(collidable Collider)
	Clear()
	GetAABB(collidable Collider) *geometry.AABB
	DetectAABB(aabb *geometry.AABB) []collision.Collider
	Raycast(ray *geometry.Ray, length float64) []collision.Collider
	Detect(a, b interface{}) bool
	DetectConvexTransform(convex1 *geometry.Convexer, transform1 *geometry.Transform, convex2 *geometry.Convexer, transform2 *geometry.Transform) bool
	GetAABBExpansion() float64
	SetAABBExpansion(expansion float64)
	ShiftCoordinates(shift *geometry.Vector2)
}
