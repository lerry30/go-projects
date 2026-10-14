package tracker

import (
	"fmt"
)

// NOTE: ON CONFLICT (employee_id) requires a unique constraint or index on employee_id.
var upsertTracker = fmt.Sprintf(`
		INSERT INTO %s(%s, %s, %s, %s) VALUES (@%s, @%s, @%s, @%s)
		ON CONFLICT (%s) DO UPDATE SET %s = EXCLUDED.%s, %s = EXCLUDED.%s, %s = EXCLUDED.%s
		RETURNING *;
	`,
	TableName,
	TrackerCols.EmployeeID,
	TrackerCols.SpreadSheetID,
	TrackerCols.Year,
	TrackerCols.Month,
	TrackerCols.EmployeeID,
	TrackerCols.SpreadSheetID,
	TrackerCols.Year,
	TrackerCols.Month,
	TrackerCols.EmployeeID,
	TrackerCols.SpreadSheetID,
	TrackerCols.SpreadSheetID,
	TrackerCols.Year,
	TrackerCols.Year,
	TrackerCols.Month,
	TrackerCols.Month,
)

/*
INSERT INTO tracker(employee_id, spread_sheet_id, year, month) VALUES (1, 'j49382932egrg', 2026, 10)
ON CONFLICT (employee_id) DO UPDATE SET spread_sheet_id = EXCLUDED.spread_sheet_id, year = EXCLUDED.year, month = EXCLUDED.month
RETURNING *;
*/