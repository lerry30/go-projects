package in_process_scheduling

type SchedulesRepository interface {
	Users() (*[]UsersScheduleEntity, error)
	User(userID int32) (*UsersScheduleEntity, error)
}