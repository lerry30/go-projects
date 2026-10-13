package appscripts

import (
	"fmt"
)

// NOTE: ON CONFLICT (employee_id) requires a unique constraint or index on employee_id.
var upsertAppScriptURL = fmt.Sprintf(`
		INSERT INTO %s(%s, %s) VALUES(@%s, @%s)
		ON CONFLICT (%s) DO UPDATE SET %s = EXCLUDED.%s
		RETURNING *;
	`, 
	TableName, 
	AppScriptsCols.EmployeeID,
	AppScriptsCols.Endpoint,
	AppScriptsCols.EmployeeID,
	AppScriptsCols.Endpoint,
	AppScriptsCols.EmployeeID,
	AppScriptsCols.Endpoint,
	AppScriptsCols.Endpoint,
)