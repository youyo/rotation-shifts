package queries

import (
	"errors"
	"time"

	"github.com/gocraft/dbr"
)

type (
	RotationDetail struct {
		RotationId int       `json:"id"`
		OrderId    int       `json:"order_id"`
		ShiftId    int       `json:"shift_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
	RotationDetails []RotationDetail
)

func NewRotationDetail() *RotationDetail {
	return new(RotationDetail)
}

func NewRotationDetails() *RotationDetails {
	return new(RotationDetails)
}

func (r *RotationDetails) SelectRotationDetailsWithOrderId(db *dbr.Session, rotationId, orderId int) (*RotationDetails, error) {
	query := "select * from rotation_details where rotation_id=? and order_id=?"
	if row, err := db.SelectBySql(query, rotationId, orderId).Load(r); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}
	return r, nil
}

func (r *RotationDetails) SelectRotationDetails(db *dbr.Session, rotationId int) (*RotationDetails, error) {
	query := "select * from rotation_details where rotation_id=?"
	if row, err := db.SelectBySql(query, rotationId).Load(r); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}
	return r, nil
}

func (r *RotationDetail) PutRotationDetail(tx *dbr.Tx, rotationId, orderId, shiftId int) error {
	query := "insert into rotation_details (`rotation_id`, `order_id`, `shift_id`) values (?,?,?) on duplicate key update `order_id`=?, `shift_id`=?"
	_, err := tx.InsertBySql(query, rotationId, orderId, shiftId, orderId, shiftId).Exec()
	return err
}

func (r *RotationDetail) DeleteRotationDetail(tx *dbr.Tx, rotationId int) error {
	query := "delete from rotation_details where rotation_id=?"
	_, err := tx.DeleteBySql(query, rotationId).Exec()
	return err
}

func (r *RotationDetail) DeleteRotationDetailWithOrderId(tx *dbr.Tx, rotationId, orderId int) error {
	query := "delete from rotation_details where rotation_id=? and order_id=?"
	_, err := tx.DeleteBySql(query, rotationId, orderId).Exec()
	return err
}

func (r *RotationDetail) DeleteRotationDetailWithOrderIdAndShiftId(tx *dbr.Tx, rotationId, orderId, shiftId int) error {
	query := "delete from rotation_details where rotation_id=? and order_id=? and shift_id=?"
	_, err := tx.DeleteBySql(query, rotationId, orderId, shiftId).Exec()
	return err
}

func (r *RotationDetail) SelectMaxOrderId(db *dbr.Session, rotationId int) (int, error) {
	var maxOrderId int
	query := "select max(order_id) as order_id from rotation_details where rotation_id=? group by rotation_id"
	if row, err := db.SelectBySql(query, rotationId).Load(&maxOrderId); err != nil {
		return 0, err
	} else if row == 0 {
		return 0, errors.New("No match records.")
	}
	return maxOrderId, nil
}

func (r *RotationDetail) SelectShiftIds(db *dbr.Session, rotationId, orderId int) ([]int, error) {
	var shiftIds []int

	query := "select shift_id from rotation_details where rotation_id=? and order_id=?"
	if row, err := db.SelectBySql(query, rotationId, orderId).Load(&shiftIds); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}
	return shiftIds, nil
}
