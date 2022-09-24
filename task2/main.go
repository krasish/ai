package main

import (
	"fmt"
	"utils"
)

var (
	k          int
	N          uint
	m          *utils.BoolMatrix
	start, end utils.Coord
)

func init() {
	utils.ReadUnsigned(&N, "N = ")
	utils.ReadInt(&k, "k = ")
	m = utils.NewBoolMatrix(N)
	if k < 0 {
		m.Read()
	} else {
		m.Generate(k)
	}

	fmt.Println()
	fmt.Println(m)
	fmt.Println()

	utils.ReadUnsigned(&start.X, "startX = ")
	utils.ReadUnsigned(&start.Y, "startY = ")
	utils.ReadUnsigned(&end.X, "endX = ")
	utils.ReadUnsigned(&end.Y, "endY = ")
}

func IDS(current, end utils.Coord) []utils.Coord {
	maxDepth := 10
	for i := 1; i <= maxDepth; i++ {
		path := DLS(current, end, i, map[utils.Coord]struct{}{})
		if len(path) != 0 {
			fmt.Printf("Found for depth %d\n", i)
			return path
		}
	}
	return []utils.Coord{}
}

func DLS(current, end utils.Coord, limit int, vis map[utils.Coord]struct{}) []utils.Coord {
	if current == end {
		return []utils.Coord{end}
	}
	if limit == 0 {
		return []utils.Coord{}
	}

	vis[current] = struct{}{}
	neighbours := current.GetNeighbours(N)
	for _, n := range neighbours {
		if !(*m)[n.X][n.Y] {
			continue
		}
		if _, alreadyVisited := vis[n]; alreadyVisited {
			continue
		}

		res := DLS(n, end, limit-1, vis)
		if len(res) != 0 {
			return append([]utils.Coord{current}, res...)
		}
	}
	return []utils.Coord{}
}

func main() {
	for {
		path := IDS(start, end)
		if len(path) == 0 {
			fmt.Println("There is no path.")
		}
		matrixWithPath := m.StringWithPath(path)
		fmt.Println(matrixWithPath)

		fmt.Println()
		fmt.Println("--------------------------------")
		fmt.Println()
		utils.ReadUnsigned(&start.X, "startX = ")
		utils.ReadUnsigned(&start.Y, "startY = ")
		utils.ReadUnsigned(&end.X, "endX = ")
		utils.ReadUnsigned(&end.Y, "endY = ")
	}
}
