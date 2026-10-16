package timelogs

import (
	"fmt"
)

// NOTE: ON CONFLICT (employee_id) requires a unique constraint or index on employee_id.
var upsertTimeLogs = fmt.Sprintf(`
		INSERT INTO %s(%s, %s, %s, %s, %s, %s, %s, %s, %s) VALUES(@%s, @%s, @%s, @%s, @%s, @%s, @%s, @%s, @%s)
		ON CONFLICT (%s) DO UPDATE SET %s = EXCLUDED.%s, %s = EXCLUDED.%s, %s = EXCLUDED.%s, %s = EXCLUDED.%s, %s = EXCLUDED.%s, %s = EXCLUDED.%s, %s = EXCLUDED.%s, %s = EXCLUDED.%s
		RETURNING *;
	`,
	TableName,
	TimeLogsCols.EmployeeID,
	TimeLogsCols.Login,
	TimeLogsCols.FirstBreakOut,
	TimeLogsCols.FirstBreakIn,
	TimeLogsCols.LunchOut,
	TimeLogsCols.LunchIn,
	TimeLogsCols.SecondBreakOut,
	TimeLogsCols.SecondBreakIn,
	TimeLogsCols.Logout,
	TimeLogsCols.EmployeeID,
	TimeLogsCols.Login,
	TimeLogsCols.FirstBreakOut,
	TimeLogsCols.FirstBreakIn,
	TimeLogsCols.LunchOut,
	TimeLogsCols.LunchIn,
	TimeLogsCols.SecondBreakOut,
	TimeLogsCols.SecondBreakIn,
	TimeLogsCols.Logout,
	TimeLogsCols.EmployeeID,
	TimeLogsCols.Login,
	TimeLogsCols.Login,
	TimeLogsCols.FirstBreakOut,
	TimeLogsCols.FirstBreakOut,
	TimeLogsCols.FirstBreakIn,
	TimeLogsCols.FirstBreakIn,
	TimeLogsCols.LunchOut,
	TimeLogsCols.LunchOut,
	TimeLogsCols.LunchIn,
	TimeLogsCols.LunchIn,
	TimeLogsCols.SecondBreakOut,
	TimeLogsCols.SecondBreakOut,
	TimeLogsCols.SecondBreakIn,
	TimeLogsCols.SecondBreakIn,
	TimeLogsCols.Logout,
	TimeLogsCols.Logout,
)

/*
INSERT INTO time_logs(employee_id, login, first_break_out, first_break_in, lunch_out, lunch_in, second_break_out, second_break_in, logout) VALUES(@employee_id, @login, @first_break_out, @first_break_in, @lunch_out, @lunch_in, @second_break_out, @second_break_in, @logout)
ON CONFLICT (employee_id) DO UPDATE SET login = EXCLUDED.login, first_break_out = EXCLUDED.first_break_out, first_break_in = EXCLUDED.first_break_in, lunch_out = EXCLUDED.lunch_out, lunch_in = EXCLUDED.lunch_in, second_break_out = EXCLUDED.second_break_out, second_break_in = EXCLUDED.second_break_in, logout = EXCLUDED.logout, %!s(MISSING) = EXCLUDED.%!s(MISSING)
RETURNING *;
*/