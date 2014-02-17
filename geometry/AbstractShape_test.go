package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

type TestShape AbstractShape

func (t *TestShape) Contains(point *Vector2, transform *Transform) bool { return false }
func (t *TestShape) CreateAABB(transform *Transform) *AABB              { return nil }
func (t *TestShape) CreateMass(density float64) *Mass                   { return new(Mass) }
func (t *TestShape) GetRadius(center *Vector2) float64                  { return 0.0 }
func (t *TestShape) Project(n *Vector2, transform *Transform) *Interval { return nil }

func TestGetID(t *testing.T) {
	s := new(TestShape)
	dyn4go.AssertNotEqual(s.GetID(), "")
}

func TestSetUserData(t *testing.T) {
	s := new(TestShape)
	dyn4go.AssertNil(s.GetUserData())
	obj := "hello"
	s.SetUserData(obj)
	dyn4go.AssertNotNil(s.GetUserData())
	dyn4go.AssertEqual(s.GetUserData(), obj)
}
