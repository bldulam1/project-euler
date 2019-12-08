package problem19

import (
	"time"
)

// GetDay returns the day of the week
func GetDay(year int, month uint8, day uint8) time.Weekday {

	return time.Date(
		year,
		time.Month(month),
		int(day),
		0, 0, 0, 0,
		time.Local).Weekday()

	// fmt.Println(date.Weekday(), date)
}
