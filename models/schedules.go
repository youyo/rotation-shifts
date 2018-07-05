package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/k0kubun/pp"
	"github.com/youyo/rotation-shifts/db"
)

type GetRotation struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetScheduleNow(rotationId int) (*UsersDetail, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	tx, err := c.Begin()
	if err != nil {
		return nil, err
	}

	rotation := &GetRotation{}
	pp.Println("start_date")
	if row, err := tx.Select("start_date").From("rotations").Where("id=?", rotationId).Load(rotation); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}

	pp.Println("time.Parse")
	startDate, err := time.Parse("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", rotation.StartDate))
	if err != nil {
		return nil, err
	}
	pp.Println(startDate)

	today := time.Now()
	//today := time.Date(2018, time.August, 7, 19, 0, 0, 0, time.Local)
	pp.Println(today)
	now := today.Format("15")
	today.Format("2006-01-02")
	pp.Println(today)

	duration := today.Sub(startDate)
	days := int(duration.Hours()) / 24
	pp.Println("days: ", days)
	amari := int(duration.Hours()) % 24
	pp.Println(amari)
	if amari > 0 {
		days++
	}
	pp.Println("days: ", days)

	rotationDetail := &RotationDetail{}
	query := fmt.Sprintf("select rotation_id, max(order_id) as order_id from rotations_detail where rotation_id=%d group by rotation_id", rotationId)
	if row, err := tx.SelectBySql(query).Load(rotationDetail); err != nil {
		return nil, err
	} else if row == 0 {
		return nil, errors.New("No match records.")
	}

	pp.Println(rotationDetail)
	totalDays := rotationDetail.OrderId * 7
	pp.Println(totalDays)
	pp.Println(days)

	week := ""
	if days < totalDays {
		pp.Println("今週")
		switch days {
		case 1:
			pp.Println("月")
			week = "monday_"
		case 2:
			pp.Println("火")
			week = "tuesday_"
		case 3:
			pp.Println("水")
			week = "wendesday_"
		case 4:
			pp.Println("木")
			week = "thursday_"
		case 5:
			pp.Println("金")
			week = "friday_"
		case 6:
			pp.Println("土")
			week = "saturday_"
		case 7:
			pp.Println("日")
			week = "sunday_"
		default:
			return nil, errors.New("いつ?")
		}
		pp.Println("order: ", 1)

		// get shift_id
		rotationsDetail := &RotationsDetail{}
		if row, err := tx.Select("shift_id").From("rotations_detail").Where("rotation_id=? and order_id=1", rotationId).Load(rotationsDetail); err != nil {
			return nil, err
		} else if row == 0 {
			return nil, errors.New("No match records.")
		}
		pp.Println(rotationsDetail)
		pp.Println(now)
		pp.Println(week)

		// get user_ids
		shiftIds := []string{}
		for _, v := range *rotationsDetail {
			shiftIds = append(shiftIds, strconv.Itoa(v.ShiftId))
		}
		pp.Println(shiftIds)
		shifts := &Shifts{}
		query = fmt.Sprintf("select user_id from shifts where %s=1 and id in (%s)", week+string(now)+"00", strings.Join(shiftIds, ","))
		pp.Println(query)
		if row, err := tx.SelectBySql(query).Load(shifts); err != nil {
			return nil, err
		} else if row == 0 {
			return nil, errors.New("No match records.")
		}
		pp.Println(shifts)

		// get users_detail
		userIds := []string{}
		for _, v := range *shifts {
			userIds = append(userIds, strconv.Itoa(v.UserId))
		}

		usersDetail := &UsersDetail{}
		query = fmt.Sprintf("select * from users_detail where user_id in (%s)", strings.Join(userIds, ","))
		pp.Println(query)
		if row, err := tx.SelectBySql(query).Load(usersDetail); err != nil {
			return nil, err
		} else if row == 0 {
			return nil, errors.New("No match records.")
		}
		pp.Println(usersDetail)
		return usersDetail, nil

	} else if days > totalDays {
		pp.Println("今後")
		remainder := days%7 + 1
		pp.Println(remainder)
		switch remainder {
		case 1:
			pp.Println("月")
		case 2:
			pp.Println("火")
		case 3:
			pp.Println("水")
		case 4:
			pp.Println("木")
		case 5:
			pp.Println("金")
		case 6:
			pp.Println("土")
		case 7:
			pp.Println("日")
		default:
			return nil, errors.New("いつ?")
		}
		quotient := days / totalDays
		pp.Println(quotient)
		order := (quotient + 1) % rotationDetail.OrderId
		pp.Println("order: ", order)
	} else {
		pp.Println("いつ？")
		return nil, errors.New("いつ?")
	}

	tx.Commit()
	return nil, nil
}
