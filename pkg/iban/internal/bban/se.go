package bban

import "regexp"

type SE struct {
	base
}

func NewSE() SE {
	b := SE{}
	b.format = regexp.MustCompile("^[0-9]{20}$")
	return b
}

func (b SE) Validate(input string) error {
	if err := b.base.Validate(input); err != nil {
		return err
	}

	// don't know how checksum is calculated

	return nil
}
