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

/*

CREATE TABLE appscripts(
	employee_id INTEGER REFERENCES users(id) UNIQUE,
	appscript_endpoint VARCHAR(255) NOT NULL,
	action VARCHAR(10) NOT NULL DEFAULT 'stop' CHECK (action IN ('stop', 'auto')),
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	id SERIAL PRIMARY KEY
);

INSERT INTO appscripts(employee_id, appscript_endpoint) VALUES(1, 'test value')
ON CONFLICT (employee_id) DO UPDATE SET appscript_endpoint = EXCLUDED.appscript_endpoint
RETURNING *;
*/