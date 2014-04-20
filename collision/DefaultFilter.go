package collision

type DefaultFilter struct{}

func NewDefaultFilter() Filterer {
	return new(DefaultFilter)
}

func (d *DefaultFilter) IsAllowed(filter Filterer) bool {
	return true
}
