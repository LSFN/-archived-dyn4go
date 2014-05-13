package narrowphase

type TypedFallbackCondition struct {
	AbstractFallbackCondition
}

func (t *TypedFallbackCondition) InitTypedFallbackCondition() {
	t.InitAbstractFallbackCondition(0)
}

func (t *TypedFallbackCondition) InitTypedFallbackConditionInt(sortIndex int) {
	t.InitAbstractFallbackCondition(sortIndex)
}

/*
func (t *TypedFallbackCondition) IsMatch(convex1 geometry.Convexer, convex2 geometry.Convexer) bool {
	return IsMatch(reflect.TypeOf(convex1), reflect.TypeOf(convex2))
}
*/
