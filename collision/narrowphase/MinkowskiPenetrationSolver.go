package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
)

type MinkowskiPenetrationSolver interface {
	GetPenetration(simplex []*geometry.Vector2, minkowskiSum *MinkowskiSum, penetration *Penetration)
}
