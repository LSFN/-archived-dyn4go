package collision

import (
	"code.google.com/p/uuid"
	"github.com/LSFN/dyn4go/geometry"
)

type Fixture struct {
	id       string
	shape    *geometry.Convex
	filter   Filterer
	sensor   bool
	userData interface{}
}

func NewFixture(shape *geometry.Convex) *Fixture {
	if shape == nil {
		panic("Cannot create fixture from nil shape")
	}
	f := new(Fixture)
	f.id = uuid.New()
	f.shape = shape
	f.filter = NewDefaultFilter()
	f.sensor = false
}

func (f *Fixture) GetID() string {
	return f.id
}

func (f *Fixture) GetShape() *geometry.Convex {
	return f.shape
}

func (f *Fixture) GetFilter() Filterer {
	return f.filter
}

func (f *Fixture) SetFilter(filter Filterer) {
	if filter == nil {
		panic("Cannot set filter to nil")
	}
	f.filter = filter
}

func (f *Fixture) IsSensor() bool {
	return f.sensor
}

func (f *Fixture) SetSensor(sensor bool) {
	f.sensor = sensor
}

func (f *Fixture) GetUserData() interface{} {
	return f.userData
}

func (f *Fixture) SetUserData(data interface{}) {
	f.userData = data
}
