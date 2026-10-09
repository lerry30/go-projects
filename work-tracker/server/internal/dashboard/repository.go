package dashboard

type DashboardRepository interface {
	GetAll(userID int64) (*EmployeeTimeLogsEntity, error)
}