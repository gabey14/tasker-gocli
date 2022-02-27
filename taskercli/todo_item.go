package taskercli

import (
	"reflect"
	"time"
)

// iso8601TimestampFormat is the timestamp format to include date, time with timezone support. Easy to parse.
const iso8601TimestampFormat = "2006-01-02T15:04:05Z07:00"

// Todo is the struct for a todo item

type Todo struct {
	ID                int      `json:"id"`
	UUID              string   `json:"uuid"`
	Subject           string   `json:"subject"`
	Projects          []string `json:projects`
	Contexts          []string `json:"contexts"`
	Due               string   `json:"due"`
	Completed         bool     `json:"completed"`
	CompletedDate     string   `json:"completed_date"`
	Status            string   `json:"status"`
	Archived          bool     `json:"archived"`
	IsPriority        bool     `json:"is_priority"`
	Notes             []string `json:"notes"`
	Recur             string   `json:"recur"`
	RecurUntil        string   `json:"recur_until"`
	PrevRecurTodoUUID string   `json:"prev_recur_todo_uuid"`
}

// NewTodo is creating a new todo item
func NewTodo() *Todo {
	return &Todo{UUID: newUUID(), Completed: false, Archived: false, IsPriority: false}
}

// Valid checks if a new todo is valid or not
func (t Todo) Valid() bool {
	return t.Subject != ""
}

// CalculateDueDate calculates the due date for a todo item
func (t Todo) CalculateDueTime() time.Time {
	if t.Due != "" {
		// comma, ok sytax
		parsedTime, _ := time.Parse(DATE_FORMAT, t.Due)
		return parsedTime
	}
	parsedTime, _ := time.Parse(DATE_FORMAT, "1900-01-01")
	return parsedTime
}

// Complete marks a todo item as completed with the current time
func (t *Todo) Complete() {
	t.Completed = true
	t.Status = "completed"
	t.CompletedDate = timestamp(time.Now()).Format(iso8601TimestampFormat)
}

func (t *Todo) Uncomplete() {
	t.Completed = false
	t.Status = "active"
	t.CompletedDate = ""
}

// Archive archives a todo item
func (t *Todo) Archive() {
	t.Archived = true
}

// Unarchive unarchives a todo item
func (t *Todo) Unarchive() {
	t.Archived = false
}

// Priority prioritizes a todo item
func (t *Todo) Prioritize() {
	t.IsPriority = true
}

// Unpriority unprioritizes a todo item
func (t *Todo) Unprioritize() {
	t.IsPriority = false
}

// CompletedDate returns the completed date of a todo item
func (t Todo) CompletedDateToDate() string {
	parsedTime, _ := time.Parse(iso8601TimestampFormat, t.CompletedDate)
	return parsedTime.Format(DATE_FORMAT)
}

// HasNotes returns true if a todo item has notes
func (t Todo) HasNotes() bool {
	return len(t.Notes) > 0
}

// Equals returns true if two todo items are equal
func (t Todo) Equals(other *Todo) bool {
	if t.ID != other.ID ||
		t.UUID != other.UUID ||
		t.Subject != other.Subject ||
		!reflect.DeepEqual(t.Projects, other.Projects) ||
		!reflect.DeepEqual(t.Contexts, other.Contexts) ||
		t.Due != other.Due ||
		t.Completed != other.Completed ||
		t.Status != other.Status ||
		t.CompletedDate != other.CompletedDate ||
		t.Archived != other.Archived ||
		t.IsPriority != other.IsPriority ||
		t.Recur != other.Recur ||
		t.RecurUntil != other.RecurUntil ||
		!reflect.DeepEqual(t.Notes, other.Notes) {
		return false
	}
	return true
}
