package bban

import (
	"regexp"
)

type IT struct {
	base
}

func NewIT() IT {
	b := IT{}
	b.format = regexp.MustCompile("^[A-Z][0-9]{10}[0-9A-Z]{12}$")
	return b
}

func (b IT) Validate(input string) error {
	if err := b.base.Validate(input); err != nil {
		return err
	}

	// don't know how checksum is calculated

	return nil
}
