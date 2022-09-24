package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"utils"
)

var (
	maxDepth     = uint(3)
	computerPlay bool
	board        *Board
)

func init() {
	board = NewBoard()
	utils.ReadYesNo(&computerPlay, "Should I play first?\n")

}

func printBoardSeparated() {
	fmt.Println()
	fmt.Println(board)
	fmt.Println()
}

func playFirstMove() {
	if computerPlay {
		rand.Seed(time.Now().UnixNano())
		board.PlayO(utils.Coord{
			X: uint(rand.Intn(BoardSize)),
			Y: uint(rand.Intn(BoardSize)),
		})
		computerPlay = !computerPlay
		printBoardSeparated()
	}
}

func scanPlayerMove() {
	c := utils.Coord{}
	fmt.Println("Enter a move:")
	for {
		for {
			utils.ReadUnsigned(&c.X, "[")
			if c.X > 2 {
				fmt.Println("\nEnter an integer value in [0, 2]")
			} else {
				break
			}
		}
		for {
			utils.ReadUnsigned(&c.Y, fmt.Sprintf("[%d,", c.X))
			if c.X > 2 {
				fmt.Println("\nEnter an integer value in [0, 2]")
			} else {
				break
			}
		}
		if !board.CanPlay(c) {
			fmt.Printf("Cannot play [%d,%d]\n", c.X, c.Y)
		} else {
			board.PlayX(c)
			break
		}
	}

}

func MiniMaxABAlgorithm(b *Board, depth uint, alpha, beta int, maximizingPlayer bool) int {
	if depth == 0 || b.IsGameOver() {
		score := b.GetScore()
		switch score {
		case WinningScore:
			return score + int(depth)
		case LosingScore:
			return score - int(depth)
		default:
			return score
		}
	}

	if maximizingPlayer {
		var (
			maxVal    = -100
			nextMoves = b.GetPossibleMovesForPlayer(O_SET)
		)
		for i, _ := range nextMoves {
			val := MiniMaxABAlgorithm(nextMoves[i], depth-1, alpha, beta, false)
			maxVal = int(math.Max(float64(maxVal), float64(val)))
			alpha = int(math.Max(float64(alpha), float64(val)))
			if beta <= alpha {
				break
			}
		}
		return maxVal
	} else {
		var (
			minVal    = 100
			nextMoves = b.GetPossibleMovesForPlayer(X_SET)
		)
		for i, _ := range nextMoves {
			val := MiniMaxABAlgorithm(nextMoves[i], depth-1, alpha, beta, true)
			minVal = int(math.Min(float64(minVal), float64(val)))
			beta = int(math.Min(float64(beta), float64(val)))
			if beta <= alpha {
				break
			}
		}
		return minVal
	}
}

func GameWithAlphaBetaAlgorithm() {
	playFirstMove()
	for {
		if computerPlay {
			var (
				nextMoves = board.GetPossibleMovesForPlayer(O_SET)
				best      = -100
				bestIndex = 0
			)
			for i, move := range nextMoves {
				val := MiniMaxABAlgorithm(move, maxDepth, -100, 100, false)
				if val > best {
					bestIndex = i
					best = val
				}
			}
			board = nextMoves[bestIndex]
		} else {
			scanPlayerMove()
		}
		printBoardSeparated()
		if board.IsWinningBoardFor(O_SET) {
			fmt.Println("You lose!")
			break
		} else if board.IsWinningBoardFor(X_SET) {
			fmt.Println("You win!")
			break
		} else if board.IsTieGame() {
			fmt.Println("Draw!")
			break
		}
		computerPlay = !computerPlay
	}
}

func main() {
	GameWithAlphaBetaAlgorithm()
}
