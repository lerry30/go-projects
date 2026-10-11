package appscripts

import (
	"net/http"

	"tracker/internal/shared/httpresponse"
	"tracker/internal/shared/account"
	"tracker/internal/platform/logger"
)

type AppScriptsHandler struct {
	appScriptsRepo AppScriptsRepository
}

func NewAppScriptsHandler(appScriptsRepo AppScriptsRepository) *AppScriptsHandler {
	return &AppScriptsHandler{
		appScriptsRepo: appScriptsRepo,
	}
}

func (a *AppScriptsHandler) AppScriptHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	var resErr httpresponse.ErrorResponse

	userID, err := account.CurrentUserID(r)
	if err != nil {
		logger.LogError("user not found", err)
		resErr.WriteMessage(http.StatusBadRequest, "user not found")
		return resErr
	}

	dbAppScript, err := a.appScriptsRepo.AppScript(int32(userID))
	if err != nil {
		logger.LogError("failed to retrieve appscript data in db", err)
		resErr.WriteMessage(http.StatusInternalServerError, "internal server error")
		return resErr
	}

	resAppScript := AppScriptsResponse{
		ID: dbAppScript.ID,
		Endpoint: dbAppScript.Endpoint,
		CreatedAt: dbAppScript.CreatedAt,
		UpdatedAt: dbAppScript.UpdatedAt,
	}

	var res httpresponse.DataResponse
	res.WriteBody(http.StatusOK, resAppScript)
	return res
}