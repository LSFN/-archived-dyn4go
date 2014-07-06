package test

import (
	"github.com/LSFN/dyn4go/collision"
	"github.com/LSFN/dyn4go/dynamics"
	"github.com/LSFN/dyn4go/geometry"
)

type CollidableTest struct {
	id        string
	fixtures  []*dynamics.BodyFixture
	transform *geometry.Transform
}

func NewCollidableTest(fixtures []*dynamics.BodyFixture) *CollidableTest {
	c := new(CollidableTest)
	c.fixtures = fixtures
	c.transform = geometry.NewTransform()
	return c
}

func NewCollidableTestShape(shape geometry.Convexer) *CollidableTest {
	c := new(CollidableTest)
	c.fixtures = []*dynamics.BodyFixture{dynamics.NewBodyFixture(shape)}
	c.transform = geometry.NewTransform()
	return c
}

func (c *CollidableTest) CreateAABB() *geometry.AABB {
	if len(c.fixtures) > 0 {
		aabb := c.fixtures[0].GetShape().CreateAABBTransform(c.transform)
		for i := 1; i < len(c.fixtures); i++ {
			faabb := c.fixtures[i].GetShape().CreateAABBTransform(c.transform)
			aabb.Union(faabb)
		}
		return aabb
	}
	return geometry.NewAABBFromFloats(0, 0, 0, 0)
}

func (c *CollidableTest) GetID() string {
	return c.id
}

func (c *CollidableTest) GetFixture(index int) collision.Fixturer {
	if len(c.fixtures) > 0 && index < len(c.fixtures) {
		return c.fixtures[index]
	}
	panic("No fixture for that index")
}

func (c *CollidableTest) GetFixtureCount() int {
	return len(c.fixtures)
}

func (c *CollidableTest) GetFixtures() []collision.Fixturer {
	result := make([]collision.Fixturer, 0)
	for _, f := range c.fixtures {
		result = append(result, f)
	}
	return result
}

func (c *CollidableTest) GetTransform() *geometry.Transform {
	return c.transform
}

func (c *CollidableTest) RotateAboutOrigin(theta float64) {
	c.transform.RotateAboutOrigin(theta)
}

func (c *CollidableTest) RotateAboutVector2(theta float64, point *geometry.Vector2) {
	c.transform.RotateAboutVector2(theta, point)
}

func (c *CollidableTest) RotateAboutXY(theta, x, y float64) {
	c.transform.RotateAboutXY(theta, x, y)
}

func (c *CollidableTest) TranslateXY(x, y float64) {
	c.transform.TranslateXY(x, y)
}

func (c *CollidableTest) TranslateVector2(vector *geometry.Vector2) {
	c.transform.TranslateVector2(vector)
}
