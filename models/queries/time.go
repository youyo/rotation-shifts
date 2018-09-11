package queries

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type (
	DATE struct {
		time.Time
	}

	HOUR struct {
		time.Time
	}
)

func (d *DATE) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(data))
	if err != nil {
		return err
	}
	*d = DATE{t}
	return nil
}

func (d *DATE) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *DATE) Scan(value interface{}) error {
	d.Time = value.(time.Time)
	return nil
}

func (d *DATE) Value() (driver.Value, error) {
	return d.Time, nil
}

func (h *HOUR) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("\"15\"", string(data))
	if err != nil {
		return err
	}
	*h = HOUR{t}
	return nil
}

func (h *HOUR) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.Format("15"))
}

func (h *HOUR) Scan(value interface{}) error {
	h.Time = value.(time.Time)
	return nil
}

func (h *HOUR) Value() (driver.Value, error) {
	return h.Time, nil
}
