package tracker

type TrackerRepository interface {
	Tracker(userID int32) (*TrackerEntity, error)
	UpdateTracker(tracker *TrackerRequest, userID int32) (int32, error)
}