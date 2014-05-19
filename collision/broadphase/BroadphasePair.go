package broadphase

type BroadphasePair struct {
	a, b interface{}
}

func NewBroadphasePair(a, b interface{}) *BroadphasePair {
	c := new(BroadphasePair)
	c.a = a
	c.b = b
	return c
}

func (b *BroadphasePair) GetA() interface{} {
	return b.a
}

func (b *BroadphasePair) SetA(val interface{}) {
	b.a = val
}
func (b *BroadphasePair) GetB() interface{} {
	return b.b
}

func (b *BroadphasePair) SetB(val interface{}) {
	b.b = val
}
