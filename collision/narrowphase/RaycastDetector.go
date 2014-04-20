package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
)

type RaycastDetector interface {
	Raycast(ray *geometry.Ray, maxLength float64, convex geometry.Convexer, transform *geometry.Transform, raycast *Raycast) bool
}
