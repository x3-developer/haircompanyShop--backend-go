package valueobject

import (
	"errors"
	"regexp"
)

var hexColorRe = regexp.MustCompile(`^#[0-9A-Fa-f]{6}$`)

type ColorVO string

func NewColorVO(v string) (ColorVO, error) {
	if !hexColorRe.MatchString(v) {
		return "", errors.New("invalid color format, must be #RRGGBB")
	}
	return ColorVO(v), nil
}

func (c ColorVO) String() string {
	return string(c)
}
