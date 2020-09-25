package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetHTTPInternalServerError tests HTTP response.
func TestGetHTTPResponse(t *testing.T) {
	actual := HTTPResponse{200, "Success", nil}
	expected := GetHTTPResponse(200, "Success", nil)
	assert.Equal(t, expected, actual)

	actual = HTTPResponse{404, "Not found", 15}
	expected = GetHTTPResponse(404, "Not found", 15)
	assert.Equal(t, expected, actual)
}

// TestGetHTTPInternalServerError tests HTTP internal server error response.
func TestGetHTTPInternalServerError(t *testing.T) {
	actual := HTTPResponse{500, "Database Error", nil}
	expected := GetHTTPInternalServerError("Database Error")
	assert.Equal(t, expected, actual)

	actual = HTTPResponse{500, "Internal Server Error", nil}
	expected = GetHTTPInternalServerError("")
	assert.Equal(t, expected, actual)
}

// TestGetHTTPStatusMessage tests HTTP status message from code.
func TestGetHTTPStatusMessage(t *testing.T) {
	assert.Equal(t, GetHTTPStatusMessage(200), "OK")
	assert.Equal(t, GetHTTPStatusMessage(400), "Bad Request")
	assert.Equal(t, GetHTTPStatusMessage(500), "Internal Server Error")
	assert.Equal(t, GetHTTPStatusMessage(0), "")
	assert.Equal(t, GetHTTPStatusMessage(8930), "")
}

