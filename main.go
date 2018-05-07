package main

import (
	"fmt"

	"github.com/masaruz/engine-lib"
)

type form = struct {
	Header    string
	Body      string
	Signature string
}

type ack = func(msg form)

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

func (g game) Start() *lib.Error {
	fmt.Println("Bomberman Started")
	return nil
}

func (g game) Pause() *lib.Error {
	fmt.Println("Bomberman Paused")
	return nil
}

func (g game) Finish() *lib.Error {
	fmt.Println("Bomberman Finished")
	return nil
}

func main() {}

// Game exported to engine
var Game game
