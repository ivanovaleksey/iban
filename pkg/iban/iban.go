package iban

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidFormat   = errors.New("iban: invalid format")
	ErrInvalidChecksum = errors.New("iban: invalid checksum")
	ErrUnknownCountry  = errors.New("iban: unknown country")
)

var knownCountries = map[string]struct{}{
	"SE": {},
	"NO": {},
	"FI": {},
	"IT": {},
	"BE": {},
}

var (
	ibanRegexp    = regexp.MustCompile("^(?P<country>[A-Z]{2})(?P<check>[0-9]{2})(?P<bban>[0-9A-Z]{1,30})$")
	ibanFormatter = formatter{
		partSize:  4,
		delimiter: " ",
	}
)

type IBAN struct {
	data string
}

func (iban IBAN) String() string {
	return ibanFormatter.Format(iban.data)
}

// Country returns country code according to ISO 3166-1 alpha-2
func (iban IBAN) Country() string {
	return iban.data[:2]
}

func MustParse(value string) IBAN {
	id, err := Parse(value)
	if err != nil {
		panic(err)
	}
	return id
}

func Parse(value string) (IBAN, error) {
	const regexpParts = 4

	m := ibanRegexp.FindStringSubmatch(value)
	if len(m) < regexpParts {
		return IBAN{}, ErrInvalidFormat
	}
	if err := validate(m[1], m[2], m[3]); err != nil {
		return IBAN{}, err
	}
	return IBAN{data: value}, nil
}

func validate(country, check, bban string) error {
	const validChecksum = 1

	_, ok := knownCountries[country]
	if !ok {
		return ErrUnknownCountry
	}
	if calcChecksum(country, check, bban) != validChecksum {
		return ErrInvalidChecksum
	}

	return nil
}
