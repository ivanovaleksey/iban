package iban

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcCheck(t *testing.T) {
	testCases := []struct {
		in  string
		out int
	}{
		{
			in:  "IT60Q0123412345000000753XYZ",
			out: 1,
		},
		{
			in:  "IT63Q0123412345000000753XYZ",
			out: 4,
		},
		{
			in:  "SE7280000810340009783242",
			out: 1,
		},
		{
			in:  "SE7280",
			out: 4,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.in, func(t *testing.T) {
			iban := testCase.in

			check := calcChecksum(iban[:2], iban[2:4], iban[4:])

			assert.Equal(t, testCase.out, check)
		})
	}
}
