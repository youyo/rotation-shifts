package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/youyo/rotation-shifts/db"
)

type RotationDetail struct {
	RotationId int       `json:"id"`
	Name       string    `json:"name"`
	OrderId    int       `json:"order_id"`
	ShiftId    int       `json:"shift_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewRotationDetail() *RotationDetail {
	return new(RotationDetail)
}

func (r *RotationDetail) GetRotationDetail(rotationId, orderId int) (*RotationDetail, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").
		From("rotations_detail").
		Where("rotation_id=? and order_id=?", rotationId, orderId).
		Load(r)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return r, nil
}

func (r *RotationDetail) PutRotationDetail(rotationId, orderId, shiftId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("insert into rotations_detail (`rotation_id`, `order_id`, `shift_id`) values (%d,%d,%d) on duplicate key update `order_id`=%d, `shift_id`=%d", rotationId, orderId, shiftId, orderId, shiftId)
	if _, err := tx.InsertBySql(query).Exec(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *RotationDetail) UpdateRotationDetailOrderId(rotationId, orderIdint int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Update("rotations_detail").
		Set("order_id", orderIdint).
		Where("rotation_id=?", rotationId).
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

func (r *RotationDetail) UpdateRotationDetailShiftId(rotationId, shiftId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Update("rotations_detail").
		Set("shift_id", shiftId).
		Where("rotation_id=?", rotationId).
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

func (r *RotationDetail) DeleteRotationDetail(rotationId, orderId int) error {
	c, err := db.NewConnection()
	if err != nil {
		return err
	}
	tx, err := c.Begin()
	if err != nil {
		return err
	}
	if _, err = tx.DeleteFrom("rotations_detail").Where("rotation_id=? and order_id=?", rotationId, orderId).Exec(); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

type RotationsDetail []RotationDetail

func NewRotationsDetail() *RotationsDetail {
	return new(RotationsDetail)
}

func (r *RotationsDetail) GetRotationsDetail() (*RotationsDetail, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	row, err := c.Select("*").From("rotations_detail").Load(r)
	if err != nil {
		return nil, err
	}
	if row == 0 {
		return nil, errors.New("No match records.")
	}
	return r, nil
}
