package services

import (
	"net/http"

	"github.com/youyo/rotation-shifts/db"
	"github.com/youyo/rotation-shifts/models/queries"
)

func GetShifts() (int, *queries.Shifts, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	s := queries.NewShifts()
	result, err := s.GetShifts(conn)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func GetShift(shiftId int) (int, *queries.Shift, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	s := queries.NewShift()
	result, err := s.GetShift(conn, shiftId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func PostShift(name string, userId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	s := queries.NewShift()
	if err = s.InsertShift(tx, name, userId); err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx.Commit()

	return http.StatusCreated, "created", nil
}

func PatchShift(shiftId int, shiftDetail *queries.ShiftDetail) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	s := queries.NewShift()
	if err = s.PatchShift(tx, shiftId, shiftDetail); err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx.Commit()

	return http.StatusOK, "updated", nil
}

func DeleteShift(shiftId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	s := queries.NewShift()
	if err = s.DeleteShift(tx, shiftId); err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}
