package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests the min > max.
 */

func TestIntervalCreateMinGreaterThanMax(t *testing.T) {
	defer dyn4go.AssertPanic(t)
	NewIntervalFromMinMax(0.0, -1.0)
}

/**
 * Tests the constructor.
 */

func TestIntervalCreate(t *testing.T) {
	defer dyn4go.AssertNoPanic(t)
	NewIntervalFromMinMax(0.0, 2.0)
}

/**
 * Tests the copy constructor.
 */

func TestIntervalCreateCopy(t *testing.T) {
	i1 := NewIntervalFromMinMax(-1.0, 2.0)
	i2 := NewIntervalFromInterval(i1)

	dyn4go.AssertNotEqual(t, i2, i1)
	dyn4go.AssertEqual(t, i1.min, i2.min)
	dyn4go.AssertEqual(t, i1.max, i2.max)
}

/**
 * Tests an invalid max.
 */

func TestIntervalSetInvalidMax(t *testing.T) {
	i := NewIntervalFromMinMax(0.0, 2.0)
	i.SetMax(-1.0)
}

/**
 * Tests a valid max.
 */

func TestIntervalSetMax(t *testing.T) {
	i := NewIntervalFromMinMax(0.0, 2.0)
	i.SetMax(1.5)
}

/**
 * Tests an invalid min.
 */

func TestIntervalSetInvalidMin(t *testing.T) {
	i := NewIntervalFromMinMax(0.0, 2.0)
	i.SetMin(3.0)
}

/**
 * Tests a valid max.
 */

func TestIntervalSetMin(t *testing.T) {
	i := NewIntervalFromMinMax(0.0, 2.0)
	i.SetMin(1.5)
}

/**
 * Tests the includes methods.
 */

func TestIntervalIncludes(t *testing.T) {
	i := NewIntervalFromMinMax(-2.5, 100.521)

	dyn4go.AssertTrue(t, i.IncludesExclusive(50.0))
	dyn4go.AssertTrue(t, !i.IncludesExclusive(100.521))
	dyn4go.AssertTrue(t, !i.IncludesExclusive(-3.0))

	dyn4go.AssertTrue(t, i.IncludesInclusive(50.0))
	dyn4go.AssertTrue(t, i.IncludesInclusive(-2.5))
	dyn4go.AssertTrue(t, !i.IncludesInclusive(-3.0))

	dyn4go.AssertTrue(t, i.IncludesInclusiveMax(50.0))
	dyn4go.AssertTrue(t, i.IncludesInclusiveMax(100.521))
	dyn4go.AssertTrue(t, !i.IncludesInclusiveMax(-2.5))
	dyn4go.AssertTrue(t, !i.IncludesInclusiveMax(200.0))

	dyn4go.AssertTrue(t, i.IncludesInclusiveMin(50.0))
	dyn4go.AssertTrue(t, i.IncludesInclusiveMin(-2.5))
	dyn4go.AssertTrue(t, !i.IncludesInclusiveMin(100.521))
	dyn4go.AssertTrue(t, !i.IncludesInclusiveMin(-3.0))
}

/**
 * Tests the overlap methods.
 */

func TestIntervalOverlaps(t *testing.T) {
	i1 := NewIntervalFromMinMax(-2.0, 5.0)
	i2 := NewIntervalFromMinMax(-4.0, 1.0)

	dyn4go.AssertTrue(t, i1.Overlaps(i2))
	// the reverse should work also
	dyn4go.AssertTrue(t, i2.Overlaps(i1))

	// distance should return zero
	dyn4go.AssertEqual(t, 0.0, i1.Distance(i2))
	dyn4go.AssertEqual(t, 0.0, i2.Distance(i1))

	// contains should return false
	dyn4go.AssertTrue(t, !i1.Contains(i2))
	dyn4go.AssertTrue(t, !i2.Contains(i1))

	ov1 := i1.GetOverlap(i2)
	ov2 := i2.GetOverlap(i1)

	dyn4go.AssertEqual(t, 3.0, ov1)
	dyn4go.AssertEqual(t, 3.0, ov2)
}

/**
 * Tests the clamp methods.
 */

func TestIntervalClamp(t *testing.T) {
	i := NewIntervalFromMinMax(-1.0, 6.5)

	dyn4go.AssertEqual(t, 2.0, i.Clamp(2.0))
	dyn4go.AssertEqual(t, 2.0, IntervalClamp(2.0, -1.0, 6.5))
	dyn4go.AssertEqual(t, -1.0, i.Clamp(-2.0))
	dyn4go.AssertEqual(t, 6.5, i.Clamp(7.0))
}

/**
 * Tests the degenerate interval methods.
 */

func TestIntervalDegenerate(t *testing.T) {
	i := NewIntervalFromMinMax(2.0, 2.0)

	dyn4go.AssertTrue(t, i.IsDegenerate())

	i.Expand(0.000001)

	dyn4go.AssertEqual(t, 1.9999995, i.min)
	dyn4go.AssertEqual(t, 2.0000005, i.max)

	dyn4go.AssertTrue(t, !i.IsDegenerate())
	dyn4go.AssertTrue(t, i.IsDegenerateWithError(0.01))
}

/**
 * Tests the union methods.
 */

func TestIntervalUnion(t *testing.T) {
	i1 := NewIntervalFromMinMax(-2.0, 3.0)
	i2 := NewIntervalFromMinMax(-1.0, 4.0)

	u := i1.GetUnion(i2)
	dyn4go.AssertEqual(t, -2.0, u.min)
	dyn4go.AssertEqual(t, 4.0, u.max)

	// test cumulativity
	u = i2.GetUnion(i1)
	dyn4go.AssertEqual(t, -2.0, u.min)
	dyn4go.AssertEqual(t, 4.0, u.max)

	// test intervals that dont overlap
	i3 := NewIntervalFromMinMax(-3.0, -2.5)
	u = i1.GetUnion(i3)
	dyn4go.AssertEqual(t, -3.0, u.min)
	dyn4go.AssertEqual(t, 3.0, u.max)
}

/**
 * Test the intersection methods.
 */

func TestIntervalIntersection(t *testing.T) {
	i1 := NewIntervalFromMinMax(-2.0, 3.0)
	i2 := NewIntervalFromMinMax(-1.0, 4.0)

	u := i1.GetIntersection(i2)
	dyn4go.AssertEqual(t, -1.0, u.min)
	dyn4go.AssertEqual(t, 3.0, u.max)

	// test cumulativity
	u = i2.GetIntersection(i1)
	dyn4go.AssertEqual(t, -1.0, u.min)
	dyn4go.AssertEqual(t, 3.0, u.max)

	// test intervals that dont overlap
	i3 := NewIntervalFromMinMax(-3.0, -2.5)
	u = i1.GetIntersection(i3)
	dyn4go.AssertEqual(t, 0.0, u.min)
	dyn4go.AssertEqual(t, 0.0, u.max)
}

/**
 * Tests the distance method.
 * @since 3.1.0
 */

func TestIntervalDistance(t *testing.T) {
	i1 := NewIntervalFromMinMax(-2.0, 3.0)
	i2 := NewIntervalFromMinMax(-1.0, 4.0)

	// overlapping intervals should return 0
	dyn4go.AssertEqual(t, 0.0, i1.Distance(i2))

	i2 = NewIntervalFromMinMax(4.0, 6.0)

	dyn4go.AssertEqual(t, 1.0, i1.Distance(i2))
	dyn4go.AssertEqual(t, 1.0, i2.Distance(i1))
}

/**
 * Tests the expand method.
 * @since 3.1.0
 */

func TestIntervalExpand(t *testing.T) {
	i := NewIntervalFromMinMax(-2.0, 2.0)

	// test a normal expansion
	ci := i.GetExpanded(2.0)
	dyn4go.AssertEqualWithinError(t, -3.0, ci.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.0, ci.max, 1.0e-3)
	i.Expand(2.0)
	dyn4go.AssertEqualWithinError(t, -3.0, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.0, i.max, 1.0e-3)

	// test no expansion
	ci = i.GetExpanded(0.0)
	dyn4go.AssertEqualWithinError(t, -3.0, ci.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.0, ci.max, 1.0e-3)
	i.Expand(0.0)
	dyn4go.AssertEqualWithinError(t, -3.0, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.0, i.max, 1.0e-3)

	// test negative expansion
	ci = i.GetExpanded(-1.0)
	dyn4go.AssertEqualWithinError(t, -2.5, ci.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.5, ci.max, 1.0e-3)
	i.Expand(-1.0)
	dyn4go.AssertEqualWithinError(t, -2.5, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 2.5, i.max, 1.0e-3)

	// test large negative expansion (this should
	// make the interval invalid, so the interval
	// is instead reduced down to a degenerate
	// interval at the mid point
	ci = i.GetExpanded(-6.0)
	dyn4go.AssertEqualWithinError(t, 0.0, ci.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.0, ci.max, 1.0e-3)
	i.Expand(-6.0)
	dyn4go.AssertEqualWithinError(t, 0.0, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.0, i.max, 1.0e-3)

	// make a copy
	i = NewIntervalFromMinMax(-2.5, 1.5)
	ci = i.GetExpanded(-6.0)
	dyn4go.AssertEqualWithinError(t, -0.5, ci.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.5, ci.max, 1.0e-3)
	i.Expand(-6.0)
	dyn4go.AssertEqualWithinError(t, -0.5, i.min, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -0.5, i.max, 1.0e-3)
}

/**
 * Returns the length of the interval.
 */

func TestIntervalGetLength(t *testing.T) {
	i := NewIntervalFromMinMax(-2.0, 2.0)
	dyn4go.AssertEqual(t, 4.0, i.GetLength())

	i = NewIntervalFromMinMax(-1.0, 2.0)
	dyn4go.AssertEqual(t, 3.0, i.GetLength())

	i = NewIntervalFromMinMax(-3.0, -1.0)
	dyn4go.AssertEqual(t, 2.0, i.GetLength())

	i = NewIntervalFromMinMax(2.0, 3.0)
	dyn4go.AssertEqual(t, 1.0, i.GetLength())

	i = NewIntervalFromMinMax(-1.0, 2.0)
	i.Expand(-4.0)
	dyn4go.AssertEqual(t, 0.0, i.GetLength())

	i = NewIntervalFromMinMax(-1.0, -1.0)
	dyn4go.AssertEqual(t, 0.0, i.GetLength())
}
