package services

import (
	"net/http"

	"github.com/youyo/rotation-shifts/db"
	"github.com/youyo/rotation-shifts/models/queries"
)

func GetRotation(rotationId int) (int, *queries.Rotation, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	s := queries.NewRotation()
	result, err := s.GetRotation(conn, rotationId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func GetRotations() (int, *queries.Rotations, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	s := queries.NewRotations()
	result, err := s.GetRotations(conn)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func PostRotation(name string, startDate queries.DATE) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	s := queries.NewRotation()
	if err := s.PostRotation(tx, name, startDate); err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusCreated, "created", nil
}

func DeleteRotation(rotationId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	s := queries.NewRotation()
	if err := s.DeleteRotation(tx, rotationId); err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}
