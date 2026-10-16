package timelogs

type TimeLogsRepository interface {
	TimeLogs(userID int32) (*TimeLogsEntity, error)
	UpdateTimeLogs(timeLogs *TimeLogsRequest, userID int32) (int32, error)
}