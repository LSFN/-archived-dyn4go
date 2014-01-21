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
