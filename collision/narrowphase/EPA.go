package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
	"math"

	"github.com/LSFN/dyn4go"
)

type EPAEdge struct {
	distance float64
	normal   *geometry.Vector2
	index    int
}

type EPA struct {
	maxIterations   int
	distanceEpsilon float64
}

func NewEPA() *EPA {
	e := new(EPA)
	e.maxIterations = 100
	e.distanceEpsilon = math.Sqrt(dyn4go.Epsilon)
	return e
}

func (e *EPA) GetPenetration(simplex *[]*geometry.Vector2, minkowskiSum *MinkowskiSum, penetration *Penetration) {
	winding := e.getWinding(*simplex)
	var point *geometry.Vector2
	var edge *EPAEdge
	for i := 0; i < e.maxIterations; i++ {
		edge = e.findClosestEdge(*simplex, winding)
		point = minkowskiSum.Support(edge.normal)
		projection := point.DotVector2(edge.normal)
		if projection-edge.distance < e.distanceEpsilon {
			penetration.normal = edge.normal
			penetration.depth = projection
			return
		}
		*simplex = append((*simplex)[:edge.index], append([]*geometry.Vector2{point}, (*simplex)[edge.index:]...)...)
	}
	penetration.normal = edge.normal
	penetration.depth = point.DotVector2(edge.normal)
}

func (e *EPA) findClosestEdge(simplex []*geometry.Vector2, winding int) *EPAEdge {
	size := len(simplex)
	edge := new(EPAEdge)
	edge.distance = math.MaxFloat64
	edge.normal = new(geometry.Vector2)
	normal := new(geometry.Vector2)
	for i := 0; i < size; i++ {
		j := i + 1
		if j == size {
			j = 0
		}
		a := simplex[i]
		b := simplex[j]
		normal.SetToXY(b.X-a.X, b.Y-a.Y)
		if winding < 0 {
			normal.Right()
		} else {
			normal.Left()
		}
		normal.Normalize()
		d := math.Abs(a.DotVector2(normal))
		if d < edge.distance {
			edge.distance = d
			edge.normal.SetToVector2(normal)
			edge.index = j
		}
	}
	return edge
}

func (e *EPA) getWinding(simplex []*geometry.Vector2) int {
	size := len(simplex)
	for i := 0; i < size; i++ {
		j := i + 1
		if j == size {
			j = 0
		}
		a := simplex[i]
		b := simplex[j]
		cross := a.CrossVector2(b)
		if cross > 0 {
			return 1
		} else if cross < 0 {
			return -1
		}
	}
	return 0
}

func (e *EPA) GetMaxIterations() int {
	return e.GetMaxIterations()
}

func (e *EPA) SetMaxIterations(maxIterations int) {
	if maxIterations < 5 {
		panic("Too few iterations")
	}
	e.maxIterations = maxIterations
}

func (e *EPA) GetDistanceEpsilon() float64 {
	return e.distanceEpsilon
}

func (e *EPA) SetDistanceEpsilon(distanceEpsilon float64) {
	if distanceEpsilon <= 0 {
		panic("Distance epsilon is malformed")
	}
	e.distanceEpsilon = distanceEpsilon
}
