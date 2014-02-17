package geometry

type Wound struct {
	AbstractShape
	vertices []*Vector2
	normals  []*Vector2
}

/*

func (w *Wound) GetRadius(v *Vector2) {
	r2 := 0.0
	for _, v2 := range w.vertices {
		r2t := v.DistanceSquaredFromVector2(v2)
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

func (w *Wound) GetID() string {
	return w.id
}

func (w *Wound) GetCenter() *Vector2 {
	return w.center
}

func (w *Wound) GetUserData() interface{} {
	return w.userData
}

func (w *Wound) SetUserData(data interface{}) {
	w.userData = data
}

func (w *Wound) RotateAboutOrigin(theta float64) {
	w.RotateAboutXY(theta, 0, 0)
}

func (w *Wound) RotateAboutCenter(theta float64) {
	w.RotateAboutXY(theta, w.center.X, w.center.Y)
}

func (w *Wound) RotateAboutVector2(theta float64, v *Vector2) {
	w.RotateAboutXY(theta, v.X, v.Y)
}

func (w *Wound) RotateAboutXY(theta, x, y float64) {
	if !(w.center.X == x && w.center.Y == y) {
		w.center.RotateAboutXY(theta, x, y)
	}
}

func (w *Wound) TranslateXY(x, y float64) {
	w.center.AddXY(x, y)
}

func (w *Wound) TranslateVector2(v *Vector2) {
	w.TranslateXY(v.X, v.Y)
}

func (w *Wound) Project(v *Vector2) *Interval {
	return w.ProjectTransform(v, NewTransform())
}

func (w *Wound) Contains(v *Vector2) bool {
	return w.ContainsTransform(v, NewTransform())
}

func (w *Wound) CreateAABB() *AABB {
	return w.CreateAABBTransform(NewTransform())
}

*/
