package util

// CalculateLimit calculates the limit and offset for a query
func CalculateLimit(current, queryCount int) (int, int) {
	if current < 1 {
		current = 1
	}
	if queryCount < 1 {
		queryCount = 10
	}
	return (current - 1) * queryCount, queryCount
}
