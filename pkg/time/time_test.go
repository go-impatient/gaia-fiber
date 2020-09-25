package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var format = "2006-01-02 15:04:05.000000000"

// TestYesterday tests Yesterday function.
func TestYesterday(t *testing.T) {
	actual := Yesterday()
	expected := time.Now().AddDate(0, 0, -1)

	assert.Equal(t, expected.Year(), actual.Year())
	assert.Equal(t, expected.Day(), actual.Day())
	assert.Equal(t, expected.Month(), actual.Month())
}

// TestStartOfDay tests StartOfDay function.
func TestStartOfDay(t *testing.T) {
	input, _ := time.Parse(format, "2015-01-01 12:34:00.000000000")
	actual := StartOfDay(input)
	expected, _ := time.Parse(format, "2015-01-01 00:00:00.000000000")

	assert.Equal(t, expected, actual)
}

// TestEndOfDay tests EndOfDay function.
func TestEndOfDay(t *testing.T) {
	input, _ := time.Parse(format, "2015-01-01 12:34:00.000000000")
	actual := EndOfDay(input)
	expected, _ := time.Parse(format, "2015-01-01 23:59:59.999999999")

	assert.Equal(t, expected, actual)
}