package queries

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gocraft/dbr"
)

type (
	Shift struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		UserId int    `json:"user_id"`
		ShiftDetail
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	ShiftDetail struct {
		Monday0000    int `json:"monday_0000" db:"monday_0000"`
		Monday0100    int `json:"monday_0100" db:"monday_0100"`
		Monday0200    int `json:"monday_0200" db:"monday_0200"`
		Monday0300    int `json:"monday_0300" db:"monday_0300"`
		Monday0400    int `json:"monday_0400" db:"monday_0400"`
		Monday0500    int `json:"monday_0500" db:"monday_0500"`
		Monday0600    int `json:"monday_0600" db:"monday_0600"`
		Monday0700    int `json:"monday_0700" db:"monday_0700"`
		Monday0800    int `json:"monday_0800" db:"monday_0800"`
		Monday0900    int `json:"monday_0900" db:"monday_0900"`
		Monday1000    int `json:"monday_1000" db:"monday_1000"`
		Monday1100    int `json:"monday_1100" db:"monday_1100"`
		Monday1200    int `json:"monday_1200" db:"monday_1200"`
		Monday1300    int `json:"monday_1300" db:"monday_1300"`
		Monday1400    int `json:"monday_1400" db:"monday_1400"`
		Monday1500    int `json:"monday_1500" db:"monday_1500"`
		Monday1600    int `json:"monday_1600" db:"monday_1600"`
		Monday1700    int `json:"monday_1700" db:"monday_1700"`
		Monday1800    int `json:"monday_1800" db:"monday_1800"`
		Monday1900    int `json:"monday_1900" db:"monday_1900"`
		Monday2000    int `json:"monday_2000" db:"monday_2000"`
		Monday2100    int `json:"monday_2100" db:"monday_2100"`
		Monday2200    int `json:"monday_2200" db:"monday_2200"`
		Monday2300    int `json:"monday_2300" db:"monday_2300"`
		Tuesday0000   int `json:"tuesday_0000" db:"tuesday_0000"`
		Tuesday0100   int `json:"tuesday_0100" db:"tuesday_0100"`
		Tuesday0200   int `json:"tuesday_0200" db:"tuesday_0200"`
		Tuesday0300   int `json:"tuesday_0300" db:"tuesday_0300"`
		Tuesday0400   int `json:"tuesday_0400" db:"tuesday_0400"`
		Tuesday0500   int `json:"tuesday_0500" db:"tuesday_0500"`
		Tuesday0600   int `json:"tuesday_0600" db:"tuesday_0600"`
		Tuesday0700   int `json:"tuesday_0700" db:"tuesday_0700"`
		Tuesday0800   int `json:"tuesday_0800" db:"tuesday_0800"`
		Tuesday0900   int `json:"tuesday_0900" db:"tuesday_0900"`
		Tuesday1000   int `json:"tuesday_1000" db:"tuesday_1000"`
		Tuesday1100   int `json:"tuesday_1100" db:"tuesday_1100"`
		Tuesday1200   int `json:"tuesday_1200" db:"tuesday_1200"`
		Tuesday1300   int `json:"tuesday_1300" db:"tuesday_1300"`
		Tuesday1400   int `json:"tuesday_1400" db:"tuesday_1400"`
		Tuesday1500   int `json:"tuesday_1500" db:"tuesday_1500"`
		Tuesday1600   int `json:"tuesday_1600" db:"tuesday_1600"`
		Tuesday1700   int `json:"tuesday_1700" db:"tuesday_1700"`
		Tuesday1800   int `json:"tuesday_1800" db:"tuesday_1800"`
		Tuesday1900   int `json:"tuesday_1900" db:"tuesday_1900"`
		Tuesday2000   int `json:"tuesday_2000" db:"tuesday_2000"`
		Tuesday2100   int `json:"tuesday_2100" db:"tuesday_2100"`
		Tuesday2200   int `json:"tuesday_2200" db:"tuesday_2200"`
		Tuesday2300   int `json:"tuesday_2300" db:"tuesday_2300"`
		Wednesday0000 int `json:"wednesday_0000" db:"wednesday_0000"`
		Wednesday0100 int `json:"wednesday_0100" db:"wednesday_0100"`
		Wednesday0200 int `json:"wednesday_0200" db:"wednesday_0200"`
		Wednesday0300 int `json:"wednesday_0300" db:"wednesday_0300"`
		Wednesday0400 int `json:"wednesday_0400" db:"wednesday_0400"`
		Wednesday0500 int `json:"wednesday_0500" db:"wednesday_0500"`
		Wednesday0600 int `json:"wednesday_0600" db:"wednesday_0600"`
		Wednesday0700 int `json:"wednesday_0700" db:"wednesday_0700"`
		Wednesday0800 int `json:"wednesday_0800" db:"wednesday_0800"`
		Wednesday0900 int `json:"wednesday_0900" db:"wednesday_0900"`
		Wednesday1000 int `json:"wednesday_1000" db:"wednesday_1000"`
		Wednesday1100 int `json:"wednesday_1100" db:"wednesday_1100"`
		Wednesday1200 int `json:"wednesday_1200" db:"wednesday_1200"`
		Wednesday1300 int `json:"wednesday_1300" db:"wednesday_1300"`
		Wednesday1400 int `json:"wednesday_1400" db:"wednesday_1400"`
		Wednesday1500 int `json:"wednesday_1500" db:"wednesday_1500"`
		Wednesday1600 int `json:"wednesday_1600" db:"wednesday_1600"`
		Wednesday1700 int `json:"wednesday_1700" db:"wednesday_1700"`
		Wednesday1800 int `json:"wednesday_1800" db:"wednesday_1800"`
		Wednesday1900 int `json:"wednesday_1900" db:"wednesday_1900"`
		Wednesday2000 int `json:"wednesday_2000" db:"wednesday_2000"`
		Wednesday2100 int `json:"wednesday_2100" db:"wednesday_2100"`
		Wednesday2200 int `json:"wednesday_2200" db:"wednesday_2200"`
		Wednesday2300 int `json:"wednesday_2300" db:"wednesday_2300"`
		Thursday0000  int `json:"thursday_0000" db:"thursday_0000"`
		Thursday0100  int `json:"thursday_0100" db:"thursday_0100"`
		Thursday0200  int `json:"thursday_0200" db:"thursday_0200"`
		Thursday0300  int `json:"thursday_0300" db:"thursday_0300"`
		Thursday0400  int `json:"thursday_0400" db:"thursday_0400"`
		Thursday0500  int `json:"thursday_0500" db:"thursday_0500"`
		Thursday0600  int `json:"thursday_0600" db:"thursday_0600"`
		Thursday0700  int `json:"thursday_0700" db:"thursday_0700"`
		Thursday0800  int `json:"thursday_0800" db:"thursday_0800"`
		Thursday0900  int `json:"thursday_0900" db:"thursday_0900"`
		Thursday1000  int `json:"thursday_1000" db:"thursday_1000"`
		Thursday1100  int `json:"thursday_1100" db:"thursday_1100"`
		Thursday1200  int `json:"thursday_1200" db:"thursday_1200"`
		Thursday1300  int `json:"thursday_1300" db:"thursday_1300"`
		Thursday1400  int `json:"thursday_1400" db:"thursday_1400"`
		Thursday1500  int `json:"thursday_1500" db:"thursday_1500"`
		Thursday1600  int `json:"thursday_1600" db:"thursday_1600"`
		Thursday1700  int `json:"thursday_1700" db:"thursday_1700"`
		Thursday1800  int `json:"thursday_1800" db:"thursday_1800"`
		Thursday1900  int `json:"thursday_1900" db:"thursday_1900"`
		Thursday2000  int `json:"thursday_2000" db:"thursday_2000"`
		Thursday2100  int `json:"thursday_2100" db:"thursday_2100"`
		Thursday2200  int `json:"thursday_2200" db:"thursday_2200"`
		Thursday2300  int `json:"thursday_2300" db:"thursday_2300"`
		Friday0000    int `json:"friday_0000" db:"friday_0000"`
		Friday0100    int `json:"friday_0100" db:"friday_0100"`
		Friday0200    int `json:"friday_0200" db:"friday_0200"`
		Friday0300    int `json:"friday_0300" db:"friday_0300"`
		Friday0400    int `json:"friday_0400" db:"friday_0400"`
		Friday0500    int `json:"friday_0500" db:"friday_0500"`
		Friday0600    int `json:"friday_0600" db:"friday_0600"`
		Friday0700    int `json:"friday_0700" db:"friday_0700"`
		Friday0800    int `json:"friday_0800" db:"friday_0800"`
		Friday0900    int `json:"friday_0900" db:"friday_0900"`
		Friday1000    int `json:"friday_1000" db:"friday_1000"`
		Friday1100    int `json:"friday_1100" db:"friday_1100"`
		Friday1200    int `json:"friday_1200" db:"friday_1200"`
		Friday1300    int `json:"friday_1300" db:"friday_1300"`
		Friday1400    int `json:"friday_1400" db:"friday_1400"`
		Friday1500    int `json:"friday_1500" db:"friday_1500"`
		Friday1600    int `json:"friday_1600" db:"friday_1600"`
		Friday1700    int `json:"friday_1700" db:"friday_1700"`
		Friday1800    int `json:"friday_1800" db:"friday_1800"`
		Friday1900    int `json:"friday_1900" db:"friday_1900"`
		Friday2000    int `json:"friday_2000" db:"friday_2000"`
		Friday2100    int `json:"friday_2100" db:"friday_2100"`
		Friday2200    int `json:"friday_2200" db:"friday_2200"`
		Friday2300    int `json:"friday_2300" db:"friday_2300"`
		Saturday0000  int `json:"saturday_0000" db:"saturday_0000"`
		Saturday0100  int `json:"saturday_0100" db:"saturday_0100"`
		Saturday0200  int `json:"saturday_0200" db:"saturday_0200"`
		Saturday0300  int `json:"saturday_0300" db:"saturday_0300"`
		Saturday0400  int `json:"saturday_0400" db:"saturday_0400"`
		Saturday0500  int `json:"saturday_0500" db:"saturday_0500"`
		Saturday0600  int `json:"saturday_0600" db:"saturday_0600"`
		Saturday0700  int `json:"saturday_0700" db:"saturday_0700"`
		Saturday0800  int `json:"saturday_0800" db:"saturday_0800"`
		Saturday0900  int `json:"saturday_0900" db:"saturday_0900"`
		Saturday1000  int `json:"saturday_1000" db:"saturday_1000"`
		Saturday1100  int `json:"saturday_1100" db:"saturday_1100"`
		Saturday1200  int `json:"saturday_1200" db:"saturday_1200"`
		Saturday1300  int `json:"saturday_1300" db:"saturday_1300"`
		Saturday1400  int `json:"saturday_1400" db:"saturday_1400"`
		Saturday1500  int `json:"saturday_1500" db:"saturday_1500"`
		Saturday1600  int `json:"saturday_1600" db:"saturday_1600"`
		Saturday1700  int `json:"saturday_1700" db:"saturday_1700"`
		Saturday1800  int `json:"saturday_1800" db:"saturday_1800"`
		Saturday1900  int `json:"saturday_1900" db:"saturday_1900"`
		Saturday2000  int `json:"saturday_2000" db:"saturday_2000"`
		Saturday2100  int `json:"saturday_2100" db:"saturday_2100"`
		Saturday2200  int `json:"saturday_2200" db:"saturday_2200"`
		Saturday2300  int `json:"saturday_2300" db:"saturday_2300"`
		Sunday0000    int `json:"sunday_0000" db:"sunday_0000"`
		Sunday0100    int `json:"sunday_0100" db:"sunday_0100"`
		Sunday0200    int `json:"sunday_0200" db:"sunday_0200"`
		Sunday0300    int `json:"sunday_0300" db:"sunday_0300"`
		Sunday0400    int `json:"sunday_0400" db:"sunday_0400"`
		Sunday0500    int `json:"sunday_0500" db:"sunday_0500"`
		Sunday0600    int `json:"sunday_0600" db:"sunday_0600"`
		Sunday0700    int `json:"sunday_0700" db:"sunday_0700"`
		Sunday0800    int `json:"sunday_0800" db:"sunday_0800"`
		Sunday0900    int `json:"sunday_0900" db:"sunday_0900"`
		Sunday1000    int `json:"sunday_1000" db:"sunday_1000"`
		Sunday1100    int `json:"sunday_1100" db:"sunday_1100"`
		Sunday1200    int `json:"sunday_1200" db:"sunday_1200"`
		Sunday1300    int `json:"sunday_1300" db:"sunday_1300"`
		Sunday1400    int `json:"sunday_1400" db:"sunday_1400"`
		Sunday1500    int `json:"sunday_1500" db:"sunday_1500"`
		Sunday1600    int `json:"sunday_1600" db:"sunday_1600"`
		Sunday1700    int `json:"sunday_1700" db:"sunday_1700"`
		Sunday1800    int `json:"sunday_1800" db:"sunday_1800"`
		Sunday1900    int `json:"sunday_1900" db:"sunday_1900"`
		Sunday2000    int `json:"sunday_2000" db:"sunday_2000"`
		Sunday2100    int `json:"sunday_2100" db:"sunday_2100"`
		Sunday2200    int `json:"sunday_2200" db:"sunday_2200"`
		Sunday2300    int `json:"sunday_2300" db:"sunday_2300"`
	}

	Shifts []Shift
)

func NewShift() *Shift {
	return new(Shift)
}

func NewShifts() *Shifts {
	return new(Shifts)
}

func NewShiftDetail() *ShiftDetail {
	return new(ShiftDetail)
}

func (s *Shift) InsertShift(tx *dbr.Tx, name string, userId int) error {
	query := "insert into shifts (name,user_id) values (?,?)"
	_, err := tx.InsertBySql(query, name, userId).Exec()
	return err
}

func (s *Shift) PatchShift(tx *dbr.Tx, shiftId int, shiftDetail *ShiftDetail) error {
	m := make(map[string]interface{})
	b, _ := json.Marshal(shiftDetail)
	json.Unmarshal(b, &m)

	_, err := tx.Update("shifts").SetMap(m).Where("id=?", shiftId).Exec()
	return err
}

func (s *Shift) DeleteShift(tx *dbr.Tx, shiftId int) error {
	query := "delete from shifts where id=?"
	_, err := tx.DeleteBySql(query, shiftId).Exec()
	return err
}

func (s *Shifts) GetShifts(db *dbr.Session) (*Shifts, error) {
	query := "select * from shifts"
	if _, err := db.SelectBySql(query).Load(s); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Shift) GetShift(db *dbr.Session, shiftId int) (*Shift, error) {
	query := "select * from shifts where id=?"
	if _, err := db.SelectBySql(query, shiftId).Load(s); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Shift) SelectAssignedUserIds(db *dbr.Session, week, hour string, shiftIds []int) ([]int, error) {
	var userIds []int
	query := fmt.Sprintf("select user_id from shifts where `%s`=1 and id in ?", week+hour+"00")
	if _, err := db.SelectBySql(query, shiftIds).Load(&userIds); err != nil {
		return nil, err
	}
	return userIds, nil
}
