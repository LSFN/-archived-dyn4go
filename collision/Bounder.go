package collision

import (
	"github.com/LSFN/dyn4go/geometry"
)

type Bounder interface {
	geometry.Transformer
	GetTransform() *geometry.Transform
	ShiftCoordinates(shift *geometry.Vector2)
	IsOutside(collider Collider)
}
