package models

type Todo struct {
	Task string `json:"task"`
	Completed bool `json:"completed"`
}