package main

import (
	"fmt"

	"github.com/masaruz/engine-lib/core"
)

type ack = core.Ack

type game struct{}

func (g *game) Init() error {
	fmt.Println("Bomberman Initiated")
	return nil
}

func (g *game) Start() error {
	fmt.Println("Bomberman Started")
	return nil
}

func (g *game) Update(msg []byte, callback ack) error {
	fmt.Printf("Bomberman Updated With Message %s\n", string(msg))
	callback("acknowledge")
	return nil
}

func (g *game) GetState() ([]byte, error) {
	fmt.Println("Bomberman Get State")
	return []byte{}, nil
}

func main() {}

// Game exported to engine
var Game game
