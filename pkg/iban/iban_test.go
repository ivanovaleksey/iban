package iban

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {

}

func TestIBAN_String(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{
			in:  "IT60Q0123412345000000753XYZ",
			out: "IT60 Q012 3412 3450 0000 0753 XYZ",
		},
		{
			in:  "SE7280000810340009783242",
			out: "SE72 8000 0810 3400 0978 3242",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.in, func(t *testing.T) {
			iban, err := Parse(testCase.in)
			require.NoError(t, err)

			str := iban.String()

			assert.Equal(t, testCase.out, str)
		})
	}
}
