package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
)

type Raycast struct {
	normal, point *geometry.Vector2
	distance      float64
}

func NewSeparation() *Raycast {
	r := new(Raycast)
	r.normal = new(geometry.Vector2)
	r.point = new(geometry.Vector2)
	return r
}

func NewSeparationVector2Vector2Float64(point, normal *geometry.Vector2, distance float64) *Raycast {
	r := new(Raycast)
	r.point = point
	r.normal = normal
	r.distance = distance
	return r
}

func (r *Raycast) Clear() {
	r.point = nil
	r.normal = nil
	r.distance = 0
}

func (r *Raycast) GetPoint() *geometry.Vector2 {
	return r.point
}

func (r *Raycast) GetNormal() *geometry.Vector2 {
	return r.normal
}

func (r *Raycast) GetDistance() float64 {
	return r.distance
}

func (r *Raycast) SetPoint(point1 *geometry.Vector2) {
	r.point = point
}

func (r *Raycast) SetNormal(normal *geometry.Vector2) {
	r.normal = normal
}

func (r *Raycast) SetDistance(distance flaot64) {
	r.distance = distance
}
