package tracker

type TrackerRepository interface {
	Tracker(userID int32) (*TrackerEntity, error)
}