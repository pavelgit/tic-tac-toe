package game

import (
	"errors"
	"fmt"
)

const (
	EMPTY_CELL = iota
	X_CELL     = iota
	O_CELL     = iota
)

const FIELD_SIZE = 3

func isEmptyCell(cell int) bool {
	return cell == EMPTY_CELL
}

func isXCell(cell int) bool {
	return cell == X_CELL
}

func isOCell(cell int) bool {
	return cell == O_CELL
}

func getCellSign(cell int) (rune, error) {
	switch cell {
	case EMPTY_CELL:
		return '.', nil
	case X_CELL:
		return 'X', nil
	case O_CELL:
		return '0', nil
	default:
		return 0, errors.New("the parameter is not a cell value")
	}
}

type Position struct {
	X, Y int
}

type Game struct {
	board [FIELD_SIZE][FIELD_SIZE]int
}

func NewGame() Game {
	return Game{board: [FIELD_SIZE][FIELD_SIZE]int{}}
}

func (game *Game) Move(p Position, sign int) {
	game.board[p.Y][p.X] = sign
}

func (game *Game) PrintBoard() error {
	for y := 0; y < len(game.board); y++ {
		row := game.board[y]
		for x := 0; x < len(row); x++ {
			sign, e := getCellSign(row[x])
			if e != nil {
				return e
			}
			fmt.Printf("%c ", sign)
		}
		fmt.Println()
	}
	fmt.Println()
	return nil
}
