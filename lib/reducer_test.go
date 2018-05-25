package lib_test

import (
	"encoding/json"
	"testing"

	"github.com/masaruz/engine-bomberman/constant"
	"github.com/masaruz/engine-bomberman/lib"
	"github.com/masaruz/engine-bomberman/model"
)

func TestGameStateReducer(t *testing.T) {
	name := "p1"
	gamestate := lib.InitGameState()
	action := &model.Action{Command: constant.SyncCommand, Payload: []byte{}}
	gamestate = lib.GameStateReducer(gamestate, action)
	if len(gamestate.Players) > 0 || gamestate.Players[name] != nil {
		t.Error()
	}
	action.Payload, _ = json.Marshal(model.PlayerData{
		Data: model.Player{Name: "p1"},
	})
	// Sync
	gamestate = lib.GameStateReducer(gamestate, action)
	if len(gamestate.Players) == 0 || gamestate.Players[name].Name != name {
		t.Error()
	}
	// Sync again
	gamestate = lib.GameStateReducer(gamestate, action)
	if len(gamestate.Players) != 1 || gamestate.Players[name].Name != name {
		t.Error()
	}
	// Remove player
	action.Command = constant.RemovePlayerCommand
	gamestate = lib.GameStateReducer(gamestate, action)
	if len(gamestate.Players) > 0 || gamestate.Players[name] != nil {
		t.Error()
	}
	// Remove player again
	action.Command = constant.RemovePlayerCommand
	gamestate = lib.GameStateReducer(gamestate, action)
	if len(gamestate.Players) > 0 || gamestate.Players[name] != nil {
		t.Error()
	}
	// Add player
	action.Command = constant.AddPlayerCommand
	gamestate = lib.GameStateReducer(gamestate, action)
	if len(gamestate.Players) != 1 || gamestate.Players[name].Name != name {
		t.Error()
	}
}
func TestMessageReducer(t *testing.T) {
	gamestate := lib.InitGameState()
	action := &model.Action{Command: constant.SyncCommand, Payload: []byte{}}
	// Sync with no data
	messege := lib.MessageReducer(gamestate, action)
	if string(messege) != `{"command":"system_sync","data":[]}` {
		t.Error()
	}
	action.Payload, _ = json.Marshal(model.PlayerData{
		Data: model.Player{
			Name: "p1",
			X:    1.0,
			Y:    1.0},
	})
	// Sync with data
	gamestate = lib.GameStateReducer(gamestate, action)
	messege = lib.MessageReducer(gamestate, action)
	if string(messege) != `{"command":"system_sync","data":[{"name":"p1","x":1,"y":1}]}` {
		t.Error()
	}
	action.Payload, _ = json.Marshal(model.PlayersData{
		Data: []model.Player{
			model.Player{
				Name: "p2",
				X:    2.0,
				Y:    2.0},
		},
	})
	// Try sync again
	gamestate = lib.GameStateReducer(gamestate, action)
	messege = lib.MessageReducer(gamestate, action)
	if string(messege) != `{"command":"system_sync","data":[{"name":"p1","x":1,"y":1}]}` {
		t.Error()
	}
	action.Command = constant.AddPlayerCommand
	action.Payload, _ = json.Marshal(model.PlayerData{
		Data: model.Player{
			Name: "p2",
			X:    2.1,
			Y:    2.1},
	})
	// Add data
	gamestate = lib.GameStateReducer(gamestate, action)
	messege = lib.MessageReducer(gamestate, action)
	if string(messege) != `{"command":"player_add","data":{"name":"p2","x":2.1,"y":2.1}}` {
		t.Error(string(messege))
	}
	action.Command = constant.SyncCommand
	action.Payload, _ = json.Marshal(model.PlayerData{})
	// Sync after added data
	messege = lib.MessageReducer(gamestate, action)
	if string(messege) != `{"command":"system_sync","data":[{"name":"p1","x":1,"y":1},{"name":"p2","x":2.1,"y":2.1}]}` {
		t.Error()
	}
	action.Command = constant.RemovePlayerCommand
	action.Payload, _ = json.Marshal(model.PlayerData{
		Data: model.Player{
			Name: "p2"},
	})
	// Remove data
	gamestate = lib.GameStateReducer(gamestate, action)
	messege = lib.MessageReducer(gamestate, action)
	if string(messege) != `{"command":"player_remove","data":{"name":"p2"}}` {
		t.Error(string(messege))
	}
	action.Command = constant.SyncCommand
	action.Payload, _ = json.Marshal(model.PlayerData{})
	// Sync after removed data
	messege = lib.MessageReducer(gamestate, action)
	if string(messege) != `{"command":"system_sync","data":[{"name":"p1","x":1,"y":1}]}` {
		t.Error()
	}
}
