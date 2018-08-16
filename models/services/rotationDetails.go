package services

import (
	"net/http"

	"github.com/youyo/rotation-shifts/db"
	"github.com/youyo/rotation-shifts/models/queries"
)

func GetRotationDetail(rotationId, orderId int) (int, *queries.RotationDetails, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	r := queries.NewRotationDetails()
	result, err := r.SelectRotationDetailsWithOrderId(conn, rotationId, orderId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func GetRotationDetails(rotationId int) (int, *queries.RotationDetails, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	r := queries.NewRotationDetails()
	result, err := r.SelectRotationDetails(conn, rotationId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, result, nil
}

func PutRotationDetail(rotationId, orderId, shiftId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	r := queries.NewRotationDetail()
	if err := r.PutRotationDetail(tx, rotationId, orderId, shiftId); err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusCreated, "created", nil
}

func DeleteRotationDetail(rotationId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	r := queries.NewRotationDetail()
	if err := r.DeleteRotationDetail(tx, rotationId); err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}

func DeleteRotationDetailWithOrderId(rotationId, orderId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	r := queries.NewRotationDetail()
	if err := r.DeleteRotationDetailWithOrderId(tx, rotationId, orderId); err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}

func DeleteRotationDetailWithOrderIdAndShiftId(rotationId, orderId, shiftId int) (int, string, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx, err := conn.Begin()
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	defer tx.RollbackUnlessCommitted()

	r := queries.NewRotationDetail()
	if err := r.DeleteRotationDetailWithOrderIdAndShiftId(tx, rotationId, orderId, shiftId); err != nil {
		return http.StatusInternalServerError, "", err
	}

	tx.Commit()

	return http.StatusOK, "deleted", nil
}
