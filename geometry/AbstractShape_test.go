package geometry2

import (
	"testing"

	"code.google.com/p/uuid"
	"github.com/LSFN/dyn4go"
)

type testShape struct {
	AbstractShape
}

func newTestShape() *testShape {
	ts := new(testShape)
	ts.id = uuid.New()
	return ts
}

func (t *testShape) ContainsVector2Transform(point *Vector2, transform *Transform) bool { return false }
func (t *testShape) CreateAABBTransform(transform *Transform) *AABB                     { return nil }
func (t *testShape) CreateMass(density float64) *Mass                                   { return new(Mass) }
func (t *testShape) GetRadiusVector2(v *Vector2) float64                                { return 0.0 }
func (t *testShape) ProjectVector2Transform(n *Vector2, transform *Transform) *Interval { return nil }
func (t *testShape) CreateAABB() *AABB                                                  { return nil }
func (t *testShape) ProjectVector2(v *Vector2) *Interval                                { return nil }
func (t *testShape) ContainsVector2(v *Vector2) bool                                    { return false }
func (t *testShape) RotateAboutOrigin(theta float64)                                    { t.RotateAboutXY(theta, 0, 0) }
func (t *testShape) RotateAboutCenter(theta float64)                                    { t.RotateAboutXY(theta, t.center.X, t.center.Y) }
func (t *testShape) RotateAboutVector2(theta float64, v *Vector2)                       { t.RotateAboutXY(theta, v.X, v.Y) }
func (t *testShape) TranslateVector2(v *Vector2)                                        { t.TranslateXY(v.X, v.Y) }

func TestAbstractShapeInterfaces(t *testing.T) {
	s := newTestShape()
	var _ Shaper = s
}

func TestAbstractShapeGetID(t *testing.T) {
	s := newTestShape()
	dyn4go.AssertNotEqual(t, s.GetID(), "")
}

func TestAbstractShapeSetUserData(t *testing.T) {
	s := newTestShape()
	dyn4go.AssertTrue(t, s.GetUserData() == nil)
	obj := "hello"
	s.SetUserData(obj)
	dyn4go.AssertFalse(t, s.GetUserData() == nil)
	dyn4go.AssertEqual(t, s.GetUserData(), obj)
}
