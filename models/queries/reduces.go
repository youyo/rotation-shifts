package queries

import (
	"fmt"
	"time"

	"github.com/gocraft/dbr"
)

type (
	Reduce struct {
		Id         int       `json:"id"`
		Name       string    `json:"name"`
		RotationId int       `json:"rotation_id"`
		Date       DATE      `json:"date"`
		Hour       HOUR      `json:"hour"`
		UserId     int       `json:"user_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	Reduces []Reduce
)

func NewReduce() *Reduce {
	return new(Reduce)
}

func NewReduces() *Reduces {
	return new(Reduces)
}

func (u *Reduce) GetReduce(db *dbr.Session, id int) (*Reduce, error) {
	query := "select * from reduces where id=?"
	if _, err := db.SelectBySql(query, id).Load(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Reduces) GetReduces(db *dbr.Session) (*Reduces, error) {
	query := "select * from reduces"
	if _, err := db.SelectBySql(query).Load(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *Reduce) InsertReduce(tx *dbr.Tx, name string, date DATE, hour HOUR, rotationId, userId int) error {
	d, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", date))
	if err != nil {
		return err
	}

	h, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", hour))
	if err != nil {
		return err
	}

	query := "insert into reduces (name,rotation_id,date,hour,user_id) values (?,?,?,?,?)"
	_, err = tx.InsertBySql(query, name, rotationId, d, h, userId).Exec()
	return err
}

func (u *Reduce) DeleteReduce(tx *dbr.Tx, id int) error {
	query := "delete from reduces where id=?"
	_, err := tx.DeleteBySql(query, id).Exec()
	return err
}

func (u *Reduces) SelectReduceAssignedUserIds(db *dbr.Session, rotationId int, date DATE, hour HOUR) ([]int, error) {
	d := date.Format("2006-01-02")
	h := "0000-01-01 " + hour.Format("15") + ":00:00"

	var userIds []int
	query := "select user_id from reduces where rotation_id=? and date=? and hour=?"
	_, err := db.SelectBySql(query, rotationId, d, h).Load(&userIds)
	return userIds, err
}
