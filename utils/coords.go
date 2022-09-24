package utils

import (
	"errors"
	"math"
)

const MaxNeighbours = 4

type Coord struct {
	X, Y uint
}

func NewCoordFromInts(x, y int) Coord {
	return Coord{
		X: uint(x),
		Y: uint(y),
	}
}

func NewCoordFromUint(n uint, size int) Coord {
	return Coord{
		X: n / uint(size),
		Y: n % uint(size),
	}
}

func (c Coord) GetNeighbours(limit uint) []Coord {
	n := make([]Coord, 0, MaxNeighbours)
	if c.X != 0 {
		n = append(n, Coord{c.X - 1, c.Y})
	}
	if c.X < limit-1 {
		n = append(n, Coord{c.X + 1, c.Y})
	}
	if c.Y != 0 {
		n = append(n, Coord{c.X, c.Y - 1})
	}
	if c.Y < limit-1 {
		n = append(n, Coord{c.X, c.Y + 1})
	}
	return n
}

func (c Coord) Distance(c2 Coord) uint {
	cx, cy, c2x, c2y := int(c.X), int(c.Y), int(c2.X), int(c2.Y)
	return uint(math.Abs(float64(cx-c2x)) + math.Abs(float64(cy-c2y)))
}

type CoordQ []Coord

func NewCoordQ() *CoordQ {
	q := make(CoordQ, 0)
	return &q
}

func (q *CoordQ) Push(coord Coord) {
	*q = append(*q, coord)
}

func (q *CoordQ) Pop() (Coord, error) {
	if len(*q) == 0 {
		return Coord{}, errors.New("empty queue")
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res, nil
}

func (q *CoordQ) IsEmpty() bool {
	return len(*q) == 0
}
