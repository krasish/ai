package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
	"utils"
)

var (
	N      uint
	board  *utils.BoolMatrix
	queens []utils.Coord
)

func init() {
	var err error

	utils.ReadUnsigned(&N, "N = ")

	board = utils.NewBoolMatrix(N)
	board.GenerateOnePerCol()

	queens = make([]utils.Coord, N, N)

	for i := 0; i < int(N); i++ {
		queens[i], err = FindQueenOnCol(int(N), i)
		if err != nil {
			log.Fatalln("No queen on each column")
		}
	}
}

func MinConflicts(m *utils.BoolMatrix) {
	matrixSize := m.Size()
	rand.Seed(time.Now().UnixNano())

	steps := 0
	start := time.Now()
	defer func() {
		fmt.Printf("Calculation for N = %d took: %v\n", N, time.Since(start))
	}()

	for {
		if CheckForSolvedBoard2() {
			return
		}
		steps++
		if (steps % (int(N) * 9)) == 0 {
			fmt.Println("random replace")
			randomReplace()
		}

		currentCol := rand.Intn(matrixSize)
		currentRow := queens[currentCol].X

		currentConflicts := CountQueenConflicts(queens[currentCol])
		bestRows := []uint{currentRow}

		for i := uint(0); i < uint(matrixSize); i++ {
			queens[currentCol].X = i
			conflictsForRow := CountQueenConflicts(queens[currentCol])
			if conflictsForRow > currentConflicts {
				bestRows = []uint{i}
				currentConflicts = conflictsForRow
			} else if conflictsForRow == currentConflicts {
				bestRows = append(bestRows, i)
			}
		}
		selectedRow := bestRows[rand.Intn(len(bestRows))]

		board.MoveFalse(utils.Coord{currentRow, queens[currentCol].Y}, utils.Coord{selectedRow, queens[currentCol].Y})
		queens[currentCol].X = selectedRow
	}
}

func CheckForSolvedBoard2() bool {
	//solved := false
	for i := 0; i < int(N); i++ {
		for j := 0; j < int(N); j++ {
			pass := false
			switch {
			case i == j:
				continue
			case queens[i].X == queens[j].X:
				pass = true
			case queens[i].Y == queens[j].Y:
				pass = true
			case math.Abs(float64(int(queens[i].X)-int(queens[j].X))) == math.Abs(float64(int(queens[i].Y)-int(queens[j].Y))):
				pass = true
			}
			if pass != true {
				return false
			}
		}
	}
	return true
}

func CheckForSolvedBoard() bool {
	for i := 0; i < int(N); i++ {
		for j := 0; j < int(N); j++ {
			switch {
			case i == j:
				continue
			case queens[i].X == queens[j].X:
				return false
			case queens[i].Y == queens[j].Y:
				return false
			case math.Abs(float64(int(queens[i].X)-int(queens[j].X))) == math.Abs(float64(int(queens[i].Y)-int(queens[j].Y))):
				return false
			}
		}
	}
	return true
}

func CountQueenConflicts(currQueen utils.Coord) uint {
	res := uint(0)

	for j := 0; j < int(N); j++ {
		switch {
		case currQueen.X == queens[j].X && currQueen.Y == queens[j].Y:
			continue
		case currQueen.X == queens[j].X:
			res += 1
		case currQueen.Y == queens[j].Y:
			res += 1
		case math.Abs(float64(int(currQueen.X)-int(queens[j].X))) == math.Abs(float64(int(currQueen.Y)-int(queens[j].Y))):
			res += 1
		}
	}

	return res
}

func FindQueenOnCol(size, col int) (utils.Coord, error) {
	res := utils.Coord{
		Y: uint(col),
	}
	for i := 0; i < size; i++ {
		res.X = uint(i)
		if !board.Get(res) {
			return res, nil
		}
	}
	return utils.Coord{}, errors.New(fmt.Sprintf("Could not find queen on col %d\n", col))
}

func randomReplace() {
	rand.Seed(time.Now().UnixNano())
	col := rand.Intn(int(N))
	row := rand.Intn(int(N))

	board.MoveFalse(queens[col], utils.Coord{uint(row), queens[col].Y})
	queens[col].X = uint(row)
}

func CountBoardConflicts() uint {
	res := uint(0)
	for i := 0; i < int(N); i++ {
		for j := 0; j < int(N); j++ {
			switch {
			case i == j:
				continue
			case queens[i].X == queens[j].X:
				res += 1
			case queens[i].Y == queens[j].Y:
				res += 1
			case math.Abs(float64(int(queens[i].X)-int(queens[j].X))) == math.Abs(float64(int(queens[i].Y)-int(queens[j].Y))):
				res += 1
			}
		}
	}
	return res / 2
}

func main() {
	fmt.Println("Initial board")
	fmt.Println()
	fmt.Println(board)
	fmt.Println()
	fmt.Println()
	MinConflicts(board)
	fmt.Println("Final board")
	fmt.Println()
	fmt.Println(board)
}
