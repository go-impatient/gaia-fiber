package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestSQLDateToTime
func TestSQLDateToTime(t *testing.T) {
	d := "2019-10-23"
	actual, err := SQLDateToTime(d)
	expected := time.Date(2019, 10, 23, 0, 0, 0, 0, time.UTC)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

// TestSQLDatetimeToTime
func TestSQLDatetimeToTime(t *testing.T) {
	d := "2019-10-23 12:10:56"
	actual, err := SQLDatetimeToTime(d)
	expected := time.Date(2019, 10, 23, 12, 10, 56, 0, time.UTC)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

// TestTimeToSQLDate tests time to SQL date format (YYYY-MM-DD).
func TestTimeToSQLDate(t *testing.T) {
	d := time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC)
	expected := "2020-01-05"
	actual := TimeToSQLDate(d)

	assert.Equal(t, expected, actual)
}

// TestTimeToSQLDatetime tests time to SQL datetime format (YYYY-MM-DD HH:MM:SS).
func TestTimeToSQLDatetime(t *testing.T) {
	d := time.Date(2020, 1, 5, 23, 12, 9, 0, time.UTC)
	expected := "2020-01-05 23:12:09"
	actual := TimeToSQLDatetime(d)

	assert.Equal(t, expected, actual)
}
