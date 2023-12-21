package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKeyOrToken extracts an API key or token from the header of an HTTP request
// Example 1:
// Authorization: ApiKey {insert apikey here}
// Example 2:
// Authorization: Bearer {insert token here}
func GetAPIKeyOrToken(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("E#1OTELT - No authentication info found!")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("E#1OTEM0 - Malformed auth header")
	}
	if vals[0] == "Bearer" || vals[0] == "ApiKey" {
		return vals[1], nil
	}
	return "", errors.New("E#1OTEM5 - Malformed auth header")
}
