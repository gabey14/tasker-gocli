package taskercli

// MemoryStore is a store that stores the todo list in memory
type MemoryStore struct {
	Todos []*Todo
}

// NewMemoryStore creates a new MemoryStore that returns a pointer to a new MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

// Init initializes the MemoryStore
func (m *MemoryStore) Initialize() {}

// Load loads the todo list from the MemoryStore
func (m *MemoryStore) Load() ([]*Todo, error) {
	return m.Todos, nil
}

// Save saves the todo list to the MemoryStore
func (m *MemoryStore) Save(todos []*Todo) {
	m.Todos = todos
}

// Check if LocalTodoFileExists
func (m *MemoryStore) LocalTodoFileExists() bool {
	return false
}

// Get Location of the memory store
func (m *MemoryStore) GetLocation() string {
	return ""
}
