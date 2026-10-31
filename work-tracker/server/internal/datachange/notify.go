package datachange

import (
	"tracker/internal/in_process_scheduling"
)

type NotifyScheduler struct {
	usersWithSchedules *in_process_scheduling.UsersWithSchedules
}

func NewNotifyScheduler(usersWithSchedules *in_process_scheduling.UsersWithSchedules) *NotifyScheduler {
	return &NotifyScheduler{
		usersWithSchedules: usersWithSchedules,
	}
}

func (n *NotifyScheduler) Trigger(userID int32) bool {
	select {
	case n.usersWithSchedules.UserID <- userID:
		return true
	default:
		return false
	}
}