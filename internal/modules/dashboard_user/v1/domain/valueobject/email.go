package valueobject

import (
	"errors"
	"regexp"
)

type Email string

var emailRe = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

func NewEmail(v string) (Email, error) {
	if !emailRe.MatchString(v) {
		return "", errors.New("invalid email format")
	}
	return Email(v), nil
}

func (e Email) String() string {
	return string(e)
}
