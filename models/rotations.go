package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/youyo/rotation-shifts/db"
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
		*time.Time
	}
)

func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(data))
	if err != nil {
		return err
	}
	*d = Date{&t}
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func NewRotation() *Rotation {
	return new(Rotation)
}

func (r *Rotation) GetRotation(id int) (*Rotation, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("rotations").Where("id=?", id).Load(r)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return r, nil
}

func (r *Rotation) PostRotation(name string, startDate Date) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	d, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", startDate))
	if err != nil {
		return err
	}
	result, err := tx.InsertInto("rotations").
		Columns("name", "start_date").
		Values(name, d).
		Exec()
	if err != nil {
		tx.Rollback()
		return err
	}
	if row, err := result.RowsAffected(); err != nil {
		tx.Rollback()
		return err
	} else if row == 0 {
		tx.Rollback()
		return errors.New("No insert record.")
	}
	tx.Commit()
	return nil
}

func (s *Rotation) PatchRotationName(rotationId int, name string) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Update("rotations").Set("name", name).Where("id=?", rotationId).Exec(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Rotation) PatchRotationStartDate(rotationId int, startDate time.Time) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Update("rotations").Set("start_date", startDate).Where("id=?", rotationId).Exec(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Rotation) DeleteRotation(rotationId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.DeleteFrom("rotations").
		Where("id=?", rotationId).
		Exec()
	if err != nil {
		tx.Rollback()
		return err
	}
	if row, err := result.RowsAffected(); err != nil {
		tx.Rollback()
		return err
	} else if row == 0 {
		tx.Rollback()
		return errors.New("No match record.")
	}
	tx.Commit()
	return nil
}

type Rotations []Rotation

func NewRotations() *Rotations {
	return new(Rotations)
}

func (s *Rotations) GetRotations() (*Rotations, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("rotations").Load(s)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return s, nil
}
