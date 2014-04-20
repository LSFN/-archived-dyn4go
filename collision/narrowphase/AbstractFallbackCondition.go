package narrowphase

type AbstractFallbackCondition struct {
	sortIndex int
}

func (a *AbstractFallbackCondition) InitAbstractFallbackCondition(sortIndex int) {
	a.sortIndex = sortIndex
}

func (a *AbstractFallbackCondition) CompareTo(o FallbackConditioner) int {
	return a.GetSortIndex() - o.GetSortIndex()
}

func (a *AbstractFallbackCondition) GetSortIndex() int {
	return a.sortIndex
}
