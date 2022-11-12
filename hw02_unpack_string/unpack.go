package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {

	var result string
	var prev rune
	prevExists := false

	for _, s := range str {

		isDigit := (s >= 48 && s <= 57)
		prevIsDigit := (prev >= 48 && prev <= 57)

		if isDigit {
			if prevExists {
				for j := 0; j < int(s-'0'); j++ {
					result = result + string(prev)
				}
				prevExists = false

			} else {
				return "", ErrInvalidString
			}
		} else {
			if !prevIsDigit {
				if prevExists {
					result = result + string(prev)
				}
				prev = s
				prevExists = true
			}
		}
	}

	// the last symbol
	if prevExists {
		result = result + string(prev)
	}

	return result, nil
}
