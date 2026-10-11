package appscripts

type AppScriptsRepository interface {
	AppScript(userID int32) (*AppScriptsEntity, error)
}