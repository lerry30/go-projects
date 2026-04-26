package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"small-store/auth"
	"small-store/repository"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type AddressHandler struct {
	repo repository.AddressRepository
}

func NewAddressHandler(repo repository.AddressRepository) *AddressHandler {
	return &AddressHandler{repo: repo}
}

// ----------------------------------------------------------------------------

type Address struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type UserAddresses struct {
	Address
	ID        int32     `json:"id"`
	IsDefault bool      `json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ----------------------------------------------------------------------------

func (h *AddressHandler) CreateHandler(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var addr Address
	if err := json.NewDecoder(r.Body).Decode(&addr); err != nil {
		return fmt.Errorf("invalid address data")
	}

	addr.Line1 = capitalize(strings.TrimSpace(addr.Line1))
	addr.Line2 = capitalize(strings.TrimSpace(addr.Line2))
	addr.City = capitalize(strings.TrimSpace(addr.City))
	addr.State = capitalize(strings.TrimSpace(addr.State))
	addr.Zip = strings.ToUpper(strings.TrimSpace(addr.Zip)) // some countries have letters
	addr.Country = capitalize(strings.TrimSpace(addr.Country))

	if addr.Line1 == "" ||
		addr.City == "" ||
		addr.State == "" ||
		addr.Zip == "" ||
		addr.Country == "" {
		return fmt.Errorf("fields cannot be empty")
	}

	// TODO: validate country, zip, state

	userID, err := auth.GetCurrentUserID(r)
	if err != nil {
		return fmt.Errorf("invalid user credentials: \n %w", err)
	}

	newAddr, err := h.repo.Create(
		pgx.NamedArgs{
			"address_line1": addr.Line1,
			"address_line2": addr.Line2,
			"city":          addr.City,
			"region":        addr.State,
			"zip_code":      addr.Zip,
			"country_name":  addr.Country,
		},
		int32(userID),
	)
	if err != nil {
		fmt.Printf("%s\n Create new address error: %s", strings.Repeat("-", 20), err)
		return fmt.Errorf("Failed to add new address")
	}

	ResponseJSON(
		w,
		http.StatusCreated,
		UserAddresses{
			ID: newAddr.ID,
			Address: Address{
				Line1:   newAddr.Line1,
				Line2:   newAddr.Line2,
				City:    newAddr.City,
				State:   newAddr.Region,
				Zip:     fmt.Sprintf("%d", newAddr.ZIPCode),
				Country: newAddr.CountryName,
			},
			IsDefault: newAddr.IsDefault,
			CreatedAt: newAddr.CreatedAt,
			UpdatedAt: newAddr.UpdatedAt,
		})

	return nil
}

func (h *AddressHandler) GetAddressesHandler(w http.ResponseWriter, r *http.Request) error {
	queryParams := r.URL.Query()
	strLimit := queryParams.Get("limit")
	strOffset := queryParams.Get("offset")

	limit, err := strconv.ParseInt(strLimit, 10, 32)
	if err != nil {
		limit = 4
	}

	offset, err := strconv.ParseInt(strOffset, 10, 32)
	if err != nil {
		offset = 0
	}

	offset = (offset - 1) * limit

	userID, err := auth.GetCurrentUserID(r)
	if err != nil {
		fmt.Printf("%s\n Invalid User ID: %s", strings.Repeat("-", 20), err)
		return fmt.Errorf("Invalid credentials")
	}

	// usrAddrsRow []addressAggre
	usrAddrsRow, err := h.repo.GetAddresses(int32(userID), int32(limit), int32(offset))
	if err != nil {
		fmt.Printf("%s\n Failed to query user addresses: %s", strings.Repeat("-", 20), err)
		return fmt.Errorf("Can't get user addresses")
	}

	usrAddresses := make([]UserAddresses, 0, len(*usrAddrsRow))
	var totalCount float64

	for _, addr := range *usrAddrsRow {
		usrAddresses = append(usrAddresses, UserAddresses{
			ID: addr.ID,
			Address: Address{
				Line1:   addr.Line1,
				Line2:   addr.Line2,
				City:    addr.City,
				State:   addr.Region,
				Zip:     fmt.Sprintf("%d", addr.ZIPCode),
				Country: addr.CountryName,
			},
			IsDefault: addr.IsDefault,
			CreatedAt: addr.CreatedAt,
			UpdatedAt: addr.UpdatedAt,
		})
	}

	if len(*usrAddrsRow) > 0 {
		totalCount = math.Ceil(float64((*usrAddrsRow)[0].TotalCount) / float64(limit))
	}

	// reconstruct to fit the total count for pagination
	ResponseJSON(
		w,
		http.StatusOK,
		struct {
			Addresses  []UserAddresses `json:"addresses"`
			TotalCount float64         `json:"total_count"`
		}{
			Addresses:  usrAddresses,
			TotalCount: totalCount,
		},
	)

	return nil
}

func (h *AddressHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) error {
	urlParams := mux.Vars(r)
	strAddrId := urlParams["id"]

	addrId, err := strconv.ParseInt(strAddrId, 10, 32)
	if err != nil {
		return fmt.Errorf("Invalid address data")
	}

	id, err := h.repo.Delete(int32(addrId))
	if id == 0 || err != nil {
		fmt.Printf("%s\n Failed to delete address: %s", strings.Repeat("-", 20), err)
		return fmt.Errorf("Failed to delete address")
	}

	ResponseJSON(w, http.StatusOK, struct{ message string }{message: "Address successfully removed"})

	return nil
}

func (h *AddressHandler) SetDefaultHandler(w http.ResponseWriter, r *http.Request) error {
	urlParams := mux.Vars(r)
	strAddrId := urlParams["id"]

	addrID, err := strconv.ParseInt(strAddrId, 10, 32)
	if err != nil {
		return fmt.Errorf("Invalid address data")
	}

	userID, err := auth.GetCurrentUserID(r)
	if err != nil {
		fmt.Printf("%s\n Invalid User ID: %s", strings.Repeat("-", 20), err)
		return fmt.Errorf("Invalid credentials")
	}

	_, err = h.repo.SetDefault(int32(addrID), int32(userID))
	if err != nil {
		fmt.Printf("%s\n Set default address: %s", strings.Repeat("-", 20), err)
		return fmt.Errorf("Failed to set address as default")
	}

	ResponseJSON(w, http.StatusOK, struct{ message string }{message: "Address successfully set as default"})

	return nil
}
