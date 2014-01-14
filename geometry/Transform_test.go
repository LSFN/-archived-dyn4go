package geometry

import (
	"math"
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Tests the identity method.
 */
func TestIdentity(t *testing.T) {
	trans := NewTransform()
	trans.TranslateXY(5, 2)

	trans.Identity()

	dyn4go.AssertEqual(t, 0.0, trans.X)
	dyn4go.AssertEqual(t, 0.0, trans.Y)
}

/**
 * Test the translate method.
 */
func TestTranslate(t *testing.T) {
	trans := NewTransform()
	trans.TranslateXY(2, -1)

	trans.TranslateXY(4, 4)

	dyn4go.AssertEqualWithinError(t, 6.000, trans.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 3.000, trans.Y, 1.0e-3)
}

/**
 * Tests the rotate method.
 */
func TestRotate(t *testing.T) {
	trans := NewTransform()
	trans.RotateAboutOrigin(math.Pi / 6)

	r := trans.GetRotation()

	dyn4go.AssertEqualWithinError(t, 30.000, math.Floor(dyn4go.RadToDeg(r)+0.5), 1.0e-3)

	trans.Identity()

	trans.TranslateXY(5, 5)
	trans.RotateAboutOrigin(math.Pi / 2)

	v := trans.GetTranslation()
	dyn4go.AssertEqualWithinError(t, -5.000, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 5.000, v.Y, 1.0e-3)

	trans.RotateAboutOrigin(math.Pi / 2)
	v = trans.GetTranslation()
	dyn4go.AssertEqualWithinError(t, -5.000, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -5.000, v.Y, 1.0e-3)

	trans.RotateAboutXY(math.Pi*7/36, -5.0, -5.0)
	v = trans.GetTranslation()
	dyn4go.AssertEqualWithinError(t, -5.000, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -5.000, v.Y, 1.0e-3)

	trans.RotateAboutXY(math.Pi/4, -1.0, -1.0)
	v = trans.GetTranslation()
	dyn4go.AssertEqualWithinError(t, -1.000, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, -6.656, v.Y, 1.0e-3)
}

/**
 * Tests the copy method.
 */
func TestCopy(t *testing.T) {
	trans := NewTransform()
	trans.TranslateXY(2.0, -1.0)
	trans.RotateAboutXY(math.Pi/9, -2.0, 6.0)

	tc := NewTransformFromTransform(trans)

	dyn4go.AssertEqual(t, trans.m00, tc.m00)
	dyn4go.AssertEqual(t, trans.m01, tc.m01)
	dyn4go.AssertEqual(t, trans.m10, tc.m10)
	dyn4go.AssertEqual(t, trans.m11, tc.m11)
	dyn4go.AssertEqual(t, trans.X, tc.X)
	dyn4go.AssertEqual(t, trans.Y, tc.Y)
}

/**
 * Tests the getTransformed methods.
 */
func TestGetTransformed(t *testing.T) {
	trans := NewTransform()
	trans.TranslateXY(2.0, 1.0)
	trans.RotateAboutXY(math.Pi*5/36, 1.0, -1.0)

	v := NewVector2FromXY(1.0, 0.0)

	// test transformation
	vt := trans.GetTransformedVector2(v)
	dyn4go.AssertEqualWithinError(t, 1.967, vt.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.657, vt.Y, 1.0e-3)

	// test inverse transformation
	vt = trans.GetInverseTransformedVector2(vt)
	dyn4go.AssertEqualWithinError(t, 1.000, vt.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, vt.Y, 1.0e-3)

	// test just a rotation transformation
	vt = trans.GetTransformedR(v)
	dyn4go.AssertEqualWithinError(t, 0.906, vt.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.422, vt.Y, 1.0e-3)

	// test inverse rotation transformation
	vt = trans.GetInverseTransformedR(v)
	vt = trans.GetTransformedR(vt)
	dyn4go.AssertTrue(t, *vt == *v)
}

/**
 * Tests the transform methods.
 * @since 3.1.0
 */
func TestTransform(t *testing.T) {
	trans := NewTransform()
	trans.TranslateXY(2.0, 1.0)
	trans.RotateAboutXY(math.Pi*5/36, 1.0, -1.0)

	v := NewVector2FromXY(1.0, 0.0)

	// test transformation
	trans.Transform(v)
	dyn4go.AssertEqualWithinError(t, 1.967, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 1.657, v.Y, 1.0e-3)

	// test inverse transformation
	trans.InverseTransform(v)
	dyn4go.AssertEqualWithinError(t, 1.000, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.000, v.Y, 1.0e-3)

	// test just a rotation transformation
	trans.TransformR(v)
	dyn4go.AssertEqualWithinError(t, 0.906, v.X, 1.0e-3)
	dyn4go.AssertEqualWithinError(t, 0.422, v.Y, 1.0e-3)

	// test inverse rotation transformation
	trans.InverseTransformR(v)
	trans.TransformR(v)
	dyn4go.AssertTrue(t, v == v)
}

/**
 * Tests the setTransform method.
 */
func TestSetTransform(t *testing.T) {
	tx := NewTransform()
	tx.RotateAboutOrigin(math.Pi / 6)
	tx.TranslateXY(2.0, 0.5)
	tx2 := NewTransform()
	tx2.Set(tx)

	// shouldnt be the same object reference
	dyn4go.AssertNotEqual(t, tx2, tx)

	// should be the same transformation
	dyn4go.AssertEqual(t, tx.m00, tx2.m00)
	dyn4go.AssertEqual(t, tx.m01, tx2.m01)
	dyn4go.AssertEqual(t, tx.m10, tx2.m10)
	dyn4go.AssertEqual(t, tx.m11, tx2.m11)
	dyn4go.AssertEqual(t, tx.X, tx2.X)
	dyn4go.AssertEqual(t, tx.Y, tx2.Y)
}

/**
 * Tests the setTranslation methods.
 */
func TestSetTranslation(t *testing.T) {
	tx := NewTransform()
	tx.TranslateXY(1.0, 2.0)
	tx.RotateAboutOrigin(math.Pi / 4)
	tx.SetTranslationFromXY(0.0, 0.0)

	dyn4go.AssertEqual(t, 0.0, tx.X)
	dyn4go.AssertEqual(t, 0.0, tx.Y)
	dyn4go.AssertEqualWithinError(t, math.Pi/4, tx.GetRotation(), 1.0e-3)

	tx.X = 2.0
	dyn4go.AssertEqual(t, 2.0, tx.X)
	dyn4go.AssertEqual(t, 0.0, tx.Y)
	dyn4go.AssertEqualWithinError(t, math.Pi/4, tx.GetRotation(), 1.0e-3)

	tx.Y = 3.0
	dyn4go.AssertEqual(t, 2.0, tx.X)
	dyn4go.AssertEqual(t, 3.0, tx.Y)
	dyn4go.AssertEqualWithinError(t, math.Pi/4, tx.GetRotation(), 1.0e-3)
}

/**
 * Tests the setRotation method.
 */
func TestSetRotation(t *testing.T) {
	tx := NewTransform()
	tx.RotateAboutOrigin(math.Pi / 4)
	tx.TranslateXY(1.0, 0.0)

	tx.SetRotation(math.Pi / 6)
	dyn4go.AssertEqualWithinError(t, 30.000, dyn4go.RadToDeg(tx.GetRotation()), 1.0e-3)
	dyn4go.AssertEqual(t, 1.0, tx.X)
	dyn4go.AssertEqual(t, 0.0, tx.Y)
}

/**
 * Tests the linear interpolation methods.
 */
func TestLerp(t *testing.T) {
	p := new(Vector2)

	start := NewTransform()
	start.TranslateXY(1.0, 0.0)
	start.RotateAboutOrigin(math.Pi / 4)

	end := NewTransform()
	end.Set(start)
	end.TranslateXY(3.0, 2.0)
	end.RotateAboutOrigin(math.Pi / 9)

	s := start.GetTransformedVector2(p)
	e := end.GetTransformedVector2(p)

	alpha := 0.5

	mid := NewTransform()
	start.LerpInDestination(end, alpha, mid)
	start.Lerp(end, alpha)

	m := mid.GetTransformedVector2(p)
	// this test only works this way for the mid point
	// otherwise we would have to replicate the lerp method
	dyn4go.AssertEqual(t, (s.X+e.X)*alpha, m.X)
	dyn4go.AssertEqual(t, (s.Y+e.Y)*alpha, m.Y)

	m = start.GetTransformedVector2(p)
	// this test only works this way for the mid point
	// otherwise we would have to replicate the lerp method
	dyn4go.AssertEqual(t, (s.X+e.X)*alpha, m.X)
	dyn4go.AssertEqual(t, (s.Y+e.Y)*alpha, m.Y)

	// test opposing sign angles
	start.Identity()
	start.RotateAboutOrigin(math.Pi * 29 / 30)

	end.Identity()
	end.RotateAboutOrigin(-math.Pi * 28 / 30)

	l := start.Lerped(end, alpha)
	dyn4go.AssertEqualWithinError(t, -3.089, l.GetRotation(), 1.0e-3)

	// test opposing sign angles
	start.Identity()
	start.RotateAboutOrigin(math.Pi * 59 / 60)

	end.Identity()
	end.RotateAboutOrigin(math.Pi / 90)

	l = start.Lerped(end, alpha)
	dyn4go.AssertEqualWithinError(t, -0.034, l.GetRotation(), 1.0e-3)
}

/**
 * Tests the getValues method.
 * @since 3.0.1
 */
func TestValues(t *testing.T) {
	trans := NewTransform()
	trans.TranslateXY(2.0, -1.0)

	values := trans.GetValues()
	dyn4go.AssertEqual(t, 1.0, values[0])
	dyn4go.AssertEqual(t, 0.0, values[1])
	dyn4go.AssertEqual(t, 2.0, values[2])
	dyn4go.AssertEqual(t, 0.0, values[3])
	dyn4go.AssertEqual(t, 1.0, values[4])
	dyn4go.AssertEqual(t, -1.0, values[5])
}
