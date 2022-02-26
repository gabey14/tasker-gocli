package taskercli

// Store is the interface for the taskercli todos.
type Store interface {
	GetLocation() string
	LocalTodosFileExists() bool
	Initialize()
	Load() ([]*Todo, error)
	Save(todos []*Todo)
}
