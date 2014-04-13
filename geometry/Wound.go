package geometry2

import (
	"math"
)

type Wound struct {
	AbstractShape
	vertices []*Vector2
	normals  []*Vector2
}

func (w *Wound) GetRadiusVector2(center *Vector2) float64 {
	r2 := 0.0
	for _, v := range w.vertices {
		r2t := center.DistanceSquaredFromVector2(v)
		r2 = math.Max(r2, r2t)
	}
	return math.Sqrt(r2)
}

func (w *Wound) GetVertices() []*Vector2 {
	return w.vertices
}

func (w *Wound) GetNormals() []*Vector2 {
	return w.normals
}
