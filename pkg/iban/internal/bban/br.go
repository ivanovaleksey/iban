package bban

import (
	"regexp"
)

type BR struct {
	base
}

func NewBR() BR {
	b := BR{}
	b.format = regexp.MustCompile("^[0-9]{23}[A-Z][0-9A-Z]$")
	return b
}

func (b BR) Validate(input string) error {
	if err := b.base.Validate(input); err != nil {
		return err
	}

	// don't know how checksum is calculated

	return nil
}
