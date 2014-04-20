package collision

import (
	"github.com/LSFN/dyn4go"
)

type BoundsListener interface {
	dyn4go.Listener
	Outside(collidable Collider)
}
