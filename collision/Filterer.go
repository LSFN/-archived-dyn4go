package collision

type Filterer interface {
	isAllowed(filter Filterer) bool
}
