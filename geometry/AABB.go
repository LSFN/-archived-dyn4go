package geometry

import (
	"math"
)

type AABB struct {
	min, max *Vector2
}

func NewAABBFromFloats(minX, minY, maxX, maxY float64) *AABB {
	a := new(AABB)
	a.min.X = minX
	a.min.Y = minY
	a.max.X = maxX
	a.max.Y = maxY
	return a
}

func NewAABBFromVector2(min, max *Vector2) *AABB {
	a := new(AABB)
	if min.X > max.X || min.Y > max.Y {
		panic("min and max are invalid")
	}
	a.min = NewVector2FromVector2(min)
	a.max = NewVector2FromVector2(max)
	return a
}

func NewAABBFromRadius(radius float64) *AABB {
	return NewAABBFromRadiusAndCenter(nil, radius)
}

func NewAABBFromRadiusAndCenter(center *Vector2, radius float64) *AABB {
	if radius < 0 {
		panic("invalid radius")
	}
	a := new(AABB)
	if center == nil {
		a.min = NewVector2FromXY(-radius, -radius)
		a.max = NewVector2FromXY(radius, radius)
	} else {
		a.min = NewVector2FromXY(center.X-radius, center.Y-radius)
		a.max = NewVector2FromXY(center.X+radius, center.Y+radius)
	}
	return a
}

func NewAABBFromAABB(a *AABB) *AABB {
	a2 := new(AABB)
	a2.min = NewVector2FromVector2(a.min)
	a2.max = NewVector2FromVector2(a.max)
	return a2
}

func (a *AABB) Translate(translation *Vector2) {
	a.max.AddVector2(translation)
	a.min.AddVector2(translation)
}

func (a *AABB) GetTranslated(translation *Vector2) *AABB {
	a2 := new(AABB)
	a2.min.SumVector2(translation)
	a2.max.SumVector2(translation)
	return a2
}

func (a *AABB) GetWidth() float64 {
	return a.max.X - a.min.X
}

func (a *AABB) GetHeight() float64 {
	return a.max.Y - a.min.Y
}

func (a *AABB) GetPerimeter() float64 {
	return 2 * (a.max.X - a.min.X + a.max.Y - a.min.Y)
}

func (a *AABB) GetArea() float64 {
	return (a.max.X - a.min.X) * (a.max.Y - a.min.Y)
}

func (a *AABB) Union(a2 *AABB) {
	a.min.X = math.Min(a.min.X, a2.min.X)
	a.min.Y = math.Min(a.min.Y, a2.min.Y)
	a.max.X = math.Max(a.max.X, a2.max.X)
	a.max.Y = math.Max(a.max.Y, a2.max.Y)
}

func (a *AABB) GetUnion(a2 *AABB) *AABB {
	min := new(Vector2)
	max := new(Vector2)
	min.X = math.Min(a.min.X, a2.min.X)
	min.Y = math.Min(a.min.Y, a2.min.Y)
	max.X = math.Max(a.max.X, a2.max.X)
	max.Y = math.Max(a.max.Y, a2.max.Y)
	return NewAABBFromVector2(min, max)
}

func (a *AABB) Intersection(a2 *AABB) {
	a.min.X = math.Max(a.min.X, a2.min.X)
	a.min.Y = math.Max(a.min.Y, a2.min.Y)
	a.max.X = math.Min(a.max.X, a2.max.X)
	a.max.Y = math.Min(a.max.Y, a2.max.Y)
	if a.min.X > a.max.X || a.min.Y > a.max.Y {
		a.min.X = 0
		a.min.Y = 0
		a.max.X = 0
		a.max.Y = 0
	}
}

func (a *AABB) GetIntersection(a2 *AABB) *AABB {
	min := new(Vector2)
	max := new(Vector2)
	min.X = math.Max(a.min.X, a2.min.X)
	min.Y = math.Max(a.min.Y, a2.min.Y)
	max.X = math.Min(a.max.X, a2.max.X)
	max.Y = math.Min(a.max.Y, a2.max.Y)
	if min.X > max.X || min.Y > max.Y {
		min.X = 0
		min.Y = 0
		max.X = 0
		max.Y = 0
	}
	return NewAABBFromVector2(min, max)
}

func (a *AABB) Expand(expansion float64) {
	e := expansion * 0.5
	a.min.X -= e
	a.min.Y -= e
	a.max.X += e
	a.max.Y += e
	if expansion < 0 {
		if a.min.X > a.max.X {
			mid := (a.min.X + a.max.X) * 0.5
			a.min.X = mid
			a.max.X = mid
		}
		if a.min.Y > a.max.Y {
			mid := (a.min.Y + a.max.Y) * 0.5
			a.min.Y = mid
			a.max.Y = mid
		}
	}
}

func (a *AABB) GetExpanded(expansion float64) *AABB {
	e := expansion * 0.5
	minX := a.min.X - e
	minY := a.min.Y - e
	maxX := a.max.X + e
	maxY := a.max.Y + e
	if expansion < 0 {
		if minX > maxX {
			mid := (minX + maxX) * 0.5
			minX = mid
			maxX = mid
		}
		if minY > maxY {
			mid := (minY + maxY) * 0.5
			minY = mid
			maxY = mid
		}
	}
	return NewAABBFromFloats(minX, minY, maxX, maxY)
}

func (a *AABB) Overlaps(a2 *AABB) bool {
	return !(a.min.X > a2.max.X || a.max.X < a2.min.X || a.min.Y > a2.max.Y || a.max.Y < a2.min.Y)
}

func (a *AABB) ContainsAABB(a2 *AABB) bool {
	return a.min.X <= a2.min.X && a.max.X >= a2.max.X && a.min.Y <= a2.min.Y && a.max.Y >= a2.max.Y
}

func (a *AABB) ContainsVector2(v *Vector2) bool {
	return a.min.X <= v.X && a.max.X >= v.X && a.min.Y <= v.Y && a.max.Y >= v.Y
}

func (a *AABB) ContainsXY(x, y float64) bool {
	return a.min.X <= x && a.max.X >= x && a.min.Y <= y && a.max.Y >= y
}

func (a *AABB) IsDegenerate() bool {
	return a.min.X == a.max.X || a.min.Y == a.max.Y
}

func (a *AABB) IsDegenerateWithError(e float64) bool {
	return math.Abs(a.max.X-a.min.X) <= e || math.Abs(a.max.Y-a.min.Y) <= e
}

func (a *AABB) getMinX() float64 {
	return a.min.X
}

func (a *AABB) getMaxX() float64 {
	return a.max.X
}

func (a *AABB) getMinY() float64 {
	return a.min.Y
}

func (a *AABB) getMaxY() float64 {
	return a.max.Y
}
