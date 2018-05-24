package model

// GameState represent current game situation
type GameState struct {
	Players    map[string]*Player
	Spectators map[string]*Player
}
