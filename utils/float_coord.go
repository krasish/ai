package utils

import (
	"math"
)

type FloatCoord struct {
	X, Y float64
}

func NewFloatCoordFromInts(x, y int) FloatCoord {
	return FloatCoord{
		X: float64(x),
		Y: float64(y),
	}
}

//func NewFloatCoordFromUint(n uint, size int) FloatCoord {
//	return FloatCoord{
//		X: float64(n / uint(size)),
//		Y: float64(n % uint(size)),
//	}
//}

func (c FloatCoord) Distance(c2 FloatCoord) float64 {
	return math.Sqrt(math.Pow(c.X-c2.X, 2) + math.Pow(c.Y-c2.Y, 2))
}
