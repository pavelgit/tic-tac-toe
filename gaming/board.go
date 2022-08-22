package gaming

import (
	"fmt"
)

const BOARD_SIZE = 3

type Position struct {
	X, Y int
}

type Board struct {
	signs [BOARD_SIZE][BOARD_SIZE]int
}

func NewBoard() Board {
	return Board{
		signs: [BOARD_SIZE][BOARD_SIZE]int{},
	}
}

func (board *Board) Get(x int, y int) int {
	return board.signs[y][x]
}

func (board *Board) Move(x int, y int, sign int) {
	board.signs[y][x] = sign
}

func (board *Board) ForEachCell(callback func(x int, y int, sign int) bool) {
	for y := 0; y < len(board.signs); y++ {
		row := board.signs[y]
		for x := 0; x < len(row); x++ {
			if callback(x, y, row[x]) {
				break
			}
		}
	}
}

func (board *Board) PrintBoard() {
	board.ForEachCell(func(x int, y int, sign int) bool {
		char := GetSignChar(sign)
		fmt.Printf("%c ", char)
		if x == BOARD_SIZE-1 {
			fmt.Println()
		}
		return false
	})
	fmt.Println()
}
