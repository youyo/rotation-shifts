package services

import (
	"net/http"

	"github.com/youyo/rotation-shifts/db"
	"github.com/youyo/rotation-shifts/models/queries"
)

func GetUser(userId int) (int, *queries.User, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	u := queries.NewUser()
	user, err := u.GetUser(conn, userId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, user, nil
}

func GetUsers() (int, *queries.Users, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	u := queries.NewUsers()
	users, err := u.GetUsers(conn)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, users, nil
}

func CreateUser(name, phoneNumber string) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	u := queries.NewUser()
	err = u.InsertUser(tx, name, phoneNumber)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusCreated, "created", nil
}

func DeleteUser(userId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	u := queries.NewUser()
	err = u.DeleteUser(tx, userId)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}
