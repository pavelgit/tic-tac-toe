package gaming

type Game struct {
	Board Board
	Turn  int
}

func NewGame() Game {
	return Game{
		Board: NewBoard(),
		Turn:  X_SIGN,
	}
}

func (game *Game) switchTurn() {
	switch game.Turn {
	case X_SIGN:
		game.Turn = O_SIGN
	case O_SIGN:
		game.Turn = X_SIGN
	}
}

func (game *Game) Move(x int, y int) (bool, int) {
	game.Board.Move(x, y, game.Turn)
	game.switchTurn()
	return game.isOver()
}

func (game *Game) getProjectionSign(mask []Position) (bool, int) {
	var foundSign = NO_SIGN
	for _, p := range mask {
		sign := game.Board.Get(p.X, p.Y)
		if sign == NO_SIGN {
			return false, 0
		}
		if foundSign == NO_SIGN {
			foundSign = sign
		} else if foundSign != sign {
			return false, 0
		}
	}
	return true, foundSign
}

func (game *Game) getWinner() (bool, int) {
	for _, mask := range winningMaskCells {
		winnerFound, winnerSign := game.getProjectionSign(mask)
		if winnerFound {
			return true, winnerSign
		}
	}
	return false, 0
}

func (game *Game) isOver() (bool, int) {
	winnerFound, winnerSign := game.getWinner()
	if winnerFound {
		return true, winnerSign
	}

	return !game.hasEmptyCells(), 0
}

func (game *Game) hasEmptyCells() bool {
	var emptySignFound = false
	game.Board.ForEachCell(func(x int, y int, sign int) bool {
		if sign == NO_SIGN {
			emptySignFound = true
			return true
		}
		return false
	})
	return emptySignFound
}
