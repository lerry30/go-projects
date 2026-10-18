package appscripts

import (
	"strings"
	"net/http"
	"encoding/json"
	"slices"

	"tracker/internal/shared/httpresponse"
	"tracker/internal/shared/account"
	"tracker/internal/platform/logger"
)

type AppScriptsHttpHandler struct {
	appScriptsRepo AppScriptsRepository
}

func NewAppScriptsHttpHandler(appScriptsRepo AppScriptsRepository) *AppScriptsHttpHandler {
	return &AppScriptsHttpHandler{
		appScriptsRepo: appScriptsRepo,
	}
}

func (a *AppScriptsHttpHandler) AppScriptHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	var resErr httpresponse.ErrorResponse

	userID, err := account.CurrentUserID(r)
	if err != nil {
		logger.LogError("user not found", err)
		resErr.WriteMessage(http.StatusNotFound, "user not found")
		return resErr
	}

	dbAppScript, err := a.appScriptsRepo.AppScript(int32(userID))
	if err != nil {
		logger.LogError("failed to retrieve appscript data in db", err)
		resErr.WriteMessage(http.StatusInternalServerError, "internal server error, please try again later")
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

func (a *AppScriptsHttpHandler) UpdateAppScriptEndpointHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	defer r.Body.Close()

	var appScriptReq AppScriptsEndpointRequest
	var resErr httpresponse.ErrorResponse

	if err := json.NewDecoder(r.Body).Decode(&appScriptReq); err != nil {
		logger.LogError("failed to decode json request, invalid http request body", err)
		resErr.WriteMessage(http.StatusBadRequest, "invalid JSON")
		return resErr
	}

	endpoint := strings.TrimSpace(appScriptReq.Endpoint)
	isAnAppScriptURLPrefix := strings.HasPrefix(endpoint, "https://script.google.com/macros/s")
	isAnAppScriptURLSuffix := strings.HasSuffix(endpoint, "exec")

	if !isAnAppScriptURLPrefix || !isAnAppScriptURLSuffix {
		logger.LogError("invalid appscript url", nil)
		resErr.WriteMessage(http.StatusBadRequest, "invalid appscript url")
		return resErr
	}

	userID, err := account.CurrentUserID(r)
	if err != nil {
		logger.LogError("user id doesn't exists", err)
		resErr.WriteMessage(http.StatusNotFound, "user not found")
		return resErr
	}

	_, err = a.appScriptsRepo.UpdateAppScriptEndpoint(endpoint, int32(userID))
	if err != nil {
		logger.LogError("updating appscript url endpoint failed", err)
		resErr.WriteMessage(http.StatusInternalServerError, "failed to add appscript url")
		return resErr
	}

	var res httpresponse.OKResponse
	res.WriteMessage(http.StatusOK, "successfully added the appscript url")
	return res
}

func (a *AppScriptsHttpHandler) UpdateAppScriptStatusHandler(w http.ResponseWriter, r *http.Request) httpresponse.Response {
	defer r.Body.Close()

	var appScriptReq AppScriptsStatusRequest
	var resErr httpresponse.ErrorResponse

	if err := json.NewDecoder(r.Body).Decode(&appScriptReq); err != nil {
		logger.LogError("failed to decode json request, invalid http request body", err)
		resErr.WriteMessage(http.StatusBadRequest, "invalid JSON")
		return resErr
	}

	var status = strings.TrimSpace(appScriptReq.Status)
	status = strings.ToLower(status)

	validState := []string{"stopped", "running"}
	if !slices.Contains(validState, status) {
		logger.LogError("invalid status", nil)
		resErr.WriteMessage(http.StatusBadRequest, "invalid status")
		return resErr
	}

	userID, err := account.CurrentUserID(r)
	if err != nil {
		logger.LogError("user not found", err)
		resErr.WriteMessage(http.StatusBadRequest, "user not found")
		return resErr
	}

	dbAppScript, err := a.appScriptsRepo.AppScript(int32(userID))
	if dbAppScript.ID == 0 {
		logger.LogError("appscript url not found", nil)
		resErr.WriteMessage(http.StatusNotFound, "appscript url not found")
		return resErr
	}
	if err != nil {
		logger.LogError("failed to retrieve appscript data in db", err)
		resErr.WriteMessage(http.StatusInternalServerError, "internal server error, please try again later")
		return resErr
	}

	_, err = a.appScriptsRepo.UpdateAppScriptStatus(status, int32(userID))
	if err != nil {
		logger.LogError("failed to update appscript status", err)
		resErr.WriteMessage(http.StatusInternalServerError, "failed to update status")
		return resErr
	}

	var res httpresponse.OKResponse
	res.WriteMessage(http.StatusOK, "successfully updated the appscript status")
	return res
}