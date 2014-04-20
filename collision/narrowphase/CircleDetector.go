package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
	"math"
)

func DetectCirclePenetration(circle1 *geometry.Circle, transform1 *geometry.Transform, circle2 *geometry.Circle, transform2 *geometry.Transform, penetration *Penetration) bool {
	ce1 := transform1.GetTransformedVector2(circle1.GetCenter())
	ce2 := transform2.GetTransformedVector2(circle2.GetCenter())
	v := ce1.HereToVector2(ce2)
	radii := circle1.GetRadius() + circle2.GetRadius()
	mag := v.GetMagnitude()
	if mag < radii {
		penetration.normal = v
		penetration.depth = radii - v.Normalize()
		return true
	}
	return false
}

func DetectCircle(circle1 *geometry.Circle, transform1 *geometry.Transform, circle2 *geometry.Circle, transform2 *geometry.Transform) bool {
	ce1 := transform1.GetTransformedVector2(circle1.GetCenter())
	ce2 := transform2.GetTransformedVector2(circle2.GetCenter())
	v := ce1.HereToVector2(ce2)
	radii := circle1.GetRadius() + circle2.GetRadius()
	mag := v.GetMagnitude()
	if mag < radii {
		return true
	}
	return false
}

func DistanceCircle(circle1 *geometry.Circle, transform1 *geometry.Transform, circle2 *geometry.Circle, transform2 *geometry.Transform, separation *Separation) bool {
	ce1 := transform1.GetTransformedVector2(circle1.GetCenter())
	ce2 := transform2.GetTransformedVector2(circle2.GetCenter())
	v := ce1.HereToVector2(ce2)
	r1 := circle1.GetRadius()
	r2 := circle2.GetRadius()
	radii := r1 + r2
	mag := v.GetMagnitude()
	if mag < radii {
		separation.normal = v
		separation.distance = v.Normalize() - radii
		separation.point1 = ce1.AddXY(v.X*r1, v.Y*r1)
		separation.point2 = ce2.AddXY(v.X*r2, v.Y*r2)
		return true
	}
	return false
}

func RaycastCircle(ray *geometry.Ray, maxLength float64, circle *geometry.Circle, transform *geometry.Transform, raycast *Raycast) bool {
	s := ray.GetStart()
	d := ray.GetDirectionVector2()
	ce := transform.GetTransformedVector2(circle.GetCenter())
	r := circle.GetRadius()
	if circle.ContainsVector2Transform(s, transform) {
		return false
	}
	sMinusC := s.DifferenceVector2(ce)
	a := d.DotVector2(d)
	b := 2 * d.DotVector2(sMinusC)
	c := sMinusC.DotVector2(sMinusC) - r*r
	inv2a := 1 / (2 * a)
	b24ac := b*b - 4*a*c
	if b24ac < 0 {
		return false
	}
	sqrt := math.Sqrt(b24ac)
	t0 := (-b + sqrt) * inv2a
	t1 := (-b - sqrt) * inv2a
	t := 0.0
	if t0 < 0 {
		if t1 < 0 {
			return false
		} else {
			t = t1
		}
	} else {
		if t1 < 0 {
			t = t0
		} else if t0 < t1 {
			t = t0

		} else {
			t = t1
		}
	}
	if maxLength > 0 && t > maxLength {
		return false
	}
	p := d.Product(t).AddVector2(s)
	n := ce.HereToVector2(p)
	n.Normalize()
	raycast.point = p
	raycast.normal = n
	raycast.distance = t
	return true
}
