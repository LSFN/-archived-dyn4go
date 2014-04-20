package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
	"reflect"
)

type SAT struct{}

func (s *SAT) DetectPenetration(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, penetration *Penetration) bool {
	if reflect.TypeOf(convex1) == "*geometry.Circle" && reflect.TypeOf(convex2) == "*geometry.Circle" {
		return DetectCirclePenetration(convex1.(*geometry.Circle), transform1, convex2.(*geometry.Circle), transform2, penetration)
	}
}
