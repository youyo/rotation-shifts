package queries

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr"
)

type (
	Override struct {
		Id         int       `json:"id"`
		Name       string    `json:"name"`
		RotationId int       `json:"rotation_id"`
		Date       DATE      `json:"date"`
		Hour       HOUR      `json:"hour"`
		UserId     int       `json:"user_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	Overrides []Override
)

func NewOverride() *Override {
	return new(Override)
}

func NewOverrides() *Overrides {
	return new(Overrides)
}

func (u *Override) GetOverride(db *dbr.Session, id int) (*Override, error) {
	query := "select * from overrides where id=?"
	if _, err := db.SelectBySql(query, id).Load(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Overrides) GetOverrides(db *dbr.Session) (*Overrides, error) {
	query := "select * from overrides"
	if _, err := db.SelectBySql(query).Load(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Override) InsertOverride(tx *dbr.Tx, name string, date DATE, hour HOUR, rotationId, userId int) error {
	d, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", date))
	if err != nil {
		return err
	}

	h, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", hour))
	if err != nil {
		return err
	}

	query := "insert into overrides (name,rotation_id,date,hour,user_id) values (?,?,?,?,?)"
	_, err = tx.InsertBySql(query, name, rotationId, d, h, userId).Exec()
	return err
}

func (u *Override) DeleteOverride(tx *dbr.Tx, id int) error {
	query := "delete from overrides where id=?"
	_, err := tx.DeleteBySql(query, id).Exec()
	return err
}

func (u *Overrides) SelectOverrideAssignedUserIds(db *dbr.Session, rotationId int, date DATE, hour HOUR) ([]int, error) {
	d := date.Format("2006-01-02")
	h := "0000-01-01 " + hour.Format("15") + ":00:00"

	var userIds []int
	query := "select user_id from overrides where rotation_id=? and date=? and hour=?"
	_, err := db.SelectBySql(query, rotationId, d, h).Load(&userIds)
	return userIds, err
}
