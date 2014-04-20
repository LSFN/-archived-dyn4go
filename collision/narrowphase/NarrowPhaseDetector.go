package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
)

type NarrowphaseDetector interface {
	DetectPenetration(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, penetration *Penetration) bool
	Detect(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, penetration *Penetration) bool
}
