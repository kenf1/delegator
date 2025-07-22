package auth

import (
	"errors"
	"regexp"
)

func SanitizeQueryParam(queryParam string) (string, error) {
	var acceptedValues = regexp.MustCompile(`[^a-zA-Z0-9\-]`)

	sanitized := acceptedValues.ReplaceAllString(queryParam, "")
	if sanitized != queryParam {
		return "", errors.New("invalid input: contains non-allowed characters")
	}

	return sanitized, nil
}
