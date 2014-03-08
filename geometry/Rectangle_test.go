package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests the constructor with an invalid width.
 */
func TestCreateInvalidWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewRectangle(-1.0, 3.0)
}

/**
 * Tests the constructor with an invalid height.
 */
func TestCreateInvalidHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewRectangle(2.0, 0.0)
}

/**
 * Tests a successful creation.
 */
func TestCreateSuccess(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	r := NewRectangle(2.0, 2.0)
	// make sure the center is 0,0
	dyn4go.AssertEqualWithinError(t, 0.000, r.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, r.center.Y, 1.0e-3)
	// make sure the points are correct
	dyn4go.AssertEqualWithinError(t, -1.000, r.vertices[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, r.vertices[0].Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 1.000, r.vertices[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, r.vertices[1].Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 1.000, r.vertices[2].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, r.vertices[2].Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, -1.000, r.vertices[3].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, r.vertices[3].Y, 1.0e-3)
}

/**
 * Tests the getAxes method.
 */
func TestGetAxes(t *testing.T) {
	r := NewRectangle(1.0, 1.0)
	transform := NewTransform()
	axes := r.GetAxes(nil, transform)

	// make sure there is only two
	dyn4go.AssertEqual(t, 2, len(axes))

	// make sure the axes are perpendicular to the edges
	ab := r.vertices[0].HereToVector2(r.vertices[1])
	ad := r.vertices[0].HereToVector2(r.vertices[3])

	dyn4go.AssertEqualWithinError(t, 0.000, ab.DotVector2(axes[1]), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, ad.DotVector2(axes[0]), 1.0e-3)

	// test a focal point
	pt := NewVector2FromXY(2.0, -1.0)
	axes = r.GetAxes([]*Vector2{pt}, transform)

	// make sure there are 4 more axes
	dyn4go.AssertEqual(t, 3, len(axes))
	// make sure they are parallel to the vector from a vertex to the focal point
	dyn4go.AssertEqualWithinError(t, 0.000, r.vertices[1].HereToVector2(pt).CrossVector2(axes[2]), 1.0e-3)
}

/**
 * Test the contains method.
 */
func TestContains(t *testing.T) {
	r := NewRectangle(1.0, 2.0)
	transform := NewTransform()

	pt := NewVector2FromXY(2.0, 0.5)

	dyn4go.AssertTrue(t, !r.ContainsVector2Transform(pt, transform))

	// move the rectangle a bit
	transform.TranslateXY(2.0, 0.0)
	transform.RotateAboutVector2(dyn4go.DegToRad(30), r.center)

	dyn4go.AssertTrue(t, r.ContainsVector2Transform(pt, transform))

	// check for on the edge
	transform.Identity()
	pt.SetToXY(0.5, 0.5)

	dyn4go.AssertTrue(t, r.ContainsVector2Transform(pt, transform))
}

/**
 * Tests the project method.
 */
func TestProject(t *testing.T) {
	r := NewRectangle(2.0, 1.0)
	transform := NewTransform()

	axis := NewVector2FromXY(1.0, 0.0)

	i := r.ProjectVector2Transform(axis, transform)

	dyn4go.AssertEqualWithinError(t, -1.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, i.max, 1.0e-3)

	// move the rectangle a bit
	transform.TranslateXY(1.0, 1.0)
	transform.RotateAboutVector2(dyn4go.DegToRad(30), transform.GetTransformedVector2(r.center))

	i = r.ProjectVector2Transform(axis, transform)

	dyn4go.AssertEqualWithinError(t, -0.116, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.116, i.max, 1.0e-3)

	axis.SetToXY(0.0, 1.0)

	i = r.ProjectVector2Transform(axis, transform)

	dyn4go.AssertEqualWithinError(t, 0.066, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.933, i.max, 1.0e-3)
}
