package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"

	"small-store/auth"
	"small-store/repository"

	"github.com/jackc/pgx/v5"
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

// ---------------------------------------

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

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

	// check if user already exists in the database
	user, err := h.repo.GetByEmail(userInfo.Email)
	if user != nil || err == nil {
		return fmt.Errorf("Invalid credentials")
	}

	hash, err := repository.HashPassword(userInfo.Password)
	if err != nil {
		return fmt.Errorf("securing user password: %w", err)
	}

	id, err := h.repo.Create(
		pgx.NamedArgs{
			"first_name":    userInfo.FirstName,
			"last_name":     userInfo.LastName,
			"email_address": userInfo.Email,
			"phone_number":  userInfo.Phone,
			"password":      hash,
		},
	)

	if err != nil {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Printf("Database error: %s", err)

		return fmt.Errorf("failed to create user account")
	}

	strId := fmt.Sprintf("%d", id)
	token, err := auth.GenerateToken(strId, userInfo.Email)
	if err != nil {
		return fmt.Errorf("could not generate token")
	}

	w.Header().Set("Authorization", "Bearer "+token)
	ResponseJSON(w, http.StatusCreated, struct{ Message string }{Message: "user account created"})

	return nil
}

func (h *UserHandler) SignInHandler(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var userCred Credentials
	if err := json.NewDecoder(r.Body).Decode(&userCred); err != nil {
		return fmt.Errorf("Invalid json.")
	}

	if !isValidEmail(userCred.Email) {
		return fmt.Errorf("%q is not a valid email address", userCred.Email)
	}

	user, err := h.repo.GetByEmail(userCred.Email)
	if err != nil {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Printf("db query error: %s\n", err)

		return fmt.Errorf("User doesn't exists.")
	}

	hash := user.Password
	if hash == "" {
		return fmt.Errorf("User doesn't exists")
	}

	var match bool = repository.CheckPasswordHash(userCred.Password, hash)
	if !match {
		return fmt.Errorf("Invalid credentials")
	}

	//fmt.Printf("user: %T \n", row["id"])

	userId := strconv.FormatInt(int64(user.ID), 10) // to string

	token, err := auth.GenerateToken(userId, userCred.Email)
	if err != nil {
		return fmt.Errorf("could not generate token")
	}

	w.Header().Set("Authorization", "Bearer "+token)
	ResponseJSON(w, http.StatusOK, struct{ Message string }{Message: "successfully logged in"})

	return nil
}

func (h *UserHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	user, _ := auth.GetCurrentUser(r)
	// Blacklist this token so it can't be reused
	auth.Store.Revoke(user.ID, user.ExpiresAt.Time)

	ResponseJSON(w, http.StatusOK, struct{ Message string }{Message: "logged out"})
	return nil
}

func (h *UserHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) error {
	user, ok := auth.GetCurrentUser(r)
	if !ok || user.UserID == "" {
		return fmt.Errorf("unknown user")
	}

	userRow, err := h.repo.GetByID(user.UserID)
	if err != nil {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Printf("db query error: %s\n", err)

		return fmt.Errorf("user doesn't exist")
	}

	ResponseJSON(
		w, http.StatusOK,
		struct {
			FirstName string    `json:"first_name"`
			LastName  string    `json:"last_name"`
			Email     string    `json:"email"`
			Phone     string    `json:"phone"`
			CreatedAt time.Time `json:"created_at"`
		}{
			FirstName: userRow.FirstName,
			LastName:  userRow.LastName,
			Email:     userRow.Email,
			Phone:     userRow.Phone,
			CreatedAt: userRow.CreatedAt,
		},
	)
	return nil
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		return fmt.Errorf("Invalid json")
	}

	userInfo.FirstName = capitalize(strings.TrimSpace(userInfo.FirstName))
	userInfo.LastName = capitalize(strings.TrimSpace(userInfo.LastName))
	userInfo.Email = strings.TrimSpace(userInfo.Email)
	userInfo.Phone = strings.TrimSpace(userInfo.Phone)

	if userInfo.FirstName == "" ||
		userInfo.LastName == "" ||
		userInfo.Email == "" ||
		userInfo.Phone == "" {
		return fmt.Errorf("Fields cannot be empty")
	}

	if !isValidEmail(userInfo.Email) {
		return fmt.Errorf("%q is not a valid email address", userInfo.Email)
	}

	if !isValidPhone(userInfo.Phone) {
		return fmt.Errorf("%q is not a valid phone number", userInfo.Phone)
	}

	currUser, ok := auth.GetCurrentUser(r)
	if !ok {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println("user request header error, unknown current user data")

		return fmt.Errorf("undefined user")
	}

	userRow, err := h.repo.GetByID(currUser.UserID)
	if err != nil {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Printf("db query error: %s\n", err)

		return fmt.Errorf("user doesn't exist")
	}

	if userInfo.FirstName == userRow.FirstName &&
		userInfo.LastName == userRow.LastName &&
		userInfo.Email == userRow.Email &&
		userInfo.Phone == userRow.Phone {
		ResponseJSON(w, http.StatusOK,
			struct {
				FirstName string    `json:"first_name"`
				LastName  string    `json:"last_name"`
				Email     string    `json:"email"`
				Phone     string    `json:"phone"`
				CreatedAt time.Time `json:"created_at"`
			}{
				FirstName: userRow.FirstName,
				LastName:  userRow.LastName,
				Email:     userRow.Email,
				Phone:     userRow.Phone,
				CreatedAt: userRow.CreatedAt,
			},
		)
		return nil
	}

	userRow, err = h.repo.Update(
		pgx.NamedArgs{
			"first_name":    userInfo.FirstName,
			"last_name":     userInfo.LastName,
			"email_address": userInfo.Email,
			"phone_number":  userInfo.Phone,
		},
		userRow.ID,
	)

	if err != nil {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Printf("db update error: %s\n", err)

		return fmt.Errorf("failed to update user info")
	}

	ResponseJSON(
		w, http.StatusOK,
		struct {
			FirstName string    `json:"first_name"`
			LastName  string    `json:"last_name"`
			Email     string    `json:"email"`
			Phone     string    `json:"phone"`
			CreatedAt time.Time `json:"created_at"`
		}{
			FirstName: userRow.FirstName,
			LastName:  userRow.LastName,
			Email:     userRow.Email,
			Phone:     userRow.Phone,
			CreatedAt: userRow.CreatedAt,
		},
	)

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
