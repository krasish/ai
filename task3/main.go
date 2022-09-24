package main

import (
	"container/heap"
	"fmt"
	"log"
	"utils"
)

var (
	N          uint
	I          int
	startBoard *Board
	zeroCoord  utils.Coord
)

func init() {
	utils.ReadUnsigned(&N, "N = ")
	utils.ReadInt(&I, "I = ")

	startBoard = NewBoard(N)
	if I == -1 {
		zeroCoord = utils.NewCoordFromUint(N, startBoard.Size())
	} else {
		zeroCoord = utils.NewCoordFromUint(uint(I), startBoard.Size())
	}

	startBoard.Read()
	if !startBoard.Solvable() {
		log.Fatalln("the entered board is not solvable")
	}

}

func Solve() []*Board {
	var (
		pq     = make(PriorityQueue, 0)
		boards = make([]*Board, 0)
	)

	pq.Push(startBoard.Item())
	heap.Init(&pq)
	current := pq.Pop().(*Item)

	for ; !current.value.Solved(); current = pq.Pop().(*Item) {
		neighbours := current.Neighbours()
		for i, _ := range neighbours {
			pq.Push(neighbours[i])
		}
	}

	for current.Parent != nil {
		boards = append([]*Board{&current.value}, boards...)
		current = current.Parent
	}
	boards = append([]*Board{&current.value}, boards...)
	return boards
}

func main() {
	solution := Solve()
	fmt.Println("---- NUMBER OF STEPS IN SOLUTION ----")
	fmt.Println(len(solution) - 1) //the first node is the starting node
	fmt.Println("---- SOLUTION ----")
	for i, _ := range solution {
		fmt.Println()
		fmt.Println(solution[i].UintMatrix)
		fmt.Println()
	}
}
