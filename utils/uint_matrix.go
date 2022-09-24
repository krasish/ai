package utils

import (
	"errors"
	"strconv"
	"strings"
)

type UintMatrix [][]uint

func NewUintMatrix(size uint) *UintMatrix {
	res := make(UintMatrix, size)
	for i, _ := range res {
		res[i] = make([]uint, size)
	}
	return &res
}

func (m *UintMatrix) Size() int {
	return len(*m)
}

func (m *UintMatrix) Get(c Coord) uint {
	return (*m)[c.X][c.Y]
}

func (m *UintMatrix) GetByInts(x, y int) uint {
	return (*m)[x][y]
}

func (m *UintMatrix) Find(u uint) (Coord, error) {
	s := m.Size()
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			if m.GetByInts(i, j) == u {
				return Coord{uint(i), uint(j)}, nil
			}
		}
	}
	return Coord{}, errors.New("not found")
}

func (m *UintMatrix) Duplicate() *UintMatrix {
	s := m.Size()
	m2 := make(UintMatrix, s)
	for i, curr := range *(m) {
		m2[i] = make([]uint, s)
		copy(m2[i], curr)
	}
	return &m2
}

func (m *UintMatrix) Read() {
	s := m.Size()
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			readUntilCorrectU(&(*m)[i][j])
		}
	}
}

func (m *UintMatrix) String() string {
	s := len(*m)
	b := strings.Builder{}
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			b.WriteString(strconv.Itoa(int((*m)[i][j])))
			if j < s-1 {
				b.WriteString("\t")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
