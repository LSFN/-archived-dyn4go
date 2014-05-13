package narrowphase

import (
	"reflect"
)

type PairwiseTypedFallbackCondition struct {
	TypedFallbackCondition
	type1, type2 reflect.Type
}

func NewPairwiseTypedFallbackCondition(type1, type2 reflect.Type) *PairwiseTypedFallbackCondition {
	return NewPairwiseTypedFallbackConditionInt(type1, type2, 0)
}

func NewPairwiseTypedFallbackConditionInt(type1, type2 reflect.Type, sortIndex int) *PairwiseTypedFallbackCondition {
	p := new(PairwiseTypedFallbackCondition)
	p.type1 = type1
	p.type2 = type2
	p.sortIndex = sortIndex
	return p
}

func (p *PairwiseTypedFallbackCondition) IsMatch(type1, type2 reflect.Type) bool {
	return (p.type1 == type1 && p.type2 == type2) || (p.type1 == type2 && p.type2 == type1)
}
