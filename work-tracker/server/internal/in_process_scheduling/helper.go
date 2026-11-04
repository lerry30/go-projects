package in_process_scheduling

import (
	"time"
	"fmt"

	"tracker/internal/shared/utils"
)

func PopulateUserData(rawUserData *UsersScheduleEntity) (*User, error) {
	if rawUserData.UserID == nil ||
		rawUserData.FirstName == nil ||
		rawUserData.LastName == nil ||
		rawUserData.Username == nil ||
		rawUserData.AppScriptID == nil ||
		rawUserData.Endpoint == nil ||
		rawUserData.Status == nil ||
		rawUserData.TrackerID == nil ||
		rawUserData.SpreadSheetID == nil ||
		rawUserData.Year == nil ||
		rawUserData.Month == nil ||
		rawUserData.TimeLogsID == nil ||
		rawUserData.Login == nil ||
		rawUserData.FirstBreakOut == nil ||
		rawUserData.FirstBreakIn == nil ||
		rawUserData.LunchOut == nil ||
		rawUserData.LunchIn == nil ||
		rawUserData.SecondBreakOut == nil ||
		rawUserData.SecondBreakIn == nil ||
		rawUserData.Logout == nil {
		return nil, fmt.Errorf("user data must be completed to create schedule")
	}

	var user User
	user.UserID = utils.Deref(rawUserData.UserID, 0)
	user.FirstName = utils.Deref(rawUserData.FirstName, "")
	user.LastName = utils.Deref(rawUserData.LastName, "")
	user.Username = utils.Deref(rawUserData.Username, "")
	user.AppScriptID = utils.Deref(rawUserData.AppScriptID, 0)
	user.Endpoint = utils.Deref(rawUserData.Endpoint, "")
	user.Status = utils.Deref(rawUserData.Status, "")
	user.TrackerID = utils.Deref(rawUserData.TrackerID, 0)
	user.SpreadSheetID = utils.Deref(rawUserData.SpreadSheetID, "")
	user.Year = utils.Deref(rawUserData.Year, 0)
	user.Month = utils.Deref(rawUserData.Month, 0)
	user.TimeLogsID = utils.Deref(rawUserData.TimeLogsID, 0)

	timeLogs := []time.Time{
		utils.Deref(rawUserData.Login, time.Time{}),
		utils.Deref(rawUserData.FirstBreakOut, time.Time{}),
		utils.Deref(rawUserData.FirstBreakIn, time.Time{}),
		utils.Deref(rawUserData.LunchOut, time.Time{}),
		utils.Deref(rawUserData.LunchIn, time.Time{}),
		utils.Deref(rawUserData.SecondBreakOut, time.Time{}),
		utils.Deref(rawUserData.SecondBreakIn, time.Time{}),
		utils.Deref(rawUserData.Logout, time.Time{}),
	}
	user.CreateSchedules(timeLogs)
	
	return &user, nil
}