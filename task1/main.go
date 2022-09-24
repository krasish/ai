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

func DFS(current, end utils.Coord, vis map[utils.Coord]struct{}) []utils.Coord {
	if current == end {
		return []utils.Coord{end}
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

		res := DFS(n, end, vis)
		if len(res) != 0 {
			return append([]utils.Coord{current}, res...)
		}
	}
	return []utils.Coord{}
}

func BFS(current, end utils.Coord) []utils.Coord {
	var (
		q     = utils.NewCoordQ()
		added = make(map[utils.Coord]utils.Coord, 0)
		err   error
	)
	if !(*m)[current.X][current.Y] {
		return []utils.Coord{}
	}

	q.Push(current)
	added[current] = current

	for current, err = q.Pop(); err == nil; current, err = q.Pop() {
		if current == end {
			return returnPath(end, added)
		}
		neighbours := current.GetNeighbours(N)
		for _, n := range neighbours {
			if _, alreadyAdded := added[n]; (*m)[n.X][n.Y] && !alreadyAdded {
				q.Push(n)
				added[n] = current
			}
		}
		if q.IsEmpty() {
			return []utils.Coord{}
		}
	}
	return []utils.Coord{}
}

func returnPath(current utils.Coord, added map[utils.Coord]utils.Coord) []utils.Coord {
	res := []utils.Coord{current}
	for parent := added[current]; parent != current; parent = added[current] {
		res = append([]utils.Coord{parent}, res...)
		current = parent
	}

	return res
}

func main() {

	for {
		path := DFS(start, end, map[utils.Coord]struct{}{})
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

//func main() {
//	for {
//		path := BFS(start, end)
//		if len(path) == 0 {
//			fmt.Println("There is no path.")
//		}
//		matrixWithPath := m.StringWithPath(path)
//		fmt.Println(matrixWithPath)
//
//		fmt.Println()
//		fmt.Println("--------------------------------")
//		fmt.Println()
//		utils.ReadUnsigned(&start.X, "startX = ")
//		utils.ReadUnsigned(&start.Y, "startY = ")
//		utils.ReadUnsigned(&end.X, "endX = ")
//		utils.ReadUnsigned(&end.Y, "endY = ")
//	}
//}
