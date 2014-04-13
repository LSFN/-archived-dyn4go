package geometry2

import (
	"math"
)

type Triangle Polygon

func NewTriangle(p1, p2, p3 *Vector2) *Triangle {
	t := new(Triangle)
	p := NewPolygon(p1, p2, p3)
	t.vertices = p.vertices
	t.normals = p.normals
	t.id = p.id
	t.center = p.center
	t.radius = p.radius
	t.userData = p.userData
	return t
}

func (t *Triangle) ContainsVector2Transform(point *Vector2, transform *Transform) bool {
	p := transform.GetInverseTransformedVector2(point)
	p1 := t.vertices[0]
	p2 := t.vertices[1]
	p3 := t.vertices[2]
	ab := p1.HereToVector2(p2)
	ac := p1.HereToVector2(p3)
	pa := p1.HereToVector2(p)
	dot00 := ac.DotVector2(ac)
	dot01 := ac.DotVector2(ab)
	dot02 := ac.DotVector2(pa)
	dot11 := ab.DotVector2(ab)
	dot12 := ab.DotVector2(pa)
	invD := 1.0 / (dot00*dot11 - dot01*dot01)
	u := (dot11*dot02 - dot01*dot12) * invD
	v := (dot00*dot12 - dot01*dot02) * invD
	return u > 0 && v > 0 && (u+v <= 1)
}

func (t *Triangle) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	arrLen := len(t.vertices)
	if foci != nil {
		arrLen += len(foci)
	}
	axes := make([]*Vector2, arrLen)
	n := 0
	for _, v := range t.normals {
		axes[n] = transform.GetTransformedR(v)
		n++
	}
	for _, f := range foci {
		closest := transform.GetTransformedVector2(t.vertices[0])
		d := f.DistanceSquaredFromVector2(closest)
		for _, v := range t.vertices {
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

func (t *Triangle) GetFoci(transform *Transform) []*Vector2 {
	return nil
}

func (t *Triangle) RotateAboutXY(theta, x, y float64) {
	if !(t.center.X == x && t.center.Y == y) {
		t.center.RotateAboutXY(theta, x, y)
	}
	for i := range t.vertices {
		t.vertices[i].RotateAboutXY(theta, x, y)
		t.normals[i].RotateAboutXY(theta, x, y)
	}
}

func (t *Triangle) TranslateXY(x, y float64) {
	t.center.AddXY(x, y)
	for _, v := range t.vertices {
		v.AddXY(x, y)
	}
}

func (t *Triangle) ProjectVector2Transform(n *Vector2, transform *Transform) *Interval {
	point := transform.GetInverseTransformedVector2(t.vertices[0])
	min := n.DotVector2(point)
	max := min
	for i, a := range t.vertices {
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

func (t *Triangle) GetFarthestFeature(n *Vector2, transform *Transform) Featurer {
	localn := transform.GetInverseTransformedR(n)
	maximum := new(Vector2)
	max := math.Inf(-1)
	index := 0
	for i, v := range t.vertices {
		projection := localn.DotVector2(v)
		if projection > max {
			maximum.SetToVector2(v)
			max = projection
			index = i
		}
	}
	l := index + 1
	if l == len(t.vertices) {
		l = 0
	}
	r := index - 1
	if r == -1 {
		r = len(t.vertices) - 1
	}
	c := index - 1
	if index == 0 {
		c = len(t.vertices) - 1
	}
	leftN := t.normals[c]
	rightN := t.normals[index]
	transform.Transform(maximum)
	vm := NewVertexVector2Int(maximum, index)
	if leftN.DotVector2(localn) < rightN.DotVector2(localn) {
		left := transform.GetTransformedVector2(t.vertices[l])
		vl := NewVertexVector2Int(left, l)
		return NewEdge(vm, vl, vm, maximum.HereToVector2(left), index+1)
	} else {
		right := transform.GetTransformedVector2(t.vertices[r])
		vr := NewVertexVector2Int(right, r)
		return NewEdge(vr, vm, vm, right.HereToVector2(maximum), index)
	}
}

func (t *Triangle) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	localn := transform.GetInverseTransformedR(n)
	point := NewVector2FromVector2(t.vertices[0])
	max := localn.DotVector2(t.vertices[0])
	for _, v := range t.vertices {
		projection := localn.DotVector2(v)
		if projection > max {
			point.SetToVector2(v)
			max = projection
		}
	}
	transform.Transform(point)
	return point
}

func (t *Triangle) CreateMass(density float64) *Mass {
	center := new(Vector2)
	var area, I float64
	ac := new(Vector2)
	for _, v := range t.vertices {
		ac.AddVector2(v)
	}
	ac.Multiply(1.0 / float64(len(t.vertices)))
	for i, v := range t.vertices {
		a := i + 1
		if a >= len(t.vertices) {
			a = 0
		}
		p1 := v.DifferenceVector2(ac)
		p2 := t.vertices[a].DifferenceVector2(ac)
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

func (t *Triangle) CreateAABBTransform(transform *Transform) *AABB {
	point := transform.GetTransformedVector2(t.vertices[0])
	minX := NewVector2FromVector2(&X_AXIS).DotVector2(point)
	maxX := minX
	minY := NewVector2FromVector2(&Y_AXIS).DotVector2(point)
	maxY := minY
	for _, v := range t.vertices {
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

func (t *Triangle) ContainsVector2(v *Vector2) bool {
	return t.ContainsVector2Transform(v, NewTransform())
}

func (t *Triangle) ProjectVector2(v *Vector2) *Interval {
	return t.ProjectVector2Transform(v, NewTransform())
}

func (t *Triangle) CreateAABB() *AABB {
	return t.CreateAABBTransform(NewTransform())
}

func (t *Triangle) RotateAboutOrigin(theta float64) {
	t.RotateAboutXY(theta, 0, 0)
}

func (t *Triangle) RotateAboutCenter(theta float64) {
	t.RotateAboutXY(theta, t.center.X, t.center.Y)
}

func (t *Triangle) RotateAboutVector2(theta float64, v *Vector2) {
	t.RotateAboutXY(theta, v.X, v.Y)
}

func (t *Triangle) TranslateVector2(v *Vector2) {
	t.TranslateXY(v.X, v.Y)
}
