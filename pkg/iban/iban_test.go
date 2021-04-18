package iban

import (
	"github.com/ivanovaleksey/iban/pkg/iban/internal/bban"
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
		err error
	}{
		{
			in:  "IT60Q0123412345000000753XYZ",
			out: "IT60 Q012 3412 3450 0000 0753 XYZ",
		},
		{
			in:  "SE7280000810340009783242",
			out: "SE72 8000 0810 3400 0978 3242",
		},
		{
			in:  "CZ6508000000192000145399",
			out: "CZ65 0800 0000 1920 0014 5399",
		},
		{
			in:  "BR9700360305000010009795493P1",
			out: "BR97 0036 0305 0000 1000 9795 493P 1",
		},
		{
			in:  "BE68539007547034",
			out: "BE68 5390 0754 7034",
		},
		{
			in:  "BE41539007547035",
			err: bban.ErrInvalidChecksum,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.in, func(t *testing.T) {
			iban, err := Parse(testCase.in)

			if testCase.err != nil {
				require.Equal(t, testCase.err, err)
				assert.Empty(t, iban)
				return
			}

			require.NoError(t, err)
			str := iban.String()

			assert.Equal(t, testCase.out, str)
		})
	}
}
