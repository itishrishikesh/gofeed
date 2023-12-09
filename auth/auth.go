package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the header of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("E#1OTELT - No authentication info found!")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("E#1OTEM0 - Malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("E#1OTEM5 - Malformed auth header")
	}
	return vals[1], nil
}
