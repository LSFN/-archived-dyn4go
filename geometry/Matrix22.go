package geometry2

import (
	"math"

	"github.com/LSFN/dyn4go"
)

type Matrix22 struct {
	m00, m01, m10, m11 float64
}

func NewMatrix22FromFloats(m00, m01, m10, m11 float64) *Matrix22 {
	m := new(Matrix22)
	m.m00 = m00
	m.m01 = m01
	m.m10 = m10
	m.m11 = m11
	return m
}

func NewMatrix22FromFloatSlice(values []float64) *Matrix22 {
	if values == nil || len(values) != 4 {
		panic("2x2 Matrices must be created with exactly 4 floats")
	}
	m := new(Matrix22)
	m.m00 = values[0]
	m.m01 = values[1]
	m.m10 = values[2]
	m.m11 = values[3]
	return m
}

func NewMatrix22FromMatrix22(m2 *Matrix22) *Matrix22 {
	m := new(Matrix22)
	*m = *m2
	return m
}

func (m *Matrix22) AddMatrix22(m2 *Matrix22) *Matrix22 {
	m.m00 += m2.m00
	m.m01 += m2.m01
	m.m10 += m2.m10
	m.m11 += m2.m11
	return m
}

func (m *Matrix22) SumMatrix22(m2 *Matrix22) *Matrix22 {
	m3 := NewMatrix22FromMatrix22(m)
	m3.m00 += m2.m00
	m3.m01 += m2.m01
	m3.m10 += m2.m10
	m3.m11 += m2.m11
	return m3
}

func (m *Matrix22) SubtractMatrix22(m2 *Matrix22) *Matrix22 {
	m.m00 -= m2.m00
	m.m01 -= m2.m01
	m.m10 -= m2.m10
	m.m11 -= m2.m11
	return m
}

func (m *Matrix22) DifferenceMatrix22(m2 *Matrix22) *Matrix22 {
	m3 := NewMatrix22FromMatrix22(m)
	m3.m00 -= m2.m00
	m3.m01 -= m2.m01
	m3.m10 -= m2.m10
	m3.m11 -= m2.m11
	return m3
}

func (m *Matrix22) MultiplyMatrix22(m2 *Matrix22) *Matrix22 {
	m00 := m.m00
	m01 := m.m01
	m10 := m.m10
	m11 := m.m11
	m.m00 = m00*m2.m00 + m01*m2.m10
	m.m01 = m00*m2.m01 + m01*m2.m11
	m.m10 = m10*m2.m00 + m11*m2.m10
	m.m11 = m10*m2.m01 + m11*m2.m11
	return m
}

func (m *Matrix22) ProductMatrix22(m2 *Matrix22) *Matrix22 {
	m3 := new(Matrix22)
	m3.m00 = m.m00*m2.m00 + m.m01*m2.m10
	m3.m01 = m.m00*m2.m01 + m.m01*m2.m11
	m3.m10 = m.m10*m2.m00 + m.m11*m2.m10
	m3.m11 = m.m10*m2.m01 + m.m11*m2.m11
	return m3
}

func (m *Matrix22) MultiplyVector2(v *Vector2) *Vector2 {
	x := v.X
	y := v.Y
	v.X = m.m00*x + m.m01*y
	v.Y = m.m10*x + m.m11*y
	return v
}

func (m *Matrix22) ProductVector2(v *Vector2) *Vector2 {
	v2 := new(Vector2)
	v2.X = m.m00*v.X + m.m01*v.Y
	v2.Y = m.m10*v.X + m.m11*v.Y
	return v2
}

func (m *Matrix22) MultiplyTVector2(v *Vector2) *Vector2 {
	x := v.X
	y := v.Y
	v.X = m.m00*x + m.m10*y
	v.Y = m.m01*x + m.m11*y
	return v
}

func (m *Matrix22) ProductTVector2(v *Vector2) *Vector2 {
	v2 := new(Vector2)
	v2.X = m.m00*v.X + m.m10*v.Y
	v2.Y = m.m01*v.X + m.m11*v.Y
	return v2
}

func (m *Matrix22) MultiplyScalar(s float64) *Matrix22 {
	m.m00 *= s
	m.m01 *= s
	m.m10 *= s
	m.m11 *= s
	return m
}

func (m *Matrix22) ProductScalar(s float64) *Matrix22 {
	m2 := NewMatrix22FromMatrix22(m)
	m2.m00 *= s
	m2.m01 *= s
	m2.m10 *= s
	m2.m11 *= s
	return m2
}

func (m *Matrix22) Identity() *Matrix22 {
	m.m00 = 1
	m.m01 = 0
	m.m10 = 0
	m.m11 = 1
	return m
}

func (m *Matrix22) Transpose() *Matrix22 {
	m.m01, m.m10 = m.m10, m.m01
	return m
}

func (m *Matrix22) GetTranspose() *Matrix22 {
	m2 := NewMatrix22FromMatrix22(m)
	return m2.Transpose()
}

func (m *Matrix22) Determinant() float64 {
	return m.m00*m.m11 - m.m01*m.m10
}

func (m *Matrix22) Invert() *Matrix22 {
	det := m.Determinant()
	if math.Abs(det) > dyn4go.Epsilon {
		det = 1 / det
	}
	a := m.m00
	b := m.m01
	c := m.m10
	d := m.m11
	m.m00 = det * d
	m.m01 = -det * b
	m.m10 = -det * c
	m.m11 = det * a
	return m
}

func (m *Matrix22) GetInverse() *Matrix22 {
	det := m.Determinant()
	if math.Abs(det) > dyn4go.Epsilon {
		det = 1 / det
	}
	m2 := new(Matrix22)
	m2.m00 = det * m.m11
	m2.m01 = -det * m.m01
	m2.m10 = -det * m.m10
	m2.m11 = det * m.m00
	return m2
}

func (m *Matrix22) Solve(v *Vector2) *Vector2 {
	det := m.Determinant()
	if math.Abs(det) > dyn4go.Epsilon {
		det = 1 / det
	}
	v2 := new(Vector2)
	v2.X = det * (m.m11*v.X - m.m01*v.Y)
	v2.Y = det * (m.m00*v.Y - m.m10*v.X)
	return v2
}
