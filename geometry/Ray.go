package geometry2

type Ray struct {
	start, direction *Vector2
}

func NewRayFromFloat(direction float64) *Ray {
	return NewRayFromVector2(NewVector2FromDirection(direction))
}

func NewRayFromVector2(direction *Vector2) *Ray {
	return NewRayFromVector2Vector2(new(Vector2), direction)
}

func NewRayFromVector2Float(start *Vector2, direction float64) *Ray {
	return NewRayFromVector2Vector2(start, NewVector2FromDirection(direction))
}

func NewRayFromVector2Vector2(start, direction *Vector2) *Ray {
	if start == nil || direction == nil || direction.IsZero() {
		panic("Ray cannot be created from nil arguments or zero direction")
	}
	r := new(Ray)
	r.start = start
	r.direction = direction
	return r
}

func (r *Ray) GetStart() *Vector2 {
	return r.start
}

func (r *Ray) SetStart(start *Vector2) {
	if start == nil {
		panic("Cannot set Ray to have a nil start point")
	}
	r.start = start
}

func (r *Ray) GetDirectionFloat() float64 {
	return r.direction.GetDirection()
}

func (r *Ray) SetDirectionFloat(direction float64) {
	r.direction = NewVector2FromDirection(direction)
}

func (r *Ray) GetDirectionVector2() *Vector2 {
	return r.direction
}

func (r *Ray) SetDirectionVector2(direction *Vector2) {
	if direction == nil || direction.IsZero() {
		panic("Ray direction cannot be created from nil argument or be zero")
	}
	r.direction = direction
}
