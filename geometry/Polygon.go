package geometry

import (
	"math"

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
	for i, v := range vertices {
		var p0, p1, p2 *Vector2
		if i-1 < 0 {
			p0 = vertices[len(vertices)-1]
		} else {
			p0 = vertices[i]
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
		if cross > 0 {
			cross = 1
		} else if cross < 0 {
			cross = -1
		} else {
			cross = 0
		}
		if sign != 0 && cross != sign {
			panic("Convex polygons are not allowed")
		}
		sign = cross
	}
	if area < 0 {
		panic("Invalid polygon winding")
	}
	p := new(Polygon)
	p.vertices = vertices
	p.normals = make([]*Vector2, len(vertices))
	for i, p1 := range p.vertices {
		var p2 *Vector2
		if i+1 == len(p.vertices) {
			p2 = vertices[0]
		} else {
			p2 = vertices[i+1]
		}
		n := p1.HereToVector2(p2).Left().Normalize()
		p.normals[i] = n
	}
	p.center = GetAreaWeightedCenterFromList(p.vertices)
	r2 := 0.0
	for _, v := range vertices {
		r2 = math.Max(r2, p.center.DistanceSquaredFromVector2(v))
	}
	p.radius = math.Sqrt(r2)
}

func (p *Polygon) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	arrLen := len(p.vertices)
	if foci != nil {
		arrLen += len(foci)
	}
	axes = make([]*Vector2, arrLen)
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
		axis := f.HereToVector2(closest).Normalize()
		axes[n] = axis
		n++
	}
	return axes
}

func (p *Polygon) GetFoci(transform *Transform) []*Vector2 {
	return nil
}

func (p *Polygon) Contains(point *Vector2, transform *Transform) bool {
	p0 := transform.GetInverseTransformedVector2(point)
	p1 := p.vertices[0]
	p2 := p.vertices[1]
	last := GetLocation(p0, p1, p2)
	for i := range p.vertices {
		p1 = p2
		if i+1 == size {
			p2 = p.vertices[0]
		} else {
			p2 = p.vertices[i+1]
		}
		if p0.EqualsVector2(p1) {
			return true
		}
		if last * GetLocation(p0, p1, p2) {
			return false
		}
	}
	return true
}

func (p *Polygon) Rotate(theta, x, y float64) {
	p.Rotate(theta, x, y)
	for i := range p.vertices {
		p.vertices[i].RotateAboutXY(theta, x, y)
		p.normals[i].RotateAboutXY(theta, x, y)
	}
}

func (p *Polygon) Translate(x, y float64) {
	p.Translate(x, y)
	for _, v := range p.vertices {
		v.AddXY(x, y)
	}
}

func (p *Polygon) Project(n *Vector2, transform *Transform) *Interval {
	point := transform.GetInverseTransformedVector2(p.vertices[0])
	min := n.DotVector2(point)
	max := min
	for _, a := range p.vertices {
		v := n.DotVector2(a)
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return NewIntervalFromMinMax(min, max)
}

func (p *Polygon) GetFarthestFeature(n *Vector2, transform *Transform) *Edge {
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
	if c == 0 {
		c = len(p.vertices) - 1
	}
	leftN := p.normals[c]
	rightN := p.normals[index]
	transform.TranslateVector2(maximum)
	vm := NewVertexFromPointIndex(maximum, index)
	if leftN.DotVector2(localn) < rightN.DotVector2(localn) {
		left := transform.GetTransformedVector2(p.vertices[l])
		vl = NewVertexFromPointIndex(left, l)
		return NewEdge(vm, vl, vm, maximum.HereToVector2(left), index+1)
	} else {
		right := transform.GetTransformedVector2(p.vertices[r])
		vr = NewVertexFromPointIndex(right, r)
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
	ac.Multiply(1.0 / n)
	for i, v := range p.vertices {
		a := i + 1
		if a >= n {
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
	return NewMassFromCenterMassInertia(c, m, i)
}

func (p *Polygon) CreateAABB(transform *Transform) *AABB {
	point := transform.GetTransformedVector2(p.vertices[0])
	minX := NewVector2FromVector2(X_AXIS).DotVector2(point)
	maxX := minX
	minY := NewVector2FromVector2(Y_AXIS).DotVector2(point)
	maxY := minY
	for _, v := range p.vertices {
		point = transform.GetTransformedVector2(v)
		vx := NewVector2FromVector2(X_AXIS).DotVector2(point)
		vy := NewVector2FromVector2(Y_AXIS).DotVector2(point)
		minX = math.Min(minX, vx)
		maxX = math.Max(maxX, vx)
		minY = math.Min(minY, vy)
		maxY = math.Max(maxY, vy)
	}
	return NewAABBFromFloats(minX, minY, maxX, maxY)
}
