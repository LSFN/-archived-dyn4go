package collision

import (
	"github.com/LSFN/dyn4go/geometry"
)

type Collider interface {
	geometry.Transformer
	GetID() string
	CreateAABB() *geometry.AABB
	GetFixture(index int) *Fixture
	GetFixtureCount() int
	GetFixtures() []*Fixture
	GetTransform() *geometry.Transform
}
