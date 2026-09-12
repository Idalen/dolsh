package transport

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// APIError describes a non-2xx response from the server.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}

func newAPIError(res *http.Response) error {
	body, _ := io.ReadAll(res.Body)

	if msg := friendlyStatus(res.StatusCode); msg != "" && len(body) == 0 {
		return &APIError{StatusCode: res.StatusCode, Message: msg}
	}

	if msg := extractErrorMessage(body); msg != "" {
		return &APIError{StatusCode: res.StatusCode, Message: msg}
	}

	if msg := friendlyStatus(res.StatusCode); msg != "" {
		return &APIError{StatusCode: res.StatusCode, Message: msg}
	}

	return &APIError{StatusCode: res.StatusCode, Message: res.Status}
}

func extractErrorMessage(body []byte) string {
	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" {
		return ""
	}

	var s string
	if err := json.Unmarshal(body, &s); err == nil {
		if strings.TrimSpace(s) != "" {
			return s
		}
	}

	var m map[string]any
	if err := json.Unmarshal(body, &m); err == nil {
		for _, key := range []string{"title", "message", "error", "detail"} {
			if v, ok := m[key].(string); ok && strings.TrimSpace(v) != "" {
				return v
			}
		}
	}

	return ""
}

func friendlyStatus(code int) string {
	switch code {
	case http.StatusBadRequest:
		return "Invalid request"
	case http.StatusUnauthorized:
		return "Invalid username or password"
	case http.StatusForbidden:
		return "Access forbidden"
	case http.StatusNotFound:
		return "Server not found"
	case http.StatusRequestTimeout:
		return "Request timed out"
	case http.StatusTooManyRequests:
		return "Too many attempts, try again later"
	case http.StatusInternalServerError:
		return "Server error"
	case http.StatusBadGateway:
		return "Server unavailable"
	case http.StatusServiceUnavailable:
		return "Server unavailable"
	case http.StatusGatewayTimeout:
		return "Server timed out"
	default:
		return ""
	}
}
