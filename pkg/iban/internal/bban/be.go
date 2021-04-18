package bban

import (
	"regexp"
	"strconv"
)

type BE struct {
	base
}

func NewBE() BE {
	b := BE{}
	b.format = regexp.MustCompile("^[0-9]{12}$")
	return b
}

func (b BE) Validate(input string) error {
	if err := b.base.Validate(input); err != nil {
		return err
	}

	num, err := strconv.ParseUint(input[:10], 10, 64)
	if err != nil {
		return err
	}

	var checksum string
	if modulo := num % 97; modulo == 0 {
		checksum = "97"
	} else {
		checksum = strconv.FormatUint(modulo, 10)
	}
	if checksum != input[10:12] {
		return ErrInvalidChecksum
	}

	return nil
}
