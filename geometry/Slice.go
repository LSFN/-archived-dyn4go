package geometry

import (
	"math"
)

type Slice struct {
	AbstractShape
	theta, alpha, sliceRadius float64
	vertices, normals         []*Vector2
	localXAxis                *Vector2
}

func NewSlice(radius, theta float64) *Slice {
	if radius == 0 || theta <= 0 || theta > math.Pi {
		panic("Cannot create Slice from zero radius, zero theta or theta greater than Pi")
	}
	s := new(Slice)
	s.radius = radius
	s.theta = theta
	s.alpha = theta * 0.5
	cx := 2 * radius * math.Sin(s.alpha) / (3 * s.alpha)
	s.center = NewVector2FromXY(cx, 0)
	x := radius * math.Cos(s.alpha)
	y := radius * math.Sin(s.alpha)
	s.vertices = []*Vector2{
		new(Vector2),
		NewVector2FromXY(x, y),
		NewVector2FromXY(x, -y),
	}
	v1 := s.vertices[1].HereToVector2(s.vertices[0])
	v2 := s.vertices[0].HereToVector2(s.vertices[2])
	v1.Left().Normalize()
	v2.Left().Normalize()
	s.normals = []*Vector2{v1, v2}
	cToOrigin := s.center.GetMagnitudeSquared()
	cToTop := s.center.DistanceSquaredFromVector2(s.vertices[1])
	s.radius = math.Sqrt(math.Max(cToOrigin, cToTop))
	s.localXAxis = NewVector2FromXY(1, 0)
	return s
}

func (s *Slice) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	fociSize := 0
	if foci != nil {
		fociSize = len(foci)
	}
	axes := make([]*Vector2, 2+fociSize)
	axes[0] = transform.GetTransformedR(s.normals[0])
	axes[1] = transform.GetTransformedR(s.normals[1])
	n := 2
	focus := transform.GetTransformedVector2(s.vertices[0])
	for _, f := range foci {
		closest := focus
		d := f.DistanceSquaredFromVector2(closest)
		for _, p := range s.vertices {
			p = transform.GetTransformedVector2(p)
			dt := f.DistanceSquaredFromVector2(p)
			if dt < d {
				closest = p
				d = dt
			}
		}
		axis := f.HereToVector2(closest)
		axis.Normalize()
		axes[n] = axis
		n++
	}
	return axes
}

func (s *Slice) GetFoci(transform *Transform) []*Vector2 {
	return []*Vector2{transform.GetTransformedVector2(s.vertices[0])}
}

func (s *Slice) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	localn := transform.GetInverseTransformedR(n)
	if math.Abs(localn.GetAngleBetween(s.localXAxis)) > s.alpha {
		point := new(Vector2)
		point.SetToVector2(s.vertices[0])
		max := localn.DotVector2(s.vertices[0])
		for _, v := range s.vertices {
			projection := localn.DotVector2(v)
			if projection > max {
				point.SetToVector2(v)
				max = projection
			}
		}
		transform.Transform(point)
		return point
	} else {
		localn.Normalize()
		localn.Multiply(s.sliceRadius).AddVector2(s.vertices[0])
		transform.Transform(localn)
		return localn
	}
}

func (s *Slice) GetFarthestFeature(n *Vector2, transform *Transform) *Feature {
	localAxis := transform.GetInverseTransformedR(n)
	if math.Abs(localAxis.GetAngleBetween(s.localXAxis)) <= s.alpha {
		point := s.GetFarthestFeature(n, transform)
		return NewVertexFromPoint(point)
	} else {
		if math.Pi-s.theta <= 1.0e-6 {
			return GetFarthestFeature(s.vertices[1], s.vertices[2], n, transform)
		}
		if localAxis.Y > 0 {
			return GetFarthestFeature(s.vertices[0], s.vertices[1], n, transform)
		} else if localAxis.Y < 0 {
			return GetFarthestFeature(s.vertices[0], s.vertices[2], n, transform)
		} else {
			return NewVertexFromPoint(transform.GetTransformedVector2(s.vertices[0]))
		}
	}
}

func (s *Slice) Project(n *Vector2, transform *Transform) *Interval {
	p1 := s.GetFarthestPoint(n, transform)
	p2 := s.GetFarthestPoint(n.GetNegative(), transform)
	d1 := p1.DotVector2(n)
	d2 := p2.DotVector2(n)
	return NewIntervalFromMinMax(d2, d1)
}

func (s *Slice) CreateAABB(transform *Transform) *AABB {
	x := s.Project(X_AXIS, transform)
	y := s.Project(Y_AXIS, transform)
	return NewAABBFromFloats(x.GetMin(), y.GetMin(), x.GetMax(), y.GetMax())
}

func (s *Slice) CreateMass(density float64) *Mass {
	r2 := s.sliceRadius * s.sliceRadius
	m := density * r2 * s.alpha
	sina := math.Sin(s.alpha)
	I := 1.0 / 18.0 * r2 * r2 * (9.0*s.alpha*s.alpha - 8.0*sina*sina) / s.alpha
	return NewMassFromCenterMassInertia(s.center, m, I)
}

func (s *Slice) GetRadius(center *Vector2) float64 {
	return s.radius + center.DistanceFromVector2(s.center)
}

func (s *Slice) Contains(point *Vector2, transform *Transform) bool {
	lp := transform.GetInverseTransformedVector2(point)
	radiusSquared := s.sliceRadius * s.sliceRadius
	v := s.vertices[0].HereToVector2(lp)
	if v.GetMagnitudeSquared() <= radiusSquared {
		if GetLocation(lp, s.vertices[0], s.vertices[1]) <= 0 &&
			GetLocation(lp, s.vertices[0], s.vertices[2]) >= 0 {
			return true
		}
	}
	return false
}

func (s *Slice) Rotate(theta, x, y float64) {
	s.Rotate(theta, x, y)
	for _, v := range s.vertices {
		v.RotateAboutXY(theta, x, y)
	}
	for _, n := range s.normals {
		n.RotateAboutOrigin(theta)
	}
	s.localXAxis.RotateAboutOrigin(theta)
}

func (s *Slice) Translate(x, y float64) {
	s.Translate(x, y)
	for _, v := range s.vertices {
		v.AddXY(x, y)
	}
}

func (s *Slice) GetRotation() float64 {
	return X_AXIS.GetAngleBetween(s.localXAxis)
}

func (s *Slice) GetTheta() float64 {
	return s.theta
}

func (s *Slice) GetSliceRadius() float64 {
	return s.sliceRadius
}

func (s *Slice) GetCircleCenter() *Vector2 {
	return s.vertices[0]
}
