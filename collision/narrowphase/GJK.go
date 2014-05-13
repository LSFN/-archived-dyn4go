package narrowphase

import (
	"github.com/LSFN/dyn4go"
	"github.com/LSFN/dyn4go/geometry"
	"math"
	"reflect"
)

type GJK struct {
	minkowskiPenetrationSolver MinkowskiPenetrationSolver
	maxIterations              int
	distanceEpsilon            float64
}

func NewGJK() *GJK {
	g := new(GJK)
	g.minkowskiPenetrationSolver = NewEPA()
	g.maxIterations = 30
	g.distanceEpsilon = math.Sqrt(dyn4go.Epsilon)
	return g
}

func NewGJKMinkowskiPenetrationSolver(minkowskiPenetrationSolver MinkowskiPenetrationSolver) *GJK {
	g := new(GJK)
	g.minkowskiPenetrationSolver = minkowskiPenetrationSolver
	g.maxIterations = 30
	g.distanceEpsilon = math.Sqrt(dyn4go.Epsilon)
	return g
}

func (g *GJK) getInitialDirection(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform) *geometry.Vector2 {
	c1 := transform1.GetTransformedVector2(convex1.GetCenter())
	c2 := transform2.GetTransformedVector2(convex2.GetCenter())
	return c1.HereToVector2(c2)
}

func (g *GJK) DetectPenetration(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, penetration *Penetration) bool {
	if reflect.TypeOf(convex1) == reflect.TypeOf(new(geometry.Circle)) && reflect.TypeOf(convex2) == reflect.TypeOf(new(geometry.Circle)) {
		return DetectCirclePenetration(convex1.(*geometry.Circle), transform1, convex2.(*geometry.Circle), transform2, penetration)
	}
	simplex := make([]*geometry.Vector2, 0, 3)
	ms := NewMinkowskiSum(convex1, transform1, convex2, transform2)
	d := g.getInitialDirection(convex1, transform1, convex2, transform2)
	if g.detect(ms, simplex, d) {
		g.minkowskiPenetrationSolver.GetPenetration(simplex, ms, penetration)
		return true
	}

	return false
}

func (g *GJK) Detect(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform) bool {
	if reflect.TypeOf(convex1) == reflect.TypeOf(new(geometry.Circle)) && reflect.TypeOf(convex2) == reflect.TypeOf(new(geometry.Circle)) {
		return DetectCircle(convex1.(*geometry.Circle), transform1, convex2.(*geometry.Circle), transform2)
	}
	simplex := make([]*geometry.Vector2, 0, 3)
	ms := NewMinkowskiSum(convex1, transform1, convex2, transform2)
	d := g.getInitialDirection(convex1, transform1, convex2, transform2)
	return g.detect(ms, simplex, d)
}

func (g *GJK) detect(ms *MinkowskiSum, simplex []*geometry.Vector2, d *geometry.Vector2) bool {
	if d.IsZero() {
		d.SetToXY(1, 0)
	}
	simplex = append(simplex, ms.Support(d))
	if simplex[0].DotVector2(d) <= 0 {
		return false
	}
	d.Negate()
	for true {
		simplex = append(simplex, ms.Support(d))
		if simplex[len(simplex)-1].DotVector2(d) <= 0 {
			return false
		} else {
			if g.checkSimplex(simplex, d) {
				return true
			}
		}
	}
	return false
}

func (g *GJK) checkSimplex(simplex []*geometry.Vector2, direction *geometry.Vector2) bool {
	a := simplex[len(simplex)-1]
	ao := a.GetNegative()
	if len(simplex) == 3 {
		b := simplex[1]
		c := simplex[0]
		ab := a.HereToVector2(b)
		ac := a.HereToVector2(c)
		abPerp := geometry.Vector2TripleProduct(ac, ab, ab)
		acPerp := geometry.Vector2TripleProduct(ab, ac, ac)
		acLocation := acPerp.DotVector2(ao)
		if acLocation >= 0 {
			simplex = append(simplex[:1], simplex[2:]...)
			direction.SetToVector2(acPerp)
		} else {
			abLocation := abPerp.DotVector2(ao)
			if abLocation < 0 {
				return true
			} else {
				simplex = simplex[1:]
				direction.SetToVector2(abPerp)
			}
		}
	} else {
		b := simplex[0]
		ab := a.HereToVector2(b)
		direction.SetToVector2(geometry.Vector2TripleProduct(ab, ao, ab))
		if direction.GetMagnitudeSquared() <= dyn4go.Epsilon {
			direction.SetToVector2(ab.Left())
		}
	}
	return false
}

func (g *GJK) Distance(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, separation *Separation) bool {
	if reflect.TypeOf(convex1) == reflect.TypeOf(new(geometry.Circle)) && reflect.TypeOf(convex2) == reflect.TypeOf(new(geometry.Circle)) {
		return DistanceCircle(convex1.(*geometry.Circle), transform1, convex2.(*geometry.Circle), transform2, separation)
	}
	ms := NewMinkowskiSum(convex1, transform1, convex2, transform2)
	a := NewMinkowskiSumPoint()
	b := NewMinkowskiSumPoint()
	c := NewMinkowskiSumPoint()
	c1 := transform1.GetTransformedVector2(convex1.GetCenter())
	c2 := transform2.GetTransformedVector2(convex2.GetCenter())
	d := c1.HereToVector2(c2)
	if d.IsZero() {
		return false
	}
	ms.SupportMinkowskiSumPoint(d, a)
	d.Negate()
	ms.SupportMinkowskiSumPoint(d, b)
	d = geometry.GetPointOnSegmentClosestToPoint(new(geometry.Vector2), b.p, a.p)
	for i := 0; i < g.maxIterations; i++ {
		d.Negate()
		if d.GetMagnitudeSquared() <= dyn4go.Epsilon {
			return false
		}
		ms.SupportMinkowskiSumPoint(d, c)
		if g.containsOrigin(a.p, b.p, c.p) {
			return false
		}
		projection := c.p.DotVector2(d)
		if projection-a.p.DotVector2(d) < g.distanceEpsilon {
			d.Normalize()
			separation.normal = d
			separation.distance = -c.p.DotVector2(d)
			g.findClosestPoints(a, b, separation)
			return true
		}
		p1 := geometry.GetPointOnSegmentClosestToPoint(new(geometry.Vector2), a.p, c.p)
		p2 := geometry.GetPointOnSegmentClosestToPoint(new(geometry.Vector2), c.p, b.p)
		p1Mag := p1.GetMagnitudeSquared()
		p2Mag := p2.GetMagnitudeSquared()
		if p1Mag <= dyn4go.Epsilon {
			d.Normalize()
			separation.distance = p1.Normalize()
			separation.normal = d
			g.findClosestPoints(c, b, separation)
			return true
		}
		if p1Mag < p2Mag {
			b.SetMinkowskiSumPoint(c)
			d = p1
		} else {
			a.SetMinkowskiSumPoint(c)
			d = p2
		}
	}
	d.Normalize()
	separation.normal = d
	separation.distance = -c.p.DotVector2(d)
	g.findClosestPoints(a, b, separation)
	return true
}

func (g *GJK) findClosestPoints(a, b *MinkowskiSumPoint, s *Separation) {
	p1 := new(geometry.Vector2)
	p2 := new(geometry.Vector2)
	l := a.p.HereToVector2(b.p)
	if l.IsZero() {
		p1.SetToVector2(a.p1)
		p2.SetToVector2(a.p2)
	} else {
		ll := l.DotVector2(l)
		l2 := -l.DotVector2(a.p) / ll
		l1 := 1 - l2
		if l1 < 0 {
			p1.SetToVector2(b.p1)
			p2.SetToVector2(b.p2)
		} else if l2 < 0 {
			p1.SetToVector2(a.p1)
			p2.SetToVector2(a.p2)
		} else {
			p1.X = a.p1.X*l1 + b.p1.X*l2
			p1.Y = a.p1.Y*l1 + b.p1.Y*l2
			p2.X = a.p2.X*l1 + b.p2.X*l2
			p2.Y = a.p2.Y*l1 + b.p2.Y*l2
		}
	}
	s.point1 = p1
	s.point2 = p2
}

func (g *GJK) containsOrigin(a, b, c *geometry.Vector2) bool {
	sa := a.CrossVector2(b)
	sb := b.CrossVector2(c)
	sc := c.CrossVector2(a)
	return (sa*sb > 0) && (sa*sc > 0)
}

func (g *GJK) Raycast(ray *geometry.Ray, maxLength float64, convex geometry.Convexer, transform *geometry.Transform, raycast *Raycast) bool {
	if reflect.TypeOf(convex) == reflect.TypeOf(new(geometry.Circle)) {
		return RaycastCircle(ray, maxLength, convex.(*geometry.Circle), transform, raycast)
	}
	if reflect.TypeOf(convex) == reflect.TypeOf(new(geometry.Segment)) {
		return RaycastSegment(ray, maxLength, convex.(*geometry.Segment), transform, raycast)
	}
	λ := 0.0
	lengthCheck := (maxLength > 0)
	var a, b *geometry.Vector2
	start := ray.GetStart()
	x := start
	r := ray.GetDirectionVector2()
	n := new(geometry.Vector2)
	if convex.ContainsVector2Transform(start, transform) {
		return false
	}
	c := transform.GetTransformedVector2(convex.GetCenter())
	d := c.HereToVector2(x)
	distanceSqrd := math.Inf(1)
	iterations := 0
	for distanceSqrd > g.distanceEpsilon {
		p := convex.GetFarthestPoint(d, transform)
		w := p.HereToVector2(x)
		dDotW := d.DotVector2(w)
		if dDotW > 0 {
			dDotR := d.DotVector2(r)
			if dDotR >= 0 {
				return false
			} else {
				λ = λ - dDotW/dDotR
				if lengthCheck && λ > maxLength {
					return false
				}
				x = r.Product(λ).AddVector2(start)
				n.SetToVector2(d)
			}
		}
		if a != nil {
			if b != nil {
				p1 := geometry.GetPointOnSegmentClosestToPoint(x, a, p)
				p2 := geometry.GetPointOnSegmentClosestToPoint(x, p, b)
				if p1.DistanceSquaredFromVector2(x) < p2.DistanceSquaredFromVector2(x) {
					b.SetToVector2(p)
					distanceSqrd = p1.DistanceSquaredFromVector2(x)
				} else {
					a.SetToVector2(p)
					distanceSqrd = p2.DistanceSquaredFromVector2(x)
				}
				ab := a.HereToVector2(b)
				ax := a.HereToVector2(x)
				d = geometry.Vector2TripleProduct(ab, ax, ab)
			} else {
				b = p
				ab := a.HereToVector2(b)
				ax := a.HereToVector2(x)
				d = geometry.Vector2TripleProduct(ab, ax, ab)
			}
		} else {
			a = p
			d.Negate()
		}
		if iterations == g.maxIterations {
			return false
		}
		iterations++
	}
	raycast.point = x
	raycast.normal = n
	n.Normalize()
	raycast.distance = λ

	return true
}

func (g *GJK) GetMaxIterations() int {
	return g.maxIterations
}

func (g *GJK) SetMaxIterations(maxIterations int) {
	if maxIterations < 5 {
		panic("Cannot set maximum number of iterations this low")
	}
	g.maxIterations = maxIterations
}

func (g *GJK) GetDistanceEpsilon() float64 {
	return g.distanceEpsilon
}

func (g *GJK) SetDistanceEpsilon(distanceEpsilon float64) {
	if distanceEpsilon <= 0 {
		panic("Distance epislon must be strictly positive")
	}
	g.distanceEpsilon = distanceEpsilon
}

func (g *GJK) GetMinkowskiPenetrationSolver() MinkowskiPenetrationSolver {
	return g.minkowskiPenetrationSolver
}

func (g *GJK) SetMinkowskiPenetrationSolver(minkowskiPenetrationSolver MinkowskiPenetrationSolver) {
	if minkowskiPenetrationSolver == nil {
		panic("Cannot set |Minkowski penetration solver to nil")
	}
	g.minkowskiPenetrationSolver = minkowskiPenetrationSolver
}
