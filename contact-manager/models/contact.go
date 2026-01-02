package models

import (
	"fmt"
	"bytes"
	"os"
	"encoding/json"
)

type Contact struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type ContactBook struct {
	Contacts []Contact `json:"contacts"`
}

const filepath = "records/contacts.json"

func (cb *ContactBook) AddContact(c Contact) {
	cb.Contacts = append(cb.Contacts, c)
}

func (cb *ContactBook) DisplayContacts() {
	for _, c := range cb.Contacts {
		fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s", 
			c.ID, c.Name, c.Email, c.Phone)
	}
}

func (cb *ContactBook) Save() error {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(cb); err != nil {
		return fmt.Errorf("failed to encode: %w", err)
	}

	if err := os.WriteFile(filepath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write contacts: %w", err)
	}

	return nil
}

func (cb *ContactBook) Load() error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read fille %s: %w", filepath, err)
	}

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(cb); err != nil {
		return fmt.Errorf("failed to decode: %w", err)
	}

	return nil
}

func (cb *ContactBook) GetLastInsertedId() int {
	length := len(cb.Contacts)
	if length == 0 {
		return 0
	}
	return cb.Contacts[length-1].ID
}