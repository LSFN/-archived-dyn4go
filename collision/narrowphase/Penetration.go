package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
	"strconv"
)

type Penetration struct {
	normal *geometry.Vector2
	depth  float64
}

func NewPenetration() *Penetration {
	p := new(Penetration)
	p.normal = new(geometry.Vector2)
	return p
}

func (p *Penetration) Clear() {
	p.normal = nil
	p.depth = 0
}

func (p *Penetration) GetNormal() *geometry.Vector2 {
	return p.normal
}

func (p *Penetration) GetDepth() float64 {
	return p.depth
}

func (p *Penetration) SetNormal(normal *geometry.Vector2) {
	p.normal = normal
}

func (p *Penetration) SetDepth(depth float64) {
	p.depth = depth
}

func (p *Penetration) String() string {
	return "Penetration[normal=" + p.normal.String() + ",depth=" + strconv.FormatFloat(p.depth, 'g', -1, 64) + "]"
}
