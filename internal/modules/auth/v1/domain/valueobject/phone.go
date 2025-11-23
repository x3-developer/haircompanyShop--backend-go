package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

var (
	reDigits = regexp.MustCompile(`\D+`)
)

type Phone string

func NewPhone(raw string) (Phone, error) {
	normalized := normalizeRuPhone(raw)

	if !isValidRuPhone(normalized) {
		return "", errors.New("invalid phone number")
	}

	return Phone(normalized), nil
}

func (p Phone) String() string {
	return string(p)
}

func isValidRuPhone(p string) bool {
	return len(p) == 11 && p[0] == '7'
}

func normalizeRuPhone(s string) string {
	s = reDigits.ReplaceAllString(s, "")

	if len(s) == 11 && strings.HasPrefix(s, "8") {
		s = "7" + s[1:]
		return s
	}

	if len(s) == 11 && strings.HasPrefix(s, "7") {
		return s
	}

	if len(s) == 10 {
		return "7" + s
	}

	return s
}
