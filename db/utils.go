package db

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const (
	ContextKeyUsername = "username"
	RepeatCountDefault = 15
)

type ProcessSequence struct {
	ID             uuid.UUID `db:"id" json:"process_id" validate:"required"`
	SequenceNumber int64     `db:"sequence_num" json:"sequence_num" validate:"required,gte=1"`
}

type WellsSlice struct {
	sort.IntSlice // Wells
	Heights       []int
}

func (s WellsSlice) Swap(i, j int) {
	s.IntSlice.Swap(i, j)
	s.Heights[i], s.Heights[j] = s.Heights[j], s.Heights[i]
}

func NewWellsSlice(n, m []int) *WellsSlice {
	if len(m) != len(n) {
		return nil
	}

	s := &WellsSlice{IntSlice: sort.IntSlice(n), Heights: m}
	sort.Sort(s)
	return s
}

func CalculateTimeInSeconds(t string) (totalTime int64, err error) {

	var hours, minutes, seconds int64
	timeArr := strings.Split(t, ":")
	if len(timeArr) != 3 {
		err = fmt.Errorf("time format isn't of the form HH:MM:SS")
		return 0, err
	}

	hours, err = parseIntRange(timeArr[0], "hours", 0, 24)
	if err != nil {
		return 0, err
	}

	minutes, err = parseIntRange(timeArr[1], "minutes", 0, 59)
	if err != nil {
		return 0, err
	}

	seconds, err = parseIntRange(timeArr[2], "seconds", 0, 59)
	if err != nil {
		return 0, err
	}

	totalTime = hours*60*60 + minutes*60 + seconds

	return
}

func parseIntRange(timeString, unit string, min, max int64) (value int64, err error) {
	value, err = strconv.ParseInt(timeString, 10, 64)
	if err != nil || value > max || value < min {
		err = fmt.Errorf("please check %v format, valid range: [%d,%d]", unit, min, max)
		return 0, err
	}
	return
}
