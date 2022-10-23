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
		if unicode.IsDigit(s) {
			if prevExists {
				for j := 0; j < int(s-'0'); j++ {
					result = result + string(prev)
				}
				prevExists = false
			}
		} else {
			if !unicode.IsDigit(prev) {
				if prevExists {
					result = result + string(prev)
				}
				prev = s
				prevExists = true
			}
		}
	}

	if prevExists {
		result = result + string(prev)
	}

	return result, nil
}
