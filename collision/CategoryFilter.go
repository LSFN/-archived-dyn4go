package collision

import (
	"math"
)

type CategoryFilter struct {
	category, mask int
}

func NewCategoryFilter() {
	c := new(CategoryFilter)
	c.category = 1
	c.mask = int(^uint(0) >> 1)
}

func (c *CategoryFilter) IsAllowed(filter Filterer) bool {
	if filter == nil {
		return true
	}
	if cf, ok := filter.(*CategoryFilter); ok {
		return (c.category&cf.mask) > 0 && (cf.category&c.mask) > 0
	}
	return true
}

func (c *CategoryFilter) GetCategory() int {
	return c.category
}

func (c *CategoryFilter) GetMask() int {
	return c.mask
}

func (c *CategoryFilter) SetCategory(category int) {
	c.category = category
}

func (c *CategoryFilter) SetMask(mask int) {
	c.mask = mask
}
