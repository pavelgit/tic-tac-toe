package gaming

var winningMaskCells = [...][]Position{
	{
		Position{X: 0, Y: 0},
		Position{X: 1, Y: 0},
		Position{X: 2, Y: 0},
	},
	{
		Position{X: 0, Y: 1},
		Position{X: 1, Y: 1},
		Position{X: 2, Y: 1},
	},
	{
		Position{X: 0, Y: 2},
		Position{X: 1, Y: 2},
		Position{X: 2, Y: 2},
	},
	{
		Position{X: 0, Y: 0},
		Position{X: 0, Y: 1},
		Position{X: 0, Y: 2},
	},
	{
		Position{X: 1, Y: 0},
		Position{X: 1, Y: 1},
		Position{X: 1, Y: 2},
	},
	{
		Position{X: 2, Y: 0},
		Position{X: 2, Y: 1},
		Position{X: 2, Y: 2},
	},
	{
		Position{X: 0, Y: 0},
		Position{X: 1, Y: 1},
		Position{X: 2, Y: 2},
	},
	{
		Position{X: 2, Y: 0},
		Position{X: 1, Y: 1},
		Position{X: 0, Y: 2},
	},
}
