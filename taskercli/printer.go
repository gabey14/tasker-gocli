package taskercli

import "regexp"

var (
	projectRegex, _ = regexp.Compile(`\+[a-zA-z\d_-]+`)
	contextRegex, _ = regexp.Compile(`\@[a-zA-z\d_-]+`)
)

// Printer is an interface for printing the output of grouped todos
type Printer interface {
	// TASK - Work on GroupedTodos
	Print(*GroupedTodos, bool, bool)
}
