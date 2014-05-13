package narrowphase

import (
	"reflect"
)

type SingleTypedFallbackCondition struct {
	TypedFallbackCondition
	compareType reflect.Type
}

func NewSingleTypedFallbackCondition(compareType reflect.Type) *SingleTypedFallbackCondition {
	return NewSingleTypedFallbackConditionInt(compareType, 0)
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
