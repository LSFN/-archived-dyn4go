package geometry

import (
	"github.com/LSFN/dyn4go"
)

type Mass struct {
	massType   int
	center     *Vector2
	mass       float64
	inertia    float64
	invMass    float64
	invInertia float64
}

const (
	NORMAL = iota
	INFINITE
	FIXED_ANGULAR_VELOCITY
	FIXED_LINEAR_VELOCITY
)

func NewMass() *Mass {
	m := new(Mass)
	m.massType = INFINITE
	m.center = new(Vector2)
	return m
}

func NewMassFromCenterMassInertia(center *Vector2, mass, inertia float64) *Mass {
	if center == nil {
		panic("Center of mass may not be nil")
	}
	if mass < 0 {
		panic("Mass may not be negative")
	}
	if inertia < 0 {
		panic("Intertia may not be negative")
	}
	m := new(Mass)
	m.massType = NORMAL
	m.center = NewVector2FromVector2(center)
	m.mass = mass
	m.inertia = inertia
	if m.mass > dyn4go.Epsilon {
		m.invMass = 1 / mass
	} else {
		m.invMass = 0
		m.massType = FIXED_LINEAR_VELOCITY
	}
	if m.inertia > dyn4go.Epsilon {
		m.invInertia = 1 / m.inertia
	} else {
		m.invInertia = 0
		m.massType = FIXED_ANGULAR_VELOCITY
	}
	if m.mass <= dyn4go.Epsilon && m.inertia <= dyn4go.Epsilon {
		m.massType = INFINITE
	}
	return m
}

func NewMassFromMass(m *Mass) *Mass {
	if m == nil {
		panic("Cannot create mass from nil")
	}
	m2 := *m
	m2.center = NewVector2FromVector2(m.center)
	return &m2
}

func CreateMass(masses []*Mass) *Mass {
	if masses == nil || len(masses) == 0 {
		panic("Mass must be created from at least one other mass")
	}
	if len(masses) == 1 {
		if masses[0] == nil {
			panic("Masses given may not be nil")
		} else {
			return NewMassFromMass(masses[0])
		}
	}
	c := new(Vector2)
	var m, i float64
	for _, mass := range masses {
		if mass == nil {
			panic("Cannot create mass from nil")
		}
		c.AddVector2(mass.center.Product(mass.mass))
		m += mass.mass
	}
	if m > 0 {
		c.Multiply(1 / m)
	}
	for _, mass := range masses {
		i += mass.inertia + mass.mass*mass.center.DistanceSquaredFromVector2(c)
	}
	return NewMassFromCenterMassInertia(c, m, i)
}

func (m *Mass) IsInfinite() bool {
	return m.massType == INFINITE
}

func (m *Mass) SetType(massType int) {
	if !(massType == NORMAL || massType == INFINITE || massType == FIXED_LINEAR_VELOCITY || massType == FIXED_ANGULAR_VELOCITY) {
		panic("Not a valid mass type")
	}
	m.massType = massType
}

func (m *Mass) GetType() int {
	return m.massType
}

func (m *Mass) GetCenter() *Vector2 {
	return m.center
}

func (m *Mass) GetMass() float64 {
	if m.massType == INFINITE || m.massType == FIXED_LINEAR_VELOCITY {
		return 0
	} else {
		return m.mass
	}
}

func (m *Mass) GetInertia() float64 {
	if m.massType == INFINITE || m.massType == FIXED_ANGULAR_VELOCITY {
		return 0
	} else {
		return m.inertia
	}
}

func (m *Mass) GetInverseMass() float64 {
	if m.massType == INFINITE || m.massType == FIXED_LINEAR_VELOCITY {
		return 0
	} else {
		return m.invMass
	}
}

func (m *Mass) GetInverseInertia() float64 {
	if m.massType == INFINITE || m.massType == FIXED_ANGULAR_VELOCITY {
		return 0
	} else {
		return m.invInertia
	}
}
