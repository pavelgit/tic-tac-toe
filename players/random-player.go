package players

import (
	"math/rand"
	"tictactoe/gaming"
	"time"
)

type RandomPlayer struct {
}

func getEmptyCellCount(game *gaming.Game) int {
	var emptyCellCount = 0
	game.Board.ForEachCell(func(x int, y int, sign int) bool {
		if sign == gaming.NO_SIGN {
			emptyCellCount++
		}
		return false
	})
	return emptyCellCount
}

func getRandom(max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(max)
}

func (p *RandomPlayer) GetNextMove(game *gaming.Game) gaming.Position {
	var emptyCellCount = getEmptyCellCount(game)
	cellNumber := getRandom(emptyCellCount)

	var resultingPosition gaming.Position
	game.Board.ForEachCell(func(x int, y int, sign int) bool {
		if sign == gaming.NO_SIGN {
			cellNumber--
			if cellNumber < 0 {
				resultingPosition = gaming.Position{X: x, Y: y}
				return true
			}
		}
		return false
	})

	return resultingPosition
}
