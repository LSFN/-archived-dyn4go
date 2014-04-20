package narrowphase

import (
	"github.com/LSFN/dyn4go"
	"github.com/LSFN/dyn4go/geometry"
	"math"
)

func RaycastSegment(ray *geometry.Ray, maxLength float64, segment *geometry.Segment, transform *geometry.Transform, raycast *Raycast) bool {
	p0 := ray.GetStart()
	d0 := ray.GetDirectionVector2()
	p1 := transform.GetTransformedVector2(segment.GetPoint1())
	p2 := transform.GetTransformedVector2(segment.GetPoint2())
	d1 := p1.HereToVector2(p2)
	if segment.ContainsVector2Transform(p0, transform) {
		return false
	}
	p0ToP1 := p1.DifferenceVector2(p0)
	num := d1.CrossVector2(p0ToP1)
	den := d1.CrossVector2(d0)
	if math.Abs(den) <= dyn4go.Epsilon {
		d0DotP0 := d0.DotVector2(p0)
		d0DotP1 := d0.DotVector2(p1)
		d0DotP2 := d0.DotVector2(p2)
		if d0DotP1 < 0 || d0DotP2 < 0 {
			return false
		}
		d := 0.0
		var p *geometry.Vector2 = nil
		if d0DotP1 < d0DotP2 {
			d = d0DotP1 - d0DotP0
			p = geometry.NewVector2FromVector2(p1)
		} else {
			d = d0DotP1 - d0DotP0
			p = geometry.NewVector2FromVector2(p2)
		}
		if maxLength > 0 && d > maxLength {
			return false
		}
		raycast.distance = d
		raycast.point = p
		raycast.normal = d0.GetNegative()
		return true
	} else {
		return false
	}
	t := num / den
	if t < 0 {
		return false
	}
	if maxLength > 0 && t > maxLength {
		return false
	}
	s := (t*d0.X + p0.X - p1.X) / d1.X
	if s < 0 || s > 1 {
		return false
	}
	p := d0.Product(t).AddVector2(p0)
	l := p1.HereToVector2(p2)
	l.Normalize()
	l.Right()
	lDotD := l.DotVector2(d0)
	if lDotD > 0 {
		l.Negate()
	}
	raycast.point = p
	raycast.normal = l
	raycast.distance = t
	return true
}
