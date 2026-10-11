-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW IS DISTINCT FROM OLD THEN
    NEW.updated_at = CURRENT_TIMESTAMP;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER trigger_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER trigger_tracker_updated_at
BEFORE UPDATE ON tracker
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER trigger_appscripts_updated_at
BEFORE UPDATE ON appscripts
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER trigger_time_logs_updated_at
BEFORE UPDATE ON time_logs
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
-- +goose StatementEnd

-- +goose Down
DROP TRIGGER IF EXISTS trigger_users_updated_at ON users;
DROP TRIGGER IF EXISTS trigger_tracker_updated_at ON tracker;
DROP TRIGGER IF EXISTS trigger_appscripts_updated_at ON appscripts;
DROP TRIGGER IF EXISTS trigger_time_logs_updated_at ON time_logs;

DROP FUNCTION IF EXISTS set_updated_at();


--INSERT INTO tracker(employee_id, spread_sheet_id, "year", "month") VALUES(1, '1ttTYPbzvW6JGi98lnHEUCI5FKrCMx538Ecm2kaQTfSs', 2026, 10);
--INSERT INTO time_logs(employee_id, tracker_id, login, first_break_out, first_break_in, lunch_out, lunch_in, second_break_out, second_break_in, logout) VALUES (1, 1, '20:40:00', '23:00:00', '23:13:00', '01:00:00', '01:56:00', '04:00:00', '04:13:00', '06:00:00');
--INSERT INTO appscripts(employee_id, appscript_endpoint) VALUES(1, '1ttTYPbzvW6JGi98lnHEUCI5FKrCMx538Ecm2kaQTfSs');