package iban

import "strings"

// formatter formats input by splitting it into parts with a given size separated by delimiter
type formatter struct {
	partSize  int
	delimiter string
}

func (f formatter) Format(input string) string {
	numParts := len(input) / f.partSize
	if len(input)%f.partSize != 0 {
		numParts++
	}

	var builder strings.Builder
	builder.Grow(numParts*f.partSize + (numParts - 1)) // chars and spaces

	for i := 0; i < numParts; i++ {
		left := i * f.partSize
		right := (i + 1) * f.partSize
		if len(input) < right {
			right = len(input)
		}
		part := input[left:right]
		builder.WriteString(part)
		if i != numParts-1 {
			builder.WriteString(f.delimiter)
		}
	}

	return builder.String()
}
