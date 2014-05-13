package collision

type Filterer interface {
	IsAllowed(filter Filterer) bool
}
