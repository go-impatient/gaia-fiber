package http

import (
	"net/http"
)

var (
	httpStatusMessages = map[int]string{
		http.StatusContinue:                      "Continue",                        // 100
		http.StatusSwitchingProtocols:            "Switching Protocols",             // 101
		http.StatusProcessing:                    "Processing",                      // 102
		http.StatusEarlyHints:                    "Early Hints",                     // 103
		http.StatusOK:                            "OK",                              // 200
		http.StatusCreated:                       "Created",                         // 201
		http.StatusAccepted:                      "Accepted",                        // 202
		http.StatusNonAuthoritativeInfo:          "Non Authoritative Info",          // 203
		http.StatusNoContent:                     "No Content",                      // 204
		http.StatusResetContent:                  "Reset Content",                   // 205
		http.StatusPartialContent:                "PartialC ontent",                 // 206
		http.StatusMultiStatus:                   "Multi Status",                    // 207
		http.StatusAlreadyReported:               "Already Reported",                // 208
		http.StatusIMUsed:                        "IM Used",                         // 226
		http.StatusMultipleChoices:               "Multiple Choices",                // 300
		http.StatusMovedPermanently:              "Moved Permanently",               // 301
		http.StatusFound:                         "Found",                           // 302
		http.StatusSeeOther:                      "See Other",                       // 303
		http.StatusNotModified:                   "Not Modified",                    // 304
		http.StatusUseProxy:                      "Use Proxy",                       // 305
		http.StatusTemporaryRedirect:             "Temporary Redirect",              // 307
		http.StatusPermanentRedirect:             "Permanent Redirect",              // 308
		http.StatusBadRequest:                    "Bad Request",                     // 400
		http.StatusUnauthorized:                  "Unauthorized",                    // 401
		http.StatusPaymentRequired:               "Payment Required",                // 402
		http.StatusForbidden:                     "Forbidden",                       // 403
		http.StatusNotFound:                      "Not Found",                       // 404
		http.StatusMethodNotAllowed:              "Method Not Allowed",              // 405
		http.StatusNotAcceptable:                 "Not Acceptable",                  // 406
		http.StatusProxyAuthRequired:             "Proxy Auth Required",             // 407
		http.StatusRequestTimeout:                "Request Timeout",                 // 408
		http.StatusConflict:                      "Conflict",                        // 409
		http.StatusGone:                          "Gone",                            // 410
		http.StatusLengthRequired:                "Length Required",                 // 411
		http.StatusPreconditionFailed:            "Precondition Failed",             // 412
		http.StatusRequestEntityTooLarge:         "Request Entity Too Large",        // 413
		http.StatusRequestURITooLong:             "Request URI Too Long",            // 414
		http.StatusUnsupportedMediaType:          "Unsupported Media Type",          // 415
		http.StatusRequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable", // 416
		http.StatusExpectationFailed:             "Expectation Failed",              // 417
		http.StatusTeapot:                        "Teapot",                          // 418
		http.StatusMisdirectedRequest:            "Misdirected Request",             // 421
		http.StatusUnprocessableEntity:           "Unprocessable Entity",            // 422
		http.StatusLocked:                        "Locked",                          // 423
		http.StatusFailedDependency:              "Failed Dependency",               // 424
		http.StatusTooEarly:                      "Too Early",                       // 425
		http.StatusUpgradeRequired:               "Upgrade Required",                // 426
		http.StatusPreconditionRequired:          "Precondition Required",           // 428
		http.StatusTooManyRequests:               "Too Many Requests",               // 429
		http.StatusRequestHeaderFieldsTooLarge:   "Request Header Fields Too Large", // 431
		http.StatusUnavailableForLegalReasons:    "Unavailable For Legal Reasons",   // 451
		http.StatusInternalServerError:           "Internal Server Error",           // 500
		http.StatusNotImplemented:                "Not Implemented",                 // 501
		http.StatusBadGateway:                    "Bad Gateway",                     // 502
		http.StatusServiceUnavailable:            "Service Unavailable",             // 503
		http.StatusGatewayTimeout:                "Gateway Timeout",                 // 504
		http.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",      // 505
		http.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",         // 506
		http.StatusInsufficientStorage:           "Insufficient Storage",            // 507
		http.StatusLoopDetected:                  "Loop Detected",                   // 508
		http.StatusNotExtended:                   "Not Extended",                    // 510
		http.StatusNetworkAuthenticationRequired: "Network Authentication Required", // 511
	}
)

// HTTPResponse type
type HTTPResponse struct {
	Code    int         `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data,omitempty" xml:"data"`
}

// GetHTTPResponse returns HTTPResponse type.
func GetHTTPResponse(code int, message string, data interface{}) HTTPResponse {
	return HTTPResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// GetHTTPInternalServerError returns 500 error.
func GetHTTPInternalServerError(message string) HTTPResponse {
	if message == "" {
		message = "Internal Server Error"
	}

	return HTTPResponse{
		Code:    500,
		Message: message,
		Data:    nil,
	}
}

// GetHTTPStatusMessage returns a message for a HTTP status code.
func GetHTTPStatusMessage(code int) string {
	msg, ok := httpStatusMessages[code]
	if !ok {
		return ""
	}
	return msg
}
