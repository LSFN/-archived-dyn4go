package geometry

import (
	"math"
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests the getAverageCenter method.
 * <p>
 * This test also shows that the average method can produce an incorrect
 * center of mass when vertices are more dense at any place along the perimeter.
 */

func TestGeometryGetAverageCenterArray(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(-2.0, 1.0),
		NewVector2FromXY(-1.0, 2.0),
		NewVector2FromXY(1.2, 0.5),
		NewVector2FromXY(1.3, 0.3),
		NewVector2FromXY(1.4, 0.2),
		NewVector2FromXY(0.0, -1.0),
	}

	c := GetAverageCenterFromList(vertices)

	dyn4go.AssertEqualWithinError(t, 0.150, c.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, c.Y, 1.0e-3)
}

/**
 * Tests the getAverageCenter method passing a nil array.
 * @since 2.0.0
 */

func TestGeometryGetAverageCenterNullArray(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	GetAverageCenterFromList(nil)
}

/**
 * Tests the getAverageCenter method passing an empty array.
 * @since 3.1.0
 */

func TestGeometryGetAverageCenterEmptyArray(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	GetAverageCenterFromList([]*Vector2{})
}

/**
 * Tests the getAverageCenter method passing an array with nil elements.
 * @since 3.1.0
 */

func TestGeometryGetAverageCenterArrayNullElements(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	GetAverageCenterFromList([]*Vector2{
		NewVector2FromXY(1.0, 0.0),
		nil,
		NewVector2FromXY(4.0, 3.0),
		NewVector2FromXY(-2.0, -1.0),
		nil,
	})
}

/**
 * Tests the getAreaWeightedCenter method.
 */

func TestGeometryGetAreaWeightedCenter(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(-2.0, 1.0),
		NewVector2FromXY(-1.0, 2.0),
		// test dense area of points
		NewVector2FromXY(1.2, 0.5),
		NewVector2FromXY(1.3, 0.3),
		NewVector2FromXY(1.4, 0.2),
		NewVector2FromXY(0.0, -1.0),
	}

	c := GetAreaWeightedCenterFromList(vertices)

	// note the x is closer to the "real" center of the object
	dyn4go.AssertEqualWithinError(t, -0.318, c.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.527, c.Y, 1.0e-3)
}

/**
 * Tests the getAreaWeightedCenter method with a polygon that is not centered
 * about the origin.
 * @since 3.1.4
 */

func TestGeometryGetAreaWeightedCenterOffset(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(-1.0, 2.0),
		NewVector2FromXY(0.0, 3.0),
		// test dense area of points
		NewVector2FromXY(2.2, 1.5),
		NewVector2FromXY(2.3, 1.3),
		NewVector2FromXY(2.4, 1.2),
		NewVector2FromXY(1.0, 0.0),
	}
	c := GetAreaWeightedCenterFromList(vertices)

	// note the x is closer to the "real" center of the object
	dyn4go.AssertEqualWithinError(t, 0.682, c.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.527, c.Y, 1.0e-3)
}

/**
 * Tests the getAreaWeightedCenter method passing a nil array.
 * @since 2.0.0
 */

func TestGeometryGetAreaWeightedCenterNullArray(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	GetAreaWeightedCenterFromList(nil)
}

/**
 * Tests the getAreaWeightedCenter method passing an empty array.
 * @since 3.1.0
 */

func TestGeometryGetAreaWeightedCenterEmptyArray(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	GetAreaWeightedCenterFromList([]*Vector2{})
}

/**
 * Tests the getAreaWeightedCenter method passing an array with nil elements.
 * @since 3.1.0
 */

func TestGeometryGetAreaWeightedCenterArrayNullElements(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	GetAreaWeightedCenterFromList([]*Vector2{
		NewVector2FromXY(1.0, 0.0),
		nil,
		NewVector2FromXY(4.0, 3.0),
		NewVector2FromXY(-2.0, -1.0),
		nil,
	})
}

/**
 * Tests the getAreaWeightedCenter method passing a list of
 * points who are all the same yielding zero area.
 * @since 2.0.0
 */

func TestGeometryGetAreaWeightedCenterZeroAreaArray(t *testing.T) {
	points := []*Vector2{
		NewVector2FromXY(2.0, 1.0),
		NewVector2FromXY(2.0, 1.0),
		NewVector2FromXY(2.0, 1.0),
		NewVector2FromXY(2.0, 1.0),
	}

	c := GetAreaWeightedCenterFromList(points)

	dyn4go.AssertEqualWithinError(t, 2.000, c.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, c.Y, 1.0e-3)
}

/**
 * Test case for the unitCirclePolygon methods.
 * @since 3.1.0
 */

func TestGeometryCreateUnitCirclePolygon(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	p := CreateUnitCirclePolygon(5, 0.5)
	// no exception indicates the generated polygon is valid
	// test that the correct vertices are created
	dyn4go.AssertEqualWithinError(t, 0.154, p.vertices[4].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.475, p.vertices[4].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.404, p.vertices[3].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.293, p.vertices[3].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.404, p.vertices[2].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.293, p.vertices[2].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.154, p.vertices[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.475, p.vertices[1].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, p.vertices[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.vertices[0].Y, 1.0e-3)

	v11 := p.vertices[0]

	p = CreateUnitCirclePolygonTheta(5, 0.5, math.Pi/2.0)
	// no exception indicates the generated polygon is valid
	// test that the correct vertices are created
	dyn4go.AssertEqualWithinError(t, 0.475, p.vertices[4].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.154, p.vertices[4].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.293, p.vertices[3].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.404, p.vertices[3].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.293, p.vertices[2].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.404, p.vertices[2].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.475, p.vertices[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.154, p.vertices[1].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.vertices[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, p.vertices[0].Y, 1.0e-3)

	v21 := p.vertices[0]

	// the angle between any two vertices of the two polygons should be PI / 2
	angle := v11.GetAngleBetween(v21)
	dyn4go.AssertEqualWithinError(t, math.Pi/2.0, angle, 1.0e-3)
}

/**
 * Tests the failed creation of a negative radius unit circle polygon.
 * @since 3.1.0
 */

func TestGeometryCreateNegativeRadiusUnitCirclePolygon(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateUnitCirclePolygon(5, -0.5)
}

/**
 * Tests the failed creation of a zero radius unit circle polygon.
 * @since 3.1.0
 */

func TestGeometryCreateZeroRadiusUnitCirclePolygon(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateUnitCirclePolygon(5, 0.0)
}

/**
 * Tests the failed creation of a unit circle polygon with less than 3 points.
 * @since 3.1.0
 */

func TestGeometryCreateLessThan3PointsUnitCirclePolygon(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateUnitCirclePolygon(2, 0.5)
}

/**
 * Tests the successful creation of a circle.
 */

func TestGeometryCreateCircle(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	CreateCircle(1.0)
}

/**
 * Tests the failed creation of a circle using a negative radius.
 * @since 3.1.0
 */

func TestGeometryCreateNegativeRadiusCircle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateCircle(-1.0)
}

/**
 * Tests the failed creation of a circle using a zero radius.
 * @since 3.1.0
 */

func TestGeometryCreateZeroRadiusCircle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateCircle(0.0)
}

/**
 * Tests the creation of a polygon with a nil array.
 */

func TestGeometryCreatePolygonNullArray(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	// should fail since the vertices list contains nil items
	CreatePolygon(nil)
}

/**
 * Tests the creation of a polygon with a nil point.
 */

func TestGeometryCreatePolygonNullPoint(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	vertices := []*Vector2{}
	// should fail since the vertices list contains nil items
	CreatePolygon(vertices...)
}

/**
 * Tests the successful creation of a polygon using vertices.
 */

func TestGeometryCreatePolygon(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	vertices := []*Vector2{
		NewVector2FromXY(1.0, 0.0),
		NewVector2FromXY(0.5, 1.0),
		NewVector2FromXY(-0.5, 1.0),
		NewVector2FromXY(-1.0, 0.0),
		NewVector2FromXY(0.0, -1.0),
	}
	// should fail since the vertices list contains nil items
	p := CreatePolygon(vertices...)

	// the array should not be the same object
	dyn4go.AssertFalse(t, &p.vertices == &vertices)
	// the points should also be copies
	for i := 0; i < 5; i++ {
		dyn4go.AssertFalse(t, &p.vertices[0] == &vertices[0])
	}
}

/**
 * Tests the successful creation of a polygon using vertices.
 */
//HERE TODO
func TestGeometryCreatePolygonAtOrigin(t *testing.T) {
	vertices := []*Vector2{
		NewVector2FromXY(1.0, 0.0),
		NewVector2FromXY(0.5, 1.0),
		NewVector2FromXY(-0.5, 1.0),
		NewVector2FromXY(-1.0, 0.0),
		NewVector2FromXY(0.0, -1.0),
	}
	// should fail since the vertices list contains nil items
	p := CreatePolygonAtOrigin(vertices...)

	// the array should not be the same object
	dyn4go.AssertFalse(t, &p.vertices == &vertices)
	// the points should also be copies
	for i, v := range p.vertices {
		dyn4go.AssertFalse(t, &v == &vertices[i])
	}

	// make sure the center is at the origin
	c := p.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, c.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, c.Y, 1.0e-3)
}

/**
 * Tests the creation of a square with a zero size.
 */
func TestGeometryCreateZeroSizeSquare(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateSquare(0.0)
}

/**
 * Tests the creation of a square with a negative size.
 */
func TestGeometryCreateNegativeSizeSquare(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateSquare(-1.0)
}

/**
 * Tests the successful creation of a square.
 */

func TestGeometryCreateSquare(t *testing.T) {
	r := CreateSquare(1.0)
	dyn4go.AssertEqualWithinError(t, 1.000, r.GetWidth(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, r.GetHeight(), 1.0e-3)
}

/**
 * Tests the successful creation of a rectangle.
 */

func TestGeometryCreateRectangle(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	CreateRectangle(1.0, 2.0)
}

/**
 * Tests the failed creation of a rectangle with a negative width.
 */

func TestGeometryCreateNegativeWidthRectangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateRectangle(-1.0, 2.0)
}

/**
 * Tests the failed creation of a rectangle with a negative height.
 */

func TestGeometryCreateNegativeHeightRectangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateRectangle(1.0, -2.0)
}

/**
 * Tests the failed creation of a rectangle with a zero width.
 */
func TestGeometryCreateZeroWidthRectangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateRectangle(0.0, 2.0)
}

/**
 * Tests the failed creation of a rectangle with a zero height.
 */
func TestGeometryCreateZeroHeightRectangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateRectangle(1.0, 0.0)
}

/**
 * Tests the creation of a triangle using a nil point.
 */
func TestGeometryCreateTriangleNullPoint(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	p1 := NewVector2FromXY(1.0, 0.0)
	p2 := NewVector2FromXY(0.5, 1.0)
	// should fail since the vertices list contains nil items
	CreateTriangle(p1, p2, nil)
}

/**
 * Tests the successful creation of a triangle using points.
 */

func TestGeometryCreateTriangle(t *testing.T) {
	p1 := NewVector2FromXY(1.0, 0.0)
	p2 := NewVector2FromXY(0.5, 1.0)
	p3 := NewVector2FromXY(-0.5, 1.0)
	triangle := CreateTriangle(p1, p2, p3)

	// the points should not be the same instances
	dyn4go.AssertFalse(t, &triangle.vertices[0] == &p1)
	dyn4go.AssertFalse(t, &triangle.vertices[1] == &p2)
	dyn4go.AssertFalse(t, &triangle.vertices[2] == &p3)
}

/**
 * Tests the successful creation of a triangle using points.
 */

func TestGeometryCreateTriangleAtOrigin(t *testing.T) {
	p1 := NewVector2FromXY(1.0, 0.0)
	p2 := NewVector2FromXY(0.5, 1.0)
	p3 := NewVector2FromXY(-0.5, 1.0)
	triangle := CreateTriangle(p1, p2, p3)

	// the points should not be the same instances
	dyn4go.AssertFalse(t, &triangle.vertices[0] == &p1)
	dyn4go.AssertFalse(t, &triangle.vertices[1] == &p2)
	dyn4go.AssertFalse(t, &triangle.vertices[2] == &p3)

	// make sure the center is at the origin
	c := triangle.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, c.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, c.Y, 1.0e-3)
}

/**
 * Tests the create right triangle method with a zero width.
 */
func TestGeometryCreateZeroWidthRightTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateRightTriangle(0.0, 2.0)
}

/**
 * Tests the create right triangle method with a zero height.
 */
func TestGeometryCreateZeroHeightRightTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateRightTriangle(1.0, 0.0)
}

/**
 * Tests the create right triangle method with a negative width.
 */
func TestGeometryCreateNegativeWidthRightTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateRightTriangle(-1.0, 2.0)
}

/**
 * Tests the create right triangle method with a negative height.
 */
func TestGeometryCreateNegativeHeightRightTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateRightTriangle(2.0, -2.0)
}

/**
 * Tests the successful creation of a right angle triangle.
 */

func TestGeometryCreateRightTriangle(t *testing.T) {
	triangle := CreateRightTriangle(1.0, 2.0)

	// test that the center is the origin
	center := triangle.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, center.Y, 1.0e-3)

	// get the vertices
	v1 := triangle.vertices[0]
	v2 := triangle.vertices[1]
	v3 := triangle.vertices[2]

	// create the edges
	e1 := v1.HereToVector2(v2)
	e2 := v2.HereToVector2(v3)
	e3 := v3.HereToVector2(v1)

	// one of the follow dot products must be zero
	// indicating a 90 degree angle
	if e1.DotVector2(e2) < 0.00001 && e1.DotVector2(e2) > -0.00001 {
		dyn4go.AssertTrue(t, true)
		return
	}

	if e2.DotVector2(e3) < 0.00001 && e2.DotVector2(e3) > -0.00001 {
		dyn4go.AssertTrue(t, true)
		return
	}

	if e3.DotVector2(e1) < 0.00001 && e3.DotVector2(e1) > -0.00001 {
		dyn4go.AssertTrue(t, true)
		return
	}

	// if we get here we didn't find a 90 degree angle
	dyn4go.AssertFalse(t, true)
}

/**
 * Tests the successful creation of a right angle triangle.
 */

func TestGeometryCreateRightTriangleMirror(t *testing.T) {
	triangle := CreateRightTriangleMirror(1.0, 2.0, true)

	// test that the center is the origin
	center := triangle.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, center.Y, 1.0e-3)

	// get the vertices
	v1 := triangle.vertices[0]
	v2 := triangle.vertices[1]
	v3 := triangle.vertices[2]

	// create the edges
	e1 := v1.HereToVector2(v2)
	e2 := v2.HereToVector2(v3)
	e3 := v3.HereToVector2(v1)

	// one of the follow dot products must be zero
	// indicating a 90 degree angle
	if e1.DotVector2(e2) < 0.00001 && e1.DotVector2(e2) > -0.00001 {
		dyn4go.AssertTrue(t, true)
		return
	}

	if e2.DotVector2(e3) < 0.00001 && e2.DotVector2(e3) > -0.00001 {
		dyn4go.AssertTrue(t, true)
		return
	}

	if e3.DotVector2(e1) < 0.00001 && e3.DotVector2(e1) > -0.00001 {
		dyn4go.AssertTrue(t, true)
		return
	}

	// if we get here we didn't find a 90 degree angle
	dyn4go.AssertFalse(t, true)
}

/**
 * Tests the create equilateral triangle method with a zero height.
 */
func TestGeometryCreateZeroHeightEquilateralTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateEquilateralTriangle(0.0)
}

/**
 * Tests the create equilateral triangle method with a negative height.
 */
func TestGeometryCreateNegativeHeightEquilateralTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateEquilateralTriangle(-1.0)
}

/**
 * Tests the successful creation of an equilateral angle triangle.
 */

func TestGeometryCreateEquilateralTriangle(t *testing.T) {
	triangle := CreateEquilateralTriangle(2.0)

	// test that the center is the origin
	center := triangle.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, center.Y, 1.0e-3)

	// compute the first angle
	previousA := triangle.vertices[0].GetAngleBetween(triangle.vertices[1])
	// put the angle between 0 and 180
	previousA = math.Abs(math.Pi - math.Abs(previousA))
	// compute the first distance
	previousD := triangle.vertices[0].DistanceFromVector2(triangle.vertices[1])
	// make sure all the angles are the same
	for i := range triangle.vertices {
		v1 := triangle.vertices[i]
		v2 := triangle.vertices[0]
		if i+1 < 3 {
			v2 = triangle.vertices[i+1]
		}
		// test the angle between the vectors
		angle := v1.GetAngleBetween(v2)
		// put the angle between 0 and 180
		angle = math.Abs(math.Pi - math.Abs(angle))
		if angle < previousA*0.9999 || angle > previousA*1.0001 {
			// its not the same as the last so we fail
			dyn4go.AssertFalse(t, true)
		}
		// test the distance between the points
		distance := v1.DistanceFromVector2(v2)
		if distance < previousD*0.9999 || distance > previousD*1.0001 {
			// its not the same as the last so we fail
			dyn4go.AssertFalse(t, true)
		}
	}
	// if we get here we didn't find a 90 degree angle
	dyn4go.AssertTrue(t, true)
}

/**
 * Tests the create right triangle method with a zero width.
 */
func TestGeometryCreateZeroWidthIsoscelesTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateIsoscelesTriangle(0.0, 1.0)
}

/**
 * Tests the create right triangle method with a zero height.
 */
func TestGeometryCreateZeroHeightIsoscelesTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateIsoscelesTriangle(1.0, 0.0)
}

/**
 * Tests the create right triangle method with a negative width.
 */
func TestGeometryCreateNegativeWidthIsoscelesTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateIsoscelesTriangle(-1.0, 2.0)
}

/**
 * Tests the create right triangle method with a negative height.
 */
func TestGeometryCreateNegativeHeightIsoscelesTriangle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateIsoscelesTriangle(2.0, -2.0)
}

/**
 * Tests the successful creation of an isosceles triangle.
 */

func TestGeometryCreateIsoscelesTriangle(t *testing.T) {
	triangle := CreateIsoscelesTriangle(2.0, 1.0)

	// test that the center is the origin
	center := triangle.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, center.Y, 1.0e-3)

	// get the vertices
	v1 := triangle.vertices[0]
	v2 := triangle.vertices[1]
	v3 := triangle.vertices[2]

	// create the edges
	e1 := v1.HereToVector2(v2)
	e2 := v2.HereToVector2(v3)
	e3 := v3.HereToVector2(v1)

	// the length of e1 and e3 should be identical
	dyn4go.AssertEqualWithinError(t, e1.GetMagnitude(), e3.GetMagnitude(), 1.0e-3)

	// then angles between e1 and e2 and e2 and e3 should be identical
	dyn4go.AssertEqualWithinError(t, e1.GetAngleBetween(e2), e2.GetAngleBetween(e3), 1.0e-3)
}

/**
 * Tests the creation of a segment passing a nil point.
 */
func TestGeometryCreateSegmentNullPoint1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateSegment(nil, new(Vector2))
}

/**
 * Tests the creation of a segment passing a nil point.
 * @since 3.1.0
 */
func TestGeometryCreateSegmentNullPoint2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateSegment(new(Vector2), nil)
}

/**
 * Tests the successful creation of a segment given two points.
 */

func TestGeometryCreateSegment(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	CreateSegment(NewVector2FromXY(1.0, 1.0), NewVector2FromXY(2.0, -1.0))
}

/**
 * Tests the successful creation of a segment given two points at the origin.
 */

func TestGeometryCreateSegmentAtOrigin(t *testing.T) {
	s := CreateSegmentAtOrigin(NewVector2FromXY(1.0, 1.0), NewVector2FromXY(2.0, -1.0))

	// test that the center is the origin
	center := s.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, center.Y, 1.0e-3)
}

/**
 * Tests the successful creation of a segment given an end point.
 */

func TestGeometryCreateSegmentEnd(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	CreateSegmentEnd(NewVector2FromXY(1.0, 1.0))
}

/**
 * Tests the creation of a segment passing a zero length.
 */
func TestGeometryCreateZeroLengthHorizontalSegment(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateHorizontalSegment(0.0)
}

/**
 * Tests the creation of a segment passing a negative length.
 */
func TestGeometryCreateNegativeLengthHorizontalSegment(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateHorizontalSegment(-1.0)
}

/**
 * Tests the successful creation of a segment given a length.
 */

func TestGeometryCreateHorizontalSegment(t *testing.T) {
	s := CreateHorizontalSegment(5.0)

	// test that the center is the origin
	center := s.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, center.Y, 1.0e-3)
}

/**
 * Tests the creation of a segment passing a zero length.
 * @since 2.2.3
 */
func TestGeometryCreateZeroLengthVerticalSegment(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateVerticalSegment(0.0)
}

/**
 * Tests the creation of a segment passing a negative length.
 * @since 2.2.3
 */
func TestGeometryCreateNegativeLengthVerticalSegment(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreateVerticalSegment(-1.0)
}

/**
 * Tests the successful creation of a segment given a length.
 * @since 2.2.3
 */

func TestGeometryCreateVerticalSegment(t *testing.T) {
	s := CreateVerticalSegment(5.0)

	// test that the center is the origin
	center := s.GetCenter()
	dyn4go.AssertEqualWithinError(t, 0.000, center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, center.Y, 1.0e-3)
}

/**
 * Tests the getWinding method passing a list.
 */

func TestGeometryGetWindingList(t *testing.T) {
	points := []*Vector2{
		NewVector2FromXY(-1.0, -1.0),
		NewVector2FromXY(1.0, -1.0),
		NewVector2FromXY(1.0, 1.0),
		NewVector2FromXY(-1.0, 1.0),
	}
	dyn4go.AssertTrue(t, GetWindingFromList(points) > 0)

	ReverseSliceVector2(points)
	dyn4go.AssertTrue(t, GetWindingFromList(points) < 0)
}

/**
 * Tests the getWinding method passing a nil list.
 */
func TestGeometryGetWindingNullList(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	GetWindingFromList(nil)
}

/**
 * Tests the getWinding method passing a list with 1 point.
 */
func TestGeometryGetWindingListLessThan2Points(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	GetWindingFromList([]*Vector2{})
}

/**
 * Tests the getWinding method passing a list that contains a nil point.
 */
func TestGeometryGetWindingListNullPoint(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	points := []*Vector2{
		new(Vector2),
		nil,
		nil,
	}
	GetWindingFromList(points)
}

/**
 * Tests the reverse winding method passing a nil array.
 */

func TestGeometryReverseWindingNullArray(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ReverseWindingFromList(nil)
}

/**
 * Tests the cleanse method passing a nil array.
 * @since 2.2.3
 */
func TestGeometryCleanseNullArray(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	Cleanse(nil)
}

/**
 * Tests the cleanse method passing a nil array.
 * @since 2.2.3
 */
func TestGeometryCleanseArrayWithNullElements(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	array := []*Vector2{
		new(Vector2),
		nil,
		nil,
		new(Vector2),
		new(Vector2),
	}
	Cleanse(array)
}

/**
 * Tests the cleanse list method.
 */

func TestGeometryCleanseList(t *testing.T) {
	points := []*Vector2{
		NewVector2FromXY(1.0, 0.0),
		NewVector2FromXY(1.0, 0.0),
		NewVector2FromXY(0.5, -0.5),
		NewVector2FromXY(0.0, -0.5),
		NewVector2FromXY(-0.5, -0.5),
		NewVector2FromXY(-2.0, -0.5),
		NewVector2FromXY(2.1, 0.5),
		NewVector2FromXY(1.0, 0.0),
	}

	result := Cleanse(points)

	dyn4go.AssertTrue(t, GetWindingFromList(result) > 0.0)
	dyn4go.AssertEqual(t, 4, len(result))
}

/**
 * Tests the createPolygonalEllipse method.
 * @since 3.1.5
 */

func TestGeometryCreatePolygonalEllipse(t *testing.T) {
	dyn4go.AssertNoPanic(t)
	// this method should succeed
	p := CreatePolygonalEllipse(10, 2, 1)
	// and the center should be the origin
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().Y, 1.0e-3)
}

/**
 * Tests the createPolygonalEllipse method with an odd count.
 * @since 3.1.5
 */

func TestGeometryCreatePolygonalEllipseOddCount(t *testing.T) {
	dyn4go.AssertNoPanic(t)
	// this method should succeed
	p := CreatePolygonalEllipse(5, 2, 1)
	// and the center should be the origin
	dyn4go.AssertEqual(t, 4, len(p.GetVertices()))

	// this method should succeed
	p = CreatePolygonalEllipse(11, 2, 1)
	// and the center should be the origin
	dyn4go.AssertEqual(t, 10, len(p.GetVertices()))
}

/**
 * Tests the createPolygonalEllipse method with less than 4 vertices.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalEllipseLessCount(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalEllipse(3, 2, 1)
}

/**
 * Tests the createPolygonalEllipse method with a zero width.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalEllipseZeroWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalEllipse(10, 0, 1)
}

/**
 * Tests the createPolygonalEllipse method with a zero height.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalEllipseZeroHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalEllipse(10, 2, 0)
}

/**
 * Tests the createPolygonalEllipse method with a negative width.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalEllipseNegativeWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalEllipse(10, -1, 1)
}

/**
 * Tests the createPolygonalEllipse method with a negative height.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalEllipseNegativeHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalEllipse(10, 2, -1)
}

/**
 * Tests the createPolygonalSlice method.
 * @since 3.1.5
 */

func TestGeometryCreatePolygonalSlice(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	// this method should succeed
	p := CreatePolygonalSlice(5, 1.0, dyn4go.DegToRad(30))
	// the center should not be at the origin
	dyn4go.AssertEqualWithinError(t, 0.658, p.GetCenter().X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().Y, 1.0e-3)
}

/**
 * Tests the createPolygonalSliceAtOrigin method.
 * @since 3.1.5
 */

func TestGeometryCreatePolygonalSliceAtOrigin(t *testing.T) {
	// this method should succeed
	p := CreatePolygonalSliceAtOrigin(5, 1.0, dyn4go.DegToRad(30))
	// and the center should be the origin
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().Y, 1.0e-3)
}

/**
 * Tests the createPolygonalSlice method with an invalid count.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalSliceInvalidCount(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalSlice(0, 1.0, dyn4go.DegToRad(30))
}

/**
 * Tests the createPolygonalSlice method with a negative radius.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalSliceNegativeRadius(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalSlice(5, -1, dyn4go.DegToRad(30))
}

/**
 * Tests the createPolygonalSlice method with a zero radius.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalSliceZeroRadius(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalSlice(5, 0, dyn4go.DegToRad(30))
}

/**
 * Tests the createPolygonalSlice method with a negative theta.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalSliceThetaLessThanZero(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalSlice(5, 1.0, -dyn4go.DegToRad(30))
}

/**
 * Tests the createPolygonalSlice method with theta equal to zero.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalSliceThetaLessZero(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalSlice(5, 1.0, 0)
}

/**
 * Tests the createPolygonalSlice method with theta greater than 180 degrees.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalSliceThetaGreaterThan180(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalSlice(5, 1.0, dyn4go.DegToRad(190))
}

/**
 * Tests the createPolygonalHalfEllipse method.
 * @since 3.1.5
 */

func TestGeometryCreatePolygonalHalfEllipse(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	// this method should succeed
	p := CreatePolygonalHalfEllipse(5, 1.0, 0.5)
	// the center should not be at the origin
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.103, p.GetCenter().Y, 1.0e-3)
}

/**
 * Tests the createPolygonalHalfEllipseAtOrigin method.
 * @since 3.1.5
 */

func TestGeometryCreatePolygonalHalfEllipseAtOrigin(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	// this method should succeed
	p := CreatePolygonalHalfEllipseAtOrigin(5, 1.0, 0.5)
	// the center should be at the origin
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().Y, 1.0e-3)
}

/**
 * Tests the createPolygonalHalfEllipse method with an invalid count.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalHalfEllipseInvalidCount(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalHalfEllipse(0, 1.0, 0.5)
}

/**
 * Tests the createPolygonalHalfEllipse method with a negative width.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalHalfEllipseZeroWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalHalfEllipse(5, 0, 0.5)
}

/**
 * Tests the createPolygonalHalfEllipse method with zero width.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalHalfEllipseNegativeWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalHalfEllipse(5, -1, 0.5)
}

/**
 * Tests the createPolygonalHalfEllipse method with a negative height.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalHalfEllipseNegativeHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalHalfEllipse(5, 1.0, -0.5)
}

/**
 * Tests the createPolygonalHalfEllipse method with zero height.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalHalfEllipseZeroHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalHalfEllipse(5, 1.0, 0)
}

/**
 * Tests the createPolygonalCapsule method.
 * @since 3.1.5
 */

func TestGeometryCreatePolygonalCapsule(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	// this method should succeed
	p := CreatePolygonalCapsule(5, 1.0, 0.5)
	// the center should be at the origin
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.GetCenter().Y, 1.0e-3)
}

/**
 * Tests the createPolygonalCapsule method with an invalid count.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalCapsuleInvalidCount(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalCapsule(0, 1.0, 0.5)
}

/**
 * Tests the createPolygonalCapsule method with zero width.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalCapsuleZeroWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalCapsule(5, 0, 0.5)
}

/**
 * Tests the createPolygonalCapsule method with a negative width.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalCapsuleNegativeWidth(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalCapsule(5, -1, 0.5)
}

/**
 * Tests the createPolygonalCapsule method with zero height.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalCapsuleZeroHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalCapsule(5, 1.0, 0)
}

/**
 * Tests the createPolygonalCapsule method with zero width.
 * @since 3.1.5
 */
func TestGeometryCreatePolygonalCapsuleNegativeHeight(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	CreatePolygonalCapsule(5, 1.0, -0.5)
}

/**
 * Tests the flip polygon method.
 * @since 3.1.4
 */

func TestGeometryFlip(t *testing.T) {
	p := CreateUnitCirclePolygon(5, 1.0)

	// flip about an arbitrary vector and point (line)
	flipped := FlipVector2(p, NewVector2FromXY(1.0, 1.0), NewVector2FromXY(0.0, 2.0))

	vertices := flipped.GetVertices()
	dyn4go.AssertEqualWithinError(t, -2.951, vertices[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.309, vertices[0].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -2.587, vertices[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.190, vertices[1].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.412, vertices[2].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.190, vertices[2].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.048, vertices[3].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.309, vertices[3].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -2.000, vertices[4].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, vertices[4].Y, 1.0e-3)
}

/**
 * Tests the flip polygon method with a nil polygon.
 * @since 3.1.4
 */
func TestGeometryFlipNullPolygon(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	FlipVector2(nil, NewVector2FromXY(1.0, 1.0), nil)
}

/**
 * Tests the flip polygon method with a nil axis.
 * @since 3.1.4
 */
func TestGeometryFlipNullAxis(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	FlipVector2(CreateSquare(1.0), nil, nil)
}

/**
 * Tests the flip polygon method with a zero vector axis.
 * @since 3.1.4
 */
func TestGeometryFlipZeroAxis(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	FlipVector2(CreateSquare(1.0), new(Vector2), nil)
}

/**
 * Tests the flip polygon method with a nil point.
 * @since 3.1.4
 */

func TestGeometryFlipNullPoint(t *testing.T) {
	// it should use the center
	FlipVector2(CreateSquare(1.0), NewVector2FromXY(1.0, 1.0), nil)
}

/**
 * Test the minkowski sum method.
 * @since 3.1.5
 */

func TestGeometryMinkowskiSum(t *testing.T) {
	// verify the generation of the polygon works
	p := MinkowskiSumPolygonCircleInt(CreateUnitCirclePolygon(5, 0.5), CreateCircle(0.2), 3)
	// verify the new vertex count
	dyn4go.AssertEqual(t, 25, len(p.vertices))

	// verify the generation of the polygon works
	p = MinkowskiSumPolygonFloat64Int(CreateUnitCirclePolygon(5, 0.5), 0.2, 3)
	// verify the new vertex count
	dyn4go.AssertEqual(t, 25, len(p.vertices))

	// verify the generation of the polygon works
	p = MinkowskiSum(CreateSquare(1.0), CreateUnitCirclePolygon(5, 0.2))
	dyn4go.AssertEqual(t, 8, len(p.vertices))

	// verify the generation of the polygon works
	p = MinkowskiSum(CreateSegmentEnd(NewVector2FromXY(1.0, 0.0)), CreateUnitCirclePolygon(5, 0.2))
	dyn4go.AssertEqual(t, 5, len(p.vertices))

	// verify the generation of the polygon works
	p = MinkowskiSum(CreateSegmentEnd(NewVector2FromXY(1.0, 0.0)), CreateSegmentEnd(NewVector2FromXY(0.5, 0.5)))
	dyn4go.AssertEqual(t, 4, len(p.vertices))
}

/**
 * Test the minkowski sum method with invalid segments.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumInvalidSegments(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSum(CreateSegmentEnd(NewVector2FromXY(1.0, 0.0)), CreateSegmentEnd(NewVector2FromXY(-0.5, 0.0)))
}

/**
 * Test the minkowski sum method given a nil shape.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumNullWound1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSum(nil, CreateUnitCirclePolygon(5, 0.5))
}

/**
 * Test the minkowski sum method given a nil shape.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumNullWound2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSum(CreateUnitCirclePolygon(5, 0.5), nil)
}

/**
 * Test the minkowski sum method given a nil shape.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumNullShape1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonCircleInt(nil, CreateCircle(0.2), 3)
}

/**
 * Test the minkowski sum method given a nil shape.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumNullShape2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonCircleInt(CreateUnitCirclePolygon(5, 0.5), nil, 3)
}

/**
 * Test the minkowski sum method given a nil shape.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumNullShape3(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonFloat64Int(nil, 0.2, 3)
}

/**
 * Test the minkowski sum method given an invalid count.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumInvalidCount1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonFloat64Int(CreateUnitCirclePolygon(5, 0.5), 0.2, 0)
}

/**
 * Test the minkowski sum method given an invalid count.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumInvalidCount2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonFloat64Int(CreateUnitCirclePolygon(5, 0.5), 0.2, -2)
}

/**
 * Test the minkowski sum method given an invalid count.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumInvalidCount3(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonCircleInt(CreateUnitCirclePolygon(5, 0.5), CreateCircle(0.5), 0)
}

/**
 * Test the minkowski sum method given an invalid count.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumInvalidCount4(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonCircleInt(CreateUnitCirclePolygon(5, 0.5), CreateCircle(0.5), -2)
}

/**
 * Test the minkowski sum method given an invalid radius.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumInvalidRadius1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonFloat64Int(CreateUnitCirclePolygon(5, 0.5), 0, 3)
}

/**
 * Test the minkowski sum method given an invalid radius.
 * @since 3.1.5
 */
func TestGeometryMinkowskiSumInvalidRadius2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	MinkowskiSumPolygonFloat64Int(CreateUnitCirclePolygon(5, 0.5), -2.0, 3)
}

/**
 * Tests that the scale methods work as expected.
 * @since 3.1.5
 */

func TestGeometryScale(t *testing.T) {
	s1 := ScaleCircle(CreateCircle(0.5), 2)
	s2 := ScaleCapsule(CreateCapsule(1.0, 0.5), 2)
	s3 := ScaleEllipse(CreateEllipse(1.0, 0.5), 2)
	s4 := ScaleHalfEllipse(CreateHalfEllipse(1.0, 0.25), 2)
	s5 := ScaleSlice(CreateSlice(0.5, dyn4go.DegToRad(30)), 2)
	s6 := ScalePolygon(CreateUnitCirclePolygon(5, 0.5), 2)
	s7 := ScaleSegment(CreateSegmentEnd(NewVector2FromXY(1.0, 0.0)), 2)

	dyn4go.AssertEqualWithinError(t, 1.000, s1.radius, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, s2.length, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, s2.capRadius*2.0, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, s3.width, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, s3.height, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, s4.width, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, s4.height, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, s5.sliceRadius, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, s6.radius, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.000, s7.length, 1.0e-3)

	s1 = ScaleCircle(CreateCircle(0.5), 0.5)
	s2 = ScaleCapsule(CreateCapsule(1.0, 0.5), 0.5)
	s3 = ScaleEllipse(CreateEllipse(1.0, 0.5), 0.5)
	s4 = ScaleHalfEllipse(CreateHalfEllipse(1.0, 0.25), 0.5)
	s5 = ScaleSlice(CreateSlice(0.5, dyn4go.DegToRad(30)), 0.5)
	s6 = ScalePolygon(CreateUnitCirclePolygon(5, 0.5), 0.5)
	s7 = ScaleSegment(CreateSegmentEnd(NewVector2FromXY(1.0, 0.0)), 0.5)

	dyn4go.AssertEqualWithinError(t, 0.250, s1.radius, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, s2.length, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, s2.capRadius*2.0, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, s3.width, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, s3.height, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, s4.width, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.125, s4.height, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, s5.sliceRadius, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.250, s6.radius, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.500, s7.length, 1.0e-3)
}

/**
 * Tests that the scale method fails if given a nil shape.
 * @since 3.1.5
 */
func TestGeometryScaleNullCircle(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleCircle(nil, 1.2)
}

/**
 * Tests that the scale method fails if given a nil shape.
 * @since 3.1.5
 */
func TestGeometryScaleNullCapsule(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleCapsule(nil, 1.2)
}

/**
 * Tests that the scale method fails if given a nil shape.
 * @since 3.1.5
 */
func TestGeometryScaleNullEllipse(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleEllipse(nil, 1.2)
}

/**
 * Tests that the scale method fails if given a nil shape.
 * @since 3.1.5
 */
func TestGeometryScaleNullHalfEllipse(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleHalfEllipse(nil, 1.2)
}

/**
 * Tests that the scale method fails if given a nil shape.
 * @since 3.1.5
 */
func TestGeometryScaleNullSlice(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleSlice(nil, 1.2)
}

/**
 * Tests that the scale method fails if given a nil shape.
 * @since 3.1.5
 */
func TestGeometryScaleNullPolygon(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScalePolygon(nil, 1.2)
}

/**
 * Tests that the scale method fails if given a nil shape.
 * @since 3.1.5
 */
func TestGeometryScaleNullSegment(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleSegment(nil, 1.2)
}

/**
 * Tests that the scale method fails if given an invalid scale factor.
 * @since 3.1.5
 */
func TestGeometryScaleCircleInvalid(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleCircle(CreateCircle(0.5), 0)
}

/**
 * Tests that the scale method fails if given an invalid scale factor.
 * @since 3.1.5
 */
func TestGeometryScaleCapsuleInvalid(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleCapsule(CreateCapsule(1.0, 0.5), 0)
}

/**
 * Tests that the scale method fails if given an invalid scale factor.
 * @since 3.1.5
 */
func TestGeometryScaleEllipseInvalid(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleEllipse(CreateEllipse(1.0, 0.5), 0)
}

/**
 * Tests that the scale method fails if given an invalid scale factor.
 * @since 3.1.5
 */
func TestGeometryScaleHalfEllipseInvalid(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleHalfEllipse(CreateHalfEllipse(1.0, 0.25), 0)
}

/**
 * Tests that the scale method fails if given an invalid scale factor.
 * @since 3.1.5
 */
func TestGeometryScaleSliceInvalid(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleSlice(CreateSlice(0.5, dyn4go.DegToRad(30)), 0)
}

/**
 * Tests that the scale method fails if given an invalid scale factor.
 * @since 3.1.5
 */
func TestGeometryScalePolygonInvalid(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScalePolygon(CreateUnitCirclePolygon(5, 0.5), 0)
}

/**
 * Tests that the scale method fails if given an invalid scale factor.
 * @since 3.1.5
 */
func TestGeometryScaleSegmentInvalid(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	ScaleSegment(CreateSegmentEnd(NewVector2FromXY(1.0, 1.0)), 0)
}
