package utils

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type BoolMatrix [][]bool

func NewBoolMatrix(size uint) *BoolMatrix {
	res := make(BoolMatrix, size)
	for i, _ := range res {
		res[i] = make([]bool, size)
	}
	return &res
}

func (m *BoolMatrix) Read() {
	s := len(*m)
	var temp uint
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			readUntilCorrectU(&temp)
			if temp == 1 {
				(*m)[i][j] = true
			}
		}
	}
}

func (m *BoolMatrix) Generate(numUnreachable int) {
	s := len(*m)
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			(*m)[i][j] = true
		}
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numUnreachable; i++ {
		coord := NewCoordFromUint(uint(rand.Intn(s*s)), s)
		for !(*m)[coord.X][coord.Y] {
			coord = NewCoordFromUint(uint(rand.Intn(s*s)), s)
		}
		(*m)[coord.X][coord.Y] = false
	}
}

func (m *BoolMatrix) GenerateOnePerCol() {
	s := len(*m)
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			(*m)[i][j] = true
		}
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < s; i++ {
		j := uint(rand.Intn(s))
		(*m)[j][i] = false
	}
}

func (m *BoolMatrix) MoveFalse(from, to Coord) {
	(*m)[from.X][from.Y], (*m)[to.X][to.Y] = true, false
}

func (m *BoolMatrix) Size() int {
	return len(*m)
}

func (m *BoolMatrix) String() string {
	s := len(*m)
	b := strings.Builder{}
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			if (*m)[i][j] {
				b.WriteString("1")
			} else {
				b.WriteString("0")
			}
			if j < s-1 {
				b.WriteString("\t")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m *BoolMatrix) StringWithPath(path []Coord) string {
	pm := map[Coord]struct{}{}
	for _, c := range path {
		pm[c] = struct{}{}
	}

	s := len(*m)
	b := strings.Builder{}
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			if (*m)[i][j] {
				if _, inPath := pm[Coord{uint(i), uint(j)}]; inPath {
					b.WriteString("*")
				} else {
					b.WriteString("1")
				}
			} else {
				b.WriteString("0")
			}
			if j < s-1 {
				b.WriteString("\t")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m *BoolMatrix) Get(c Coord) bool {
	return (*m)[c.X][c.Y]
}

func (m *BoolMatrix) SafeGet(c Coord) (bool, error) {
	s := uint(len(*m))
	if c.X < s && c.Y < s {
		return (*m)[c.X][c.Y], nil
	}
	return false, errors.New("unreachable coordinates")
}
