package in_process_scheduling

var selectUsersDetailsByStatus = `
	SELECT
		-- user entry
		u.id AS user_id,
		u.first_name,
		u.last_name,
		u.username,

		-- appscripts
		sc.id AS appscript_id,
		sc.appscript_endpoint AS endpoint,
		sc.status,

		-- tracker (monthly spreadsheet)
		tr.id AS tracker_id,
		tr.spread_sheet_id,
		tr.year,
		tr.month,

		-- time logs entry
		tm.id AS time_logs_id,
		tm.login,
		tm.first_break_out,
		tm.first_break_in,
		tm.lunch_out,
		tm.lunch_in,
		tm.second_break_out,
		tm.second_break_in,
		tm.logout
	FROM appscripts AS sc
	LEFT JOIN users AS u
		ON u.id = sc.employee_id
	LEFT JOIN tracker AS tr
		ON tr.employee_id = sc.employee_id
	LEFT JOIN time_logs AS tm
		ON tm.employee_id = sc.employee_id
	WHERE sc.status = @status;
`;

var selectUser = `
	SELECT
		-- user entry
		u.id AS user_id,
		u.first_name,
		u.last_name,
		u.username,

		-- appscripts
		sc.id AS appscript_id,
		sc.appscript_endpoint AS endpoint,
		sc.status,

		-- tracker (monthly spreadsheet)
		tr.id AS tracker_id,
		tr.spread_sheet_id,
		tr.year,
		tr.month,

		-- time logs entry
		tm.id AS time_logs_id,
		tm.login,
		tm.first_break_out,
		tm.first_break_in,
		tm.lunch_out,
		tm.lunch_in,
		tm.second_break_out,
		tm.second_break_in,
		tm.logout
	FROM appscripts AS sc
	LEFT JOIN users AS u
		ON u.id = sc.employee_id
	LEFT JOIN tracker AS tr
		ON tr.employee_id = sc.employee_id
	LEFT JOIN time_logs AS tm
		ON tm.employee_id = sc.employee_id
	WHERE sc.status = @status AND u.id = @id;
`;