package appscripts

type AppScriptsRepository interface {
	AppScript(userID int32) (*AppScriptsEntity, error)
	UpdateAppScript(endpoint string, userID int32) (int32, error)
}