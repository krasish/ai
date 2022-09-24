package main

import "utils"

type Board struct {
	*utils.UintMatrix
}

func NewBoard(numDigits uint) *Board {
	return &Board{UintMatrix: utils.NewUintMatrix(utils.UintSqrt(numDigits + 1))}
}

func (b *Board) ManhattanDistance() (res uint) {
	s := b.Size()
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			val := b.GetByInts(i, j)
			if val == 0 {
				continue
			}
			if uint(I) > val { //Note that in case of -1 the cast will overflow and the "I" will definitely be bigger which is the desired state.
				val = val - 1 // In case the zero is after the current value, subtract 1 to get the correct spot
			}
			correct := utils.NewCoordFromUint(val, s)
			res += correct.Distance(utils.NewCoordFromInts(i, j))
		}
	}
	return
}

func (b *Board) Solvable() bool {
	if (b.Size() % 2) == 1 {
		return (b.inversions() % 2) == 0
	} else {
		blankCoords, err := b.Find(0)
		if err != nil {
			return false
		}
		return ((b.inversions() + blankCoords.X) % 2) == 1
	}
}

func (b *Board) Neighbours() []*Board {
	var (
		res     = make([]*Board, 0, 4)
		zero, _ = b.Find(0)
		s       = uint(b.Size())
	)
	if zero.X != 0 {
		n := &Board{b.Duplicate()}
		(*n.UintMatrix)[zero.X][zero.Y], (*n.UintMatrix)[zero.X-1][zero.Y] = (*n.UintMatrix)[zero.X-1][zero.Y], (*n.UintMatrix)[zero.X][zero.Y]
		res = append(res, n)
	}
	if zero.Y != 0 {
		n := &Board{b.Duplicate()}
		(*n.UintMatrix)[zero.X][zero.Y], (*n.UintMatrix)[zero.X][zero.Y-1] = (*n.UintMatrix)[zero.X][zero.Y-1], (*n.UintMatrix)[zero.X][zero.Y]
		res = append(res, n)
	}
	if zero.X < s-1 {
		n := &Board{b.Duplicate()}
		(*n.UintMatrix)[zero.X][zero.Y], (*n.UintMatrix)[zero.X+1][zero.Y] = (*n.UintMatrix)[zero.X+1][zero.Y], (*n.UintMatrix)[zero.X][zero.Y]
		res = append(res, n)
	}
	if zero.Y < s-1 {
		n := &Board{b.Duplicate()}
		(*n.UintMatrix)[zero.X][zero.Y], (*n.UintMatrix)[zero.X][zero.Y+1] = (*n.UintMatrix)[zero.X][zero.Y+1], (*n.UintMatrix)[zero.X][zero.Y]
		res = append(res, n)
	}
	return res
}

func (b *Board) Solved() bool {
	var (
		s        = b.Size()
		expected = uint(1)
	)
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			actual := b.Get(utils.NewCoordFromInts(i, j))
			if actual == 0 {
				if uint(i) == zeroCoord.X && uint(j) == zeroCoord.Y {
					continue
				}
				return false
			}
			if actual != expected {
				return false
			}
			expected++
		}
	}
	return true
}

func (b *Board) Item() *Item {
	return &Item{
		value:    *b,
		priority: int(b.ManhattanDistance()),
	}
}

func (b *Board) inversions() (inversions uint) {
	var curr, next uint
	s := b.Size()
	s2 := s * s
	for i := 0; i < s2; i++ {
		if curr = b.Get(utils.NewCoordFromUint(uint(i), s)); curr == 0 {
			continue
		}
		for j := i; j < s2; j++ {
			if next = b.Get(utils.NewCoordFromUint(uint(j), s)); next == 0 {
				continue
			}
			if curr > next {
				inversions++
			}
		}
	}
	return
}
