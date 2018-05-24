package model

// Player represent position and name
type Player struct {
	Name string  `json:"name"`
	X    float64 `json:"x,omitempty"`
	Y    float64 `json:"y,omitempty"`
}
