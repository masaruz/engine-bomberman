package lib

import (
	"engine-bomberman/model"
)

// InitGameState return initialized gamestate
func InitGameState() model.GameState {
	return model.GameState{
		Players:    make(map[string]*model.Player),
		Spectators: make(map[string]*model.Player),
	}
}

// CloneGameState to brand new object
func CloneGameState(gamestate model.GameState) model.GameState {
	_gamestate := gamestate
	return _gamestate
}

// DefaultGameState init new gamestate or clone from existed
func DefaultGameState(gamestate model.GameState) model.GameState {
	var _gamestate model.GameState
	if gamestate.Players == nil || gamestate.Spectators == nil {
		_gamestate = InitGameState()
	} else {
		_gamestate = CloneGameState(gamestate)
	}
	return _gamestate
}
