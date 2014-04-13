package geometry

import (
	"math"

	"github.com/LSFN/dyn4go"
)

type Matrix33 struct {
	m00, m01, m02, m10, m11, m12, m20, m21, m22 float64
}

func NewMatrix33FromFloats(m00, m01, m02, m10, m11, m12, m20, m21, m22 float64) *Matrix33 {
	m := new(Matrix33)
	m.m00 = m00
	m.m01 = m01
	m.m02 = m02
	m.m10 = m10
	m.m11 = m11
	m.m12 = m12
	m.m20 = m20
	m.m21 = m21
	m.m22 = m22
	return m
}

func NewMatrix33FromFloatSlice(values []float64) *Matrix33 {
	if values == nil || len(values) != 9 {
		panic("3 x 3 matrix must be created from exactly 9 floats")
	}
	m := new(Matrix33)
	m.m00 = values[0]
	m.m01 = values[1]
	m.m02 = values[2]
	m.m10 = values[3]
	m.m11 = values[4]
	m.m12 = values[5]
	m.m20 = values[6]
	m.m21 = values[7]
	m.m22 = values[8]
	return m
}

func NewMatrix33FromMatrix33(m2 *Matrix33) *Matrix33 {
	m := new(Matrix33)
	m.m00 = m2.m00
	m.m01 = m2.m01
	m.m02 = m2.m02
	m.m10 = m2.m10
	m.m11 = m2.m11
	m.m12 = m2.m12
	m.m20 = m2.m20
	m.m21 = m2.m21
	m.m22 = m2.m22
	return m
}

func (m *Matrix33) Add(m2 *Matrix33) *Matrix33 {
	m.m00 += m2.m00
	m.m01 += m2.m01
	m.m02 += m2.m02
	m.m10 += m2.m10
	m.m11 += m2.m11
	m.m12 += m2.m12
	m.m20 += m2.m20
	m.m21 += m2.m21
	m.m22 += m2.m22
	return m
}

func (m *Matrix33) Sum(m2 *Matrix33) *Matrix33 {
	m3 := new(Matrix33)
	m3.m00 = m.m00 + m2.m00
	m3.m01 = m.m01 + m2.m01
	m3.m02 = m.m02 + m2.m02
	m3.m10 = m.m10 + m2.m10
	m3.m11 = m.m11 + m2.m11
	m3.m12 = m.m12 + m2.m12
	m3.m20 = m.m20 + m2.m20
	m3.m21 = m.m21 + m2.m21
	m3.m22 = m.m22 + m2.m22
	return m3
}

func (m *Matrix33) Subtract(m2 *Matrix33) *Matrix33 {
	m.m00 -= m2.m00
	m.m01 -= m2.m01
	m.m02 -= m2.m02
	m.m10 -= m2.m10
	m.m11 -= m2.m11
	m.m12 -= m2.m12
	m.m20 -= m2.m20
	m.m21 -= m2.m21
	m.m22 -= m2.m22
	return m
}

func (m *Matrix33) Difference(m2 *Matrix33) *Matrix33 {
	m3 := new(Matrix33)
	m3.m00 = m.m00 - m2.m00
	m3.m01 = m.m01 - m2.m01
	m3.m02 = m.m02 - m2.m02
	m3.m10 = m.m10 - m2.m10
	m3.m11 = m.m11 - m2.m11
	m3.m12 = m.m12 - m2.m12
	m3.m20 = m.m20 - m2.m20
	m3.m21 = m.m21 - m2.m21
	m3.m22 = m.m22 - m2.m22
	return m3
}

func (m *Matrix33) MultiplyMatrix33(m2 *Matrix33) *Matrix33 {
	m00 := m.m00
	m01 := m.m01
	m02 := m.m02
	m10 := m.m10
	m11 := m.m11
	m12 := m.m12
	m20 := m.m20
	m21 := m.m21
	m22 := m.m22
	m.m00 = m00*m2.m00 + m01*m2.m10 + m02*m2.m20
	m.m01 = m00*m2.m01 + m01*m2.m11 + m02*m2.m21
	m.m02 = m00*m2.m02 + m01*m2.m12 + m02*m2.m22
	m.m10 = m10*m2.m00 + m11*m2.m10 + m12*m2.m20
	m.m11 = m10*m2.m01 + m11*m2.m11 + m12*m2.m21
	m.m12 = m10*m2.m02 + m11*m2.m12 + m12*m2.m22
	m.m20 = m20*m2.m00 + m21*m2.m10 + m22*m2.m20
	m.m21 = m20*m2.m01 + m21*m2.m11 + m22*m2.m21
	m.m22 = m20*m2.m02 + m21*m2.m12 + m22*m2.m22
	return m
}

func (m *Matrix33) ProductMatrix33(m2 *Matrix33) *Matrix33 {
	m3 := new(Matrix33)
	m3.m00 = m.m00*m2.m00 + m.m01*m2.m10 + m.m02*m2.m20
	m3.m01 = m.m00*m2.m01 + m.m01*m2.m11 + m.m02*m2.m21
	m3.m02 = m.m00*m2.m02 + m.m01*m2.m12 + m.m02*m2.m22
	m3.m10 = m.m10*m2.m00 + m.m11*m2.m10 + m.m12*m2.m20
	m3.m11 = m.m10*m2.m01 + m.m11*m2.m11 + m.m12*m2.m21
	m3.m12 = m.m10*m2.m02 + m.m11*m2.m12 + m.m12*m2.m22
	m3.m20 = m.m20*m2.m00 + m.m21*m2.m10 + m.m22*m2.m20
	m3.m21 = m.m20*m2.m01 + m.m21*m2.m11 + m.m22*m2.m21
	m3.m22 = m.m20*m2.m02 + m.m21*m2.m12 + m.m22*m2.m22
	return m3
}

func (m *Matrix33) MultiplyVector3(v *Vector3) *Vector3 {
	x := v.X
	y := v.Y
	z := v.Z
	v.X = m.m00*x + m.m01*y + m.m02*z
	v.Y = m.m10*x + m.m11*y + m.m12*z
	v.Z = m.m20*x + m.m21*y + m.m22*z
	return v
}

func (m *Matrix33) ProductVector3(v *Vector3) *Vector3 {
	v2 := new(Vector3)
	v2.X = m.m00*v.X + m.m01*v.Y + m.m02*v.Z
	v2.Y = m.m10*v.X + m.m11*v.Y + m.m12*v.Z
	v2.Z = m.m20*v.X + m.m21*v.Y + m.m22*v.Z
	return v2
}

func (m *Matrix33) MultiplyTVector3(v *Vector3) *Vector3 {
	x := v.X
	y := v.Y
	z := v.Z
	v.X = m.m00*x + m.m10*y + m.m20*z
	v.Y = m.m01*x + m.m11*y + m.m21*z
	v.Z = m.m02*x + m.m12*y + m.m22*z
	return v
}

func (m *Matrix33) ProductTVector3(v *Vector3) *Vector3 {
	v2 := new(Vector3)
	v2.X = m.m00*v.X + m.m10*v.Y + m.m20*v.Z
	v2.Y = m.m01*v.X + m.m11*v.Y + m.m21*v.Z
	v2.Z = m.m02*v.X + m.m12*v.Y + m.m22*v.Z
	return v2
}

func (m *Matrix33) MultiplyScalar(s float64) *Matrix33 {
	m.m00 *= s
	m.m01 *= s
	m.m02 *= s
	m.m10 *= s
	m.m11 *= s
	m.m12 *= s
	m.m20 *= s
	m.m21 *= s
	m.m22 *= s
	return m
}

func (m *Matrix33) ProductScalar(s float64) *Matrix33 {
	m2 := new(Matrix33)
	m2.m00 = m.m00 * s
	m2.m01 = m.m01 * s
	m2.m02 = m.m02 * s
	m2.m10 = m.m10 * s
	m2.m11 = m.m11 * s
	m2.m12 = m.m12 * s
	m2.m20 = m.m20 * s
	m2.m21 = m.m21 * s
	m2.m22 = m.m22 * s
	return m2
}

func (m *Matrix33) Identity() *Matrix33 {
	m.m00 = 1
	m.m01 = 0
	m.m02 = 0
	m.m10 = 0
	m.m11 = 1
	m.m12 = 0
	m.m20 = 0
	m.m21 = 0
	m.m22 = 1
	return m
}

func (m *Matrix33) Transpose() *Matrix33 {
	m.m01, m.m10 = m.m10, m.m01
	m.m02, m.m20 = m.m20, m.m02
	m.m12, m.m21 = m.m21, m.m12
	return m
}

func (m *Matrix33) GetTranspose() *Matrix33 {
	m2 := new(Matrix33)
	m2.m00, m2.m11, m2.m22 = m.m00, m.m11, m.m22
	m2.m01, m2.m10 = m.m10, m.m01
	m2.m02, m2.m20 = m.m20, m.m02
	m2.m12, m2.m21 = m.m21, m.m12
	return m2
}

func (m *Matrix33) Determinant() float64 {
	return m.m00*m.m11*m.m22 +
		m.m01*m.m12*m.m20 +
		m.m02*m.m10*m.m21 -
		m.m20*m.m11*m.m02 -
		m.m21*m.m12*m.m00 -
		m.m22*m.m10*m.m01
}

func (m *Matrix33) Invert() *Matrix33 {
	det := m.Determinant()
	if math.Abs(det) > dyn4go.Epsilon {
		det = 1.0 / det
	}

	m00 := det * (m.m11*m.m22 - m.m12*m.m21)
	m01 := -det * (m.m01*m.m22 - m.m21*m.m02)
	m02 := det * (m.m01*m.m12 - m.m11*m.m02)
	m10 := -det * (m.m10*m.m22 - m.m20*m.m12)
	m11 := det * (m.m00*m.m22 - m.m20*m.m02)
	m12 := -det * (m.m00*m.m12 - m.m10*m.m02)
	m20 := det * (m.m10*m.m21 - m.m20*m.m11)
	m21 := -det * (m.m00*m.m21 - m.m20*m.m01)
	m22 := det * (m.m00*m.m11 - m.m10*m.m01)

	m.m00 = m00
	m.m01 = m01
	m.m02 = m02
	m.m10 = m10
	m.m11 = m11
	m.m12 = m12
	m.m20 = m20
	m.m21 = m21
	m.m22 = m22

	return m
}

func (m *Matrix33) GetInverse() *Matrix33 {
	det := m.Determinant()
	if math.Abs(det) > dyn4go.Epsilon {
		det = 1.0 / det
	}

	m3 := new(Matrix33)
	m3.m00 = det * (m.m11*m.m22 - m.m12*m.m21)
	m3.m01 = -det * (m.m01*m.m22 - m.m21*m.m02)
	m3.m02 = det * (m.m01*m.m12 - m.m11*m.m02)
	m3.m10 = -det * (m.m10*m.m22 - m.m20*m.m12)
	m3.m11 = det * (m.m00*m.m22 - m.m20*m.m02)
	m3.m12 = -det * (m.m00*m.m12 - m.m10*m.m02)
	m3.m20 = det * (m.m10*m.m21 - m.m20*m.m11)
	m3.m21 = -det * (m.m00*m.m21 - m.m20*m.m01)
	m3.m22 = det * (m.m00*m.m11 - m.m10*m.m01)
	return m3
}

func (m *Matrix33) Solve33(v *Vector3) *Vector3 {
	det := m.Determinant()
	if math.Abs(det) > dyn4go.Epsilon {
		det = 1.0 / det
	}

	m00 := m.m11*m.m22 - m.m12*m.m21
	m01 := -m.m01*m.m22 + m.m21*m.m02
	m02 := m.m01*m.m12 - m.m11*m.m02
	m10 := -m.m10*m.m22 + m.m20*m.m12
	m11 := m.m00*m.m22 - m.m20*m.m02
	m12 := -m.m00*m.m12 + m.m10*m.m02
	m20 := m.m10*m.m21 - m.m20*m.m11
	m21 := -m.m00*m.m21 + m.m20*m.m01
	m22 := m.m00*m.m11 - m.m10*m.m01

	v2 := new(Vector3)
	v2.X = det * (m00*v.X + m01*v.Y + m02*v.Z)
	v2.Y = det * (m10*v.X + m11*v.Y + m12*v.Z)
	v2.Z = det * (m20*v.X + m21*v.Y + m22*v.Z)
	return v2
}

func (m *Matrix33) Solve22(v *Vector2) *Vector2 {
	det := m.m00*m.m11 - m.m01*m.m10
	if math.Abs(det) > dyn4go.Epsilon {
		det = 1 / det
	}
	v2 := new(Vector2)
	v2.X = det * (m.m11*v.X - m.m01*v.Y)
	v2.Y = det * (m.m00*v.Y - m.m10*v.X)
	return v2
}
