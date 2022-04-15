package taskercli

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// current version and date format of taskercli
const (
	VERSION     string = "1.0"
	DATE_FORMAT string = "2006-01-02"
)

// App is giving the structure of the taskercli
type App struct {
	EventLogger *EventLogger
	TodoStore   Store
	Printer     Printer
	TodoList    *TodoList
}

// NewApp is creating a new tasker app.
func NewApp() *App {
	app := &App{
		TodoList:  &TodoList{},
		Printer:   NewScreenPrinter(true),
		TodoStore: NewFileStore(),
	}
	return app
}

// NewAppWithPrintOptions creates a new app with options for printing on screen.
func NewAppWithPrintOptions(unicodeSupport bool, colorSupport bool) *App {
	var printer Printer
	if colorSupport {
		printer = NewScreenPrinter(unicodeSupport)
	} else {
		printer = NewSimpleScreenPrinter(unicodeSupport)
	}

	app := &App{
		TodoList:  &TodoList{},
		Printer:   printer,
		TodoStore: NewFileStore(),
	}
	return app
}

// Initialize is initializing the taskercli repo
func (a *App) InitializeRepo() {
	a.TodoStore.Initialize()
	fmt.Println("TaskerCLI initialized!")
}

// AddTodo adds a new todo to the todo list
func (a *App) AddTodo(input string) {
	a.load()
	parser := &InputParser{}

	filter, err := parser.Parse(input)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("I need more information. Try something like 'a chat with @jim due tom'")
		return
	}

	todoItem, err := CreateTodo(filter)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	a.TodoList.Add(todoItem)
	a.save()
	fmt.Printf("Todo %d added.\n", todoItem.ID)
}

// load is loading the todo list from the todo store
func (a *App) load() error {
	todos, err := a.TodoStore.Load()
	if err != nil {
		return err
	}
	a.TodoList.Load(todos)
	a.EventLogger = NewEventLogger(a.TodoList, a.TodoStore)
	return nil
}

// DeleteTodo deletes a todo from the todo list
func (a *App) DeleteTodo(input string) {
	a.load()
	ids := a.getIDs(input)
	if len(ids) == 0 {
		return
	}
	a.TodoList.Delete(ids...)
	a.save()
	fmt.Printf("%s deleted.\n", pluralize(len(ids), "Todo", "Todos"))
}

// CompleteTodo completes a todo
func (a *App) CompleteTodo(input string, archive bool) {
	a.load()
	ids := a.getIDs(input)
	if len(ids) == 0 {
		return
	}
	a.TodoList.Complete(ids...)
	if archive {
		a.TodoList.Archive(ids...)
	}
	a.save()
	fmt.Println("Todo completed.")
}

// UncompleteTodo uncompletes a todo.
func (a *App) UncompleteTodo(input string) {
	a.load()
	ids := a.getIDs(input)
	if len(ids) == 0 {
		return
	}
	a.TodoList.Uncomplete(ids...)
	a.save()
	fmt.Println("Todo uncompleted.")
}

// ArchiveTodo archives a todo.
func (a *App) ArchiveTodo(input string) {
	a.load()
	ids := a.getIDs(input)
	if len(ids) == 0 {
		return
	}
	a.TodoList.Archive(ids...)
	a.save()
	fmt.Println("Todo archived.")
}

// UnarchiveTodo unarchives a todo.
func (a *App) UnarchiveTodo(input string) {
	a.load()
	ids := a.getIDs(input)
	if len(ids) == 0 {
		return
	}
	a.TodoList.Unarchive(ids...)
	a.save()
	fmt.Println("Todo unarchived.")
}

// EditTodo edits a todo with the given input.
func (a *App) EditTodo(todoID int, input string) {
	a.load()
	todo := a.TodoList.FindByID(todoID)
	if todo == nil {
		fmt.Println("No todo with that id.")
		return
	}

	parser := &InputParser{}
	filter, err := parser.Parse(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = EditTodo(todo, a.TodoList, filter); err != nil {
		fmt.Println(err.Error())
		return
	}

	a.save()
	fmt.Println("Todo updated.")
}

// AddNote adds a note to a todo.
func (a *App) AddNote(todoID int, note string) {
	a.load()

	todo := a.TodoList.FindByID(todoID)
	if todo == nil {
		fmt.Println("No todo with that id.")
		return
	}
	todo.Notes = append(todo.Notes, note)

	fmt.Println("Note added.")
	a.save()
}

// EditNote edits a todo's note.
func (a *App) EditNote(todoID int, noteID int, note string) {
	a.load()

	todo := a.TodoList.FindByID(todoID)
	if todo == nil {
		fmt.Println("No todo with that id.")
		return
	}

	if noteID >= len(todo.Notes) {
		fmt.Println("No note could be found with that ID.")
		return
	}

	todo.Notes[noteID] = note

	fmt.Println("Note edited.")
	a.save()
}

// DeleteNote deletes a note from a todo.
func (a *App) DeleteNote(todoID int, noteID int) {
	a.load()

	todo := a.TodoList.FindByID(todoID)
	if todo == nil {
		fmt.Println("No todo with that id.")
		return
	}

	if noteID >= len(todo.Notes) {
		fmt.Println("No note could be found with that ID.")
		return
	}

	todo.Notes = append(todo.Notes[:noteID], todo.Notes[noteID+1:]...)

	fmt.Println("Note deleted.")
	a.save()
}

// ArchiveCompleted will archive all completed todos.
func (a *App) ArchiveCompleted() {
	a.load()
	for _, todo := range a.TodoList.Todos() {
		if todo.Completed {
			todo.Archive()
		}
	}
	a.save()
	fmt.Println("All completed todos have been archived.")
}

// ListTodos will list all todos.
func (a *App) ListTodos(input string, showNotes bool, showStatus bool) {
	a.load()

	parser := &InputParser{}

	filter, err := parser.Parse(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	todoFilter := &TodoFilter{Todos: a.TodoList.Todos(), Filter: filter}
	grouped := a.getGroups(input, todoFilter.ApplyFilter())

	a.Printer.Print(grouped, showNotes, showStatus)
}

// PrioritizeTodo will prioritize a todo.
func (a *App) PrioritizeTodo(input string) {
	a.load()
	ids := a.getIDs(input)
	if len(ids) == 0 {
		return
	}
	a.TodoList.Prioritize(ids...)
	a.save()
	fmt.Println("Todo prioritized.")
}

// UnprioritizeTodo unprioritizes a todo.
func (a *App) UnprioritizeTodo(input string) {
	a.load()
	ids := a.getIDs(input)
	if len(ids) == 0 {
		return
	}
	a.TodoList.Unprioritize(ids...)
	a.save()
	fmt.Println("Todo un-prioritized.")
}

// StartTodo will start a todo.
func (a *App) SetTodoStatus(input string) {
	a.load()
	ids := a.getIDs(input)
	if len(ids) == 0 {
		return
	}

	splitted := strings.Split(input, " ")

	a.TodoList.SetStatus(splitted[len(splitted)-1], ids...)
	a.save()
	fmt.Println("Todo status updated.")
}

// GarbageCollect will delete all archived todos.
func (a *App) GarbageCollect() {
	a.load()
	a.TodoList.GarbageCollector()
	a.save()
	fmt.Println("Garbage collection complete.")
}

// Save the todolist to the todo store
func (a *App) save() {
	a.TodoStore.Save(a.TodoList.Data)
}

func (a *App) getID(input string) (int, error) {
	splitted := strings.Split(input, " ")
	id, err := strconv.Atoi(splitted[0])
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Invalid id: '%s'", splitted[0]))
	}
	return id, nil
}

func (a *App) getIDs(input string) (ids []int) {
	idGroups := strings.Split(input, ",")
	for _, idGroup := range idGroups {
		if rangedIds, err := a.parseRangedIds(idGroup); len(rangedIds) > 0 || err != nil {
			if err != nil {
				fmt.Printf("Invalid id group: %s.\n", input)
				continue
			}
			ids = append(ids, rangedIds...)
		} else if id, err := a.getID(idGroup); err == nil {
			ids = append(ids, id)
		} else {
			fmt.Printf("Invalid id: %s.\n", idGroup)
		}
	}
	return ids
}

func (a *App) parseRangedIds(input string) (ids []int, err error) {
	rangeNumberRE, _ := regexp.Compile(`(\d+)-(\d+)`)
	if matches := rangeNumberRE.FindStringSubmatch(input); len(matches) > 0 {
		lowerID, _ := strconv.Atoi(matches[1])
		upperID, _ := strconv.Atoi(matches[2])
		if lowerID >= upperID {
			return ids, fmt.Errorf("invalid id group: %s", input)
		}
		for id := lowerID; id <= upperID; id++ {
			ids = append(ids, id)
		}
	}
	return ids, err
}

func (a *App) getGroups(input string, todos []*Todo) *GroupedTodos {
	grouper := &Grouper{}
	contextRegex, _ := regexp.Compile(`group:c.*$`)
	projectRegex, _ := regexp.Compile(`group:p.*$`)
	statusRegex, _ := regexp.Compile(`group:s.*$`)

	var grouped *GroupedTodos

	if contextRegex.MatchString(input) {
		grouped = grouper.GroupByContext(todos)
	} else if projectRegex.MatchString(input) {
		grouped = grouper.GroupByProject(todos)
	} else if statusRegex.MatchString(input) {
		grouped = grouper.GroupByStatus(todos)
	} else {
		grouped = grouper.GroupByNothing(todos)
	}
	return grouped
}
