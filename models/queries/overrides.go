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
	Override struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		RotationId int    `json:"rotation_id"`
		Date       DATE   `json:"date"`
		UserId     int    `json:"user_id"`
		OverrideDetail
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	DATE struct {
		time.Time
	}

	OverrideDetail struct {
		Column0000 int `json:"0000" db:"0000"`
		Column0100 int `json:"0100" db:"0100"`
		Column0200 int `json:"0200" db:"0200"`
		Column0300 int `json:"0300" db:"0300"`
		Column0400 int `json:"0400" db:"0400"`
		Column0500 int `json:"0500" db:"0500"`
		Column0600 int `json:"0600" db:"0600"`
		Column0700 int `json:"0700" db:"0700"`
		Column0800 int `json:"0800" db:"0800"`
		Column0900 int `json:"0900" db:"0900"`
		Column1000 int `json:"1000" db:"1000"`
		Column1100 int `json:"1100" db:"1100"`
		Column1200 int `json:"1200" db:"1200"`
		Column1300 int `json:"1300" db:"1300"`
		Column1400 int `json:"1400" db:"1400"`
		Column1500 int `json:"1500" db:"1500"`
		Column1600 int `json:"1600" db:"1600"`
		Column1700 int `json:"1700" db:"1700"`
		Column1800 int `json:"1800" db:"1800"`
		Column1900 int `json:"1900" db:"1900"`
		Column2000 int `json:"2000" db:"2000"`
		Column2100 int `json:"2100" db:"2100"`
		Column2200 int `json:"2200" db:"2200"`
		Column2300 int `json:"2300" db:"2300"`
	}

	Overrides []Override
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

func NewOverride() *Override {
	return new(Override)
}

func NewOverrideDetail() *OverrideDetail {
	return new(OverrideDetail)
}

func NewOverrides() *Overrides {
	return new(Overrides)
}

func (u *Override) GetOverride(db *dbr.Session, id int) (*Override, error) {
	query := "select * from overrides where id=?"
	if row, err := db.SelectBySql(query, id).Load(u); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}
	return u, nil
}

func (u *Overrides) GetOverrides(db *dbr.Session) (*Overrides, error) {
	query := "select * from overrides"
	if row, err := db.SelectBySql(query).Load(u); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}
	return u, nil
}

func (u *Override) InsertOverride(tx *dbr.Tx, name string, date DATE, rotationId, userId int) error {
	query := "insert into overrides (name,rotation_id,date,user_id) values (?,?,?,?)"
	_, err := tx.InsertBySql(query, name, rotationId, date, userId).Exec()
	return err
}

func (u *Override) PatchOverrideDetail(tx *dbr.Tx, id int, detail *OverrideDetail) error {
	m := make(map[string]interface{})
	b, _ := json.Marshal(detail)
	json.Unmarshal(b, &m)

	_, err := tx.Update("overrides").SetMap(m).Where("id=?", id).Exec()
	return err
}

func (u *Override) DeleteOverride(tx *dbr.Tx, id int) error {
	query := "delete from overrides where id=?"
	_, err := tx.DeleteBySql(query, id).Exec()
	return err
}

func (u *Overrides) SelectOverrideAssignedUserIds(db *dbr.Session, rotationId int, date DATE, hour string) ([]int, error) {
	d := date.Format("2006-01-02")

	var userIds []int
	query := fmt.Sprintf("select user_id from overrides where rotation_id=? and date=? and `%s00`=1", hour)
	_, err := db.SelectBySql(query, rotationId, d).Load(&userIds)
	return userIds, err
}
