package collision

import (
	"github.com/LSFN/dyn4go/geometry"
)

type BoundsAdapter struct{}

func (b *BoundsAdapter) Outside(collidable Collider) {}
