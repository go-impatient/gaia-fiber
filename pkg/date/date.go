package date

import (
	"time"
)

// SQLDateToTime converts SQL date to time.Time.
func SQLDateToTime(d string) (time.Time, error) {
	return time.Parse("2006-01-02", d)
}

// SQLDatetimeToTime converts SQL date to time.Time.
func SQLDatetimeToTime(d string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", d)
}

// TimeToSQLDate returns a time to SQL date format (YYYY-MM-DD).
func TimeToSQLDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// TimeToSQLDatetime returns a time to SQL datetime format (YYYY-MM-DD HH:MM:SS).
func TimeToSQLDatetime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
