package models

import (
	"errors"
	"time"

	"github.com/youyo/rotation-shifts/db"
)

type DisabledUser struct {
	UserId    int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewDisabledUser() *DisabledUser {
	return new(DisabledUser)
}

func (e *DisabledUser) GetDisabledUser(userId int) (*DisabledUser, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("disabled_users").Where("user_id=?", userId).Load(e)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return e, nil
}

func (e *DisabledUser) PostDisabledUser(userId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.InsertInto("disabled_users").
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

func (e *DisabledUser) DeleteDisabledUser(userId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.DeleteFrom("disabled_users").Where("user_id=?", userId).Exec(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type DisabledUsers []DisabledUser

func NewDisabledUsers() *DisabledUsers {
	return new(DisabledUsers)
}

func (e *DisabledUsers) GetDisabledUsers() (*DisabledUsers, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("disabled_users").Load(e)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return e, nil
}
