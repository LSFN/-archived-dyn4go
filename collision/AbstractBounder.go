package collision

import (
	"github.com/LSFN/dyn4go/geometry"
)

type AbstractBounder interface {
	Bounder
	geometry.Transformer
}
