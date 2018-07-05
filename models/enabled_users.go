package models

import (
	"errors"
	"time"

	"github.com/youyo/rotation-shifts/db"
)

type EnabledUser struct {
	UserId    int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewEnabledUser() *EnabledUser {
	return new(EnabledUser)
}

/*
func (e *EnabledUser) GetEnabledUser(userId int) (*EnabledUser, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("enabled_users").Where("user_id=?", userId).Load(e)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return e, nil
}
*/

func (e *EnabledUser) PostEnabledUser(userId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.InsertInto("enabled_users").
		Columns("user_id").
		Values(userId).
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

func (e *EnabledUser) DeleteEnabledUser(userId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.DeleteFrom("enabled_users").Where("user_id=?", userId).Exec(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type EnabledUsers []EnabledUser

func NewEnabledUsers() *EnabledUsers {
	return new(EnabledUsers)
}

func (e *EnabledUsers) GetEnabledUsers() (*EnabledUsers, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("enabled_users").Load(e)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return e, nil
}
