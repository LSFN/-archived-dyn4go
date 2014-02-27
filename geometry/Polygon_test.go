package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests not enough points.
 */
func TestCreateNotEnoughPoints(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	vertices := []*Vector2{
		new(Vector2),
		new(Vector2),
	}
	NewPolygon(vertices...)
}

/**
 * Tests not CCW.
 */
func TestCreateNotCCW(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	vertices := []*Vector2{
		new(Vector2),
		NewVector2FromXY(2.0, 2.0),
		NewVector2FromXY(1.0, 0.0),
	}
	NewPolygon(vertices...)
}

/**
 * Tests that the triangle is CCW.
 */
func TestCreateCCW(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	vertices := []*Vector2{
		NewVector2FromXY(0.5, 0.5),
		NewVector2FromXY(-0.3, -0.5),
		NewVector2FromXY(1.0, -0.3),
	}
	NewPolygon(vertices...)
}

/**
 * Tests coincident points
 */
func TestCreateCoincident(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	vertices := []*Vector2{
		new(Vector2),
		NewVector2FromXY(2.0, 2.0),
		NewVector2FromXY(2.0, 2.0),
		NewVector2FromXY(1.0, 0.0),
	}
	NewPolygon(vertices...)
}

/**
 * Tests non convex.
 */
func TestCreateNonConvex(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	vertices := []*Vector2{
		NewVector2FromXY(1.0, 1.0),
		NewVector2FromXY(-1.0, 1.0),
		NewVector2FromXY(-0.5, 0.0),
		NewVector2FromXY(-1.0, -1.0),
		NewVector2FromXY(1.0, -1.0),
	}
	NewPolygon(vertices...)
}

/**
 * Tests nil point array.
 */
func TestCreateNullPoints(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewPolygon(nil)
}

/**
 * Tests an array with nil points
 */
func TestCreateNullPoint(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	vertices := []*Vector2{
		new(Vector2),
		nil,
		NewVector2FromXY(0, 2),
	}
	NewPolygon(vertices...)
}

/**
 * Tests the constructor.
 */
func TestCreateSuccess(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-2.0, -2.0),
		NewVector2FromXY(1.0, -2.0),
	}
	NewPolygon(vertices...)
}

/**
 * Tests the contains method.
 */
func TestContains(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-1.0, 0.0),
		NewVector2FromXY(1.0, 0.0),
	}
	p := NewPolygon(vertices...)

	transform := NewTransform()
	pt := NewVector2FromXY(2.0, 4.0)

	// shouldn't be in the polygon
	dyn4go.AssertTrue(t, !p.ContainsVector2Transform(pt, transform))

	// move the polygon a bit
	transform.TranslateXY(2.0, 3.5)

	// should be in the polygon
	dyn4go.AssertTrue(t, p.ContainsVector2Transform(pt, transform))

	transform.TranslateXY(0.0, -0.5)

	// should be on a vertex
	dyn4go.AssertTrue(t, p.ContainsVector2Transform(pt, transform))
}

/**
 * Tests the project method.
 */
func TestProject(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-1.0, 0.0),
		NewVector2FromXY(1.0, 0.0),
	}
	p := NewPolygon(vertices...)
	transform := NewTransform()
	x := NewVector2FromXY(1.0, 0.0)
	y := NewVector2FromXY(0.0, 1.0)

	transform.TranslateXY(1.0, 0.5)

	i := p.ProjectVector2Transform(x, transform)

	dyn4go.AssertEqualWithinError(t, 0.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, i.max, 1.0e-3)

	// rotating about the center
	transform.RotateAboutXY(dyn4go.DegToRad(90), 1.0, 0.5)

	i = p.ProjectVector2Transform(y, transform)

	dyn4go.AssertEqualWithinError(t, -0.500, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.500, i.max, 1.0e-3)
}

/**
 * Tests the farthest methods.
 */
func TestGetFarthest(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-1.0, -1.0),
		NewVector2FromXY(1.0, -1.0),
	}
	p := NewPolygon(vertices...)
	transform := NewTransform()
	y := NewVector2FromXY(0.0, -1.0)

	f := p.GetFarthestFeature(y, transform)
	// should always get an edge
	dyn4go.AssertTrue(t, f.IsEdge())
	dyn4go.AssertEqualWithinError(t, -1.000, f.max.point.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, f.max.point.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, f.vertex1.point.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, f.vertex1.point.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, f.vertex2.point.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, f.vertex2.point.Y, 1.0e-3)

	pt := p.GetFarthestPoint(y, transform)

	dyn4go.AssertEqualWithinError(t, -1.000, pt.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, pt.Y, 1.0e-3)

	// rotating about the origin
	transform.RotateAboutXY(dyn4go.DegToRad(90), 0, 0)

	pt = p.GetFarthestPoint(y, transform)

	dyn4go.AssertEqualWithinError(t, 1.000, pt.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, pt.Y, 1.0e-3)
}

/**
 * Tests the getAxes method.
 */
func TestGetAxes(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-1.0, -1.0),
		NewVector2FromXY(1.0, -1.0),
	}
	p := NewPolygon(vertices...)
	transform := NewTransform()

	axes := p.GetAxes(nil, transform)
	dyn4go.AssertNotNil(t, axes)
	dyn4go.AssertEqual(t, 3, len(axes))

	// test passing some focal points
	pt := NewVector2FromXY(-3.0, 2.0)
	axes = p.GetAxes([]*Vector2{pt}, transform)
	dyn4go.AssertEqual(t, 4, len(axes))

	// make sure the axes are perpendicular to the edges
	ab := p.vertices[0].HereToVector2(p.vertices[1])
	bc := p.vertices[1].HereToVector2(p.vertices[2])
	ca := p.vertices[2].HereToVector2(p.vertices[0])

	dyn4go.AssertEqualWithinError(t, 0.000, ab.DotVector2(axes[0]), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, bc.DotVector2(axes[1]), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, ca.DotVector2(axes[2]), 1.0e-3)

	// make sure that the focal axes are correct
	dyn4go.AssertEqualWithinError(t, 0.000, p.vertices[0].HereToVector2(pt).CrossVector2(axes[3]), 1.0e-3)
}

/**
 * Tests the getFoci method.
 */
func TestGetFoci(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-1.0, -1.0),
		NewVector2FromXY(1.0, -1.0),
	}
	p := NewPolygon(vertices...)
	transform := NewTransform()
	// should return none
	foci := p.GetFoci(transform)
	dyn4go.AssertTrue(t, foci == nil)
}

/**
 * Tests the rotate methods.
 */
func TestRotate(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-1.0, -1.0),
		NewVector2FromXY(1.0, -1.0),
	}
	p := NewPolygon(vertices...)

	// should move the points
	p.RotateAboutXY(dyn4go.DegToRad(90), 0, 0)

	dyn4go.AssertEqualWithinError(t, -1.000, p.vertices[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.vertices[0].Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 1.000, p.vertices[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, p.vertices[1].Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 1.000, p.vertices[2].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, p.vertices[2].Y, 1.0e-3)
}

/**
 * Tests the translate methods.
 */
func TestTranslate(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-1.0, -1.0),
		NewVector2FromXY(1.0, -1.0),
	}
	p := NewPolygon(vertices...)

	p.TranslateXY(1.0, -0.5)

	dyn4go.AssertEqualWithinError(t, 1.000, p.vertices[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, p.vertices[0].Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 0.000, p.vertices[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.500, p.vertices[1].Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 2.000, p.vertices[2].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.500, p.vertices[2].Y, 1.0e-3)
}

/**
 * Tests the createAABB method.
 * @since 3.1.0
 */
func TestCreateAABB(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(-1.0, -1.0),
		NewVector2FromXY(1.0, -1.0),
	}
	p := NewPolygon(vertices...)

	aabb := p.CreateAABBTransform(NewTransform())
	dyn4go.AssertEqualWithinError(t, -1.0, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.0, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.GetMaxY(), 1.0e-3)

	// try using the default method
	aabb2 := p.CreateAABB()
	dyn4go.AssertEqual(t, aabb.GetMinX(), aabb2.GetMinX())
	dyn4go.AssertEqual(t, aabb.GetMinY(), aabb2.GetMinY())
	dyn4go.AssertEqual(t, aabb.GetMaxX(), aabb2.GetMaxX())
	dyn4go.AssertEqual(t, aabb.GetMaxY(), aabb2.GetMaxY())

	tx := NewTransform()
	tx.RotateAboutOrigin(dyn4go.DegToRad(30.0))
	tx.TranslateXY(1.0, 2.0)
	aabb = p.CreateAABBTransform(tx)
	dyn4go.AssertEqualWithinError(t, 0.500, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.634, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.366, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.866, aabb.GetMaxY(), 1.0e-3)
}
