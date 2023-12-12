package game

import "errors"

var uniqId = 0

type Net struct {
	Id   int
	Grid [][]string
}

func NewNet() (*Net, error) {
	grid := make([][]string, 3)
	for i := range grid {
		grid[i] = make([]string, 3)
	}

	uniqId++
	return &Net{
		Id:   uniqId,
		Grid: grid,
	}, nil
}

func (n *Net) Set(x, y int, isFirst bool) (bool, error) {
	if err := isValid(x, y); err != nil {
		return false, err
	}
	if n.Grid[x][y] != "" && n.Grid[x][y] != " " {
		return false, errors.New("Cell is already occupied")
	}
	if isFirst {
		n.Grid[x][y] = "X"
	} else {
		n.Grid[x][y] = "O"
	}
	return n.isWinner(), nil
}

func isValid(x, y int) error {
	if x >= 0 && x < 3 && y >= 0 && y < 3 {
		return nil
	}
	return errors.New("Position is not inside the grid")
}

func (n *Net) isWinner() bool {
	for i := 0; i < 3; i++ {
		if n.checkStreak(n.Grid[i][0], n.Grid[i][1], n.Grid[i][2]) {
			return true
		}
	}

	for j := 0; j < 3; j++ {
		if n.checkStreak(n.Grid[0][j], n.Grid[1][j], n.Grid[2][j]) {
			return true
		}
	}

	if n.checkStreak(n.Grid[0][0], n.Grid[1][1], n.Grid[2][2]) ||
		n.checkStreak(n.Grid[0][2], n.Grid[1][1], n.Grid[2][0]) {
		return true
	}

	return false
}

func (n *Net) checkStreak(a, b, c string) bool {
	return a != "" && a == b && b == c
}
