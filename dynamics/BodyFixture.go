package dynamics

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/geometry"
)

type BodyFixture struct {
	collision.Fixture
	density, friction, restitution float64
}

var _ collision.Fixturer = new(BodyFixture)

func NewBodyFixture(shape geometry.Convexer) *BodyFixture {
	b := new(BodyFixture)
	b.Fixture = *collision.NewFixture(shape)
	b.density = 1
	b.friction = 0.2
	b.restitution = 0
	return b
}

func (b *BodyFixture) SetDensity(density float64) {
	if density <= 0 {
		panic("Density must be strictly positive")
	}
	b.density = density
}

func (b *BodyFixture) GetDensity() float64 {
	return b.density
}

func (b *BodyFixture) SetFriction(friction float64) {
	if friction <= 0 {
		panic("Friction must be strictly positive")
	}
	b.friction = friction
}

func (b *BodyFixture) GetFriction() float64 {
	return b.friction
}

func (b *BodyFixture) SetRestitution(restitution float64) {
	if restitution < 0 {
		panic("Restitution must not be negative")
	}
	b.restitution = restitution
}

func (b *BodyFixture) GetRestitution() float64 {
	return b.restitution
}

func (b *BodyFixture) CreateMass() *geometry.Mass {
	return b.GetShape().CreateMass(b.density)
}
