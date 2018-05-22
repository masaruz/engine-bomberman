package main

import (
	"fmt"

	"github.com/masaruz/engine-lib/core"
)

type ack = core.Ack

type game struct{}

var _msg []byte

func (g *game) Init() error {
	fmt.Println("Bomberman Initiated")
	return nil
}

func (g *game) Start() error {
	fmt.Println("Bomberman Started")
	return nil
}

func (g *game) Update(msg []byte, callback ack) error {
	_msg = msg
	fmt.Printf("Bomberman Updated With Message %s\n", string(_msg))
	callback("acknowledge")
	return nil
}

func (g *game) GetState() []byte {
	fmt.Println("Bomberman Get State")
	return _msg
}

func main() {}

// Game exported to engine
var Game game
