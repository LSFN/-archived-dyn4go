package geometry2

import (
	"testing"

	"github.com/LSFN/dyn4go"
)

/**
 * Test the creation method passing four doubles.
 */
func TestMatrix22CreateFull(t *testing.T) {
	m := NewMatrix22FromFloats(1.0, 2.0,
		-3.0, 8.0)
	dyn4go.AssertEqual(t, 1.0, m.m00)
	dyn4go.AssertEqual(t, 2.0, m.m01)
	dyn4go.AssertEqual(t, -3.0, m.m10)
	dyn4go.AssertEqual(t, 8.0, m.m11)
}

/**
 * Test the creation method passing a double array.
 */
func TestMatrix22CreateFullArray(t *testing.T) {
	m := NewMatrix22FromFloatSlice([]float64{1.0, 2.0,
		-3.0, 8.0})
	dyn4go.AssertEqual(t, 1.0, m.m00)
	dyn4go.AssertEqual(t, 2.0, m.m01)
	dyn4go.AssertEqual(t, -3.0, m.m10)
	dyn4go.AssertEqual(t, 8.0, m.m11)
}

/**
 * Tests the copy constructor.
 */
func TestMatrix22Copy(t *testing.T) {
	m1 := new(Matrix22)
	m1.m00 = 0
	m1.m01 = 2
	m1.m10 = 1
	m1.m11 = 3

	// make a copy
	m2 := NewMatrix22FromMatrix22(m1)
	// test the values
	dyn4go.AssertEqual(t, m1.m00, m2.m00)
	dyn4go.AssertEqual(t, m1.m01, m2.m01)
	dyn4go.AssertEqual(t, m1.m10, m2.m10)
	dyn4go.AssertEqual(t, m1.m11, m2.m11)
}

/**
 * Tests the add method.
 */
func TestMatrix22Add(t *testing.T) {
	m1 := NewMatrix22FromFloats(0.0, 2.0,
		3.5, 1.2)
	m2 := NewMatrix22FromFloats(1.3, 0.3,
		0.0, 4.5)
	m1.AddMatrix22(m2)
	// test the values
	dyn4go.AssertEqual(t, 1.3, m1.m00)
	dyn4go.AssertEqual(t, 2.3, m1.m01)
	dyn4go.AssertEqual(t, 3.5, m1.m10)
	dyn4go.AssertEqual(t, 5.7, m1.m11)
}

/**
 * Tests the sum method.
 */
func TestMatrix22Sum(t *testing.T) {
	m1 := NewMatrix22FromFloats(0.0, 2.0,
		3.5, 1.2)
	m2 := NewMatrix22FromFloats(1.3, 0.3,
		0.0, 4.5)
	m3 := m1.SumMatrix22(m2)
	// test the values
	dyn4go.AssertEqual(t, 1.3, m3.m00)
	dyn4go.AssertEqual(t, 2.3, m3.m01)
	dyn4go.AssertEqual(t, 3.5, m3.m10)
	dyn4go.AssertEqual(t, 5.7, m3.m11)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m3)
}

/**
 * Tests the subtract method.
 */
func TestMatrix22Subtract(t *testing.T) {
	m1 := NewMatrix22FromFloats(0.0, 2.0,
		3.5, 1.2)
	m2 := NewMatrix22FromFloats(1.3, 0.3,
		0.0, 4.5)
	m1.SubtractMatrix22(m2)
	// test the values
	dyn4go.AssertEqual(t, -1.3, m1.m00)
	dyn4go.AssertEqual(t, 1.7, m1.m01)
	dyn4go.AssertEqual(t, 3.5, m1.m10)
	dyn4go.AssertEqual(t, -3.3, m1.m11)
}

/**
 * Tests the difference method.
 */
func TestMatrix22Difference(t *testing.T) {
	m1 := NewMatrix22FromFloats(0.0, 2.0,
		3.5, 1.2)
	m2 := NewMatrix22FromFloats(1.3, 0.3,
		0.0, 4.5)
	m3 := m1.DifferenceMatrix22(m2)
	// test the values
	dyn4go.AssertEqual(t, -1.3, m3.m00)
	dyn4go.AssertEqual(t, 1.7, m3.m01)
	dyn4go.AssertEqual(t, 3.5, m3.m10)
	dyn4go.AssertEqual(t, -3.3, m3.m11)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m3)
}

/**
 * Tests the multiply matrix method.
 */
func TestMatrix22MultiplyMatrix(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m2 := NewMatrix22FromFloats(4.0, 3.0,
		2.0, 1.0)
	m1.MultiplyMatrix22(m2)
	dyn4go.AssertEqual(t, 8.0, m1.m00)
	dyn4go.AssertEqual(t, 5.0, m1.m01)
	dyn4go.AssertEqual(t, 20.0, m1.m10)
	dyn4go.AssertEqual(t, 13.0, m1.m11)
}

/**
 * Tests the product matrix method.
 */
func TestMatrix22ProductMatrix(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m2 := NewMatrix22FromFloats(4.0, 3.0,
		2.0, 1.0)
	m3 := m1.ProductMatrix22(m2)
	dyn4go.AssertEqual(t, 8.0, m3.m00)
	dyn4go.AssertEqual(t, 5.0, m3.m01)
	dyn4go.AssertEqual(t, 20.0, m3.m10)
	dyn4go.AssertEqual(t, 13.0, m3.m11)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m3)
}

/**
 * Tests the multiply vector method.
 */
func TestMatrix22MultiplyVector(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	v1 := NewVector2FromXY(1.0, -1.0)
	m1.MultiplyVector2(v1)
	dyn4go.AssertEqual(t, -1.0, v1.X)
	dyn4go.AssertEqual(t, -1.0, v1.Y)
}

/**
 * Tests the product vector method.
 */
func TestMatrix22ProductVector(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	v1 := NewVector2FromXY(1.0, -1.0)
	v2 := m1.ProductVector2(v1)
	dyn4go.AssertEqual(t, -1.0, v2.X)
	dyn4go.AssertEqual(t, -1.0, v2.Y)
	// make sure we didnt modify the first vector
	dyn4go.AssertFalse(t, *v1 == *v2)
}

/**
 * Tests the multiply vector transpose method.
 */
func TestMatrix22MultiplyVectorT(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	v1 := NewVector2FromXY(1.0, -1.0)
	m1.MultiplyTVector2(v1)
	dyn4go.AssertEqual(t, -2.0, v1.X)
	dyn4go.AssertEqual(t, -2.0, v1.Y)
}

/**
 * Tests the product vector transpose method.
 */
func TestMatrix22ProductVectorT(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	v1 := NewVector2FromXY(1.0, -1.0)
	v2 := m1.ProductTVector2(v1)
	dyn4go.AssertEqual(t, -2.0, v2.X)
	dyn4go.AssertEqual(t, -2.0, v2.Y)
	// make sure we didnt modify the first vector
	dyn4go.AssertFalse(t, *v1 == *v2)
}

/**
 * Tests the multiply by a scalar method.
 */
func TestMatrix22MultiplyScalar(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m1.MultiplyScalar(2.0)
	dyn4go.AssertEqual(t, 2.0, m1.m00)
	dyn4go.AssertEqual(t, 4.0, m1.m01)
	dyn4go.AssertEqual(t, 6.0, m1.m10)
	dyn4go.AssertEqual(t, 8.0, m1.m11)
}

/**
 * Tests the product by a scalar method.
 */
func TestMatrix22ProductScalar(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m2 := m1.ProductScalar(2.0)
	dyn4go.AssertEqual(t, 2.0, m2.m00)
	dyn4go.AssertEqual(t, 4.0, m2.m01)
	dyn4go.AssertEqual(t, 6.0, m2.m10)
	dyn4go.AssertEqual(t, 8.0, m2.m11)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m2)
}

/**
 * Tests the identity method.
 */
func TestMatrix22Identity(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m1.Identity()
	dyn4go.AssertEqual(t, 1.0, m1.m00)
	dyn4go.AssertEqual(t, 0.0, m1.m01)
	dyn4go.AssertEqual(t, 0.0, m1.m10)
	dyn4go.AssertEqual(t, 1.0, m1.m11)
}

/**
 * Tests the transpose method.
 */
func TestMatrix22Transpose(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m1.Transpose()
	dyn4go.AssertEqual(t, 1.0, m1.m00)
	dyn4go.AssertEqual(t, 3.0, m1.m01)
	dyn4go.AssertEqual(t, 2.0, m1.m10)
	dyn4go.AssertEqual(t, 4.0, m1.m11)
}

/**
 * Tests the get transpose method.
 */
func TestMatrix22GetTranspose(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m2 := m1.GetTranspose()
	dyn4go.AssertEqual(t, 1.0, m2.m00)
	dyn4go.AssertEqual(t, 3.0, m2.m01)
	dyn4go.AssertEqual(t, 2.0, m2.m10)
	dyn4go.AssertEqual(t, 4.0, m2.m11)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m2)
}

/**
 * Tests the determinant method.
 */
func TestMatrix22Determinant(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	det := m1.Determinant()
	dyn4go.AssertEqual(t, -2.0, det)
}

/**
 * Tests the invert method.
 */
func TestMatrix22Invert(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m1.Invert()
	dyn4go.AssertEqual(t, -2.0, m1.m00)
	dyn4go.AssertEqual(t, 1.0, m1.m01)
	dyn4go.AssertEqual(t, 1.5, m1.m10)
	dyn4go.AssertEqual(t, -0.5, m1.m11)
}

/**
 * Tests the get inverse method.
 */
func TestMatrix22GetInverse(t *testing.T) {
	m1 := NewMatrix22FromFloats(1.0, 2.0,
		3.0, 4.0)
	m2 := m1.GetInverse()
	dyn4go.AssertEqual(t, -2.0, m2.m00)
	dyn4go.AssertEqual(t, 1.0, m2.m01)
	dyn4go.AssertEqual(t, 1.5, m2.m10)
	dyn4go.AssertEqual(t, -0.5, m2.m11)
	// make sure we didnt modify the first matrix
	dyn4go.AssertFalse(t, *m1 == *m2)
}

/**
 * Tests the solve method.
 */
func TestMatrix22Solve(t *testing.T) {
	A := NewMatrix22FromFloats(3.0, -1.0,
		-1.0, -1.0)
	b := NewVector2FromXY(2.0, 6.0)
	x := A.Solve(b)
	dyn4go.AssertEqual(t, -1.0, x.X)
	dyn4go.AssertEqual(t, -5.0, x.Y)
}
