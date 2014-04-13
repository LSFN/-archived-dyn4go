package geometry

import (
	"math"

	"code.google.com/p/uuid"
	"github.com/LSFN/dyn4go"
)

type Polygon struct {
	Wound
}

func NewPolygon(vertices ...*Vector2) *Polygon {
	if vertices == nil || len(vertices) < 3 {
		panic("Cannot create polygon without at least 3 vertices")
	}
	for _, v := range vertices {
		if v == nil {
			panic("Cannot create polygon from nil vertices")
		}
	}
	var area, sign float64
	for i := range vertices {
		var p0, p1, p2 *Vector2
		if i-1 < 0 {
			p0 = vertices[len(vertices)-1]
		} else {
			p0 = vertices[i-1]
		}
		p1 = vertices[i]
		if i+1 == len(vertices) {
			p2 = vertices[0]
		} else {
			p2 = vertices[i+1]
		}
		area += p1.CrossVector2(p2)
		if *p1 == *p2 {
			panic("Points on polygon may not coincide")
		}
		cross := p0.HereToVector2(p1).CrossVector2(p1.HereToVector2(p2))
		if cross > 0.0 {
			cross = 1.0
		} else if cross < 0.0 {
			cross = -1.0
		}
		if math.Abs(cross) > dyn4go.Epsilon && sign != 0 && cross != sign {
			panic("Non convex polygons are not allowed")
		}
		sign = cross
	}
	if area < 0 {
		panic("Invalid polygon winding")
	}
	p := new(Polygon)
	p.id = uuid.New()
	p.vertices = vertices
	p.normals = make([]*Vector2, len(vertices))
	for i, p1 := range p.vertices {
		var p2 *Vector2
		if i+1 == len(p.vertices) {
			p2 = vertices[0]
		} else {
			p2 = vertices[i+1]
		}
		n := p1.HereToVector2(p2).Left()
		n.Normalize()
		p.normals[i] = n
	}
	p.center = GetAreaWeightedCenterFromList(p.vertices)
	r2 := 0.0
	for _, v := range vertices {
		r2 = math.Max(r2, p.center.DistanceSquaredFromVector2(v))
	}
	p.radius = math.Sqrt(r2)
	return p
}

func (p *Polygon) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	arrLen := len(p.vertices)
	if foci != nil {
		arrLen += len(foci)
	}
	axes := make([]*Vector2, arrLen)
	n := 0
	for _, v := range p.normals {
		axes[n] = transform.GetTransformedR(v)
		n++
	}
	for _, f := range foci {
		closest := transform.GetTransformedVector2(p.vertices[0])
		d := f.DistanceSquaredFromVector2(closest)
		for _, v := range p.vertices {
			v2 := transform.GetTransformedVector2(v)
			dt := f.DistanceSquaredFromVector2(v2)
			if dt < d {
				closest = v2
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

func (p *Polygon) GetFoci(transform *Transform) []*Vector2 {
	return nil
}

func (p *Polygon) ContainsVector2Transform(point *Vector2, transform *Transform) bool {
	p0 := transform.GetInverseTransformedVector2(point)
	p1 := p.vertices[0]
	p2 := p.vertices[1]
	last := GetLocation(p0, p1, p2)
	for i := range p.vertices {
		p1 = p2
		if i+1 == len(p.vertices) {
			p2 = p.vertices[0]
		} else {
			p2 = p.vertices[i+1]
		}
		if p0.EqualsVector2(p1) {
			return true
		}
		if last*GetLocation(p0, p1, p2) < 0 {
			return false
		}
	}
	return true
}

func (p *Polygon) RotateAboutXY(theta, x, y float64) {
	if !(p.center.X == x && p.center.Y == y) {
		p.center.RotateAboutXY(theta, x, y)
	}
	for i := range p.vertices {
		p.vertices[i].RotateAboutXY(theta, x, y)
		p.normals[i].RotateAboutXY(theta, x, y)
	}
}

func (p *Polygon) TranslateXY(x, y float64) {
	p.center.AddXY(x, y)
	for _, v := range p.vertices {
		v.AddXY(x, y)
	}
}

func (p *Polygon) ProjectVector2Transform(n *Vector2, transform *Transform) *Interval {
	point := transform.GetInverseTransformedVector2(p.vertices[0])
	min := n.DotVector2(point)
	max := min
	for i, a := range p.vertices {
		if i != 0 {
			v := n.DotVector2(transform.GetTransformedVector2(a))
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
		}
	}
	return NewIntervalFromMinMax(min, max)
}

func (p *Polygon) GetFarthestFeature(n *Vector2, transform *Transform) Featurer {
	localn := transform.GetInverseTransformedR(n)
	maximum := new(Vector2)
	max := math.Inf(-1)
	index := 0
	for i, v := range p.vertices {
		projection := localn.DotVector2(v)
		if projection > max {
			maximum.SetToVector2(v)
			max = projection
			index = i
		}
	}
	l := index + 1
	if l == len(p.vertices) {
		l = 0
	}
	r := index - 1
	if r == -1 {
		r = len(p.vertices) - 1
	}
	c := index - 1
	if index == 0 {
		c = len(p.vertices) - 1
	}
	leftN := p.normals[c]
	rightN := p.normals[index]
	transform.Transform(maximum)
	vm := NewVertexVector2Int(maximum, index)
	if leftN.DotVector2(localn) < rightN.DotVector2(localn) {
		left := transform.GetTransformedVector2(p.vertices[l])
		vl := NewVertexVector2Int(left, l)
		return NewEdge(vm, vl, vm, maximum.HereToVector2(left), index+1)
	} else {
		right := transform.GetTransformedVector2(p.vertices[r])
		vr := NewVertexVector2Int(right, r)
		return NewEdge(vr, vm, vm, right.HereToVector2(maximum), index)
	}
}

func (p *Polygon) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	localn := transform.GetInverseTransformedR(n)
	point := NewVector2FromVector2(p.vertices[0])
	max := localn.DotVector2(p.vertices[0])
	for _, v := range p.vertices {
		projection := localn.DotVector2(v)
		if projection > max {
			point.SetToVector2(v)
			max = projection
		}
	}
	transform.Transform(point)
	return point
}

func (p *Polygon) CreateMass(density float64) *Mass {
	center := new(Vector2)
	var area, I float64
	ac := new(Vector2)
	for _, v := range p.vertices {
		ac.AddVector2(v)
	}
	ac.Multiply(1.0 / float64(len(p.vertices)))
	for i, v := range p.vertices {
		a := i + 1
		if a >= len(p.vertices) {
			a = 0
		}
		p1 := v.DifferenceVector2(ac)
		p2 := p.vertices[a].DifferenceVector2(ac)
		D := p1.CrossVector2(p2)
		triangleArea := 0.5 * D
		area += triangleArea
		center.X += (p1.X + p2.X) * INV_3 * triangleArea
		center.Y += (p1.Y + p2.Y) * INV_3 * triangleArea
		I += triangleArea * (p2.DotVector2(p2) + p2.DotVector2(p1) + p1.DotVector2(p1))
	}
	m := density * area
	center.Multiply(1.0 / area)
	c := center.SumVector2(ac)
	I *= density / 6.0
	I -= m * center.GetMagnitudeSquared()
	return NewMassFromCenterMassInertia(c, m, I)
}

func (p *Polygon) CreateAABBTransform(transform *Transform) *AABB {
	point := transform.GetTransformedVector2(p.vertices[0])
	minX := NewVector2FromVector2(&X_AXIS).DotVector2(point)
	maxX := minX
	minY := NewVector2FromVector2(&Y_AXIS).DotVector2(point)
	maxY := minY
	for _, v := range p.vertices {
		point = transform.GetTransformedVector2(v)
		vx := NewVector2FromVector2(&X_AXIS).DotVector2(point)
		vy := NewVector2FromVector2(&Y_AXIS).DotVector2(point)
		minX = math.Min(minX, vx)
		maxX = math.Max(maxX, vx)
		minY = math.Min(minY, vy)
		maxY = math.Max(maxY, vy)
	}
	return NewAABBFromFloats(minX, minY, maxX, maxY)
}

func (p *Polygon) ContainsVector2(v *Vector2) bool {
	return p.ContainsVector2Transform(v, NewTransform())
}

func (p *Polygon) ProjectVector2(v *Vector2) *Interval {
	return p.ProjectVector2Transform(v, NewTransform())
}

func (p *Polygon) CreateAABB() *AABB {
	return p.CreateAABBTransform(NewTransform())
}

func (p *Polygon) RotateAboutOrigin(theta float64) {
	p.RotateAboutXY(theta, 0, 0)
}

func (p *Polygon) RotateAboutCenter(theta float64) {
	p.RotateAboutXY(theta, p.center.X, p.center.Y)
}

func (p *Polygon) RotateAboutVector2(theta float64, v *Vector2) {
	p.RotateAboutXY(theta, v.X, v.Y)
}

func (p *Polygon) TranslateVector2(v *Vector2) {
	p.TranslateXY(v.X, v.Y)
}
