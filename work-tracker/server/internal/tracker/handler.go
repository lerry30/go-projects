package tracker

import (
	"net/http"

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