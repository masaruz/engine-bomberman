package model

import "engine-bomberman/constant"

// Action of this server in each state
type Action struct {
	Command constant.Command `json:"command"`
	Payload []byte
}
