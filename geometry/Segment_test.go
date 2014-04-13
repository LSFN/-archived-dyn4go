package geometry2

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

func TestSegmentInterfaces(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(1.5, 3.0),
	)
	var _ Convexer = s
	var _ Wounder = s
}

/**
 * Tests a failed create using one nil point.
 * @since 3.1.0
 */
func TestSegmentCreateNullPoint1(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewSegment(nil, new(Vector2))
}

/**
 * Tests a failed create using one nil point.
 * @since 3.1.0
 */
func TestSegmentCreateNullPoint2(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewSegment(new(Vector2), nil)
}

/**
 * Tests coincident points.
 */
func TestSegmentCreateCoincident(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewSegment(new(Vector2), new(Vector2))
}

/**
 * Tests a successful creation.
 */
func TestSegmentCreatSuccess(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	s := NewSegment(
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(1.0, 2.0),
	)

	dyn4go.AssertEqualWithinError(t, 0.500, s.center.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.500, s.center.Y, 1.0e-3)
}

/**
 * Tests the length method.
 */
func TestSegmentGetLength(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(1.5, 3.0),
	)

	dyn4go.AssertEqualWithinError(t, 2.500, s.GetLength(), 1.0e-3)
}

/**
 * Tests the getLocation method.
 */
func TestSegmentGetLocation(t *testing.T) {
	// test invalid line
	loc := GetLocation(NewVector2FromXY(1.0, 1.0), new(Vector2), new(Vector2))
	dyn4go.AssertEqualWithinError(t, 0.000, loc, 1.0e-3)

	// test valid line/on line
	loc = GetLocation(NewVector2FromXY(1.0, 1.0), new(Vector2), NewVector2FromXY(2.0, 2.0))
	dyn4go.AssertEqualWithinError(t, 0.000, loc, 1.0e-3)

	// test valid line/left-above line
	loc = GetLocation(NewVector2FromXY(1.0, 1.0), new(Vector2), NewVector2FromXY(1.0, 0.5))
	dyn4go.AssertTrue(t, loc > 0)

	// test valid line/right-below line
	loc = GetLocation(NewVector2FromXY(1.0, 1.0), new(Vector2), NewVector2FromXY(1.0, 2.0))
	dyn4go.AssertTrue(t, loc < 0)
}

/**
 * Tests the get closest point methods.
 */
func TestSegmentGetPointClosest(t *testing.T) {
	pt := NewVector2FromXY(1.0, -1.0)

	// test invalid line/segment
	p := GetPointOnLineClosestToPoint(pt, NewVector2FromXY(1.0, 1.0), NewVector2FromXY(1.0, 1.0))
	dyn4go.AssertEqualWithinError(t, 1.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, p.Y, 1.0e-3)

	p = GetPointOnSegmentClosestToPoint(pt, NewVector2FromXY(1.0, 1.0), NewVector2FromXY(1.0, 1.0))
	dyn4go.AssertEqualWithinError(t, 1.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, p.Y, 1.0e-3)

	// test valid line
	p = GetPointOnLineClosestToPoint(pt, new(Vector2), NewVector2FromXY(5.0, 5.0))
	// since 0,0 is perp to pt
	dyn4go.AssertEqualWithinError(t, 0.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.Y, 1.0e-3)

	p = GetPointOnLineClosestToPoint(pt, new(Vector2), NewVector2FromXY(2.5, 5.0))
	dyn4go.AssertEqualWithinError(t, -0.200, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.400, p.Y, 1.0e-3)

	// test valid segment
	p = GetPointOnSegmentClosestToPoint(pt, NewVector2FromXY(-1.0, -1.0), NewVector2FromXY(1.0, 1.0))
	// since 0,0 is perp to pt
	dyn4go.AssertEqualWithinError(t, 0.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.Y, 1.0e-3)

	// test closest is one of the segment points
	p = GetPointOnSegmentClosestToPoint(pt, new(Vector2), NewVector2FromXY(2.5, 5.0))
	dyn4go.AssertEqualWithinError(t, 0.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.Y, 1.0e-3)
}

/**
 * Tests the getAxes method.
 */
func TestSegmentGetAxes(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(1.5, 3.0),
	)
	transform := NewTransform()

	axes := s.GetAxes(nil, transform)

	dyn4go.AssertEqual(t, 2, len(axes))

	seg := s.vertices[0].HereToVector2(s.vertices[1])
	// one should be the line itself and the other should be the perp
	dyn4go.AssertEqualWithinError(t, 0.000, seg.CrossVector2(axes[1]), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, seg.DotVector2(axes[0]), 1.0e-3)

	// perform some transformations
	transform.TranslateXY(1.0, 0.0)
	transform.RotateAboutOrigin(dyn4go.DegToRad(25))

	axes = s.GetAxes(nil, transform)

	seg = transform.GetTransformedVector2(s.vertices[0]).HereToVector2(transform.GetTransformedVector2(s.vertices[1]))
	// one should the line itself and the other should be the perp
	dyn4go.AssertEqualWithinError(t, 0.000, seg.CrossVector2(axes[1]), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, seg.DotVector2(axes[0]), 1.0e-3)

	// test for some foci
	f := NewVector2FromXY(2.0, -2.0)
	transform.Identity()

	axes = s.GetAxes([]*Vector2{f}, transform)

	dyn4go.AssertEqual(t, 3, len(axes))

	v1 := s.vertices[0].HereToVector2(f)
	v1.Normalize()

	dyn4go.AssertEqualWithinError(t, v1.X, axes[2].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, v1.Y, axes[2].Y, 1.0e-3)
}

/**
 * Tests the getFoci method.
 */
func TestSegmentGetFoci(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(1.5, 3.0),
	)
	transform := NewTransform()

	foci := s.GetFoci(transform)
	dyn4go.AssertTrue(t, len(foci) == 0)
}

/**
 * Tests the contains method.
 */
func TestSegmentContains(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(1.5, 3.0),
	)
	transform := NewTransform()

	dyn4go.AssertFalse(t, s.ContainsVector2Transform(NewVector2FromXY(2.0, 2.0), transform))
	dyn4go.AssertTrue(t, s.ContainsVector2Transform(NewVector2FromXY(0.75, 2.0), transform))
}

/**
 * Tests the contains with radius method.
 */
func TestSegmentContainsRadius(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(1.0, 1.0),
		NewVector2FromXY(-1.0, -1.0),
	)
	transform := NewTransform()

	dyn4go.AssertFalse(t, s.ContainsTransformRadius(NewVector2FromXY(2.0, 2.0), transform, 0.1))
	dyn4go.AssertTrue(t, s.ContainsTransformRadius(NewVector2FromXY(1.05, 1.05), transform, 0.1))
	dyn4go.AssertTrue(t, s.ContainsTransformRadius(NewVector2FromXY(0.505, 0.5), transform, 0.1))
}

/**
 * Tests the project method.
 */
func TestSegmentProject(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(1.5, 3.0),
	)
	transform := NewTransform()
	n := NewVector2FromXY(1.0, 0.0)

	i := s.ProjectVector2Transform(n, transform)

	dyn4go.AssertEqualWithinError(t, 0.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.500, i.max, 1.0e-3)

	n.SetToXY(1.0, 1.0)
	i = s.ProjectVector2Transform(n, transform)

	dyn4go.AssertEqualWithinError(t, 1.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 4.500, i.max, 1.0e-3)

	n.SetToXY(0.0, 1.0)
	i = s.ProjectVector2Transform(n, transform)

	dyn4go.AssertEqualWithinError(t, 1.000, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, i.max, 1.0e-3)

	// transform the segment a bit
	transform.TranslateXY(1.0, 2.0)
	transform.RotateAboutVector2(dyn4go.DegToRad(90), transform.GetTransformedVector2(s.center))

	i = s.ProjectVector2Transform(n, transform)

	dyn4go.AssertEqualWithinError(t, 3.250, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 4.750, i.max, 1.0e-3)
}

/**
 * Tests the getFarthest methods.
 */
func TestSegmentGetFarthest(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 1.0),
		NewVector2FromXY(1.5, 3.0),
	)
	transform := NewTransform()
	n := NewVector2FromXY(1.0, 0.0)

	f := s.GetFarthestFeature(n, transform)
	dyn4go.AssertTrue(t, f.IsEdge())
	e := f.(*Edge)
	dyn4go.AssertEqualWithinError(t, 1.500, e.max.point.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, e.max.point.Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 0.000, e.vertex1.point.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.000, e.vertex1.point.Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.500, e.vertex2.point.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, e.vertex2.point.Y, 1.0e-3)

	p := s.GetFarthestPoint(n, transform)
	dyn4go.AssertEqualWithinError(t, 1.500, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, p.Y, 1.0e-3)

	// move the segment a bit
	transform.TranslateXY(0.0, -1.0)
	transform.RotateAboutOrigin(dyn4go.DegToRad(45))

	p = s.GetFarthestPoint(n, transform)
	dyn4go.AssertEqualWithinError(t, 0.000, p.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, p.Y, 1.0e-3)
}

/**
 * Tests the rotate method.
 */
func TestSegmentRotate(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 0.0),
		NewVector2FromXY(1.0, 1.0),
	)
	s.RotateAboutXY(dyn4go.DegToRad(45), 0, 0)

	dyn4go.AssertEqualWithinError(t, 0.000, s.vertices[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, s.vertices[0].Y, 1.0e-3)

	dyn4go.AssertEqualWithinError(t, 0.000, s.vertices[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.414, s.vertices[1].Y, 1.0e-3)
}

/**
 * Tests the translate method.
 */
func TestSegmentTranslate(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 0.0),
		NewVector2FromXY(1.0, 1.0),
	)
	s.TranslateXY(2.0, -1.0)

	dyn4go.AssertEqualWithinError(t, 2.000, s.vertices[0].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -1.000, s.vertices[0].Y, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, s.vertices[1].X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, s.vertices[1].Y, 1.0e-3)
}

/**
 * Tests the createAABB method.
 * @since 3.1.0
 */
func TestSegmentCreateAABB(t *testing.T) {
	s := NewSegment(
		NewVector2FromXY(0.0, 0.0),
		NewVector2FromXY(1.0, 1.0),
	)

	aabb := s.CreateAABBTransform(NewTransform())
	dyn4go.AssertEqualWithinError(t, 0.0, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.0, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.GetMaxY(), 1.0e-3)

	// try using the default method
	aabb2 := s.CreateAABB()
	dyn4go.AssertEqual(t, aabb.GetMinX(), aabb2.GetMinX())
	dyn4go.AssertEqual(t, aabb.GetMinY(), aabb2.GetMinY())
	dyn4go.AssertEqual(t, aabb.GetMaxX(), aabb2.GetMaxX())
	dyn4go.AssertEqual(t, aabb.GetMaxY(), aabb2.GetMaxY())

	tx := NewTransform()
	tx.RotateAboutOrigin(dyn4go.DegToRad(30.0))
	tx.TranslateXY(1.0, 2.0)
	aabb = s.CreateAABBTransform(tx)
	dyn4go.AssertEqualWithinError(t, 1.0, aabb.GetMinX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.0, aabb.GetMinY(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.366, aabb.GetMaxX(), 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.366, aabb.GetMaxY(), 1.0e-3)
}

/**
 * Tests the getLineIntersection method.
 * @since 3.1.1
 */
func TestSegmentGetLineIntersection(t *testing.T) {
	// normal case
	p := GetLineIntersection(
		NewVector2FromXY(-1.0, -1.0), NewVector2FromXY(2.0, 0.0),
		NewVector2FromXY(-1.0, 0.0), NewVector2FromXY(1.0, 0.5))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 11.0, p.X)
	dyn4go.AssertEqual(t, 3.0, p.Y)

	// try horizontal line
	p = GetLineIntersection(
		NewVector2FromXY(-1.0, 1.0), NewVector2FromXY(2.0, 1.0),
		NewVector2FromXY(-1.0, 0.0), NewVector2FromXY(1.0, 0.5))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 3.0, p.X)
	dyn4go.AssertEqual(t, 1.0, p.Y)

	// try a vertical line
	p = GetLineIntersection(
		NewVector2FromXY(3.0, 0.0), NewVector2FromXY(3.0, 1.0),
		NewVector2FromXY(-1.0, 0.0), NewVector2FromXY(1.0, 0.5))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 3.0, p.X)
	dyn4go.AssertEqual(t, 1.0, p.Y)

	// try a vertical and horizontal line
	p = GetLineIntersection(
		NewVector2FromXY(3.0, 0.0), NewVector2FromXY(3.0, -2.0),
		NewVector2FromXY(0.0, 1.0), NewVector2FromXY(4.0, 1.0))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 3.0, p.X)
	dyn4go.AssertEqual(t, 1.0, p.Y)

	// try two parallel lines
	p = GetLineIntersection(
		NewVector2FromXY(-2.0, -1.0), NewVector2FromXY(-1.0, 0.0),
		NewVector2FromXY(-1.0, -1.0), NewVector2FromXY(0.0, 0.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try two vertical lines (parallel)
	p = GetLineIntersection(
		NewVector2FromXY(3.0, 0.0), NewVector2FromXY(3.0, 1.0),
		NewVector2FromXY(2.0, 0.0), NewVector2FromXY(2.0, 1.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try two horizontal lines (parallel)
	p = GetLineIntersection(
		NewVector2FromXY(3.0, 1.0), NewVector2FromXY(4.0, 1.0),
		NewVector2FromXY(2.0, 2.0), NewVector2FromXY(4.0, 2.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try colinear lines
	p = GetLineIntersection(
		NewVector2FromXY(-1.0, -1.0), NewVector2FromXY(1.0, 1.0),
		NewVector2FromXY(-2.0, -2.0), NewVector2FromXY(-1.5, -1.5))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try colinear vertical lines
	p = GetLineIntersection(
		NewVector2FromXY(3.0, 0.0), NewVector2FromXY(3.0, 1.0),
		NewVector2FromXY(3.0, 2.0), NewVector2FromXY(3.0, 7.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try colinear horizontal lines
	p = GetLineIntersection(
		NewVector2FromXY(4.0, 1.0), NewVector2FromXY(5.0, 1.0),
		NewVector2FromXY(-1.0, 1.0), NewVector2FromXY(1.0, 1.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}
}

/**
 * Tests the getLineIntersection method.
 * @since 3.1.1
 */
func TestSegmentGetSegmentIntersection(t *testing.T) {
	p := GetSegmentIntersection(
		NewVector2FromXY(-3.0, -1.0), NewVector2FromXY(3.0, 1.0),
		NewVector2FromXY(-1.0, -2.0), NewVector2FromXY(1.0, 2.0))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 0.0, p.X)
	dyn4go.AssertEqual(t, 0.0, p.Y)

	// normal case, no intersection
	p = GetSegmentIntersection(
		NewVector2FromXY(-1.0, -1.0), NewVector2FromXY(2.0, 0.0),
		NewVector2FromXY(-1.0, 0.0), NewVector2FromXY(1.0, 0.5))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try horizontal segment
	p = GetSegmentIntersection(
		NewVector2FromXY(-1.0, 1.0), NewVector2FromXY(2.0, 1.0),
		NewVector2FromXY(-1.0, 0.0), NewVector2FromXY(1.0, 2.0))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 0.0, p.X)
	dyn4go.AssertEqual(t, 1.0, p.Y)

	// try a vertical segment
	p = GetSegmentIntersection(
		NewVector2FromXY(3.0, 0.0), NewVector2FromXY(3.0, 3.0),
		NewVector2FromXY(4.0, 0.0), NewVector2FromXY(1.0, 3.0))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 3.0, p.X)
	dyn4go.AssertEqual(t, 1.0, p.Y)

	// try a vertical and horizontal segment
	p = GetSegmentIntersection(
		NewVector2FromXY(3.0, 2.0), NewVector2FromXY(3.0, -2.0),
		NewVector2FromXY(0.0, 1.0), NewVector2FromXY(4.0, 1.0))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 3.0, p.X)
	dyn4go.AssertEqual(t, 1.0, p.Y)

	// try two parallel segments
	p = GetSegmentIntersection(
		NewVector2FromXY(-2.0, -1.0), NewVector2FromXY(-1.0, 0.0),
		NewVector2FromXY(-1.0, -1.0), NewVector2FromXY(0.0, 0.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try two vertical segments (parallel)
	p = GetSegmentIntersection(
		NewVector2FromXY(3.0, 0.0), NewVector2FromXY(3.0, 1.0),
		NewVector2FromXY(2.0, 0.0), NewVector2FromXY(2.0, 1.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try two horizontal segments (parallel)
	p = GetSegmentIntersection(
		NewVector2FromXY(3.0, 1.0), NewVector2FromXY(4.0, 1.0),
		NewVector2FromXY(3.0, 2.0), NewVector2FromXY(4.0, 2.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try colinear segments
	p = GetSegmentIntersection(
		NewVector2FromXY(-1.0, -1.0), NewVector2FromXY(1.0, 1.0),
		NewVector2FromXY(-2.0, -2.0), NewVector2FromXY(-1.5, -1.5))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try colinear vertical segments
	p = GetSegmentIntersection(
		NewVector2FromXY(3.0, 0.0), NewVector2FromXY(3.0, 1.0),
		NewVector2FromXY(3.0, -1.0), NewVector2FromXY(3.0, 7.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try colinear horizontal segments
	p = GetSegmentIntersection(
		NewVector2FromXY(-1.0, 1.0), NewVector2FromXY(5.0, 1.0),
		NewVector2FromXY(-1.0, 1.0), NewVector2FromXY(1.0, 1.0))

	if p != nil {
		t.Error("Value is not nil in assertion")
	}

	// try intersection at end point
	p = GetSegmentIntersection(
		NewVector2FromXY(1.0, 0.0), NewVector2FromXY(3.0, -2.0),
		NewVector2FromXY(-1.0, -1.0), NewVector2FromXY(1.0, 0.0))

	if p == nil {
		t.Error("Value is nil in assertion")
	}
	dyn4go.AssertEqual(t, 1.0, p.X)
	dyn4go.AssertEqual(t, 0.0, p.Y)

	// test segment intersection perpendicular
	s1 := NewSegment(NewVector2FromXY(-10, 10), NewVector2FromXY(10, 10))
	s2 := NewSegment(NewVector2FromXY(0, 0), NewVector2FromXY(0, 5))
	p = s2.GetSegmentIntersection(s1)
	if p != nil {
		t.Error("Value is not nil in assertion")
	}
}
