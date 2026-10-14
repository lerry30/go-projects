package tracker

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"time"

	"tracker/internal/shared/httpresponse"
	"tracker/internal/shared/account"
	"tracker/internal/platform/logger"
)

type TrackerHttpHandler struct {
	trackerRepo TrackerRepository
}

func NewTrackerHttpHandler(trackerRepo TrackerRepository) *TrackerHttpHandler {
	return &TrackerHttpHandler{
		trackerRepo: trackerRepo,
	}
}

func (t *TrackerHttpHandler) TrackerHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	var resErr httpresponse.ErrorResponse

	userID, err := account.CurrentUserID(r)
	if err != nil {
		logger.LogError("user not found", err)
		resErr.WriteMessage(http.StatusNotFound, "user not found")
		return resErr
	}

	dbTracker, err := t.trackerRepo.Tracker(int32(userID))
	if err != nil {
		logger.LogError("failed to retrieve tracker data in db", err)
		resErr.WriteMessage(http.StatusInternalServerError, "internal server error, please try again later")
		return resErr
	}

	resTracker := TrackerResponse{
		ID: dbTracker.ID,
		SpreadSheetID: dbTracker.SpreadSheetID,
		Year: dbTracker.Year,
		Month: dbTracker.Month,
		CreatedAt: dbTracker.CreatedAt,
		UpdatedAt: dbTracker.UpdatedAt,
	}

	var res httpresponse.DataResponse
	res.WriteBody(http.StatusOK, resTracker)
	return res
}

func (t *TrackerHttpHandler) UpdateTrackerHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	defer r.Body.Close()

	var trackerReq TrackerRequest
	var resErr httpresponse.ErrorResponse

	if err := json.NewDecoder(r.Body).Decode(&trackerReq); err != nil {
		logger.LogError("failed to decode json data from tracker http request", err)
		resErr.WriteMessage(http.StatusBadRequest, "invalid JSON")
		return resErr
	}

	trackerReq.SpreadSheetID = strings.TrimSpace(trackerReq.SpreadSheetID)
	year := int(trackerReq.Year)
	month := int(trackerReq.Month)
	
	// Month needs to be type cast to turn it to numeric representation, 
	// because it implements String() that returns month in text(November)
	dateNow := time.Now()
	currentYear := dateNow.Year()
	currentMonth := int(dateNow.Month())

	fmt.Println(currentYear)

	currentYearOffset := year-currentYear
	if year < 2000 || currentYearOffset > 1 || currentYearOffset < -1 {
		validRangeMessage := fmt.Sprintf("invalid year. Allowed range: %d-%d", currentYear-1, currentYear+1)
		logger.LogError(validRangeMessage, nil)
		resErr.WriteMessage(http.StatusBadRequest, validRangeMessage)
		return resErr
	}
	
	currentMonthOffset := month-currentMonth
	if month < 1 || month > 12 || currentMonthOffset > 1 || currentMonthOffset < -1 {
		validRangeMessage := fmt.Sprintf("invalid month. Allowed range: %d-%d", currentMonth-1, currentMonth+1)
		logger.LogError(validRangeMessage, nil)
		resErr.WriteMessage(http.StatusBadRequest, validRangeMessage)
		return resErr
	}

	userID, err := account.CurrentUserID(r)
	if err != nil {
		logger.LogError("user not found", err)
		resErr.WriteMessage(http.StatusNotFound, "user not found")
		return resErr
	}

	_, err = t.trackerRepo.UpdateTracker(&trackerReq, int32(userID))
	if err != nil {
		logger.LogError("updating tracker endpoint failed", err)
		resErr.WriteMessage(http.StatusInternalServerError, "failed to add tracker")
		return resErr
	}

	var res httpresponse.OKResponse
	res.WriteMessage(http.StatusOK, "successfully added the tracker")
	return res
}