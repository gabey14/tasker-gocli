package taskercli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// name of the todo file
const TodosJSONFile = ".todos.json"

// FileStore is the struct for the file store
type FileStore struct {
	Loaded bool
}

// NewFileStore is creating a new file store
func NewFileStore() *FileStore {
	return &FileStore{Loaded: false}

}

// Initialize is initializing a new .toos.json file
func (f *FileStore) Initialize() {
	if f.LocalTodosFileExists() {
		fmt.Println("Local .todos.json file already exists! Doing nothing.")
		os.Exit(0)
	}
	if err := ioutil.WriteFile(TodosJSONFile, []byte("[]"), 0644); err != nil {
		fmt.Println("Error writing json file", err)
		os.Exit(1)
	}
}

// Returns true if the .todos.json file exists in the current dir
func (f *FileStore) LocalTodosFileExists() bool {
	dir, _ := os.Getwd()
	localrepo := filepath.Join(dir, TodosJSONFile)
	_, err := os.Stat(localrepo)
	return err == nil
}

// Load is loading the todos from the .todos.json file either from the current dir or from the home dir
func (f *FileStore) Load() ([]*Todo, error) {
	data, err := ioutil.ReadFile(f.GetLocation())
	if err != nil {
		fmt.Println("No todo file found!")
		fmt.Println("Please run 'tasker init' to create a new todo file.")
		os.Exit(0)
		return nil, err
	}

	var todos []*Todo

	if err := json.Unmarshal(data, &todos); err != nil {
		fmt.Println("Error reading json data", err)
		os.Exit(1)
		return nil, err
	}
	f.Loaded = true

	return todos, nil
}

// Save is saving the todos to the .todos.json file
func (f *FileStore) Save(todos []*Todo) {
	// make sure UUID is set for todos at save time
	for _, todo := range todos {
		if todo.UUID == "" {
			todo.UUID = newUUID()
		}
	}

	data, _ := json.Marshal(todos)
	if err := ioutil.WriteFile(f.GetLocation(), []byte(data), 0644); err != nil {
		fmt.Println("Error writing json file", err)
	}
}

// GetLocation is getting the location of the .todos.json file
func (f *FileStore) GetLocation() string {
	if f.LocalTodosFileExists() {
		dir, _ := os.Getwd()
		localrepo := filepath.Join(dir, TodosJSONFile)
		return localrepo
	}
	return fmt.Sprintf("%s/%s", UserHomeDir(), TodosJSONFile)
}
