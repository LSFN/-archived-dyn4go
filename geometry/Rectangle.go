package geometry

import (
	"math"
)

type Rectangle struct {
	Polygon
	width, height float64
}

func NewRectangle(width, height float64) *Rectangle {
	if width <= 0 || height <= 0 {
		panic("Width and height must both be positive")
	}
	r := new(Rectangle)
	r.vertices = []*Vector2{
		NewVector2FromXY(-width*0.5, -height*0.5),
		NewVector2FromXY(width*0.5, -height*0.5),
		NewVector2FromXY(width*0.5, height*0.5),
		NewVector2FromXY(-width*0.5, height*0.5),
	}
	r.normals = []*Vector2{
		NewVector2FromXY(0, -1),
		NewVector2FromXY(1, 0),
		NewVector2FromXY(0, 1),
		NewVector2FromXY(-1, 0),
	}
	r.center = GetAverageCenter(r.vertices...)
	r.radius = r.center.DistanceFromVector2(r.vertices[0])
	r.width = width
	r.height = height
	return r
}

func (r *Rectangle) GetWidth() float64 {
	return r.width
}

func (r *Rectangle) GetHeight() float64 {
	return r.height
}

func (r *Rectangle) GetRotation() float64 {
	return r.normals[1].GetAngleBetween(&X_AXIS)
}

func (r *Rectangle) GetAxes(foci []*Vector2, transform *Transform) []*Vector2 {
	arrLen := 2
	if foci != nil {
		arrLen += len(foci)
	}
	axes := make([]*Vector2, arrLen)
	axes[0] = transform.GetTransformedR(r.normals[1])
	axes[1] = transform.GetTransformedR(r.normals[2])
	n := 2
	for _, f := range foci {
		closest := transform.GetTransformedVector2(r.vertices[0])
		d := f.DistanceSquaredFromVector2(closest)
		var i int
		for i = 1; i < 4; i++ {
			v2 := transform.GetTransformedVector2(r.vertices[i])
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

func (r *Rectangle) GetFoci(transform *Transform) []*Vector2 {
	return nil
}

func (r *Rectangle) ContainsVector2Transform(point *Vector2, transform *Transform) bool {
	p0 := transform.GetInverseTransformedVector2(point)
	c := r.center
	p1 := r.vertices[0]
	p2 := r.vertices[1]
	p4 := r.vertices[3]
	widthSquared := p1.DistanceSquaredFromVector2(p2)
	heightSquared := p1.DistanceSquaredFromVector2(p4)
	projectAxis0 := p1.HereToVector2(p2)
	projectAxis1 := p1.HereToVector2(p4)
	toPoint := c.HereToVector2(p0)
	if toPoint.Project(projectAxis0).GetMagnitudeSquared() <= widthSquared*0.25 {
		if toPoint.Project(projectAxis1).GetMagnitudeSquared() <= heightSquared*0.25 {
			return true
		}
	}
	return false
}

func (r *Rectangle) RotateAboutXY(theta, x, y float64) {
	if !(r.center.X == x && r.center.Y == y) {
		r.center.RotateAboutXY(theta, x, y)
	}
	for i := range r.vertices {
		r.vertices[i].RotateAboutXY(theta, x, y)
		r.normals[i].RotateAboutXY(theta, x, y)
	}
}

func (r *Rectangle) TranslateXY(x, y float64) {
	r.center.AddXY(x, y)
	for _, v := range r.vertices {
		v.AddXY(x, y)
	}
}

func (r *Rectangle) ProjectVector2Transform(axis *Vector2, transform *Transform) *Interval {
	center := transform.GetTransformedVector2(r.center)
	projectAxis0 := transform.GetTransformedR(r.normals[1])
	projectAxis1 := transform.GetTransformedR(r.normals[2])
	c := center.DotVector2(axis)
	e := (r.width*0.5)*math.Abs(projectAxis0.DotVector2(axis)) + (r.height*0.5)*math.Abs(projectAxis1.DotVector2(axis))
	return NewIntervalFromMinMax(c-e, c+e)
}

func (r *Rectangle) GetFarthestFeature(n *Vector2, transform *Transform) *Edge {
	localn := transform.GetInverseTransformedR(n)
	maximum := new(Vector2)
	max := math.Inf(-1)
	index := 0
	for i, v := range r.vertices {
		projection := localn.DotVector2(v)
		if projection > max {
			maximum.SetToVector2(v)
			max = projection
			index = i
		}
	}
	l := index + 1
	if l == len(r.vertices) {
		l = 0
	}
	r2 := index - 1
	if r2 == -1 {
		r2 = len(r.vertices) - 1
	}
	c := index - 1
	if index == 0 {
		c = len(r.vertices) - 1
	}
	leftN := r.normals[c]
	rightN := r.normals[index]
	transform.Transform(maximum)
	vm := NewVertexFromVector2Int(maximum, index)
	if leftN.DotVector2(localn) < rightN.DotVector2(localn) {
		left := transform.GetTransformedVector2(r.vertices[l])
		vl := NewVertexFromVector2Int(left, l)
		return NewEdge(vm, vl, vm, maximum.HereToVector2(left), index+1)
	} else {
		right := transform.GetTransformedVector2(r.vertices[r2])
		vr := NewVertexFromVector2Int(right, r2)
		return NewEdge(vr, vm, vm, right.HereToVector2(maximum), index)
	}
}

func (r *Rectangle) GetFarthestPoint(n *Vector2, transform *Transform) *Vector2 {
	localn := transform.GetInverseTransformedR(n)
	point := NewVector2FromVector2(r.vertices[0])
	max := localn.DotVector2(r.vertices[0])
	for _, v := range r.vertices {
		projection := localn.DotVector2(v)
		if projection > max {
			point.SetToVector2(v)
			max = projection
		}
	}
	transform.Transform(point)
	return point
}

func (r *Rectangle) CreateMass(density float64) *Mass {
	center := new(Vector2)
	var area, I float64
	ac := new(Vector2)
	for _, v := range r.vertices {
		ac.AddVector2(v)
	}
	ac.Multiply(1.0 / float64(len(r.vertices)))
	for i, v := range r.vertices {
		a := i + 1
		if a >= len(r.vertices) {
			a = 0
		}
		p1 := v.DifferenceVector2(ac)
		p2 := r.vertices[a].DifferenceVector2(ac)
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

func (r *Rectangle) CreateAABBTransform(transform *Transform) *AABB {
	point := transform.GetTransformedVector2(r.vertices[0])
	minX := NewVector2FromVector2(&X_AXIS).DotVector2(point)
	maxX := minX
	minY := NewVector2FromVector2(&Y_AXIS).DotVector2(point)
	maxY := minY
	for _, v := range r.vertices {
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

func (r *Rectangle) GetVertices() []*Vector2 {
	return r.vertices
}

func (r *Rectangle) GetNormals() []*Vector2 {
	return r.normals
}

func (r *Rectangle) GetID() string {
	return r.id
}

func (r *Rectangle) GetCenter() *Vector2 {
	return r.center
}

func (r *Rectangle) GetUserData() interface{} {
	return r.userData
}

func (r *Rectangle) SetUserData(data interface{}) {
	r.userData = data
}

func (r *Rectangle) RotateAboutOrigin(theta float64) {
	r.RotateAboutXY(theta, 0, 0)
}

func (r *Rectangle) RotateAboutCenter(theta float64) {
	r.RotateAboutXY(theta, r.center.X, r.center.Y)
}

func (r *Rectangle) RotateAboutVector2(theta float64, v *Vector2) {
	r.RotateAboutXY(theta, v.X, v.Y)
}

func (r *Rectangle) TranslateVector2(v *Vector2) {
	r.TranslateXY(v.X, v.Y)
}

func (r *Rectangle) ProjectVector2(v *Vector2) *Interval {
	return r.ProjectVector2Transform(v, NewTransform())
}

func (r *Rectangle) ContainsVector2(v *Vector2) bool {
	return r.ContainsVector2Transform(v, NewTransform())
}

func (r *Rectangle) CreateAABB() *AABB {
	return r.CreateAABBTransform(NewTransform())
}
