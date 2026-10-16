package types

import (
	"fmt"
	"time"
	"encoding/json"
	"errors"

	"tracker/internal/shared/utils"
)

type ClockTime struct{
	time.Time
}

func (c *ClockTime) UnmarshalJSON(b []byte) error {
	var strTime string
	if err := json.Unmarshal(b, &strTime); err != nil {
		return errors.New("time must be a string")
	}

	t, err := utils.ParseStrTo24HourFormat(strTime)
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}

func (c ClockTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.To24HourFormat(c.Time))
}

func (c *ClockTime) Scan(v any) error {
	t, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("unexpected type %T", v)
	}
	c.Time = t
	return nil
}