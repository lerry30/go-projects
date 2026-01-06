package main

import (
	"fmt"
	"net/http"
	"log"

	"backend/objects"
)

func main() {
	var todoList objects.TodoList
	todoList.Load()

	http.HandleFunc("/newtask", todoList.NewTodo)
	http.HandleFunc("/todos", todoList.GetTodos)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}