package routes

import (
	"fmt"
	"net/http"
	"encoding/json"

	"backend/config"
	"backend/models"
)

func NewTodo(todos *[]models.Todo, w http.ResponseWriter, r *http.Request) (models.Todo, error) {
	config.EnableCORS(&w)

	defer r.Body.Close()

	if r.Method == "OPTIONS" {
		return models.Todo{}, fmt.Errorf("preflight")
	}

	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Method not allowed."})
		return models.Todo{}, fmt.Errorf("Method not allowed.")
	}

	var newTask models.Todo

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid json."})
		return models.Todo{}, fmt.Errorf("Invalid json.")
	}

	if newTask.Task == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Task is empty."})
		return models.Todo{}, fmt.Errorf("Task is empty.")
	}

	id := 1
	if len(*todos) > 0 {
		id = (*todos)[len(*todos)-1].ID
	}
	newTask.ID = id
	newTask.Completed = false

	// create response
	resp := models.Response{
		Message: "Todo added successfully.",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	return newTask, nil
}

func GetTodos(todos *[]models.Todo, w http.ResponseWriter, r *http.Request) error {
	config.EnableCORS(&w)

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
		return nil
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return fmt.Errorf("Method not allowed")
}

func SetTodoStatus(todos *[]models.Todo, w http.ResponseWriter, r *http.Request) error {
	config.EnableCORS(&w)

	defer r.Body.Close()

	if r.Method == "OPTIONS" {
		return fmt.Errorf("preflight")
	}

	if r.Method != "PUT" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Method not allowed."})
		return fmt.Errorf("Method not allowed.")
	}

	var updateTask models.Todo

	err := json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid JSON."})
		return fmt.Errorf("Invalid JSON.")
	}

	if updateTask.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid JSON fields."})
		return fmt.Errorf("Invalid JSON fields.")
	}

	doesExists := false
	
	for i, todo := range *todos {
		if todo.ID == updateTask.ID {
			(*todos)[i].Completed = updateTask.Completed
			doesExists = true
			break
		}
	}

	if !doesExists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Todo not found."})
		return fmt.Errorf("Todo not found.")
	}

	resp := models.Response{
		Message: "Todo status updated.",
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	return nil
}

func RemoveTodo(todos *[]models.Todo, w http.ResponseWriter, r *http.Request) error {
	config.EnableCORS(&w)

	defer r.Body.Close()

	if r.Method == "OPTIONS" {
		return fmt.Errorf("preflight")
	}

	if r.Method != "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Method not allowed."})
		return fmt.Errorf("Method not allowed")
	}

	var removeTask models.Todo

	err := json.NewDecoder(r.Body).Decode(&removeTask)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid JSON."})
		return fmt.Errorf("Invalid JSON.")
	}

	if removeTask.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid JSON fields."})
		return fmt.Errorf("Invalid JSON fields.")
	}

	doesExists := false

	for i, todo := range *todos {
		if todo.ID == removeTask.ID {
			*todos = append((*todos)[:i], (*todos)[i+1:]...)
			doesExists = true
			break
		}
	}

	if !doesExists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Todo not found."})
		return fmt.Errorf("Todo not found.")
	}

	resp := models.Response{
		Message: "Todo successfully removed",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	return nil
}