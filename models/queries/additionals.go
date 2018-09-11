package queries

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr"
)

type (
	Additional struct {
		Id         int       `json:"id"`
		Name       string    `json:"name"`
		RotationId int       `json:"rotation_id"`
		Date       DATE      `json:"date"`
		Hour       HOUR      `json:"hour"`
		UserId     int       `json:"user_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	Additionals []Additional
)

func NewAdditional() *Additional {
	return new(Additional)
}

func NewAdditionals() *Additionals {
	return new(Additionals)
}

func (u *Additional) GetAdditional(db *dbr.Session, id int) (*Additional, error) {
	query := "select * from additionals where id=?"
	if _, err := db.SelectBySql(query, id).Load(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Additionals) GetAdditionals(db *dbr.Session) (*Additionals, error) {
	query := "select * from additionals"
	if _, err := db.SelectBySql(query).Load(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Additional) InsertAdditional(tx *dbr.Tx, name string, date DATE, hour HOUR, rotationId, userId int) error {
	d, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", date))
	if err != nil {
		return err
	}

	h, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", hour))
	if err != nil {
		return err
	}

	query := "insert into additionals (name,rotation_id,date,hour,user_id) values (?,?,?,?,?)"
	_, err = tx.InsertBySql(query, name, rotationId, d, h, userId).Exec()
	return err
}

func (u *Additional) DeleteAdditional(tx *dbr.Tx, id int) error {
	query := "delete from additionals where id=?"
	_, err := tx.DeleteBySql(query, id).Exec()
	return err
}

func (u *Additionals) SelectAdditionalAssignedUserIds(db *dbr.Session, rotationId int, date DATE, hour HOUR) ([]int, error) {
	d := date.Format("2006-01-02")
	h := "0000-01-01 " + hour.Format("15") + ":00:00"

	var userIds []int
	query := "select user_id from additionals where rotation_id=? and date=? and hour=?"
	_, err := db.SelectBySql(query, rotationId, d, h).Load(&userIds)
	return userIds, err
}
