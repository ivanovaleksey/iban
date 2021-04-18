package bban

import (
	"regexp"
)

type CZ struct {
	base
}

func NewCZ() CZ {
	b := CZ{}
	b.format = regexp.MustCompile("^[0-9]{20}$")
	return b
}

func (b CZ) Validate(input string) error {
	if err := b.base.Validate(input); err != nil {
		return err
	}

	// don't know how checksum is calculated

	return nil
}
