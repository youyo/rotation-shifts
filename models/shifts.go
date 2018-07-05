package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/youyo/rotation-shifts/db"
)

type (
	Shift struct {
		Id        int       `json:"id"`
		Name      string    `json:"name"`
		UserId    int       `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Schedule
	}
	Schedule struct {
		Monday0000    int `json:"monday_0000"`
		Monday0100    int `json:"monday_0100"`
		Monday0200    int `json:"monday_0200"`
		Monday0300    int `json:"monday_0300"`
		Monday0400    int `json:"monday_0400"`
		Monday0500    int `json:"monday_0500"`
		Monday0600    int `json:"monday_0600"`
		Monday0700    int `json:"monday_0700"`
		Monday0800    int `json:"monday_0800"`
		Monday0900    int `json:"monday_0900"`
		Monday1000    int `json:"monday_1000"`
		Monday1100    int `json:"monday_1100"`
		Monday1200    int `json:"monday_1200"`
		Monday1300    int `json:"monday_1300"`
		Monday1400    int `json:"monday_1400"`
		Monday1500    int `json:"monday_1500"`
		Monday1600    int `json:"monday_1600"`
		Monday1700    int `json:"monday_1700"`
		Monday1800    int `json:"monday_1800"`
		Monday1900    int `json:"monday_1900"`
		Monday2000    int `json:"monday_2000"`
		Monday2100    int `json:"monday_2100"`
		Monday2200    int `json:"monday_2200"`
		Monday2300    int `json:"monday_2300"`
		Tuesday0000   int `json:"tuesday_0000"`
		Tuesday0100   int `json:"tuesday_0100"`
		Tuesday0200   int `json:"tuesday_0200"`
		Tuesday0300   int `json:"tuesday_0300"`
		Tuesday0400   int `json:"tuesday_0400"`
		Tuesday0500   int `json:"tuesday_0500"`
		Tuesday0600   int `json:"tuesday_0600"`
		Tuesday0700   int `json:"tuesday_0700"`
		Tuesday0800   int `json:"tuesday_0800"`
		Tuesday0900   int `json:"tuesday_0900"`
		Tuesday1000   int `json:"tuesday_1000"`
		Tuesday1100   int `json:"tuesday_1100"`
		Tuesday1200   int `json:"tuesday_1200"`
		Tuesday1300   int `json:"tuesday_1300"`
		Tuesday1400   int `json:"tuesday_1400"`
		Tuesday1500   int `json:"tuesday_1500"`
		Tuesday1600   int `json:"tuesday_1600"`
		Tuesday1700   int `json:"tuesday_1700"`
		Tuesday1800   int `json:"tuesday_1800"`
		Tuesday1900   int `json:"tuesday_1900"`
		Tuesday2000   int `json:"tuesday_2000"`
		Tuesday2100   int `json:"tuesday_2100"`
		Tuesday2200   int `json:"tuesday_2200"`
		Tuesday2300   int `json:"tuesday_2300"`
		Wednesday0000 int `json:"wednesday_0000"`
		Wednesday0100 int `json:"wednesday_0100"`
		Wednesday0200 int `json:"wednesday_0200"`
		Wednesday0300 int `json:"wednesday_0300"`
		Wednesday0400 int `json:"wednesday_0400"`
		Wednesday0500 int `json:"wednesday_0500"`
		Wednesday0600 int `json:"wednesday_0600"`
		Wednesday0700 int `json:"wednesday_0700"`
		Wednesday0800 int `json:"wednesday_0800"`
		Wednesday0900 int `json:"wednesday_0900"`
		Wednesday1000 int `json:"wednesday_1000"`
		Wednesday1100 int `json:"wednesday_1100"`
		Wednesday1200 int `json:"wednesday_1200"`
		Wednesday1300 int `json:"wednesday_1300"`
		Wednesday1400 int `json:"wednesday_1400"`
		Wednesday1500 int `json:"wednesday_1500"`
		Wednesday1600 int `json:"wednesday_1600"`
		Wednesday1700 int `json:"wednesday_1700"`
		Wednesday1800 int `json:"wednesday_1800"`
		Wednesday1900 int `json:"wednesday_1900"`
		Wednesday2000 int `json:"wednesday_2000"`
		Wednesday2100 int `json:"wednesday_2100"`
		Wednesday2200 int `json:"wednesday_2200"`
		Wednesday2300 int `json:"wednesday_2300"`
		Thursday0000  int `json:"thursday_0000"`
		Thursday0100  int `json:"thursday_0100"`
		Thursday0200  int `json:"thursday_0200"`
		Thursday0300  int `json:"thursday_0300"`
		Thursday0400  int `json:"thursday_0400"`
		Thursday0500  int `json:"thursday_0500"`
		Thursday0600  int `json:"thursday_0600"`
		Thursday0700  int `json:"thursday_0700"`
		Thursday0800  int `json:"thursday_0800"`
		Thursday0900  int `json:"thursday_0900"`
		Thursday1000  int `json:"thursday_1000"`
		Thursday1100  int `json:"thursday_1100"`
		Thursday1200  int `json:"thursday_1200"`
		Thursday1300  int `json:"thursday_1300"`
		Thursday1400  int `json:"thursday_1400"`
		Thursday1500  int `json:"thursday_1500"`
		Thursday1600  int `json:"thursday_1600"`
		Thursday1700  int `json:"thursday_1700"`
		Thursday1800  int `json:"thursday_1800"`
		Thursday1900  int `json:"thursday_1900"`
		Thursday2000  int `json:"thursday_2000"`
		Thursday2100  int `json:"thursday_2100"`
		Thursday2200  int `json:"thursday_2200"`
		Thursday2300  int `json:"thursday_2300"`
		Friday0000    int `json:"friday_0000"`
		Friday0100    int `json:"friday_0100"`
		Friday0200    int `json:"friday_0200"`
		Friday0300    int `json:"friday_0300"`
		Friday0400    int `json:"friday_0400"`
		Friday0500    int `json:"friday_0500"`
		Friday0600    int `json:"friday_0600"`
		Friday0700    int `json:"friday_0700"`
		Friday0800    int `json:"friday_0800"`
		Friday0900    int `json:"friday_0900"`
		Friday1000    int `json:"friday_1000"`
		Friday1100    int `json:"friday_1100"`
		Friday1200    int `json:"friday_1200"`
		Friday1300    int `json:"friday_1300"`
		Friday1400    int `json:"friday_1400"`
		Friday1500    int `json:"friday_1500"`
		Friday1600    int `json:"friday_1600"`
		Friday1700    int `json:"friday_1700"`
		Friday1800    int `json:"friday_1800"`
		Friday1900    int `json:"friday_1900"`
		Friday2000    int `json:"friday_2000"`
		Friday2100    int `json:"friday_2100"`
		Friday2200    int `json:"friday_2200"`
		Friday2300    int `json:"friday_2300"`
		Saturday0000  int `json:"saturday_0000"`
		Saturday0100  int `json:"saturday_0100"`
		Saturday0200  int `json:"saturday_0200"`
		Saturday0300  int `json:"saturday_0300"`
		Saturday0400  int `json:"saturday_0400"`
		Saturday0500  int `json:"saturday_0500"`
		Saturday0600  int `json:"saturday_0600"`
		Saturday0700  int `json:"saturday_0700"`
		Saturday0800  int `json:"saturday_0800"`
		Saturday0900  int `json:"saturday_0900"`
		Saturday1000  int `json:"saturday_1000"`
		Saturday1100  int `json:"saturday_1100"`
		Saturday1200  int `json:"saturday_1200"`
		Saturday1300  int `json:"saturday_1300"`
		Saturday1400  int `json:"saturday_1400"`
		Saturday1500  int `json:"saturday_1500"`
		Saturday1600  int `json:"saturday_1600"`
		Saturday1700  int `json:"saturday_1700"`
		Saturday1800  int `json:"saturday_1800"`
		Saturday1900  int `json:"saturday_1900"`
		Saturday2000  int `json:"saturday_2000"`
		Saturday2100  int `json:"saturday_2100"`
		Saturday2200  int `json:"saturday_2200"`
		Saturday2300  int `json:"saturday_2300"`
		Sunday0000    int `json:"sunday_0000"`
		Sunday0100    int `json:"sunday_0100"`
		Sunday0200    int `json:"sunday_0200"`
		Sunday0300    int `json:"sunday_0300"`
		Sunday0400    int `json:"sunday_0400"`
		Sunday0500    int `json:"sunday_0500"`
		Sunday0600    int `json:"sunday_0600"`
		Sunday0700    int `json:"sunday_0700"`
		Sunday0800    int `json:"sunday_0800"`
		Sunday0900    int `json:"sunday_0900"`
		Sunday1000    int `json:"sunday_1000"`
		Sunday1100    int `json:"sunday_1100"`
		Sunday1200    int `json:"sunday_1200"`
		Sunday1300    int `json:"sunday_1300"`
		Sunday1400    int `json:"sunday_1400"`
		Sunday1500    int `json:"sunday_1500"`
		Sunday1600    int `json:"sunday_1600"`
		Sunday1700    int `json:"sunday_1700"`
		Sunday1800    int `json:"sunday_1800"`
		Sunday1900    int `json:"sunday_1900"`
		Sunday2000    int `json:"sunday_2000"`
		Sunday2100    int `json:"sunday_2100"`
		Sunday2200    int `json:"sunday_2200"`
		Sunday2300    int `json:"sunday_2300"`
	}
)

func NewShift() *Shift {
	return new(Shift)
}

func (s *Shift) GetShift(id int) (*Shift, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("shifts").Where("id=?", id).Load(s)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return s, nil
}

func (s *Shift) PostShift(name string, userId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.InsertInto("shifts").
		Columns("name", "user_id").
		Values(name, userId).
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

func (s *Shift) PatchShift(shiftId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	m := make(map[string]interface{})
	b, _ := json.Marshal(s.Schedule)
	json.Unmarshal(b, &m)
	if _, err := tx.Update("shifts").SetMap(m).Where("id=?", shiftId).Exec(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Shift) DeleteShift(shiftId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.DeleteFrom("shifts").
		Where("id=?", shiftId).
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

type Shifts []Shift

func NewShifts() *Shifts {
	return new(Shifts)
}

func (s *Shifts) GetShifts() (*Shifts, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("shifts").Load(s)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return s, nil
}
