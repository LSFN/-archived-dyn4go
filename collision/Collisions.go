package collision

var ESTIMATED_COLLISIONS_PER_BODY int = 4

func GetEstimatedCollisionPairs(n int) int {
	return n * ESTIMATED_COLLISIONS_PER_BODY
}

func GetEstimatedCollisions() int {
	return ESTIMATED_COLLISIONS_PER_BODY
}
