package geometry

import (
	"testing"

	"code.google.com/p/uuid"
	"github.com/LSFN/dyn4go"
)

type TestShape AbstractShape

func NewTestShape() *TestShape {
	ts := new(TestShape)
	ts.id = uuid.New()
	return ts
}

func (t *TestShape) ContainsTransform(point *Vector2, transform *Transform) bool { return false }
func (t *TestShape) CreateAABBTransform(transform *Transform) *AABB              { return nil }
func (t *TestShape) CreateMass(density float64) *Mass                            { return new(Mass) }
func (t *TestShape) GetRadius(center *Vector2) float64                           { return 0.0 }
func (t *TestShape) ProjectTransform(n *Vector2, transform *Transform) *Interval { return nil }
func (t *TestShape) GetID() string {
	return t.id
}

func (t *TestShape) GetCenter() *Vector2 {
	return t.center
}

func (t *TestShape) GetUserData() interface{} {
	return t.userData
}

func (t *TestShape) SetUserData(data interface{}) {
	t.userData = data
}

func (t *TestShape) RotateAboutOrigin(theta float64) {
	t.RotateAboutXY(theta, 0, 0)
}

func (t *TestShape) RotateAboutCenter(theta float64) {
	t.RotateAboutXY(theta, t.center.X, t.center.Y)
}

func (t *TestShape) RotateAboutVector2(theta float64, v *Vector2) {
	t.RotateAboutXY(theta, v.X, v.Y)
}

func (t *TestShape) RotateAboutXY(theta, x, y float64) {
	if !(t.center.X == x && t.center.Y == y) {
		t.center.RotateAboutXY(theta, x, y)
	}
}

func (t *TestShape) TranslateXY(x, y float64) {
	t.center.AddXY(x, y)
}

func (t *TestShape) TranslateVector2(v *Vector2) {
	t.TranslateXY(v.X, v.Y)
}

func (t *TestShape) Project(v *Vector2) *Interval {
	return t.ProjectTransform(v, NewTransform())
}

func (t *TestShape) Contains(v *Vector2) bool {
	return t.ContainsTransform(v, NewTransform())
}

func (t *TestShape) CreateAABB() *AABB {
	return t.CreateAABBTransform(NewTransform())
}

func TestGetID(t *testing.T) {
	s := NewTestShape()
	dyn4go.AssertNotEqual(t, s.GetID(), "")
}

func TestSetUserData(t *testing.T) {
	s := new(TestShape)
	dyn4go.AssertNil(t, s.GetUserData())
	obj := "hello"
	s.SetUserData(obj)
	dyn4go.AssertNotNil(t, s.GetUserData())
	dyn4go.AssertEqual(t, s.GetUserData(), obj)
}
