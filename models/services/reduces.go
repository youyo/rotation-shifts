package services

import (
	"net/http"

	"github.com/youyo/rotation-shifts/db"
	"github.com/youyo/rotation-shifts/models/queries"
)

func GetReduce(id int) (int, *queries.Reduce, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	u := queries.NewReduce()
	override, err := u.GetReduce(conn, id)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, override, nil
}

func GetReduces() (int, *queries.Reduces, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	u := queries.NewReduces()
	overrides, err := u.GetReduces(conn)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, overrides, nil
}

func CreateReduce(name string, date queries.DATE, hour queries.HOUR, rotationId, userId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	u := queries.NewReduce()
	err = u.InsertReduce(tx, name, date, hour, rotationId, userId)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusCreated, "created", nil
}

func DeleteReduce(id int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	u := queries.NewReduce()
	err = u.DeleteReduce(tx, id)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}
