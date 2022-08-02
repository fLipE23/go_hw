package hw02_unpack_string

import (
	"errors"
	"fmt"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {

	fmt.Println(len([]rune(str)))
	fmt.Println(str)

	var result string

	var prev rune

	for i, s := range str {

		fmt.Println("i: ", i, "s: ", s, string(s), unicode.IsDigit(s), " ", prev, result)

		if unicode.IsDigit(s) {

			for j := 0; j < int(s-'0'); j++ {
				result = result + string(prev)

				fmt.Println("prev: ", prev)
			}

		} else {
			if !unicode.IsDigit(prev) {
				result = result + string(prev)
			}
			if i == len(str)-1 {
				result += string(s)
			}
		}

		prev = s
	}

	fmt.Println("result: ", result)

	return result, nil
}
