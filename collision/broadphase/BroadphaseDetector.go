package broadphase

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
)

var DEFAULT_AABB_EXPANSION float64 = 0.2

type BroadphaseDetector interface {
	Add(collidable collision.Collider)
	Remove(collidable collision.Collider)
	Update(collidable collision.Collider)
	Clear()
	GetAABB(collidable collision.Collider) *geometry.AABB
	DetectAABB(aabb *geometry.AABB) []collision.Collider
	Raycast(ray *geometry.Ray, length float64) []collision.Collider
	Detect(a, b interface{}) bool
	DetectConvexTransform(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform) bool
	GetAABBExpansion() float64
	SetAABBExpansion(expansion float64)
	ShiftCoordinates(shift *geometry.Vector2)
}
