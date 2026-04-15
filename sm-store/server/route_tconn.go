package main

import (
	"fmt"
	"net/http"

	"small-store/auth"
)

func (api *APIServer) TestHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("got reach")

	//row, err := api.db.Create(
	//	"test",
	//	[]string{"first_name", "last_name", "email", "phone"},
	//	"lerry", "samson", "lerry@email.com", "09631595556",
	//)

	// ------------

	//row, err := api.db.Update(
	//	"test",
	//	[]string{"email"},
	//	"johnlerrysamson@gmail.com",
	//)

	//if err != nil {
	//	return err
	//}

	//fmt.Println(row)

	// -------------

	//row, err := api.db.QueryRow("test", "first_name", "lerry")
	//if err != nil {
	//	return err
	//}
	//fmt.Println(row)

	// -------------

	//id, err := api.db.Delete("test", "first_name", "lerry")
	//if err != nil {
	//	return err
	//}
	//fmt.Println("row deleted", id)

	// --------------

	//rows, err := api.db.QueryAll("test")
	//if err != nil {
	//	return err
	//}
	//fmt.Println(rows)

	// -----------jwt----------

	token, err := auth.GenerateToken("user-123", "lerry@mail.com")
	if err != nil {
		return fmt.Errorf("could not generate token")
	}

	ResponseJSON(w, http.StatusOK, struct {
		Token   string `json:"token"`
		Message string `json:"message"`
	}{Message: "Test okay", Token: token})

	return nil
}
