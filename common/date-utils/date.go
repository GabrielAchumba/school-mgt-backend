package dateutils

import (
	"strings"
	"time"
)

func GetDate(dateValue string) time.Time {

	const shortForm = "2006-01-02"
	_d := strings.Split(dateValue, "/")

	// reorder date to year:month:day
	if len(_d[0]) < 2 {
		_d[0] = "0" + _d[0]
	}
	if len(_d[1]) < 2 {
		_d[1] = "0" + _d[1]
	}

	year := _d[2]
	_d[2] = _d[0]
	_d[0] = year

	newDate := strings.Join(_d, "-")
	result, _ := time.Parse(shortForm, newDate)

	return result
}
