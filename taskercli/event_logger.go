package taskercli

// the different types of events that can occur
const (
	AddEvent    = "EventAdded"
	UpdateEvent = "EventUpdated"
	DeleteEvent = "EventDeleted"
)

// EventLogger is the main struct of this file
type EventLogger struct {
	PreviousTodoList  *TodoList
	CurrentTodoList   *TodoList
	Store             Store
	SyncedLists       []*SyncedList
	CurrentSyncedList *SyncedList
	Events            []*EventLog
}

// Synced
type SyncedList struct {
	Filename string      `json:"filename"`
	UUID     string      `json:"uuid"`
	Name     string      `json:"name"`
	Events   []*EventLog `json:"events"`
}

// EventLog is a log of events that occurred, with the todo data.
type EventLog struct {
	EventType    string `json:"event_type"`
	ObjectType   string `json:"object_type"`
	TodoListUUID string `json:"todo_list_uuid"`
	Object       *Todo  `json:"object"`
}
