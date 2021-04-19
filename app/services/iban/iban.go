package iban

import "github.com/ivanovaleksey/iban/pkg/iban"

type Service struct {}

func (s Service) Validate(input string) error {
	_, err := iban.Parse(input)
	return err
}
