package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("Missing authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed authorization header")
	}

	return vals[1], nil
}
