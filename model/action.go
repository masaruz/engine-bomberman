package model

import "github.com/masaruz/engine-bomberman/constant"

// Action of this server in each state
type Action struct {
	Command constant.Command `json:"command"`
	Payload []byte
}
