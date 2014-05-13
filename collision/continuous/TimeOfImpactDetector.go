package continuous

import (
	"github.com/LSFN/dyn4go/geometry"
)

type TimeOfImpactDetector interface {
	GetTimeOfImpact(convex1 geometry.Convexer, transform1 *geometry.Transform, dp1 *geometry.Vector2, da1 float64, convex2 geometry.Convexer, transform2 *geometry.Transform, dp2 *geometry.Vector2, da2 float64, toi *TimeOfImpact) bool
	GetTimeOfImpactBounded(convex1 geometry.Convexer, transform1 *geometry.Transform, dp1 *geometry.Vector2, da1 float64, convex2 geometry.Convexer, transform2 *geometry.Transform, dp2 *geometry.Vector2, da2, t1, t2 float64, toi *TimeOfImpact) bool
}
