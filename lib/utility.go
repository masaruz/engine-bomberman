package lib

import (
	"github.com/masaruz/engine-bomberman/model"
)

// Value convert map to array
func Value(m map[string]*model.Player) []model.Player {
	values := []model.Player{}
	for _, value := range m {
		values = append(values, *value)
	}
	return values
}
