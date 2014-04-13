package geometry

import (
	"math"

	"github.com/LSFN/dyn4go"
)

var TWO_PI = 2.0 * math.Pi
var INV_3 = 1.0 / 3.0
var INV_3_SQRT = 1.0 / math.Sqrt(3.0)

func GetWindingFromList(points []*Vector2) float64 {
	if points == nil || len(points) < 2 {
		panic("List of points must not be nil")
	}
	area := 0.0
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		i2 := i + 1
		if i2 == len(points) {
			i2 = 0
		}
		p2 := points[i2]
		if p1 == nil || p2 == nil {
			panic("Points must not be nil")
		}
		area += p1.CrossVector2(p2)
	}
	return area
}

func GetWinding(points ...*Vector2) float64 {
	return GetWindingFromList(points)
}

func ReverseWindingFromList(points []*Vector2) {
	if points == nil || len(points) < 2 {
		panic("List of points must not be nil or of length less than 2")
	}
	i := 0
	j := len(points) - 1
	for j > i {
		points[i], points[j] = points[j], points[i]
		j--
		i++
	}
}

func ReverseWinding(points ...*Vector2) {
	ReverseWindingFromList(points)
}

func GetAverageCenterFromList(points []*Vector2) *Vector2 {
	if points == nil || len(points) == 0 {
		panic("Need at least one point")
	}
	if len(points) == 1 {
		return NewVector2FromVector2(points[0])
	}
	a := new(Vector2)
	for _, v := range points {
		a.AddVector2(v)
	}
	return a.Multiply(1.0 / float64(len(points)))
}

func GetAverageCenter(points ...*Vector2) *Vector2 {
	return GetAverageCenterFromList(points)
}

func GetAreaWeightedCenterFromList(points []*Vector2) *Vector2 {
	if points == nil || len(points) == 0 {
		panic("Need at least one point")
	}
	if len(points) == 1 {
		return NewVector2FromVector2(points[0])
	}
	ac := GetAverageCenterFromList(points)
	center := new(Vector2)
	var area float64
	for i := range points {
		p1 := points[i]
		var p2 *Vector2
		if i == len(points)-1 {
			p2 = points[0]
		} else {
			p2 = points[i+1]
		}
		p1 = p1.DifferenceVector2(ac)
		p2 = p2.DifferenceVector2(ac)
		triangleArea := 0.5 * p1.CrossVector2(p2)
		area += triangleArea
		center.AddVector2(p1.AddVector2(p2).Multiply(INV_3).Multiply(triangleArea))
	}
	if math.Abs(area) <= dyn4go.Epsilon {
		return NewVector2FromVector2(points[0])
	}
	center.Multiply(1 / area).AddVector2(ac)
	return center
}

func GetAreaWeightedCenter(points ...*Vector2) *Vector2 {
	return GetAverageCenterFromList(points)
}

func CreateCircle(radius float64) *Circle {
	return NewCircle(radius)
}

func CreatePolygon(vertices ...*Vector2) *Polygon {
	if vertices == nil {
		panic("Polygon cannot be created without a few vertices to go on.")
	}
	verts := make([]*Vector2, len(vertices))
	for i, v := range vertices {
		if v != nil {
			verts[i] = NewVector2FromVector2(v)
		} else {
			panic("Polygon cannot be created from nil points")
		}
	}
	return NewPolygon(verts...)
}

func CreatePolygonAtOrigin(vertices ...*Vector2) *Polygon {
	polygon := CreatePolygon(vertices...)
	center := polygon.GetCenter()
	polygon.TranslateXY(-center.X, -center.Y)
	return polygon
}

func CreateUnitCirclePolygon(count int, radius float64) *Polygon {
	return CreateUnitCirclePolygonTheta(count, radius, 0)
}

func CreateUnitCirclePolygonTheta(count int, radius, theta float64) *Polygon {
	if count < 3 {
		panic("Too few vertices")
	}
	if radius <= 0 {
		panic("Radius must be strictly positive")
	}
	return CreatePolygonalCircleTheta(count, radius, theta)
}

func CreateSquare(size float64) *Rectangle {
	if size <= 0 {
		panic("Size of square must be positive")
	}
	return NewRectangle(size, size)
}

func CreateRectangle(width, height float64) *Rectangle {
	return NewRectangle(width, height)
}

func CreateTriangle(p1, p2, p3 *Vector2) *Triangle {
	if p1 == nil || p2 == nil || p3 == nil {
		panic("Triangle cannot be created from nil vertices")
	}
	return NewTriangle(p1, p2, p3)
}

func CreateTriangleAtOrigin(p1, p2, p3 *Vector2) *Triangle {
	triangle := CreateTriangle(p1, p2, p3)
	center := triangle.GetCenter()
	triangle.TranslateXY(-center.X, -center.Y)
	return triangle
}

func CreateRightTriangle(width, height float64) *Triangle {
	return CreateRightTriangleMirror(width, height, false)
}

func CreateRightTriangleMirror(width, height float64, mirror bool) *Triangle {
	if width <= 0 || height <= 0 {
		panic("Width and height must be positive")
	}
	top := NewVector2FromXY(0, height)
	left := NewVector2FromXY(0, 0)
	right := NewVector2FromXY(width, 0)
	if mirror {
		right.X = -right.X
	}
	var triangle *Triangle
	if mirror {
		triangle = NewTriangle(top, right, left)
	} else {
		triangle = NewTriangle(top, left, right)
	}
	center := triangle.GetCenter()
	triangle.TranslateXY(-center.X, -center.Y)
	return triangle
}

func CreateEquilateralTriangle(height float64) *Triangle {
	if height <= 0 {
		panic("Height of equilateral triangle must be positive")
	}
	a := 2 * height * INV_3_SQRT
	return CreateIsoscelesTriangle(a, height)
}

func CreateIsoscelesTriangle(width, height float64) *Triangle {
	if width <= 0 || height <= 0 {
		panic("Width and height must both be strictly positive")
	}
	top := NewVector2FromXY(0, height)
	left := NewVector2FromXY(-width*0.5, 0)
	right := NewVector2FromXY(width*0.5, 0)
	triangle := NewTriangle(top, left, right)
	center := triangle.GetCenter()
	triangle.TranslateXY(-center.X, -center.Y)
	return triangle
}

func CreateSegmentAtOrigin(p1, p2 *Vector2) *Segment {
	segment := CreateSegment(p1, p2)
	center := segment.GetCenter()
	segment.TranslateXY(-center.X, -center.Y)
	return segment
}

func CreateSegment(p1, p2 *Vector2) *Segment {
	if p1 == nil || p2 == nil {
		panic("Cannot create segment from nil points")
	}
	return NewSegment(NewVector2FromVector2(p1), NewVector2FromVector2(p2))
}

func CreateSegmentEnd(end *Vector2) *Segment {
	return CreateSegment(new(Vector2), end)
}

func CreateHorizontalSegment(length float64) *Segment {
	if length <= 0 {
		panic("Length must be strictly positive")
	}
	start := NewVector2FromXY(-length*0.5, 0)
	end := NewVector2FromXY(length*0.5, 0)
	return NewSegment(start, end)
}

func CreateVerticalSegment(length float64) *Segment {
	if length <= 0 {
		panic("Length must be strictly positive")
	}
	start := NewVector2FromXY(0, -length*0.5)
	end := NewVector2FromXY(0, length*0.5)
	return NewSegment(start, end)
}

func CreateCapsule(width, height float64) *Capsule {
	return NewCapsule(width, height)
}

func CreateSlice(radius, theta float64) *Slice {
	return NewSlice(radius, theta)
}

func CreateSliceAtOrigin(radius, theta float64) *Slice {
	slice := NewSlice(radius, theta)
	slice.TranslateXY(-slice.center.X, -slice.center.Y)
	return slice
}

func CreateEllipse(width, height float64) *Ellipse {
	return NewEllipse(width, height)
}

func CreateHalfEllipse(width, height float64) *HalfEllipse {
	return NewHalfEllipse(width, height)
}

func CreateHalfEllipseAtOrigin(width, height float64) *HalfEllipse {
	half := NewHalfEllipse(width, height)
	c := half.GetCenter()
	half.TranslateXY(-c.X, -c.Y)
	return half
}

func CreatePolygonalCircle(count int, radius float64) *Polygon {
	return CreatePolygonalCircleTheta(count, radius, 0)
}

func CreatePolygonalCircleTheta(count int, radius, theta float64) *Polygon {
	if count < 3 {
		panic("Too few vertices")
	}
	if radius <= 0 {
		panic("Radius must be strictly positive")
	}
	pin := TWO_PI / float64(count)
	vertices := make([]*Vector2, count)
	c := math.Cos(pin)
	s := math.Sin(pin)
	t := 0.0
	x := radius
	y := 0.0
	if theta != 0 {
		x = radius * math.Cos(theta)
		y = radius * math.Sin(theta)
	}
	for i := 0; i < count; i++ {
		vertices[i] = NewVector2FromXY(x, y)

		//apply the rotation matrix
		t = x
		x = c*x - s*y
		y = s*t + c*y
	}
	return NewPolygon(vertices...)
}

func CreatePolygonalSlice(count int, radius, theta float64) *Polygon {
	if count < 1 {
		panic("Count must be strictly positive")
	} else if radius <= 0 || theta <= 0 {
		panic("Radius and theta must be strictly positive")
	}
	pin := theta / (float64(count) + 1.0)
	vertices := make([]*Vector2, count+3)
	c := math.Cos(pin)
	s := math.Sin(pin)
	t := 0.0
	x := radius * math.Cos(-theta*0.5)
	y := radius * math.Sin(-theta*0.5)
	vertices[0] = NewVector2FromXY(x, y)
	vertices[count+1] = NewVector2FromXY(x, -y)
	for i := 1; i < count+1; i++ {
		t = x
		x = c*x - s*y
		y = s*t + c*y
		vertices[i] = NewVector2FromXY(x, y)
	}
	vertices[count+2] = new(Vector2)
	return NewPolygon(vertices...)
}

func CreatePolygonalSliceAtOrigin(count int, radius, theta float64) *Polygon {
	p := CreatePolygonalSlice(count, radius, theta)
	p.TranslateVector2(p.GetCenter().GetNegative())
	return p
}

func CreatePolygonalEllipse(count int, width, height float64) *Polygon {
	if count < 4 {
		panic("Cannot create polgonal ellipse from less than 4 vertices")
	}
	if width <= 0 || height <= 0 {
		panic("Width and height must be strictly positive")
	}
	a := width * 0.5
	b := height * 0.5
	n2 := count / 2
	pin2 := math.Pi / float64(n2)
	vertices := make([]*Vector2, n2*2)
	j := 0
	for i := 0; i < n2+1; i++ {
		t := pin2 * float64(i)
		x := a * math.Cos(t)
		y := b * math.Sin(t)
		if i > 0 {
			vertices[len(vertices)-j] = NewVector2FromXY(x, -y)
		}
		vertices[j] = NewVector2FromXY(x, y)
		j++
	}
	return NewPolygon(vertices...)
}

func CreatePolygonalHalfEllipse(count int, width, height float64) *Polygon {
	if count < 1 {
		panic("Cannot create polgonal half-ellipse from less than 4 vertices")
	}
	if width <= 0 || height <= 0 {
		panic("Width and height must be strictly positive")
	}
	a := width * 0.5
	b := height * 0.5
	inc := math.Pi / float64(count+1)
	vertices := make([]*Vector2, count+2)
	vertices[0] = NewVector2FromXY(a, 0)
	vertices[count+1] = NewVector2FromXY(-a, 0)

	for i := 0; i < count+1; i++ {
		t := inc * float64(i)
		x := a * math.Cos(t)
		y := b * math.Sin(t)
		vertices[i] = NewVector2FromXY(x, y)
	}
	return NewPolygon(vertices...)
}

func CreatePolygonalHalfEllipseAtOrigin(count int, width, height float64) *Polygon {
	p := CreatePolygonalHalfEllipse(count, width, height)
	p.TranslateVector2(p.GetCenter().GetNegative())
	return p
}

func CreatePolygonalCapsule(count int, width, height float64) *Polygon {
	if count < 1 {
		panic("Cannot create polygonal capsule from this many vertices")
	}
	if width <= 0 || height <= 0 {
		panic("Width and height must be strictly positive")
	}
	if math.Abs(width-height) < dyn4go.Epsilon {
		return CreatePolygonalCircle(count, width)
	}
	pin := math.Pi / float64(count+1)
	vertices := make([]*Vector2, 4+2*count)
	c := math.Cos(pin)
	s := math.Sin(pin)
	t := 0.0
	major, minor := width, height
	vertical := false
	if width < height {
		major, minor = height, width
		vertical = true
	}
	radius := minor * 0.5
	offset := major*0.5 - radius
	ox := 0.0
	oy := 0.0
	if vertical {
		oy = offset
	} else {
		ox = offset
	}

	n := 0
	ao := math.Pi * 0.5
	if vertical {
		ao = 0
	}
	x := radius * math.Cos(pin-ao)
	y := radius * math.Sin(pin-ao)
	for i := 0; i < count; i++ {
		vertices[n] = NewVector2FromXY(x+ox, y+oy)
		n++
		t = x
		x = c*x - s*y
		y = s*t + c*y
	}
	if vertical {
		vertices[n] = NewVector2FromXY(-radius, oy)
		n++
		vertices[n] = NewVector2FromXY(-radius, -oy)
		n++
	} else {
		vertices[n] = NewVector2FromXY(ox, radius)
		n++
		vertices[n] = NewVector2FromXY(-ox, radius)
		n++
	}
	ao = math.Pi
	if !vertical {
		ao /= 2
	}
	x = radius * math.Cos(pin+ao)
	y = radius * math.Sin(pin+ao)
	for i := 0; i < count; i++ {
		vertices[n] = NewVector2FromXY(x-ox, y-oy)
		n++
		t = x
		x = c*x - s*y
		y = s*t + c*y
	}
	if vertical {
		vertices[n] = NewVector2FromXY(radius, -oy)
		n++
		vertices[n] = NewVector2FromXY(radius, oy)
		n++
	} else {
		vertices[n] = NewVector2FromXY(-ox, -radius)
		n++
		vertices[n] = NewVector2FromXY(ox, -radius)
		n++
	}
	return NewPolygon(vertices...)
}

func Cleanse(points []*Vector2) []*Vector2 {
	if points == nil {
		panic("Cannot cleanse a nil slice of points")
	}
	size := len(points)
	if size == 0 {
		return points
	}
	result := make([]*Vector2, 0, size)
	winding := 0.0
	for i := 0; i < size; i++ {
		point := points[i]
		index := i - 1
		if index < 0 {
			index = size - 1
		}
		prev := points[index]
		index = i + 1
		if index == size {
			index = 0
		}
		next := points[index]
		if point == nil || prev == nil || next == nil {
			panic("Cannot cleanse a slice containing nil elements")
		}
		diff := point.DifferenceVector2(next)
		if diff.IsZero() {
			continue
		}
		prevToPoint := prev.HereToVector2(point)
		pointToNext := point.HereToVector2(next)
		if !prevToPoint.IsZero() {
			cross := prevToPoint.CrossVector2(pointToNext)
			if math.Abs(cross) <= dyn4go.Epsilon {
				continue
			}
		}
		winding += point.CrossVector2(next)
		result = append(result, point)
	}
	if winding <= 0 {
		ReverseWindingFromList(result)
	}
	return result
}

func FlipVector2(polygon Wounder, axis, point *Vector2) *Polygon {
	if polygon == nil || axis == nil {
		panic("Arguments to Flip cannot be nil")
	}
	if point == nil {
		point = polygon.GetCenter()
	}
	if axis.IsZero() {
		panic("axis parameter to Flip must be non-zero")
	}
	axis.Normalize()
	pv := polygon.GetVertices()
	nv := make([]*Vector2, len(pv))
	for i, v0 := range pv {
		v1 := v0.DifferenceVector2(point)
		proj := v1.DotVector2(axis)
		vp := axis.Product(proj)
		rv := vp.AddXY(vp.X-v1.X, vp.Y-v1.Y)
		nv[i] = rv.AddVector2(point)
	}
	if GetWindingFromList(nv) < 0 {
		ReverseWindingFromList(nv)
	}
	return NewPolygon(nv...)
}

func Flip(polygon Wounder, axis *Vector2) *Polygon {
	return FlipVector2(polygon, axis, nil)
}

func FlipAlongTheXAxis(polygon Wounder) *Polygon {
	return FlipVector2(polygon, &X_AXIS, nil)
}

func FlipAlongTheXAxisVector2(polygon Wounder, point *Vector2) *Polygon {
	return FlipVector2(polygon, &X_AXIS, point)
}

func FlipAlongTheYAxis(polygon Wounder) *Polygon {
	return FlipVector2(polygon, &Y_AXIS, nil)
}

func FlipAlongTheYAxisVector2(polygon Wounder, point *Vector2) *Polygon {
	return FlipVector2(polygon, &Y_AXIS, point)
}

func MinkowskiSum(p1, p2 WounderConvexer) *Polygon {
	if p1 == nil || p2 == nil {
		panic("Cannot compute Minkowski sum of nil structs")
	}
	s1t, s1Converted := p1.(*Segment)
	s2t, s2Converted := p2.(*Segment)
	if s1Converted && s2Converted {
		s1 := s1t.vertices[0].HereToVector2(s1t.vertices[1])
		s2 := s2t.vertices[0].HereToVector2(s2t.vertices[1])
		if s1.CrossVector2(s2) <= dyn4go.Epsilon {
			panic("Cannot compute Minkowski sum when cross product of two edges is zero")
		}
	}
	p1v := p1.GetVertices()
	p2v := p2.GetVertices()
	c1 := len(p1v)
	c2 := len(p2v)
	i, j := 0, 0
	min := NewVector2FromXY(math.Inf(1), math.Inf(1))
	for k, v := range p1v {
		if v.Y < min.Y || (v.Y == min.Y && v.X < min.X) {
			min.SetToVector2(v)
			i = k
		}
	}
	min.SetToXY(math.Inf(1), math.Inf(1))
	for k, v := range p2v {
		if v.Y < min.Y || (v.Y == min.Y && v.X < min.X) {
			min.SetToVector2(v)
			j = k
		}
	}
	n1 := c1 + i
	n2 := c2 + j
	sum := make([]*Vector2, 0, c1+c2)
	for i <= n1 && j <= n2 {
		v1s := p1v[i%c1]
		v1e := p1v[(i+1)%c1]
		v2s := p2v[j%c2]
		v2e := p2v[(j+1)%c2]
		sum = append(sum, v1s.SumVector2(v2s))
		e1 := v1s.HereToVector2(v1e)
		e2 := v2s.HereToVector2(v2e)
		a1 := X_AXIS.GetAngleBetween(e1)
		a2 := X_AXIS.GetAngleBetween(e2)
		if a1 < 0 {
			a1 += TWO_PI
		}
		if a2 < 0 {
			a2 += TWO_PI
		}
		if a1 < a2 {
			i++
		} else if a2 < a1 {
			j++
		} else {
			i++
			j++
		}
	}
	return NewPolygon(sum...)
}

func MinkowskiSumCirclePolygonInt(circle *Circle, polygon *Polygon, count int) *Polygon {
	return MinkowskiSumPolygonCircleInt(polygon, circle, count)
}

func MinkowskiSumPolygonCircleInt(polygon *Polygon, circle *Circle, count int) *Polygon {
	if circle == nil {
		panic("Cannot comput the Minkowski sum of a nil Circle")
	}
	return MinkowskiSumPolygonFloat64Int(polygon, circle.radius, count)
}

func MinkowskiSumPolygonFloat64Int(polygon *Polygon, radius float64, count int) *Polygon {
	if polygon == nil {
		panic("Cannot compute Minkowski sum of nil Polygon")
	}
	if radius <= 0 {
		panic("Radius must be strictly positive")
	}
	if count <= 0 {
		panic("Count must be strictly positive")
	}
	var vertices []*Vector2 = polygon.vertices
	var normals []*Vector2 = polygon.normals
	size := len(vertices)
	nVerts := make([]*Vector2, size*2+size*count)
	j := 0
	for i := 0; i < size; i++ {
		v1 := vertices[i]
		k := i + 1
		if k == size {
			k = 0
		}
		v2 := vertices[k]
		normal := normals[i]
		nv1 := normal.Product(radius).AddVector2(v1)
		nv2 := normal.Product(radius).AddVector2(v2)
		var cv1 *Vector2
		if i == 0 {
			tn := normals[size-1]
			cv1 = v1.HereToVector2(tn.Product(radius).AddVector2(v1))
		} else {
			cv1 = v1.HereToVector2(nVerts[j-1])
		}
		cv2 := v1.HereToVector2(nv1)
		theta := cv1.GetAngleBetween(cv2)
		pin := theta / float64(count+1)
		c := math.Cos(pin)
		s := math.Sin(pin)
		t := 0.0
		k = i - 1
		if k < 0 {
			k = size - 1
		}
		sTheta := X_AXIS.GetAngleBetween(normals[k])
		if sTheta < 0 {
			sTheta += TWO_PI
		}
		x := radius * math.Cos(sTheta)
		y := radius * math.Sin(sTheta)
		for k = 0; k < count; k++ {
			t = x
			x = c*x - s*y
			y = s*t + c*y
			nVerts[j] = NewVector2FromXY(x, y).AddVector2(v1)
			j++
		}
		nVerts[j] = nv1
		j++
		nVerts[j] = nv2
		j++
	}
	return NewPolygon(nVerts...)
}

func ScaleCircle(circle *Circle, scale float64) *Circle {
	if circle == nil {
		panic("Cannot scale a nil reference")
	}
	if scale <= 0 {
		panic("Scale factor must be strictly positive")
	}
	return NewCircle(circle.radius * scale)
}

func ScaleCapsule(capsule *Capsule, scale float64) *Capsule {
	if capsule == nil {
		panic("Cannot scale a nil reference")
	}
	if scale <= 0 {
		panic("Scale factor must be strictly positive")
	}
	return NewCapsule(capsule.length*scale, capsule.capRadius*2*scale)
}

func ScaleEllipse(ellipse *Ellipse, scale float64) *Ellipse {
	if ellipse == nil {
		panic("Cannot scale a nil reference")
	}
	if scale <= 0 {
		panic("Scale factor must be strictly positive")
	}
	return NewEllipse(ellipse.width*scale, ellipse.height*scale)
}

func ScaleHalfEllipse(halfEllipse *HalfEllipse, scale float64) *HalfEllipse {
	if halfEllipse == nil {
		panic("Cannot scale a nil reference")
	}
	if scale <= 0 {
		panic("Scale factor must be strictly positive")
	}
	return NewHalfEllipse(halfEllipse.width*scale, halfEllipse.height*scale)
}

func ScaleSlice(slice *Slice, scale float64) *Slice {
	if slice == nil {
		panic("Cannot scale a nil reference")
	}
	if scale <= 0 {
		panic("Scale factor must be strictly positive")
	}
	return NewSlice(slice.sliceRadius*scale, slice.theta)
}

func ScalePolygon(polygon *Polygon, scale float64) *Polygon {
	if polygon == nil {
		panic("Cannot scale a nil reference")
	}
	if scale <= 0 {
		panic("Scale factor must be strictly positive")
	}
	oVertices := polygon.vertices
	vertices := make([]*Vector2, len(oVertices))
	center := polygon.GetCenter()
	for i, v := range oVertices {
		vertices[i] = center.HereToVector2(v).Multiply(scale).AddVector2(center)
	}
	return NewPolygon(vertices...)
}

func ScaleSegment(segment *Segment, scale float64) *Segment {
	if segment == nil {
		panic("Cannot scale a nil reference")
	}
	if scale <= 0 {
		panic("Scale factor must be strictly positive")
	}
	length := segment.GetLength() * scale * 0.5
	n := segment.vertices[0].HereToVector2(segment.vertices[1])
	n.Normalize()
	n.Multiply(length)
	return NewSegment(segment.center.SumXY(n.X, n.Y), segment.center.DifferenceXY(n.X, n.Y))
}
