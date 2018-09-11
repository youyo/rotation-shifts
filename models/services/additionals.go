package services

import (
	"net/http"

	"github.com/youyo/rotation-shifts/db"
	"github.com/youyo/rotation-shifts/models/queries"
)

func GetAdditional(id int) (int, *queries.Additional, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	u := queries.NewAdditional()
	override, err := u.GetAdditional(conn, id)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, override, nil
}

func GetAdditionals() (int, *queries.Additionals, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	u := queries.NewAdditionals()
	overrides, err := u.GetAdditionals(conn)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, overrides, nil
}

func CreateAdditional(name string, date queries.DATE, hour queries.HOUR, rotationId, userId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	u := queries.NewAdditional()
	err = u.InsertAdditional(tx, name, date, hour, rotationId, userId)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusCreated, "created", nil
}

func DeleteAdditional(id int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	u := queries.NewAdditional()
	err = u.DeleteAdditional(tx, id)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}
