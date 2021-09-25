package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var sb strings.Builder
	// fmt.Println(len(s))
	fmt.Println(len(s))
	for i := 0; i <= len(s)-1; i++ {
		switch {
		case unicode.IsDigit(rune(s[i])):
			return "", ErrInvalidString
		case i < len(s)-1 && unicode.IsDigit(rune(s[i+1])):
			fmt.Println("hack", string(s[i]))
			byteToInt, _ := strconv.Atoi(string(s[i+1]))
			multipliedChar := strings.Repeat(string(s[i]), byteToInt)
			sb.WriteString(multipliedChar)
			i++
		default:
			sb.WriteByte(s[i])
		}
	}
	return sb.String(), nil
}

// * "a4bc2d5e" => "aaaabccddddde"
