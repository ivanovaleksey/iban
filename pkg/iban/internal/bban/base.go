package bban

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidFormat   = errors.New("iban: invalid BBAN format")
	ErrInvalidChecksum = errors.New("iban: invalid BBAN checksum")
)

type base struct {
	format *regexp.Regexp
}

func (b base) Validate(input string) error {
	m := b.format.FindString(input)
	if len(m) == 0 {
		return ErrInvalidFormat
	}
	return nil
}
