package taskercli

// Current version of taskercli
const (
	VERSION     string = "0.0.1"
	DATE_FORMAT string = "2022-01-05"
)

// App is giving the structure of the taskercli
type App struct {
	EventLogger *EventLogger
	TodoStore   Store
	Printer     Printer
	TodoList    *TodoList
}
