package timelogs

type TimeLogsRepository interface {
	TimeLogs(userID int32) (*TimeLogsEntity, error)
}