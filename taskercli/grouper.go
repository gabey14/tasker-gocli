package taskercli

// Grouper is the group struct
type Grouper struct{}

// GroupedTodos is the main struct storing the grouped todos
type GroupedTodos struct {
	Groups map[string][]*Todo
}

// TASK - Work on GroupedTodos
