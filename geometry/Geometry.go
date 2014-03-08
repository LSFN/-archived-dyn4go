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

/*
func CreateSquare(size float64) *Rectangle {

}*/

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
