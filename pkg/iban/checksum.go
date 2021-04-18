package iban

import (
	"strconv"
	"strings"
)

func calcChecksum(country, check, bban string) int {
	var buf strings.Builder
	buf.Grow(len(country) + len(check) + len(bban))

	for i := 0; i < len(bban); i++ {
		char := bban[i]
		if char >= 'A' && char <= 'Z' {
			buf.WriteString(letterAsNumber(char))
		} else {
			buf.WriteByte(char)
		}
	}
	for i := 0; i < len(country); i++ {
		buf.WriteString(letterAsNumber(country[i]))
	}
	buf.WriteString(check)

	return modulo97(buf.String())
}

// letterAsNumber converts letter to responding number
// A=10, B=11, ..., Z=35
func letterAsNumber(letter byte) string {
	char := letter - 'A' + 10
	return strconv.Itoa(int(char))
}

func modulo97(number string) int {
	const (
		partSize    = 9
		denominator = 97
	)

	numParts := len(number) / partSize
	if len(number)%partSize != 0 {
		numParts++
	}

	var (
		moduloInt int
		moduloStr string
	)
	for i := 0; i < numParts; i++ {
		left := i * partSize
		right := (i + 1) * partSize
		if len(number) < right {
			right = len(number)
		}

		part := number[left:right]
		num, err := strconv.ParseUint(moduloStr+part, 10, 64)
		if err != nil {
			panic(err)
		}
		moduloInt = int(num % denominator)
		moduloStr = strconv.Itoa(moduloInt)
	}
	return moduloInt
}
