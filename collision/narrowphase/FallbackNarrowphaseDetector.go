package narrowphase

import (
	"github.com/LSFN/dyn4go/geometry"
	"sort"
)

type FallbackConditionSortable []FallbackConditioner

func (f *FallbackConditionSortable) Len() int {
	return len(*f)
}

func (f *FallbackConditionSortable) Less(i, j int) bool {
	return (*f)[i].GetSortIndex() < (*f)[j].GetSortIndex()
}

func (f *FallbackConditionSortable) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

type FallbackNarrowphaseDetector struct {
	primaryNarrowphaseDetector, fallbackNarrowphaseDetector NarrowphaseDetector
	fallbackConditions                                      FallbackConditionSortable
}

func NewFallbackNarrowphaseDetector(primaryNarrowphaseDetector, fallbackNarrowphaseDetector NarrowphaseDetector) *FallbackNarrowphaseDetector {
	return NewFallbackNarrowphaseDetectorFallbackConditions(primaryNarrowphaseDetector, fallbackNarrowphaseDetector, []FallbackConditioner{})
}

func NewFallbackNarrowphaseDetectorFallbackConditions(primaryNarrowphaseDetector, fallbackNarrowphaseDetector NarrowphaseDetector, conditions []FallbackConditioner) *FallbackNarrowphaseDetector {
	if primaryNarrowphaseDetector == nil || fallbackNarrowphaseDetector == nil {
		panic("Both primary and fallback narrowphase detectors must not be nil")
	}
	f := new(FallbackNarrowphaseDetector)
	f.primaryNarrowphaseDetector = primaryNarrowphaseDetector
	f.fallbackNarrowphaseDetector = fallbackNarrowphaseDetector
	if conditions == nil {
		f.fallbackConditions = []FallbackConditioner{}
	} else {
		f.fallbackConditions = conditions
	}
	return f
}

func (f *FallbackNarrowphaseDetector) AddCondition(condition FallbackConditioner) {
	f.fallbackConditions = append(f.fallbackConditions, condition)
	sort.Sort(&f.fallbackConditions)
}

func (f *FallbackNarrowphaseDetector) RemoveCondition(condition FallbackConditioner) bool {
	for i := range f.fallbackConditions {
		if f.fallbackConditions[i] == condition {
			f.fallbackConditions = append(f.fallbackConditions[:i], f.fallbackConditions[i+1:]...)
			return true
		}
	}
	return false
}

func (f *FallbackNarrowphaseDetector) ContainsCondition(condition FallbackConditioner) bool {
	for i := range f.fallbackConditions {
		if f.fallbackConditions[i] == condition {
			return true
		}
	}
	return false
}

func (f *FallbackNarrowphaseDetector) GetConditionCount() int {
	return len(f.fallbackConditions)
}

func (f *FallbackNarrowphaseDetector) GetCondition(index int) FallbackConditioner {
	return f.fallbackConditions[index]
}

func (f *FallbackNarrowphaseDetector) IsFallbackRequired(convex1, convex2 geometry.Convexer) bool {
	for _, condition := range f.fallbackConditions {
		if condition != nil && condition.IsMatch(convex1, convex2) {
			return true
		}
	}
	return false
}

func (f *FallbackNarrowphaseDetector) Detect(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform) bool {
	if f.IsFallbackRequired(convex1, convex2) {
		return f.fallbackNarrowphaseDetector.Detect(convex1, transform1, convex2, transform2)
	} else {
		return f.primaryNarrowphaseDetector.Detect(convex1, transform1, convex2, transform2)
	}
}

func (f *FallbackNarrowphaseDetector) DetectPenetration(convex1 geometry.Convexer, transform1 *geometry.Transform, convex2 geometry.Convexer, transform2 *geometry.Transform, penetration *Penetration) bool {
	if f.IsFallbackRequired(convex1, convex2) {
		return f.fallbackNarrowphaseDetector.DetectPenetration(convex1, transform1, convex2, transform2, penetration)
	} else {
		return f.primaryNarrowphaseDetector.DetectPenetration(convex1, transform1, convex2, transform2, penetration)
	}
}

func (f *FallbackNarrowphaseDetector) GetPrimaryNarrowphaseDetector() NarrowphaseDetector {
	return f.primaryNarrowphaseDetector
}

func (f *FallbackNarrowphaseDetector) GetFallbackNarrowphaseDetector() NarrowphaseDetector {
	return f.fallbackNarrowphaseDetector
}
