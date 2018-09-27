package services

import (
	"net/http"
	"strconv"
	"time"

	"github.com/youyo/rotation-shifts/db"
)

type (
	CalendarEvent struct {
		Title string `json:"title"`
		Start string `json:"start"`
	}
	CalendarEvents []CalendarEvent
)

func NewCalendarEvents() *CalendarEvents {
	return new(CalendarEvents)
}

func (c *CalendarEvents) GetCalendarMonthly(rotationId int, yearMonth string) (int, *CalendarEvents, error) {
	// get location
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return 0, nil, err
	}

	// first day
	firstDay := 1

	// get last day
	t, err := time.ParseInLocation("2006-01", yearMonth, loc)
	if err != nil {
		return 0, nil, err
	}
	lastDayTimeFormat := t.AddDate(0, 3, 0).AddDate(0, 0, -1)
	lastDay, err := strconv.Atoi(lastDayTimeFormat.Format("02"))
	if err != nil {
		return 0, nil, err
	}

	// all times
	var days []string
	for day := firstDay; day < 10; day++ {
		for h := 0; h < 10; h++ {
			s := yearMonth + "-0" + strconv.Itoa(day) + "-0" + strconv.Itoa(h)
			days = append(days, s)
		}
		for h := 10; h < 24; h++ {
			s := yearMonth + "-0" + strconv.Itoa(day) + "-" + strconv.Itoa(h)
			days = append(days, s)
		}
	}
	for day := 10; day <= lastDay; day++ {
		for h := 0; h < 10; h++ {
			s := yearMonth + "-" + strconv.Itoa(day) + "-0" + strconv.Itoa(h)
			days = append(days, s)
		}
		for h := 10; h < 24; h++ {
			s := yearMonth + "-" + strconv.Itoa(day) + "-" + strconv.Itoa(h)
			days = append(days, s)
		}
	}

	conn, err := db.NewConnection()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	for _, day := range days {
		s, err := NewSchedule(rotationId, day)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		statusCode, resp, err := s.QuerySchedules(conn)
		if err != nil {
			return statusCode, nil, err
		}
		start, err := time.Parse("2006-01-02-15", day)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		for _, user := range *resp {
			cal := CalendarEvent{
				Title: user.Name,
				Start: start.Format("2006-01-02T15:04:05"),
			}
			*c = append(*c, cal)
		}
	}

	return http.StatusOK, c, nil
}
