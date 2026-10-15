package timelogs

import (
	"net/http"

	"tracker/internal/shared/httpresponse"
	"tracker/internal/shared/account"
	"tracker/internal/shared/utils"
	"tracker/internal/platform/logger"
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

	resTimeLogs := TimeLogsResponse{
		ID: dbTimeLogs.ID,
		Login: utils.To24HourFormat(dbTimeLogs.Login),
		FirstBreakOut: utils.To24HourFormat(dbTimeLogs.FirstBreakOut),
		FirstBreakIn: utils.To24HourFormat(dbTimeLogs.FirstBreakIn),
		LunchOut: utils.To24HourFormat(dbTimeLogs.LunchOut),
		LunchIn: utils.To24HourFormat(dbTimeLogs.LunchIn),
		SecondBreakOut: utils.To24HourFormat(dbTimeLogs.SecondBreakOut),
		SecondBreakIn: utils.To24HourFormat(dbTimeLogs.SecondBreakIn),
		Logout: utils.To24HourFormat(dbTimeLogs.Logout),
		CreatedAt: dbTimeLogs.CreatedAt,
		UpdatedAt: dbTimeLogs.UpdatedAt,
	}

	var res httpresponse.DataResponse
	res.WriteBody(http.StatusOK, resTimeLogs)
	return res
}