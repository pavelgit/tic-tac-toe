package main

import (
	"fmt"
	"tictactoe/gaming"
	"tictactoe/players"
)

func main() {
	g := gaming.NewGame()

	p := players.RandomPlayer{}

	g.Board.PrintBoard()
	for {
		position := p.GetNextMove(&g)
		isOver, winnerSign := g.Move(position.X, position.Y)
		g.Board.PrintBoard()
		if isOver {
			fmt.Printf("%c\n", gaming.GetSignChar(winnerSign))
			break
		}
	}
}
