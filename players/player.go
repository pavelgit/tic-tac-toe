package players

import (
	"tictactoe/gaming"
)

type Player interface {
	getNextMove(game *gaming.Game) gaming.Position
}
