package hw02_unpack_string

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result string
	var prev rune
	prevExists := false

	for _, s := range str {
		if unicode.IsDigit(s) && !prevExists {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(s) {
			for j := 0; j < int(s-'0'); j++ {
				result += string(prev)
			}
			prevExists = false
		} else {
			if prevExists {
				result += string(prev)
			}
			prev = s
			prevExists = true
		}
	}

	// the last symbol
	if prevExists {
		result += string(prev)
	}

	return result, nil
}
