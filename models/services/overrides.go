package services

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/youyo/rotation-shifts/db"
	"github.com/youyo/rotation-shifts/models/queries"
)

func GetOverride(id int) (int, *queries.Override, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	u := queries.NewOverride()
	override, err := u.GetOverride(conn, id)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, override, nil
}

func GetOverrides() (int, *queries.Overrides, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	u := queries.NewOverrides()
	pp.Println("queries.NewOverrides")
	overrides, err := u.GetOverrides(conn)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, overrides, nil
}

func CreateOverride(name string, date queries.DATE, hour queries.HOUR, rotationId, userId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	u := queries.NewOverride()
	err = u.InsertOverride(tx, name, date, hour, rotationId, userId)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusCreated, "created", nil
}

func DeleteOverride(id int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	u := queries.NewOverride()
	err = u.DeleteOverride(tx, id)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}
