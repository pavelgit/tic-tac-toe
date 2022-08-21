package main

import (
	"tictactoe/game"
)

func main() {
	g := game.NewGame()
	g.PrintBoard()
	g.Move(game.Position{X: 1, Y: 1}, game.X_CELL)
	g.PrintBoard()
}
