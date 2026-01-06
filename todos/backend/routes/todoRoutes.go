package routes

import (
	"fmt"
	"net/http"
	"encoding/json"

	"backend/config"
	"backend/models"
)

func NewTodo(w http.ResponseWriter, r *http.Request) (models.Todo, error) {
	config.EnableCORS(&w)

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