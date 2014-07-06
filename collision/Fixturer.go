package collision

import (
	"github.com/LSFN/dyn4go/geometry"
)

type Fixturer interface {
	GetID() string
	GetShape() geometry.Convexer
	GetFilter() Filterer
	SetFilter(filter Filterer)
	IsSensor() bool
	SetSensor(sensor bool)
	GetUserData() interface{}
	SetUserData(data interface{})
}
