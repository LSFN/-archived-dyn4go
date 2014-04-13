package geometry

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Test the creation method passing four doubles.
 */
func TestMatrix33CreateFull(t *testing.T) {
	m := NewMatrix33FromFloats(1.0, 2.0, 1.0,
		-3.0, 8.0, 2.0,
		1.0, 5.0, -1.0)
	dyn4go.AssertEqual(t, 1.0, m.m00)
	dyn4go.AssertEqual(t, 2.0, m.m01)
	dyn4go.AssertEqual(t, 1.0, m.m02)
	dyn4go.AssertEqual(t, -3.0, m.m10)
	dyn4go.AssertEqual(t, 8.0, m.m11)
	dyn4go.AssertEqual(t, 2.0, m.m12)
	dyn4go.AssertEqual(t, 1.0, m.m20)
	dyn4go.AssertEqual(t, 5.0, m.m21)
	dyn4go.AssertEqual(t, -1.0, m.m22)
}

/**
 * Test the creation method passing a double array.
 */
func TestMatrix33CreateFullArray(t *testing.T) {
	m := NewMatrix33FromFloatSlice([]float64{1.0, 2.0, 1.0,
		-3.0, 8.0, 2.0,
		1.0, 5.0, -1.0})
	dyn4go.AssertEqual(t, 1.0, m.m00)
	dyn4go.AssertEqual(t, 2.0, m.m01)
	dyn4go.AssertEqual(t, 1.0, m.m02)
	dyn4go.AssertEqual(t, -3.0, m.m10)
	dyn4go.AssertEqual(t, 8.0, m.m11)
	dyn4go.AssertEqual(t, 2.0, m.m12)
	dyn4go.AssertEqual(t, 1.0, m.m20)
	dyn4go.AssertEqual(t, 5.0, m.m21)
	dyn4go.AssertEqual(t, -1.0, m.m22)
}

/**
 * Tests the copy constructor.
 */
func TestMatrix33Copy(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 1.0, -1.0,
		2.0, -1.0, 3.0,
		3.0, 2.5, 0.5)

	// make a copy
	m2 := NewMatrix33FromMatrix33(m1)
	// test the values
	dyn4go.AssertEqual(t, m2.m00, m1.m00)
	dyn4go.AssertEqual(t, m2.m01, m1.m01)
	dyn4go.AssertEqual(t, m2.m02, m1.m02)
	dyn4go.AssertEqual(t, m2.m10, m1.m10)
	dyn4go.AssertEqual(t, m2.m11, m1.m11)
	dyn4go.AssertEqual(t, m2.m12, m1.m12)
	dyn4go.AssertEqual(t, m2.m20, m1.m20)
	dyn4go.AssertEqual(t, m2.m21, m1.m21)
	dyn4go.AssertEqual(t, m2.m22, m1.m22)
}

/**
 * Tests the add method.
 */
func TestMatrix33Add(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m2 := NewMatrix33FromFloats(1.0, 1.0, 3.0,
		0.0, 4.0, 1.0,
		2.0, 2.0, 1.0)
	m1.Add(m2)
	// test the values
	dyn4go.AssertEqual(t, 1.0, m1.m00)
	dyn4go.AssertEqual(t, 3.0, m1.m01)
	dyn4go.AssertEqual(t, 3.0, m1.m02)
	dyn4go.AssertEqual(t, 3.0, m1.m10)
	dyn4go.AssertEqual(t, 5.0, m1.m11)
	dyn4go.AssertEqual(t, 2.0, m1.m12)
	dyn4go.AssertEqual(t, 4.0, m1.m20)
	dyn4go.AssertEqual(t, 2.0, m1.m21)
	dyn4go.AssertEqual(t, 0.0, m1.m22)
}

/**
 * Tests the sum method.
 */
func TestMatrix33Sum(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m2 := NewMatrix33FromFloats(1.0, 1.0, 3.0,
		0.0, 4.0, 1.0,
		2.0, 2.0, 1.0)
	m3 := m1.Sum(m2)
	// test the values
	dyn4go.AssertEqual(t, 1.0, m3.m00)
	dyn4go.AssertEqual(t, 3.0, m3.m01)
	dyn4go.AssertEqual(t, 3.0, m3.m02)
	dyn4go.AssertEqual(t, 3.0, m3.m10)
	dyn4go.AssertEqual(t, 5.0, m3.m11)
	dyn4go.AssertEqual(t, 2.0, m3.m12)
	dyn4go.AssertEqual(t, 4.0, m3.m20)
	dyn4go.AssertEqual(t, 2.0, m3.m21)
	dyn4go.AssertEqual(t, 0.0, m3.m22)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m3)
}

/**
 * Tests the subtract method.
 */
func TestMatrix33Subtract(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m2 := NewMatrix33FromFloats(1.0, 1.0, 3.0,
		0.0, 4.0, 1.0,
		2.0, 2.0, 1.0)
	m1.Subtract(m2)
	// test the values
	dyn4go.AssertEqual(t, -1.0, m1.m00)
	dyn4go.AssertEqual(t, 1.0, m1.m01)
	dyn4go.AssertEqual(t, -3.0, m1.m02)
	dyn4go.AssertEqual(t, 3.0, m1.m10)
	dyn4go.AssertEqual(t, -3.0, m1.m11)
	dyn4go.AssertEqual(t, 0.0, m1.m12)
	dyn4go.AssertEqual(t, 0.0, m1.m20)
	dyn4go.AssertEqual(t, -2.0, m1.m21)
	dyn4go.AssertEqual(t, -2.0, m1.m22)
}

/**
 * Tests the difference method.
 */
func TestMatrix33Difference(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m2 := NewMatrix33FromFloats(1.0, 1.0, 3.0,
		0.0, 4.0, 1.0,
		2.0, 2.0, 1.0)
	m3 := m1.Difference(m2)
	// test the values
	dyn4go.AssertEqual(t, -1.0, m3.m00)
	dyn4go.AssertEqual(t, 1.0, m3.m01)
	dyn4go.AssertEqual(t, -3.0, m3.m02)
	dyn4go.AssertEqual(t, 3.0, m3.m10)
	dyn4go.AssertEqual(t, -3.0, m3.m11)
	dyn4go.AssertEqual(t, 0.0, m3.m12)
	dyn4go.AssertEqual(t, 0.0, m3.m20)
	dyn4go.AssertEqual(t, -2.0, m3.m21)
	dyn4go.AssertEqual(t, -2.0, m3.m22)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m3)
}

/**
 * Tests the multiply matrix method.
 */
func TestMatrix33MultiplyMatrix(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m2 := NewMatrix33FromFloats(1.0, 1.0, 3.0,
		0.0, 4.0, 1.0,
		2.0, 2.0, 1.0)
	m1.MultiplyMatrix33(m2)
	dyn4go.AssertEqual(t, 0.0, m1.m00)
	dyn4go.AssertEqual(t, 8.0, m1.m01)
	dyn4go.AssertEqual(t, 2.0, m1.m02)
	dyn4go.AssertEqual(t, 5.0, m1.m10)
	dyn4go.AssertEqual(t, 9.0, m1.m11)
	dyn4go.AssertEqual(t, 11.0, m1.m12)
	dyn4go.AssertEqual(t, 0.0, m1.m20)
	dyn4go.AssertEqual(t, 0.0, m1.m21)
	dyn4go.AssertEqual(t, 5.0, m1.m22)
}

/**
 * Tests the product matrix method.
 */
func TestMatrix33ProductMatrix(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m2 := NewMatrix33FromFloats(1.0, 1.0, 3.0,
		0.0, 4.0, 1.0,
		2.0, 2.0, 1.0)
	m3 := m1.ProductMatrix33(m2)
	dyn4go.AssertEqual(t, 0.0, m3.m00)
	dyn4go.AssertEqual(t, 8.0, m3.m01)
	dyn4go.AssertEqual(t, 2.0, m3.m02)
	dyn4go.AssertEqual(t, 5.0, m3.m10)
	dyn4go.AssertEqual(t, 9.0, m3.m11)
	dyn4go.AssertEqual(t, 11.0, m3.m12)
	dyn4go.AssertEqual(t, 0.0, m3.m20)
	dyn4go.AssertEqual(t, 0.0, m3.m21)
	dyn4go.AssertEqual(t, 5.0, m3.m22)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m3)
}

/**
 * Tests the multiply vector method.
 */
func TestMatrix33MultiplyVector(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	v1 := NewVector3FromFloats(1.0, -1.0, 2.0)

	m1.MultiplyVector3(v1)
	dyn4go.AssertEqual(t, -2.0, v1.X)
	dyn4go.AssertEqual(t, 4.0, v1.Y)
	dyn4go.AssertEqual(t, 0.0, v1.Z)
}

/**
 * Tests the product vector method.
 */
func TestMatrix33ProductVector(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	v1 := NewVector3FromFloats(1.0, -1.0, 2.0)

	v2 := m1.ProductVector3(v1)
	dyn4go.AssertEqual(t, -2.0, v2.X)
	dyn4go.AssertEqual(t, 4.0, v2.Y)
	dyn4go.AssertEqual(t, 0.0, v2.Z)
	// make sure we didnt modify the first vector
	dyn4go.AssertFalse(t, *v1 == *v2)
}

/**
 * Tests the multiply vector transpose method.
 */
func TestMatrix33MultiplyVectorT(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	v1 := NewVector3FromFloats(1.0, -1.0, 2.0)

	m1.MultiplyTVector3(v1)
	dyn4go.AssertEqual(t, 1.0, v1.X)
	dyn4go.AssertEqual(t, 1.0, v1.Y)
	dyn4go.AssertEqual(t, -3.0, v1.Z)
}

/**
 * Tests the product vector transpose method.
 */
func TestMatrix33ProductVectorT(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	v1 := NewVector3FromFloats(1.0, -1.0, 2.0)

	v2 := m1.ProductTVector3(v1)
	dyn4go.AssertEqual(t, 1.0, v2.X)
	dyn4go.AssertEqual(t, 1.0, v2.Y)
	dyn4go.AssertEqual(t, -3.0, v2.Z)
	// make sure we didnt modify the first vector
	dyn4go.AssertFalse(t, *v1 == *v2)
}

/**
 * Tests the multiply by a scalar method.
 */
func TestMatrix33MultiplyScalar(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m1.MultiplyScalar(2.0)
	dyn4go.AssertEqual(t, 0.0, m1.m00)
	dyn4go.AssertEqual(t, 4.0, m1.m01)
	dyn4go.AssertEqual(t, 0.0, m1.m02)
	dyn4go.AssertEqual(t, 6.0, m1.m10)
	dyn4go.AssertEqual(t, 2.0, m1.m11)
	dyn4go.AssertEqual(t, 2.0, m1.m12)
	dyn4go.AssertEqual(t, 4.0, m1.m20)
	dyn4go.AssertEqual(t, 0.0, m1.m21)
	dyn4go.AssertEqual(t, -2.0, m1.m22)
}

/**
 * Tests the product by a scalar method.
 */
func TestMatrix33ProductScalar(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m2 := m1.ProductScalar(2.0)
	dyn4go.AssertEqual(t, 0.0, m2.m00)
	dyn4go.AssertEqual(t, 4.0, m2.m01)
	dyn4go.AssertEqual(t, 0.0, m2.m02)
	dyn4go.AssertEqual(t, 6.0, m2.m10)
	dyn4go.AssertEqual(t, 2.0, m2.m11)
	dyn4go.AssertEqual(t, 2.0, m2.m12)
	dyn4go.AssertEqual(t, 4.0, m2.m20)
	dyn4go.AssertEqual(t, 0.0, m2.m21)
	dyn4go.AssertEqual(t, -2.0, m2.m22)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m2)
}

/**
 * Tests the identity method.
 */
func TestMatrix33Identity(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m1.Identity()
	dyn4go.AssertEqual(t, 1.0, m1.m00)
	dyn4go.AssertEqual(t, 0.0, m1.m01)
	dyn4go.AssertEqual(t, 0.0, m1.m02)
	dyn4go.AssertEqual(t, 0.0, m1.m10)
	dyn4go.AssertEqual(t, 1.0, m1.m11)
	dyn4go.AssertEqual(t, 0.0, m1.m12)
	dyn4go.AssertEqual(t, 0.0, m1.m20)
	dyn4go.AssertEqual(t, 0.0, m1.m21)
	dyn4go.AssertEqual(t, 1.0, m1.m22)
}

/**
 * Tests the transpose method.
 */
func TestMatrix33Transpose(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m1.Transpose()
	dyn4go.AssertEqual(t, 0.0, m1.m00)
	dyn4go.AssertEqual(t, 3.0, m1.m01)
	dyn4go.AssertEqual(t, 2.0, m1.m02)
	dyn4go.AssertEqual(t, 2.0, m1.m10)
	dyn4go.AssertEqual(t, 1.0, m1.m11)
	dyn4go.AssertEqual(t, 0.0, m1.m12)
	dyn4go.AssertEqual(t, 0.0, m1.m20)
	dyn4go.AssertEqual(t, 1.0, m1.m21)
	dyn4go.AssertEqual(t, -1.0, m1.m22)
}

/**
 * Tests the get transpose method.
 */
func TestMatrix33GetTranspose(t *testing.T) {
	m1 := NewMatrix33FromFloats(0.0, 2.0, 0.0,
		3.0, 1.0, 1.0,
		2.0, 0.0, -1.0)
	m2 := m1.GetTranspose()
	dyn4go.AssertEqual(t, 0.0, m2.m00)
	dyn4go.AssertEqual(t, 3.0, m2.m01)
	dyn4go.AssertEqual(t, 2.0, m2.m02)
	dyn4go.AssertEqual(t, 2.0, m2.m10)
	dyn4go.AssertEqual(t, 1.0, m2.m11)
	dyn4go.AssertEqual(t, 0.0, m2.m12)
	dyn4go.AssertEqual(t, 0.0, m2.m20)
	dyn4go.AssertEqual(t, 1.0, m2.m21)
	dyn4go.AssertEqual(t, -1.0, m2.m22)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m2)
}

/**
 * Tests the determinant method.
 */
func TestMatrix33Determinant(t *testing.T) {
	m1 := NewMatrix33FromFloats(1.0, 0.0, 5.0,
		2.0, 1.0, 6.0,
		3.0, 4.0, 0.0)
	det := m1.Determinant()
	dyn4go.AssertEqual(t, 1.0, det)
}

/**
 * Tests the invert method.
 */
func TestMatrix33Invert(t *testing.T) {
	m1 := NewMatrix33FromFloats(1.0, 0.0, 5.0,
		2.0, 1.0, 6.0,
		3.0, 4.0, 0.0)
	m1.Invert()
	//-24.0 20.0 -5.0 18.0 -15.0 4.0 5.0 -4.0 1.0
	dyn4go.AssertEqual(t, -24.0, m1.m00)
	dyn4go.AssertEqual(t, 20.0, m1.m01)
	dyn4go.AssertEqual(t, -5.0, m1.m02)
	dyn4go.AssertEqual(t, 18.0, m1.m10)
	dyn4go.AssertEqual(t, -15.0, m1.m11)
	dyn4go.AssertEqual(t, 4.0, m1.m12)
	dyn4go.AssertEqual(t, 5.0, m1.m20)
	dyn4go.AssertEqual(t, -4.0, m1.m21)
	dyn4go.AssertEqual(t, 1.0, m1.m22)
}

/**
 * Tests the get inverse method.
 */
func TestMatrix33GetInverse(t *testing.T) {
	m1 := NewMatrix33FromFloats(1.0, 0.0, 5.0,
		2.0, 1.0, 6.0,
		3.0, 4.0, 0.0)
	m2 := m1.GetInverse()
	//-24.0 20.0 -5.0 18.0 -15.0 4.0 5.0 -4.0 1.0
	dyn4go.AssertEqual(t, -24.0, m2.m00)
	dyn4go.AssertEqual(t, 20.0, m2.m01)
	dyn4go.AssertEqual(t, -5.0, m2.m02)
	dyn4go.AssertEqual(t, 18.0, m2.m10)
	dyn4go.AssertEqual(t, -15.0, m2.m11)
	dyn4go.AssertEqual(t, 4.0, m2.m12)
	dyn4go.AssertEqual(t, 5.0, m2.m20)
	dyn4go.AssertEqual(t, -4.0, m2.m21)
	dyn4go.AssertEqual(t, 1.0, m2.m22)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m2)
}

/**
 * Tests the solve method.
 */
func TestMatrix33Solve22(t *testing.T) {
	A := NewMatrix33FromFloats(3.0, -1.0, 0.0,
		-1.0, -1.0, 0.0,
		0.0, 0.0, 0.0)
	b := NewVector2FromXY(2.0, 6.0)
	x := A.Solve22(b)
	dyn4go.AssertEqual(t, -1.0, x.X)
	dyn4go.AssertEqual(t, -5.0, x.Y)
}

/**
 * Tests the solve method.
 */
func TestMatrix33Solve33(t *testing.T) {
	A := NewMatrix33FromFloats(1.0, -3.0, 3.0,
		2.0, 3.0, -1.0,
		4.0, -3.0, -1.0)
	b := NewVector3FromFloats(-4.0, 15.0, 19.0)
	x := A.Solve33(b)
	//(5.0, 1.0, -2.0)
	dyn4go.AssertEqual(t, 5.0, x.X)
	dyn4go.AssertEqual(t, 1.0, x.Y)
	dyn4go.AssertEqual(t, -2.0, x.Z)
}
