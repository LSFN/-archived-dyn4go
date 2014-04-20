package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
)

type DistanceDetector interface {
	Distance(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, separation *Separation) bool
}
