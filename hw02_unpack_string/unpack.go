package hw02unpackstring

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

var cache = map[string][]rune{}

func RuneAt(s string, idx int) rune {
	rs := cache[s]
	if rs == nil {
		rs = []rune(s)
		cache[s] = []rune(s)
	}
	if idx >= len(rs) {
		return 0
	}
	return rs[idx]
}

func Unpack(s string) (string, error) {
	var sb strings.Builder
	fmt.Println("input string:", s)
	for pos, char := range s {
		// shoudn't start from digit
		if pos == 0 && unicode.IsDigit(RuneAt(s, 0)) {
			return "", ErrInvalidString
		}
		switch {
		case unicode.IsLetter(char):
			// letter + digit
			if unicode.IsDigit(RuneAt(s, pos+1)) {
				// https://stackoverflow.com/questions/21322173/convert-rune-to-int
				// subtracting the value of rune '0' from any rune '0' through '9' will give you an integer 0 through 9
				byteToInt := int(RuneAt(s, pos+1) - '0')
				multipliedChar := strings.Repeat(string(char), byteToInt)
				sb.WriteString(multipliedChar)
				// letter only
			} else {
				sb.WriteString(string(char))
			}
		// digit + digit -> error
		case unicode.IsDigit(RuneAt(s, pos)) && unicode.IsDigit(RuneAt(s, pos+1)):
			return "", ErrInvalidString
		// char != letter or digital
		case !unicode.IsDigit(char) && !unicode.IsLetter(char):
			return "", ErrInvalidString
		}
	}
	return sb.String(), nil
}
