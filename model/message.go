package model

import "engine-bomberman/constant"

// SyncMessage contain players' data to send to engine
type SyncMessage struct {
	Command constant.Command `json:"command"`
	Data    []Player         `json:"data"`
}

// PlayerData from engine
type PlayerData struct {
	Data Player `json:"data"`
}

// PlayersData from engine
type PlayersData struct {
	Data []Player `json:"data"`
}
