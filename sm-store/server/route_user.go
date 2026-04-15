package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"strings"

	"small-store/auth"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	Credentials
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func (api *APIServer) SignUpHandler(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	if r.Method != "POST" {
		return fmt.Errorf("Method not allowed.")
	}

	var userInfo UserInfo
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		return fmt.Errorf("Invalid json.")
	}

	userInfo.FirstName = capitalize(strings.TrimSpace(userInfo.FirstName))
	userInfo.LastName = capitalize(strings.TrimSpace(userInfo.LastName))
	userInfo.Email = strings.TrimSpace(userInfo.Email)
	userInfo.Phone = strings.TrimSpace(userInfo.Phone)

	if userInfo.FirstName == "" ||
		userInfo.LastName == "" ||
		userInfo.Email == "" ||
		userInfo.Phone == "" ||
		userInfo.Password == "" {
		return fmt.Errorf("Fields cannot be empty")
	}

	if !isValidEmail(userInfo.Email) {
		return fmt.Errorf("%q is not a valid email address", userInfo.Email)
	}

	if !isValidPhone(userInfo.Phone) {
		return fmt.Errorf("%q is not a valid phone number", userInfo.Phone)
	}

	hash, err := HashPassword(userInfo.Password)
	if err != nil {
		return fmt.Errorf("securing user password: %w", err)
	}

	userId, err := api.db.Create(
		"site_user",
		[]string{"first_name", "last_name", "email_address", "phone_number", "password"},
		userInfo.FirstName, userInfo.LastName, userInfo.Email, userInfo.Phone, hash,
	)

	token, err := auth.GenerateToken(userId, userInfo.Email)
	if err != nil {
		return fmt.Errorf("could not generate token")
	}

	w.Header().Set("Authorization", "Bearer "+token)
	ResponseJSON(w, http.StatusCreated, struct{ Message string }{Message: "user account created"})

	return nil
}

func (api *APIServer) SignInHandler(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	if r.Method != "POST" {
		return fmt.Errorf("Method not allowed.")
	}

	var userCred Credentials
	if err := json.NewDecoder(r.Body).Decode(&userCred); err != nil {
		return fmt.Errorf("Invalid json.")
	}

	if !isValidEmail(userCred.Email) {
		return fmt.Errorf("%q is not a valid email address", userCred.Email)
	}

	row, err := api.db.QueryRow("site_user", "email_address", userCred.Email)
	if err != nil {
		return fmt.Errorf("User doesn't exists.")
	}

	hash, ok := row["password"].(string)
	if !ok {
		fmt.Println("Type assertion error: db password")
		return fmt.Errorf("Something went wrong, user doesn't exists")
	}

	if hash == "" {
		return fmt.Errorf("User doesn't exists")
	}

	var match bool = CheckPasswordHash(userCred.Password, hash)
	if !match {
		return fmt.Errorf("Invalid credentials")
	}

	//fmt.Printf("user: %T \n", row["id"])

	id, ok := row["id"].(int32)
	if !ok {
		fmt.Println("Type assertion error: db id")
		return fmt.Errorf("Something went wrong")
	}

	userId := strconv.FormatInt(int64(id), 10)

	token, err := auth.GenerateToken(userId, userCred.Email)
	if err != nil {
		return fmt.Errorf("could not generate token")
	}

	w.Header().Set("Authorization", "Bearer "+token)
	ResponseJSON(w, http.StatusOK, struct{ Message string }{Message: "successfully logged in"})

	return nil
}

func (api *APIServer) LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	user, _ := auth.GetCurrentUser(r)
	// Blacklist this token so it can't be reused
	auth.Store.Revoke(user.ID, user.ExpiresAt.Time)

	ResponseJSON(w, http.StatusOK, struct{ Message string }{Message: "logged out"})
	return nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func isValidPhone(phone string) bool {
	// Strip spaces, dashes, parentheses
	stripped := regexp.MustCompile(`[\s\-().+]`).ReplaceAllString(phone, "")
	// E.164-style: optional leading +, then 7–15 digits
	var phoneNumFormat = regexp.MustCompile(`^\+?[0-9]{7,15}$`)
	return phoneNumFormat.MatchString(stripped)
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	s = strings.ToLower(s)
	return strings.ToUpper(s[:1]) + s[1:]
}
