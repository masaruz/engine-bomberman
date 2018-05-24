package constant

// Command from engine
type Command string

// Export command string
const (
	SyncCommand         Command = "system_sync"
	AddPlayerCommand    Command = "player_add"
	MovePlayerCommand   Command = "player_move"
	RemovePlayerCommand Command = "player_remove"
)
