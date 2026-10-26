package in_process_scheduling

type SchedulesRepository interface {
	RunningAppScripts() (*[]UsersScheduleEntity, error)
}