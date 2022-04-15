package taskercli

import (
	"sort"
	"time"
)

// TodoList is the struct for a todo list with multiple todos.
type TodoList struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
	// isSynced bool
	Data []*Todo `json:"todo_items_attributes"`
}

// Load is loading the todo list
func (t *TodoList) Load(todos []*Todo) {
	t.Data = todos
}

// Add is adding a single todo to a todo list
func (t *TodoList) Add(todo *Todo) {
	todo.ID = t.NextID()
	t.Data = append(t.Data, todo)
}

// Delete is a variadic function deleting multiple todos from a todo list by their ids
func (t *TodoList) Delete(ids ...int) {
	for _, id := range ids {
		todo := t.FindByID(id)
		if todo == nil {
			continue
		}
		i := -1
		for index, todo := range t.Data {
			if todo.ID == id {
				i = index
			}
		}

		t.Data = append(t.Data[:i], t.Data[i+1:]...)
	}
}

// Complete is a variadic function completing multiple todos by their ids
func (t *TodoList) Complete(ids ...int) {
	for _, id := range ids {
		todo := t.FindByID(id)
		if todo == nil {
			continue
		}
		todo.Complete()
		t.Delete(id)
		t.Data = append(t.Data, todo)

		prevStatus := todo.Status
		r := &Recurrence{}
		if r.HasNextRecurringTodo(todo) {
			next := r.NextRecurringTodo(todo, time.Now())
			next.Status = prevStatus
			t.Add(next)
		}
	}
}

// Uncomplete is a variadic function uncompleting multiple todos by their ids
func (t *TodoList) Uncomplete(ids ...int) {
	for _, id := range ids {
		todo := t.FindByID(id)
		if todo == nil {
			continue
		}
		todo.Uncomplete()
		t.Delete(id)
		t.Data = append(t.Data, todo)
	}
}

// Archive is archiving multiple todos by their ids from todo list
func (t *TodoList) Archive(ids ...int) {
	for _, id := range ids {
		todo := t.FindByID(id)
		if todo == nil {
			continue
		}
		todo.Archive()
		t.Delete(id)
		t.Data = append(t.Data, todo)
	}
}

// Unarchive is unarchiving multiple todos by their ids from todo list
func (t *TodoList) Unarchive(ids ...int) {
	for _, id := range ids {
		todo := t.FindByID(id)
		if todo == nil {
			continue
		}
		todo.Unarchive()
		t.Delete(id)
		t.Data = append(t.Data, todo)
	}
}

// Prioritize is prioritizing multiple todos by their ids from todo list
func (t *TodoList) Prioritize(ids ...int) {
	for _, id := range ids {
		todo := t.FindByID(id)
		if todo == nil {
			continue
		}
		todo.Prioritize()
		t.Delete(id)
		t.Data = append(t.Data, todo)
	}
}

// Unprioritize is unprioritizing multiple todos by their ids from todo list
func (t *TodoList) Unprioritize(ids ...int) {
	for _, id := range ids {
		todo := t.FindByID(id)
		if todo == nil {
			continue
		}
		todo.Unprioritize()
		t.Delete(id)
		t.Data = append(t.Data, todo)
	}
}

// SetStatus is setting the status of multiple todos by their ids from todo list along with the status
func (t *TodoList) SetStatus(status string, ids ...int) {
	for _, id := range ids {
		todo := t.FindByID(id)
		if todo == nil {
			continue
		}
		todo.Status = status
		t.Delete(id)
		t.Data = append(t.Data, todo)
	}
}

// IndexOf returns the index of a todo item
func (t *TodoList) IndexOf(todoToFind *Todo) int {
	for i, todo := range t.Data {
		if todo.ID == todoToFind.ID {
			return i
		}
	}
	return -1
}

// ByDate is the date struct of todo for sorting todo items by date
type ByDate []*Todo

func (a ByDate) Len() int {
	return len(a)
}

func (a ByDate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDate) Less(i, j int) bool {
	t1Due := a[i].CalculateDueTime()
	t2Due := a[j].CalculateDueTime()
	return t1Due.Before(t2Due)
}

// Sorted list of Todos
func (t *TodoList) Todos() []*Todo {
	sort.Sort(ByDate(t.Data))
	return t.Data
}

// MaxID returns the max id for a todo item
func (t *TodoList) MaxID() int {
	var maxID int = 0
	for _, todo := range t.Data {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID
}

// NextID returns the next id for a todo item
func (t *TodoList) NextID() int {
	var found bool
	maxId := t.MaxID()

	for i := 1; i <= maxId; i++ {
		found = false
		for _, todo := range t.Data {
			if todo.ID == i {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return maxId + 1
}

// FindByID finds a todo by ID.
func (t *TodoList) FindByID(id int) *Todo {
	for _, todo := range t.Data {
		if todo.ID == id {
			return todo
		}
	}
	return nil
}

// GarbageCollector is a function that cleans up the todo list by removing archived todos
func (t *TodoList) GarbageCollector() {
	var toDelete []*Todo
	// add all todos that are archived to the toDelete slice
	for _, todo := range t.Data {
		if todo.Archived {
			toDelete = append(toDelete, todo)
		}
	}
	// delete all todos that are archived
	for _, todo := range toDelete {
		t.Delete(todo.ID)
	}
}
