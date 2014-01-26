package geometry

type Triangle Polygon

func (t *Triangle) Contains(point *Vector2, transform *Transform) bool {
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
