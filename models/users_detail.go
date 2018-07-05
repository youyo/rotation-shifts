package models

import (
	"errors"
	"time"

	"github.com/youyo/rotation-shifts/db"
)

type UserDetail struct {
	UserId      int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUserDetail() *UserDetail {
	return new(UserDetail)
}

func (u *UserDetail) GetUserDetail(userId int) (*UserDetail, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").
		From("users_detail").
		Where("user_id=?", userId).
		Load(u)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return u, nil
}

func (u *UserDetail) PostUserDetail(userId int, name string) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.InsertInto("users_detail").
		Columns("user_id", "name").
		Values(userId, name).
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

func (u *UserDetail) UpdateUserDetailEmail(userId int, email string) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Update("users_detail").
		Set("email", email).
		Where("user_id=?", userId).
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

func (u *UserDetail) UpdateUserDetailPhoneNumber(userId int, phoneNumber string) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Update("users_detail").
		Set("phone_number", phoneNumber).
		Where("user_id=?", userId).
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

func (u *UserDetail) DeleteUserDetail(userId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.DeleteFrom("users_detail").
		Where("user_id=?", userId).
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

type UsersDetail []UserDetail

func NewUsersDetail() *UsersDetail {
	return new(UsersDetail)
}

func (u *UsersDetail) GetUsersDetail() (*UsersDetail, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("users_detail").Load(u)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return u, nil
}
