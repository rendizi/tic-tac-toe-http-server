package service

import (
	"github.com/rendizi/tic-tac-toe-http-server/pkg/game/tictac"
	"math/rand"
	"time"
)

func New() (*tictac.Net, error) {
	rand.NewSource(time.Now().UTC().UnixNano())

	currentNet, err := tictac.NewNet()
	if err != nil {
		return nil, err
	}

	return currentNet, nil
}

func (ls *LifeService) NewState() *life.World {
	life.NextState(ls.currentWorld, ls.nextWorld)

	ls.currentWorld = ls.nextWorld

	return ls.currentWorld
}
