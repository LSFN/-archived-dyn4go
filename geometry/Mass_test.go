package geometry

import (
	"math"
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Test the create method.
 * <p>
 * Should throw an exception because the mass must be > 0.
 */
func TestCreateNegativeMass(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewMassFromCenterMassInertia(new(Vector2), -1.0, 1.0)
}

/**
 * Test the create method.
 * <p>
 * Should throw an exception because the inertia tensor must be > 0.
 */
func TestCreateNegativeInertia(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewMassFromCenterMassInertia(new(Vector2), 1.0, -1.0)
}

/**
 * Test the create method.
 * <p>
 * Should throw an exception because the center is null.
 * @since 2.0.0
 */
func TestCreateNullCenter(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewMassFromCenterMassInertia(nil, 1.0, 1.0)
}

/**
 * Test the create method.
 */
func TestCreateSuccess(t *testing.T) {
	m := NewMassFromCenterMassInertia(new(Vector2), 1.0, 1.0)
	dyn4go.AssertTrue(t, *m.GetCenter() == *new(Vector2))
	dyn4go.AssertEqual(t, m.GetMass(), 1.0)
	dyn4go.AssertEqual(t, m.GetInertia(), 1.0)
}

/**
 * Test the create infinite method.
 * @since 2.0.0
 */
func TestCreateInfinite(t *testing.T) {
	m := NewMassFromCenterMassInertia(new(Vector2), 0, 0)
	dyn4go.AssertTrue(t, m.IsInfinite())
	dyn4go.AssertTrue(t, *m.GetCenter() == *new(Vector2))
	dyn4go.AssertEqual(t, m.GetMass(), 0.0)
	dyn4go.AssertEqual(t, m.GetInertia(), 0.0)
}

/**
 * Test the create fixed linear velocity method.
 * @since 2.0.0
 */
func TestCreateFixedLinearVelocity(t *testing.T) {
	m := NewMassFromCenterMassInertia(new(Vector2), 0, 1.0)
	dyn4go.AssertFalse(t, m.IsInfinite())
	dyn4go.AssertEqual(t, FIXED_LINEAR_VELOCITY, m.GetType())
	dyn4go.AssertTrue(t, *m.GetCenter() == *new(Vector2))
	dyn4go.AssertEqual(t, m.GetMass(), 0.0)
	dyn4go.AssertEqual(t, m.GetInertia(), 1.0)
}

/**
 * Test the create fixed angular velocity method.
 * @since 2.0.0
 */
func TestCreateFixedAngularVelocity(t *testing.T) {
	m := NewMassFromCenterMassInertia(new(Vector2), 1.0, 0.0)
	dyn4go.AssertFalse(t, m.IsInfinite())
	dyn4go.AssertEqual(t, FIXED_ANGULAR_VELOCITY, m.GetType())
	dyn4go.AssertTrue(t, *m.GetCenter() == *new(Vector2))
	dyn4go.AssertEqual(t, m.GetMass(), 1.0)
	dyn4go.AssertEqual(t, m.GetInertia(), 0.0)
}

/**
 * Test the create method.
 * <p>
 * Should throw an exception because the mass to copy is null.
 * @since 2.0.0
 */
func TestCreateCopyNull(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewMassFromMass(nil)
}

/**
 * Test the create method.
 * @since 2.0.0
 */
func TestCreateCopy(t *testing.T) {
	m := NewMassFromCenterMassInertia(NewVector2FromXY(1.0, 0.0), 2.0, 1.0)
	m2 := NewMassFromMass(m)

	dyn4go.AssertNotEqual(t, m, m2)
	dyn4go.AssertNotEqual(t, m.center, m2.center)
	dyn4go.AssertEqual(t, m.center.X, m2.center.X)
	dyn4go.AssertEqual(t, m.center.Y, m2.center.Y)
	dyn4go.AssertEqual(t, m.GetMass(), m2.GetMass())
	dyn4go.AssertEqual(t, m.GetInertia(), m2.GetInertia())
	dyn4go.AssertEqual(t, m.GetType(), m2.GetType())
}

/**
 * Test case for the circle create method.
 */
func TestCreateCircle(t *testing.T) {
	c := NewCircle(3.0)
	m := c.CreateMass(2.0)
	// the mass should be pi * r * r * d
	dyn4go.AssertEqualWithinError(t, 56.548, m.GetMass(), 1.0e-3)
	// I should be m * r * r / 2
	dyn4go.AssertEqualWithinError(t, 254.469, m.GetInertia(), 1.0e-3)
}

/**
 * Test case for the polygon create method.
 */
func TestCreatePolygon(t *testing.T) {
	p := CreateUnitCirclePolygon(5, 0.5)
	m := p.CreateMass(1.0)
	// the polygon mass should be the area * d
	dyn4go.AssertEqualWithinError(t, 0.594, m.GetMass(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.057, m.GetInertia(), 1.0e-3)
}

/**
 * Test case for the rectangle create method.
 */
func TestCreateRectangle(t *testing.T) {
	r := NewRectangle(1.0, 1.0)
	m := r.CreateMass(1.5)
	// the mass of a rectangle should be h * w * d
	dyn4go.AssertEqual(t, 1.500, m.GetMass(), 1.0e-3)
	dyn4go.AssertEqual(t, 0.250, m.GetInertia(), 1.0e-3)
}

/**
 * Test case for the segment create method.
 */
func TestCreateSegment(t *testing.T) {
	s := NewSegment(NewVector2FromXY(-1.0, 0.0), NewVector2FromXY(1.0, 0.5))
	m := s.CreateMass(1.0)
	// the mass of a segment should be l * d
	dyn4go.AssertEqualWithinError(t, 2.061, m.GetMass(), 1.0e-3)
	// the I of a segment should be 1 / 12 * l ^ 2 * m
	dyn4go.AssertEqualWithinError(t, 0.730, m.GetInertia(), 1.0e-3)
}

/**
 * Test the create method accepting an array of {@link Mass} objects.
 * <p>
 * Renamed from createArray
 * @since 2.0.0
 */
func TestCreateList(t *testing.T) {
	masses := []*Mass{
		NewMassFromCenterMassInertia(NewVector2FromXY(1.0, 1.0), 3.00, 1.00),
		NewMassFromCenterMassInertia(NewVector2FromXY(-1.0, 0.0), 0.50, 0.02),
		NewMassFromCenterMassInertia(NewVector2FromXY(1.0, -2.0), 2.00, 3.00),
	}
	m := CreateMass(masses)

	c := m.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.818, c.x, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.181, c.y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 5.500, m.GetMass(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 16.656, m.GetInertia(), 1.0e-3)
}

/**
 * Test the create method accepting an array of infinite {@link Mass} objects.
 * @since 2.0.0
 */
func TestCreateListInfinite(t *testing.T) {
	masses := []*Mass{
		new(Mass), new(Mass), new(Mass),
	}
	m := CreateMass(masses)

	c := m.GetCenter()
	dyn4go.AssertTrue(t, m.IsInfinite())
	dyn4go.AssertEqualWithinError(t, 0.000, c.x, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, c.y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, m.GetMass(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, m.GetInertia(), 1.0e-3)
}

/**
 * Test the create method accepting a list of one mass.
 * @since 2.0.0
 */
func TestCreateListOneElement(t *testing.T) {
	m1 := NewMassFromCenterMassInertia(new(Vector2), 1.0, 2.0)
	masses := []*Mass{
		NewMassFromCenterMassInertia(new(Vector2), 1.0, 2.0),
	}
	m := CreateMass(masses)

	c := m.GetCenter()
	dyn4go.AssertFalse(t, m.IsInfinite())
	dyn4go.AssertNotEqual(t, m1, m)
	dyn4go.AssertEqualWithinError(t, 0.000, c.x, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, c.y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, m.GetMass(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, m.GetInertia(), 1.0e-3)
}

/**
 * Test the create method accepting a null list.
 * @since 3.1.0
 */
func TestCreateListNull(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateMass(nil)
}

/**
 * Test the create method accepting an empty list.
 * @since 3.1.0
 */
func TestCreateListEmpty(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateMass(make([]*Mass, 0))
}

/**
 * Test the create method accepting a list of one null mass.
 * @since 2.0.0
 */
func TestCreateListOneNullElement(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	masses := []*Mass{nil}
	CreateMass(masses)
}

/**
 * Test the create method accepting a list masses where 1 is null.
 * @since 2.0.0
 */
func TestCreateListNullElement(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	masses := []*Mass{
		NewMassFromCenterMassInertia(new(Vector2), 1.0, 2.0),
		nil,
		NewMassFromCenterMassInertia(new(Vector2), 2.0, 7.0),
	}
	CreateMass(masses)
}

/**
 * Tests setting the type of mass.
 * @since 1.0.2
 */
func TestSetType(t *testing.T) {
	c := CreateCircle(2.0)
	mi := c.CreateMass(1.0)

	// setting the type should not alter the
	// mass values
	mi.SetType(INFINITE)
	dyn4go.AssertTrue(t, mi.IsInfinite())
	dyn4go.AssertFalse(t, 0.0 == mi.mass)
	dyn4go.AssertFalse(t, 0.0 == mi.invMass)
	dyn4go.AssertFalse(t, 0.0 == mi.inertia)
	dyn4go.AssertFalse(t, 0.0 == mi.invInertia)
	// the get methods should return 0
	dyn4go.AssertEqual(t, 0.0, mi.GetMass())
	dyn4go.AssertEqual(t, 0.0, mi.GetInverseMass())
	dyn4go.AssertEqual(t, 0.0, mi.GetInertia())
	dyn4go.AssertEqual(t, 0.0, mi.GetInverseInertia())

	mi.SetType(FIXED_ANGULAR_VELOCITY)
	dyn4go.AssertFalse(t, 0.0 == mi.mass)
	dyn4go.AssertFalse(t, 0.0 == mi.invMass)
	dyn4go.AssertFalse(t, 0.0 == mi.inertia)
	dyn4go.AssertFalse(t, 0.0 == mi.invInertia)
	dyn4go.AssertEqual(t, 0.0, mi.GetInertia())
	dyn4go.AssertEqual(t, 0.0, mi.GetInverseInertia())

	mi.SetType(FIXED_LINEAR_VELOCITY)
	dyn4go.AssertFalse(t, 0.0 == mi.mass)
	dyn4go.AssertFalse(t, 0.0 == mi.invMass)
	dyn4go.AssertFalse(t, 0.0 == mi.inertia)
	dyn4go.AssertFalse(t, 0.0 == mi.invInertia)
	dyn4go.AssertEqual(t, 0.0, mi.GetMass())
	dyn4go.AssertEqual(t, 0.0, mi.GetInverseMass())
}

/**
 * Tests setting the type of mass to null.
 * @since 3.1.0
 */
func TestSetNullType(t *testing.T) {
	dyn4go.AssertPanic(t)
	m := new(Mass)
	m.SetType(nil)
}

/**
 * Tests the inertia and COM calculations for polygon shapes.
 * @since 3.1.4
 */
func TestPolygonInertiaAndCOM(t *testing.T) {
	// a polygon of a simple shape should match a simple shape's mass and inertia
	p := CreateUnitCirclePolygon(4, math.Hypot(0.5, 0.5))
	r := CreateSquare(1.0)

	pm := p.CreateMass(10.0)
	rm := r.CreateMass(10.0)

	dyn4go.AssertEqualWithinError(t, rm.mass, pm.mass, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, rm.inertia, pm.inertia, 1.0e-3)
}

/**
 * Make sure the center of mass does not effect the mass or inertia.
 * @since 3.1.5
 */
func TestPolygonInertiaAndMass(t *testing.T) {
	// a polygon of a simple shape should match a simple shape's mass and inertia
	p := CreateUnitCirclePolygon(4, math.Hypot(0.5, 0.5))
	m1 := p.CreateMass(10.0)

	p.Translate(0.5, -2.0)
	m2 := p.CreateMass(10.0)

	dyn4go.AssertEqualWithinError(t, m1.mass, m2.mass, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, m1.inertia, m2.inertia, 1.0e-3)
}
