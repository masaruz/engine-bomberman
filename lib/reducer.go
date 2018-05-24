package lib

import (
	"encoding/json"

	"github.com/masaruz/engine-bomberman/constant"
	"github.com/masaruz/engine-bomberman/model"
)

// GameStateReducer will reduce for command and produce new state
func GameStateReducer(gamestate model.GameState, action *model.Action) model.GameState {
	_gamestate := DefaultGameState(gamestate)
	player := &model.PlayerData{}
	json.Unmarshal(action.Payload, player)
	switch action.Command {
	case constant.RemovePlayerCommand:
		delete(_gamestate.Players, player.Data.Name)
	default:
		// If message is not the valid format
		// Do nothing
		if player.Data.Name != "" {
			_gamestate.Players[player.Data.Name] = &model.Player{
				Name: player.Data.Name,
				X:    player.Data.X,
				Y:    player.Data.Y,
			}
		}
	}
	return _gamestate
}

// MessageReducer return client's response message
func MessageReducer(gamestate model.GameState, action *model.Action) []byte {
	_message := action.Payload
	_gamestate := DefaultGameState(gamestate)
	player := &model.PlayerData{}
	json.Unmarshal(action.Payload, player)
	switch action.Command {
	case constant.SyncCommand:
		_message, _ = json.Marshal(model.SyncMessage{
			Command: constant.SyncCommand,
			Data:    Value(_gamestate.Players),
		})
	default:
		// If message is not the valid format
		// Do nothing
		if player.Data.Name != "" {
			_message, _ = json.Marshal(struct {
				Command constant.Command `json:"command"`
				Data    model.Player     `json:"data"`
			}{
				Command: action.Command,
				Data:    player.Data,
			})
		}
	}
	return _message
}
