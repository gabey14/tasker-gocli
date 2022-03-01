package taskercli

import "regexp"

var (
	projectRegex, _ = regexp.Compile(`\+[\p{L}\d_]+`)
	contextRegex, _ = regexp.Compile(`\@[\p{L}\d_]+`)
)

// Printer is an interface for printing the output of grouped todos
type Printer interface {
	// TASK - Work on GroupedTodos
	Print(*GroupedTodos, bool, bool)
}
