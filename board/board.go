package board

import (
	"fmt"
	"math/rand"
)

type Board [][]int

type Moves interface {
	right()
	left()
	up()
	down()
}

func getEmptyRows(board [][]int) []int {
	rows := []int{}
	for idx, row := range board {
		for _, val := range row {
			if val == 0 {
				rows = append(rows, idx)
				break
			}
		}
	}
	return rows
}

func getEmptyCols(row []int) []int {
	cols := []int{}
	for idx, val := range row {
		if val == 0 {
			cols = append(cols, idx)
		}
	}
	return cols
}

func (b Board) Spawn() {
	emptyRows := getEmptyRows(b)
	row := emptyRows[rand.Intn(len(emptyRows))]
	emptyCols := getEmptyCols(b[row])
	col := emptyCols[rand.Intn(len(emptyCols))]
	b[row][col] = 2
}

func getLeftNonZero(row []int, ptr int) int {
	for i := ptr - 1; i >= 0; i-- {
		if row[i] != 0 {
			return i
		}
	}
	return -1
}

func moveRight(row []int) {
	ptr := len(row) - 1
	for ptr > 0 {
		if row[ptr] == 0 {
			nonZeroIdx := getLeftNonZero(row, ptr)
			if nonZeroIdx != -1 {
				row[ptr], row[nonZeroIdx] = row[nonZeroIdx], row[ptr]
			}
		}
		ptr -= 1
	}
}

// Right defines swipe right action
func (b Board) Right() {
	for _, row := range b {
		// first swipe non zero to right
		moveRight(row)

		// now merge same element(s)
		for i := len(row) - 1; i > 0; i-- {
			if row[i] != 0 && row[i] == row[i-1] {
				row[i] *= 2
				row[i-1] = 0
				moveRight(row)
			}
		}
	}
}

func getRightNonZero(row []int, ptr int) int {
	for i := ptr + 1; i < len(row); i++ {
		if row[i] != 0 {
			return i
		}
	}
	return -1
}

func moveLeft(row []int) {
	ptr := 0
	for ptr < len(row)-1 {
		if row[ptr] == 0 {
			nonZeroIdx := getRightNonZero(row, ptr)
			if nonZeroIdx != -1 {
				row[ptr], row[nonZeroIdx] = row[nonZeroIdx], row[ptr]
			}
		}
		ptr += 1
	}
}

// Left defines swipe left action
func (b Board) Left() {
	for _, row := range b {
		// first swipe non zero to left
		moveLeft(row)

		// now merge same element(s)
		for i := 0; i < len(row)-1; i++ {
			if row[i] != 0 && row[i] == row[i+1] {
				row[i] *= 2
				row[i+1] = 0
				moveLeft(row)
			}
		}
	}
}

func transpose(board [][]int) {
	temp := make([][]int, 4)
	for i := 0; i < len(board[0]); i++ {
		temp[i] = make([]int, 4)
		for j := 0; j < len(board); j++ {
			temp[i][j] = board[j][i]
		}
	}
	copy(board, temp)
}

// Up defines the up swipe action
func (b Board) Up() {
	// Up swipe is same as left swipe of transpose of the board.
	transpose(b)
	b.Left()
	// revert the transpose
	transpose(b)
}

// Down defines the down swipe action
func (b Board) Down() {
	// Down swipe is same as right swipe of transpose of the board.
	transpose(b)
	b.Right()
	// revert the transpose
	transpose(b)
}

// GameOver checks whether the game is over
func (b Board) GameOver() (bool, bool) {
	// check if player has won
	for _, row := range b {
		for _, col := range row {
			if col == 2048 {
				return true, true
			}
		}
	}

	spaceLeft := false
	// check if board is full
	for _, row := range b {
		for _, col := range row {
			if col == 0 {
				spaceLeft = true
			}
		}
	}
	return !spaceLeft, false
}

// Format board for printing
func (b Board) String() string {
	str := "\n-----------------------------------------------------------------\n"
	for _, row := range b {
		for i := 0; i < len(row); i++ {
			str += "| \t\t"
		}
		str += "|\n| "
		for _, col := range row {
			str += fmt.Sprintf("\t%d\t|", col)
		}
		str += "\n"
		for i := 0; i < len(row); i++ {
			str += "| \t\t"
		}
		str += "|\n-----------------------------------------------------------------\n"
	}
	return str
}
