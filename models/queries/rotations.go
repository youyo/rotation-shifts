package queries

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gocraft/dbr"
)

type (
	Rotation struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		StartDate Date      `json:"start_date"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Date struct {
		time.Time
	}

	Rotations []Rotation
)

func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(data))
	if err != nil {
		return err
	}
	*d = Date{t}
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) Scan(value interface{}) error {
	d.Time = value.(time.Time)
	return nil
}

func (d *Date) Value() (driver.Value, error) {
	return d.Time, nil
}

func NewRotation() *Rotation {
	return new(Rotation)
}

func NewRotations() *Rotations {
	return new(Rotations)
}

func (r *Rotation) GetRotation(db *dbr.Session, rotationId int) (*Rotation, error) {
	query := "select * from rotations where id=?"
	if row, err := db.SelectBySql(query, rotationId).Load(r); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}
	return r, nil
}

func (r *Rotations) GetRotations(db *dbr.Session) (*Rotations, error) {
	query := "select * from rotations"
	if row, err := db.SelectBySql(query).Load(r); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}
	return r, nil
}

func (r *Rotation) PostRotation(tx *dbr.Tx, name string, startDate Date) error {
	d, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", startDate))
	if err != nil {
		return err
	}
	query := "insert into rotations (name,start_date) values (?,?)"
	_, err = tx.InsertBySql(query, name, d).Exec()
	return err
}

func (s *Rotation) DeleteRotation(tx *dbr.Tx, rotationId int) error {
	query := "delete from rotations where id=?"
	_, err := tx.DeleteBySql(query, rotationId).Exec()
	return err
}
