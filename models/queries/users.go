package queries

import (
	"time"

	"github.com/gocraft/dbr"
)

type (
	User struct {
		Id          int       `json:"id"`
		Name        string    `json:"name"`
		PhoneNumber string    `json:"phone_number"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	Users []User
)

func NewUser() *User {
	return new(User)
}

func (u *User) GetUser(db *dbr.Session, userId int) (*User, error) {
	query := "select * from users where id=?"
	if _, err := db.SelectBySql(query, userId).Load(u); err != nil {
		return nil, err
	}
	return u, nil
}

func NewUsers() *Users {
	return new(Users)
}

func (u *Users) GetUsers(db *dbr.Session) (*Users, error) {
	query := "select * from users"
	if _, err := db.SelectBySql(query).Load(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) InsertUser(tx *dbr.Tx, name, phoneNumber string) error {
	query := "insert into users (name,phone_number) values (?,?)"
	_, err := tx.InsertBySql(query, name, phoneNumber).Exec()
	return err
}

func (u *User) UpdateUserName(tx *dbr.Tx, id int, name string) error {
	query := "update users set name=? where id=?"
	_, err := tx.UpdateBySql(query, name, id).Exec()
	return err
}

func (u *User) UpdateUserPhoneNumber(tx *dbr.Tx, id int, phoneNumber string) error {
	query := "update users set phone_number=? where id=?"
	_, err := tx.UpdateBySql(query, phoneNumber, id).Exec()
	return err
}

func (u *User) DeleteUser(tx *dbr.Tx, id int) error {
	query := "delete from users where id=?"
	_, err := tx.DeleteBySql(query, id).Exec()
	return err
}

func (u *Users) SelectAssignedUsers(db *dbr.Session, ids []int) (*Users, error) {
	query := "select * from users where id in ?"
	if row, err := db.SelectBySql(query, ids).Load(u); row == 0 {
		return u, nil
	} else if err != nil {
		return nil, err
	}
	return u, nil
}
