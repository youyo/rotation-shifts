package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gocraft/dbr"
	"github.com/youyo/rotation-shifts/db"
	"github.com/youyo/rotation-shifts/models/queries"
)

const (
	timeZone       string = "Asia/Tokyo"
	dateTimeFormat string = "2006-01-02-15-04"
)

type Schedule struct {
	RotationId   int
	Date         time.Time
	Hour         string
	Week         string
	OverrideHour time.Time
}

func NewSchedule(rotationId int, dateTime string) (*Schedule, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return nil, err
	}

	var date time.Time
	var hour string

	if dateTime == "now" {
		date = time.Now().In(loc)
		hour = date.Format("15")
	} else {
		date, err = time.ParseInLocation("2006-01-02-15", dateTime, loc)
		if err != nil {
			return nil, err
		}
		hour = date.Format("15")
	}

	s := &Schedule{
		RotationId:   rotationId,
		Date:         date,
		Hour:         hour,
		OverrideHour: date,
	}
	return s, nil
}

func (s *Schedule) GetSchedule() (int, *queries.Users, error) {
	// database connection
	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return s.QuerySchedules(conn)
}

func (s *Schedule) QuerySchedules(conn *dbr.Session) (int, *queries.Users, error) {
	// check override shift
	overrides := queries.NewOverrides()
	overrodeUserIds, err := overrides.SelectOverrideAssignedUserIds(conn, s.RotationId, queries.DATE{Time: s.Date}, queries.HOUR{Time: s.OverrideHour})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if len(overrodeUserIds) != 0 {
		users := queries.NewUsers()
		_, err = users.SelectAssignedUsers(conn, overrodeUserIds)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		return http.StatusOK, users, nil
	}

	// 通常フロー
	r := queries.NewRotation()
	rd, err := r.GetRotation(conn, s.RotationId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	startDate, err := time.ParseInLocation("2006-01-02 15:04:05 -0700 MST", fmt.Sprintf("%s", rd.StartDate), loc)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// スケジュール開始日と取得したい日の差分
	duration := s.Date.Sub(startDate)
	if duration < 0 {
		users := queries.NewUsers()
		return http.StatusOK, users, nil
	}

	// duration の総時間から日数を取得
	days := int(duration.Hours()) / 24

	// 最大 order_id 取得
	// 週数
	rotationDetail := queries.NewRotationDetail()
	maxOrderId, err := rotationDetail.SelectMaxOrderId(conn, s.RotationId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// シフトの全日数を取得
	totalDays := maxOrderId * 7

	// 曜日取得
	switch days % 7 {
	case 0:
		s.Week = "monday_"
	case 1:
		s.Week = "tuesday_"
	case 2:
		s.Week = "wednesday_"
	case 3:
		s.Week = "thursday_"
	case 4:
		s.Week = "friday_"
	case 5:
		s.Week = "saturday_"
	case 6:
		s.Week = "sunday_"
	default:
		return http.StatusInternalServerError, nil, errors.New("Unmatch days")
	}

	// order_id 確定
	for days > totalDays {
		days = days - totalDays
	}
	orderId := 0
	for orderIdStart := 1; orderIdStart <= maxOrderId; orderIdStart++ {
		if days <= orderIdStart*7 {
			orderId = orderIdStart
			break
		}
	}

	// shift_ids 取得
	shiftIds, err := rotationDetail.SelectShiftIds(conn, s.RotationId, orderId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// 有効な user_id 取得
	shift := queries.NewShift()
	userIds, err := shift.SelectAssignedUserIds(conn, s.Week, s.Hour, shiftIds)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// ユーザー情報取得
	users := queries.NewUsers()
	users, err = users.SelectAssignedUsers(conn, userIds)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, users, nil
}
