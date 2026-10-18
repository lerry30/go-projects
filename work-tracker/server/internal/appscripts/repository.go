package appscripts

type AppScriptsRepository interface {
	AppScript(userID int32) (*AppScriptsEntity, error)
	UpdateAppScriptEndpoint(endpoint string, userID int32) (int32, error)
	UpdateAppScriptStatus(status string, userID int32) (int32, error)
}