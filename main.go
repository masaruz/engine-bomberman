package main

import (
	"fmt"

	"github.com/masaruz/engine-lib/core"
)

type form = core.Form

type ack = core.Ack

type game struct{}

func (g game) On(name string, ack func(msg form)) {
	fmt.Printf("Event: %s\n", name)
	fmt.Println("Let's play Bomberman")
	ack(form{
		Header:    "application/json",
		Body:      "encrypted_string",
		Signature: "jwt_string",
	})
}

func (g game) Init() error {
	fmt.Println("Bomberman Inited")
	return nil
}

func (g game) Start() error {
	fmt.Println("Bomberman Started")
	return nil
}

func (g game) Update(msg []byte) error {
	fmt.Printf("Bomberman Updated With Message %s\n", string(msg))
	return nil
}

func (g game) Pause() error {
	fmt.Println("Bomberman Paused")
	return nil
}

func (g game) Finish() error {
	fmt.Println("Bomberman Finished")
	return nil
}

func (g game) Shutdown() error {
	fmt.Println("Bomberman Shutdown")
	return nil
}

func main() {}

// Game exported to engine
var Game game
