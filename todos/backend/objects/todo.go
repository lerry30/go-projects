package objects

import (
	"fmt"
	"bytes"
	"os"
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/routes"
)

type TodoList struct {
	Todos []models.Todo `json:"todos"`
}

const filepath = "records/todos.json"

func (tl *TodoList) AddTodo(todo models.Todo) {
	tl.Todos = append(tl.Todos, todo)
}

func (tl *TodoList) Save() error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(tl); err != nil {
		return fmt.Errorf("failed to encode: %w", err)
	}

	if err := os.WriteFile(filepath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to save data: %w", err)
	}

	return nil
}

func (tl *TodoList) Load() error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read json file: %w", err)
	}

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(tl); err != nil {
		return fmt.Errorf("failed to decode data: %w", err)
	}

	return nil
}

func (tl *TodoList) NewTodo(w http.ResponseWriter, r *http.Request) {
	newTodo, err := routes.NewTodo(w, r)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	tl.AddTodo(newTodo)
	tl.Save()
}

func (tl *TodoList) GetTodos(w http.ResponseWriter, r *http.Request) {
	err := routes.GetTodos(&tl.Todos, w, r)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
}