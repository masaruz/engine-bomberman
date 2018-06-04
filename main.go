package main

import (
	"encoding/json"

	"github.com/masaruz/engine-bomberman/lib"
	"github.com/masaruz/engine-bomberman/model"

	"github.com/masaruz/engine-lib/common"
	"github.com/masaruz/engine-lib/core"
)

type (
	ack  = core.Ack
	game struct{}
)

var (
	_message   []byte
	_gamestate model.GameState
)

func (g *game) Init() error {
	_gamestate = lib.InitGameState()
	common.Print("Bomberman Initiated")
	return nil
}

func (g *game) Start() error {
	common.Print("Bomberman Started")
	return nil
}

func (g *game) Update(msg []byte, callback ack) error {
	action := &model.Action{
		Payload: msg,
	}
	json.Unmarshal(msg, action)
	// In case command is just sync the state
	_gamestate = lib.GameStateReducer(_gamestate, action)
	// _message = lib.MessageReducer(_gamestate, action)
	_message = msg
	common.Printf("Bomberman Updated With Message %s\n", string(_message))
	callback("acknowledge")
	return nil
}

func (g *game) GetState() []byte {
	common.Print("Bomberman Get State")
	return make([]byte, 2000)
}

func main() {}

// Game exported to engine
var Game game
