package numbers

// GCD finds the greatest common denominator of two integers,
// implemented using the classic Euclidian Algorithm.
func GCD(x int, y int) int {
	if y == 0 {
		return 0
	}
	if x == 0 {
		return y
	}
	var z int
	for {
		if z = x % y; z == 0 {
			if y < 0 {
				return y * -1
			}
			return y
		}
		x, y = y, z
	}
}
