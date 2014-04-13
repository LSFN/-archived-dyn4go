package geometry2

import (
	"math"

	"code.google.com/p/uuid"
	"github.com/LSFN/dyn4go"
)

type Segment struct {
	Wound
	length float64
}

func NewSegment(p1, p2 *Vector2) *Segment {
	if p1 == nil || p2 == nil {
		panic("Arguments to NewSegment must not be nil")
	}
	if *p1 == *p2 {
		panic("Arguments to NewSegment must not be equivalent")
	}
	s := new(Segment)
	s.vertices = make([]*Vector2, 2)
	s.vertices[0] = p1
	s.vertices[1] = p2
	s.normals = make([]*Vector2, 2)
	s.normals[0] = p1.HereToVector2(p2).Right()
	s.normals[0].Normalize()
	s.normals[1] = p1.HereToVector2(p2).Left()
	s.normals[1].Normalize()
	s.center = GetAverageCenterFromList(s.vertices)
	s.length = p1.DistanceFromVector2(p2)
	s.radius = s.length * 0.5
	s.id = uuid.New()
	return s
}

func (s *Segment) GetPoint1() *Vector2 {
	return s.vertices[0]
}

func (s *Segment) GetPoint2() *Vector2 {
	return s.vertices[1]
}

func (s *Segment) GetLength() float64 {
	return s.length
}

func GetLocation(point, linePoint1, linePoint2 *Vector2) float64 {
	return (linePoint2.X-linePoint1.X)*(point.Y-linePoint1.Y) - (point.X-linePoint1.X)*(linePoint2.Y-linePoint1.Y)
}

func GetPointOnLineClosestToPoint(point, linePoint1, linePoint2 *Vector2) *Vector2 {
	p1ToP := point.DifferenceVector2(linePoint1)
	line := linePoint2.DifferenceVector2(linePoint1)
	ab2 := line.DotVector2(line)
	if ab2 <= dyn4go.Epsilon {
		return NewVector2FromVector2(linePoint1)
	}
	ap_ab := p1ToP.DotVector2(line)
	t := ap_ab / ab2
	return line.Multiply(t).AddVector2(linePoint1)
}

func (s *Segment) GetPointOnLineClosestToPoint(point *Vector2) *Vector2 {
	return GetPointOnLineClosestToPoint(point, s.vertices[0], s.vertices[1])
}

func GetPointOnSegmentClosestToPoint(point, linePoint1, linePoint2 *Vector2) *Vector2 {
	p1ToP := point.DifferenceVector2(linePoint1)
	line := linePoint2.DifferenceVector2(linePoint1)
	ab2 := line.DotVector2(line)
	ap_ab := p1ToP.DotVector2(line)
	if ab2 <= dyn4go.Epsilon {
		return NewVector2FromVector2(linePoint1)
	}
	t := ap_ab / ab2
	t = IntervalClamp(t, 0, 1)
	return line.Multiply(t).AddVector2(linePoint1)
}

func (s *Segment) GetPointOnSegmentClosestToPoint(point *Vector2) *Vector2 {
	return GetPointOnSegmentClosestToPoint(point, s.vertices[0], s.vertices[1])
}

func GetLineIntersection(ap1, ap2, bp1, bp2 *Vector2) *Vector2 {
	a := ap1.HereToVector2(ap2)
	b := bp1.HereToVector2(bp2)
	bxa := b.CrossVector2(a)
	if math.Abs(bxa) <= dyn4go.Epsilon {
		return nil
	}
	ambxA := ap1.DifferenceVector2(bp1).CrossVector2(a)
	if math.Abs(ambxA) <= dyn4go.Epsilon {
		return nil
	}
	return b.Product(ambxA / bxa).AddVector2(bp1)
}

func (s *Segment) GetLineIntersection(segment *Segment) *Vector2 {
	return GetLineIntersection(s.vertices[0], s.vertices[1], segment.vertices[0], segment.vertices[1])
}

func GetSegmentIntersection(ap1, ap2, bp1, bp2 *Vector2) *Vector2 {
	a := ap1.HereToVector2(ap2)
	b := bp1.HereToVector2(bp2)
	bxa := b.CrossVector2(a)
	if math.Abs(bxa) <= dyn4go.Epsilon {
		return nil
	}
	ambxA := ap1.DifferenceVector2(bp1).CrossVector2(a)
	if math.Abs(ambxA) <= dyn4go.Epsilon {
		return nil
	}
	tb := ambxA / bxa
	if tb < 0 || tb > 1 {
		return nil
	}
	ip := b.Product(tb).AddVector2(bp1)
	ta := ip.DifferenceVector2(ap1).DotVector2(a) / a.DotVector2(a)
	if ta < 0 || ta > 1 {
		return nil
	}
	return ip
}

func (s *Segment) GetSegmentIntersection(segment *Segment) *Vector2 {
	return GetSegmentIntersection(s.vertices[0], s.vertices[1], segment.vertices[0], segment.vertices[1])
}

func GetFarthestFeature(v1, v2, n *Vector2, t *Transform) Featurer {
	p1 := t.GetTransformedVector2(v1)
	p2 := t.GetTransformedVector2(v2)
	dot1 := n.DotVector2(p1)
	dot2 := n.DotVector2(p2)
	max := p1
	index := 0
	if dot1 < dot2 {
		max = p2
		index = 1
	}
	vp1 := NewVertexVector2Int(p1, 0)
	vp2 := NewVertexVector2Int(p2, 1)
	vm := NewVertexVector2Int(max, index)
	if p1.HereToVector2(p2).Right().DotVector2(n) > 0 {
		return NewEdge(vp2, vp1, vm, p2.HereToVector2(p1), 0)
	} else {
		return NewEdge(vp1, vp2, vm, p1.HereToVector2(p2), 0)
	}
}

func GetFarthestPoint(v1, v2, n *Vector2, t *Transform) *Vector2 {
	p1 := t.GetTransformedVector2(v1)
	p2 := t.GetTransformedVector2(v2)
	dot1 := n.DotVector2(p1)
	dot2 := n.DotVector2(p2)
	if dot1 >= dot2 {
		return p1
	} else {
		return p2
	}
}

func (s *Segment) GetAxes(foci []*Vector2, t *Transform) []*Vector2 {
	size := 0
	if foci != nil {
		size = len(foci)
	}
	axes := make([]*Vector2, 2+size)
	n := 0
	p1 := t.GetTransformedVector2(s.vertices[0])
	p2 := t.GetTransformedVector2(s.vertices[1])
	axes[n] = t.GetTransformedR(s.normals[1])
	n++
	axes[n] = t.GetTransformedR(s.normals[0].GetLeftHandOrthogonalVector())
	n++
	var axis *Vector2
	for _, f := range foci {
		if p1.DistanceSquaredFromVector2(f) < p2.DistanceSquaredFromVector2(f) {
			axis = p1.HereToVector2(f)
		} else {
			axis = p2.HereToVector2(f)
		}
		axis.Normalize()
		axes[n] = axis
		n++
	}
	return axes
}

func (s *Segment) GetFoci(transform *Transform) []*Vector2 {
	return nil
}

func (s *Segment) ContainsVector2Transform(point *Vector2, transform *Transform) bool {
	p := transform.GetInverseTransformedVector2(point)
	value := GetLocation(p, s.vertices[0], s.vertices[1])
	if math.Abs(value) <= dyn4go.Epsilon {
		distSqrd := s.vertices[0].DistanceSquaredFromVector2(s.vertices[1])
		if p.DistanceSquaredFromVector2(s.vertices[0]) <= distSqrd &&
			p.DistanceSquaredFromVector2(s.vertices[1]) <= distSqrd {
			return true
		}
	}
	return false
}

func (s *Segment) ContainsTransformRadius(point *Vector2, transform *Transform, radius float64) bool {
	if radius <= 0 {
		return s.ContainsVector2Transform(point, transform)
	} else {
		p := transform.GetInverseTransformedVector2(point)
		if s.vertices[0].DistanceSquaredFromVector2(p) <= radius*radius ||
			s.vertices[1].DistanceSquaredFromVector2(p) <= radius*radius {
			return true
		} else {
			l := s.vertices[0].HereToVector2(s.vertices[1])
			p1 := s.vertices[0].HereToVector2(p)
			p2 := s.vertices[1].HereToVector2(p)
			if l.DotVector2(p1) > 0 && -l.DotVector2(p2) > 0 {
				dist := p1.Project(l.GetRightHandOrthogonalVector()).GetMagnitudeSquared()
				return dist <= radius*radius
			}
		}
	}
	return false
}

func (s *Segment) ProjectVector2Transform(n *Vector2, transform *Transform) *Interval {
	p1 := transform.GetTransformedVector2(s.vertices[0])
	p2 := transform.GetTransformedVector2(s.vertices[1])
	min := n.DotVector2(p1)
	max := min
	v := n.DotVector2(p2)
	if v < min {
		min = v
	} else if v > max {
		max = v
	}
	return NewIntervalFromMinMax(min, max)
}

func (s *Segment) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	return GetFarthestPoint(s.vertices[0], s.vertices[1], n, transform)
}

func (s *Segment) GetFarthestFeature(n *Vector2, transform *Transform) Featurer {
	return GetFarthestFeature(s.vertices[0], s.vertices[1], n, transform)
}

func (s *Segment) RotateAboutXY(theta, x, y float64) {
	if !(s.center.X == x && s.center.Y == y) {
		s.center.RotateAboutXY(theta, x, y)
	}
	s.vertices[0].RotateAboutXY(theta, x, y)
	s.vertices[1].RotateAboutXY(theta, x, y)
	s.normals[0].RotateAboutXY(theta, x, y)
	s.normals[1].RotateAboutXY(theta, x, y)
}

func (s *Segment) TranslateXY(x, y float64) {
	s.center.AddXY(x, y)
	s.vertices[0].AddXY(x, y)
	s.vertices[1].AddXY(x, y)
}

func (s *Segment) CreateMass(density float64) *Mass {
	length := s.length
	mass := density * length
	inertia := length * length * mass / 12.0
	return NewMassFromCenterMassInertia(s.center, mass, inertia)
}

func (s *Segment) CreateAABBTransform(transform *Transform) *AABB {
	p := transform.GetTransformedVector2(s.vertices[0])
	minX := NewVector2FromVector2(&X_AXIS).DotVector2(p)
	maxX := minX
	minY := NewVector2FromVector2(&Y_AXIS).DotVector2(p)
	maxY := minY
	p = transform.GetTransformedVector2(s.vertices[1])
	vx := NewVector2FromVector2(&X_AXIS).DotVector2(p)
	vy := NewVector2FromVector2(&Y_AXIS).DotVector2(p)
	minX = math.Min(minX, vx)
	maxX = math.Max(maxX, vx)
	minY = math.Min(minY, vy)
	maxY = math.Max(maxY, vy)
	return NewAABBFromFloats(minX, minY, maxX, maxY)
}

func (s *Segment) GetRadiusVector2(v *Vector2) float64 {
	r2 := 0.0
	for _, v2 := range s.vertices {
		r2t := v.DistanceSquaredFromVector2(v2)
		r2 = math.Max(r2, r2t)
	}
	return math.Sqrt(r2)
}

func (s *Segment) ContainsVector2(v *Vector2) bool {
	return s.ContainsVector2Transform(v, NewTransform())
}

func (s *Segment) ProjectVector2(v *Vector2) *Interval {
	return s.ProjectVector2Transform(v, NewTransform())
}

func (s *Segment) CreateAABB() *AABB {
	return s.CreateAABBTransform(NewTransform())
}

func (s *Segment) RotateAboutOrigin(theta float64) {
	s.RotateAboutXY(theta, 0, 0)
}

func (s *Segment) RotateAboutCenter(theta float64) {
	s.RotateAboutXY(theta, s.center.X, s.center.Y)
}

func (s *Segment) RotateAboutVector2(theta float64, v *Vector2) {
	s.RotateAboutXY(theta, v.X, v.Y)
}

func (s *Segment) TranslateVector2(v *Vector2) {
	s.TranslateXY(v.X, v.Y)
}
