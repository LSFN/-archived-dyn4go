package geometry

type Wound struct {
	AbstractShape
	vertices []*Vector2
	normals  []*Vector2
}

func (w *Wound) GetRadius(v *Vector2) {
	r2 := 0.0
	for _, v2 := range w.vertices {
		r2t := v.DistanceSquaredFromVector2(v2)
		r2 = math.Max(r2, r2t)
	}
	return math.Sqrt(r2)
}

func (w *Wound) GetVertices() {
	return w.vertices
}

func (w *Wound) GetNormals() {
	return w.normals
}
