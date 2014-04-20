package narrowphase

import (
	"reflect"

	"github.com/LSFN/dyn4go/geometry"
)

type SingleTypedFallbackCondition struct {
	TypedFallbackCondition
	compareType reflect.Type
}

func NewSingleTypedFallbackCondition(compareType reflect.Type) *SingleTypedFallbackCondition {
	return NewSingleTypedFallbackConditionIntBool(compareType, 0, true)
}

func NewSingleTypedFallbackConditionInt(compareType reflect.Type, sortIndex int) *SingleTypedFallbackCondition {
	s := new(SingleTypedFallbackCondition)
	s.InitTypedFallbackConditionInt(sortIndex)
	s.compareType = compareType
	return s
}

func (s *SingleTypedFallbackCondition) IsMatch(type1, type2 reflect.Type) bool {
	return s.compareType == type1 || s.compareType == type2
}
