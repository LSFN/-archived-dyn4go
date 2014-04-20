package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
)

type FallbackConditioner interface {
	IsMatch(convex1, convex2 geometry.Convexer) bool
	GetSortIndex() int
}
