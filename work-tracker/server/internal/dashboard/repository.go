package dashboard

type DashboardRepository interface {
	DashboardData(userID int64) (*EmployeeTimeLogsEntity, error)
}