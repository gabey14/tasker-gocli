package taskercli

// Printer is an interface for printing the output of grouped todos
type Printer interface {
	// TASK - Work on GroupedTodos
	Print(*GroupedTodos, bool, bool)
}
