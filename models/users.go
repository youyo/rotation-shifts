package models

import (
	"errors"
	"time"

	"github.com/youyo/rotation-shifts/db"
)

type User struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser() *User {
	return new(User)
}

/*
func (u *User) GetUser(id int) (*User, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("users").Where("id=?", id).Load(u)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return u, nil
}
*/

func (u *User) PostUser() (int, error) {
	c, err := db.NewConnection()
	if err != nil {
		return 0, err
	}
	tx, err := c.Begin()
	if err != nil {
		return 0, err
	}
	result, err := tx.InsertBySql("insert into users values ()").Exec()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if row, err := result.RowsAffected(); err != nil {
		tx.Rollback()
		return 0, err
	} else if row == 0 {
		tx.Rollback()
		return 0, errors.New("No insert record.")
	}
	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return int(id), nil
}

func (u *User) DeleteUser(userId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.DeleteFrom("users").Where("id=?", userId).Exec(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type Users []User

func NewUsers() *Users {
	return new(Users)
}

func (u *Users) GetUsers() (*Users, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("users").Load(u)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return u, nil
}
