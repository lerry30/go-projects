package timelogs

import (
	"net/http"
	"encoding/json"
	"time"

	"tracker/internal/platform/logger"
	"tracker/internal/shared/httpresponse"
	"tracker/internal/shared/account"
	"tracker/internal/shared/handler/types"
)

type TimeLogsHttpHandler struct {
	timeLogsRepo TimeLogsRepository
}

func NewTimeLogsHttpHandler(timeLogsRepo TimeLogsRepository) *TimeLogsHttpHandler {
	return &TimeLogsHttpHandler{
		timeLogsRepo: timeLogsRepo,
	}
}

func (t *TimeLogsHttpHandler) TimeLogsHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	var resErr httpresponse.ErrorResponse

	userID, err := account.CurrentUserID(r)
	if err != nil {
		logger.LogError("user not found", err)
		resErr.WriteMessage(http.StatusNotFound, "user not found")
		return resErr
	}

	dbTimeLogs, err := t.timeLogsRepo.TimeLogs(int32(userID))
	if err != nil {
		logger.LogError("failed to retrieve time logs data in db", err)
		resErr.WriteMessage(http.StatusInternalServerError, "internal server error, please try again later")
		return resErr
	}

	resTimeLogs := TimeLogsResponse{}
	resTimeLogs.ID = dbTimeLogs.ID
	resTimeLogs.Login.Scan(dbTimeLogs.Login)
	resTimeLogs.FirstBreakOut.Scan(dbTimeLogs.FirstBreakOut)
	resTimeLogs.FirstBreakIn.Scan(dbTimeLogs.FirstBreakIn)
	resTimeLogs.LunchOut.Scan(dbTimeLogs.LunchOut)
	resTimeLogs.LunchIn.Scan(dbTimeLogs.LunchIn)
	resTimeLogs.SecondBreakOut.Scan(dbTimeLogs.SecondBreakOut)
	resTimeLogs.SecondBreakIn.Scan(dbTimeLogs.SecondBreakIn)
	resTimeLogs.Logout.Scan(dbTimeLogs.Logout)
	resTimeLogs.CreatedAt = dbTimeLogs.CreatedAt
	resTimeLogs.UpdatedAt = dbTimeLogs.UpdatedAt

	var res httpresponse.DataResponse
	res.WriteBody(http.StatusOK, resTimeLogs)
	return res
}

func (t *TimeLogsHttpHandler) UpdateTimeLogsHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	defer r.Body.Close()

	var timeLogsReq TimeLogsRequest
	var resErr httpresponse.ErrorResponse
	
	if err := json.NewDecoder(r.Body).Decode(&timeLogsReq); err != nil {
		logger.LogError("failed to decode json data from time logs http request", err)
		resErr.WriteMessage(http.StatusBadRequest, "invalid JSON")
		return resErr
	}

	// time sequence validation
	timeLogsSeq := []types.ClockTime{
		timeLogsReq.Login,
		timeLogsReq.FirstBreakOut,
		timeLogsReq.FirstBreakIn,
		timeLogsReq.LunchOut,
		timeLogsReq.LunchIn,
		timeLogsReq.SecondBreakOut,
		timeLogsReq.SecondBreakIn,
		timeLogsReq.Logout,
	}
	tmSeq := make([]time.Time, len(timeLogsSeq))
	
	for i, seq := range timeLogsSeq {
		tmSeq[i] = seq.Time
	}

	resolvedLogsSeq := ResolveNightShiftSequence(tmSeq)
	// Shift should not span 24 hours
	if resolvedLogsSeq[len(resolvedLogsSeq)-1].Sub(resolvedLogsSeq[0]) >= 24*time.Hour {
		logger.LogError("times are not in a valid order", nil)
		resErr.WriteMessage(http.StatusBadRequest, "times are not in a valid order")
		return resErr
	}

	userID, err := account.CurrentUserID(r)
	if err != nil {
		logger.LogError("user not found", err)
		resErr.WriteMessage(http.StatusNotFound, "user not found")
		return resErr
	}

	_, err = t.timeLogsRepo.UpdateTimeLogs(&timeLogsReq, int32(userID))
	if err != nil {
		logger.LogError("updating time logs endpoint failed", err)
		resErr.WriteMessage(http.StatusInternalServerError, "failed to add time logs")
		return resErr
	}

	var res httpresponse.OKResponse
	res.WriteMessage(http.StatusOK, "successfully added the time logs")
	return res
}