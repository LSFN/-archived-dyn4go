package collision

import (
	"github.com/LSFN/dyn4go/geometry"
)

type AxisAlignedBounds struct {
	AbstractBounds
	aabb *geometry.AABB
}

func NewAxisAlignedBounds(width, height float64) *AxisAlignedBounds {
	a := new(AxisAlignedBounds)
	if width <= 0 || height <= 0 {
		panic("Width and height must be strictly positive")
	}
	InitAbstractBounds(&a.AbstractBounds)
	w2 := width * 0.5
	h2 := height * 0.5
	a.aabb = geometry.NewAABBFromFloats(-w2, -h2, w2, h2)
	return a
}

func (a *AxisAlignedBounds) IsOutside(collidable Collider) bool {
	tx := a.transform.GetTranslation()
	aabbBounds := a.aabb.GetTranslated(tx)
	aabbBody := collidable.CreateAABB()
	return !aabbBounds.Overlaps(aabbBody)
}

func (a *AxisAlignedBounds) GetBounds() *geometry.AABB {
	return a.aabb.GetTranslated(a.transform.GetTranslation())
}

func (a *AxisAlignedBounds) GetTranslation() *geometry.Vector2 {
	return a.transform.GetTranslation()
}

func (a *AxisAlignedBounds) GetWidth() float64 {
	return a.aabb.GetWidth()
}

func (a *AxisAlignedBounds) GetHeight() float64 {
	return a.aabb.GetHeight()
}

func (a *AxisAlignedBounds) RotateAboutOrigin(theta float64)                           {}
func (a *AxisAlignedBounds) RotateAboutVector2(theta float64, point *geometry.Vector2) {}
func (a *AxisAlignedBounds) RotateAboutXY(theta, x, y float64)                         {}
