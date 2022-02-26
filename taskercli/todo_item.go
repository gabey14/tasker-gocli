package taskercli

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
