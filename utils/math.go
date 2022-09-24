package utils

import "math"

func UintSqrt(u uint) uint {
	return uint(math.Sqrt(float64(u)))
}
