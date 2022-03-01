package taskercli

// MemoryPrinter is a printer that prints to the console
type MemoryPrinter struct {
	Groups *GroupedTodos
}

// Print prints the grouped todo list to the console
func (m *MemoryPrinter) Print(groupedTodos *GroupedTodos, printNotes bool) {
	m.Groups = groupedTodos
}
