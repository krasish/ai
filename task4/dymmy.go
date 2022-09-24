package main

import (
	"log"
	"math/rand"
	"time"
	"utils"
)

func MinConflictsDummy(m *utils.BoolMatrix) {
	matrixSize := m.Size()
	rand.Seed(time.Now().UnixNano())
	for {
		if CheckForSolvedBoardDummy(matrixSize) {
			return
		}

		currentCol := rand.Intn(matrixSize)

		currentQueen, err := FindQueenOnCol(matrixSize, currentCol)
		if err != nil {
			log.Fatalln(err)
		}

		currentConflicts := FindConflictsForQueenDummy(matrixSize, currentQueen)
		bestRows := []uint{currentConflicts}
		for i := uint(0); i < uint(matrixSize); i++ {
			conflictsForRow := FindConflictsForQueenDummy(matrixSize, utils.Coord{i, currentQueen.Y})
			if conflictsForRow < currentConflicts {
				bestRows = []uint{i}
				currentConflicts = conflictsForRow
			} else if conflictsForRow == currentConflicts {
				bestRows = append(bestRows, i)
			}
		}
		selectedRow := bestRows[rand.Intn(len(bestRows))]
		board.MoveFalse(currentQueen, utils.Coord{selectedRow, currentQueen.Y})
	}
}

func CheckForSolvedBoardDummy(size int) bool {
	for i := uint(0); i < uint(size); i++ {
		currentQueen, err := FindQueenOnCol(size, int(i))
		if err != nil {
			return false
		}
		currentConflicts := FindConflictsForQueenDummy(size, currentQueen)
		if currentConflicts != 0 {
			return false
		}
	}
	return true
}

func FindConflictsForQueenDummy(size int, q utils.Coord) uint {
	var (
		conflicts      = uint(0)
		continueSearch = true
	)
	for i := uint(1); i < uint(size) && continueSearch; i++ {
		continueSearch = false

		conflicts, continueSearch = countConflictsForQueen(conflicts, continueSearch, utils.Coord{q.X + i, q.Y + i})
		conflicts, continueSearch = countConflictsForQueen(conflicts, continueSearch, utils.Coord{q.X, q.Y + i})
		conflicts, continueSearch = countConflictsForQueen(conflicts, continueSearch, utils.Coord{q.X - i, q.Y + i})

		conflicts, continueSearch = countConflictsForQueen(conflicts, continueSearch, utils.Coord{q.X + i, q.Y})
		conflicts, continueSearch = countConflictsForQueen(conflicts, continueSearch, utils.Coord{q.X - i, q.Y})

		conflicts, continueSearch = countConflictsForQueen(conflicts, continueSearch, utils.Coord{q.X + i, q.Y - i})
		conflicts, continueSearch = countConflictsForQueen(conflicts, continueSearch, utils.Coord{q.X, q.Y - i})
		conflicts, continueSearch = countConflictsForQueen(conflicts, continueSearch, utils.Coord{q.X - i, q.Y - i})
	}
	return conflicts
}

func countConflictsForQueen(currentConflicts uint, currentState bool, c utils.Coord) (uint, bool) {
	if curr, err := board.SafeGet(c); err == nil {
		if !curr {
			return currentConflicts + 1, true
		}
		return currentConflicts, true
	}
	return currentConflicts, currentState
}
