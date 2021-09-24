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
	newS := ""
	fmt.Println(len(s))
	for i := 0; i <= len(s)-1; i++ {
		fmt.Println(i)
		switch {
		case unicode.IsDigit(rune(s[i])):
			return "", ErrInvalidString
		case i < len(s)-1 && unicode.IsDigit(rune(s[i+1])):
			byteToInt, _ := strconv.Atoi(string(s[i+1]))
			multipliedChar := strings.Repeat(string(s[i]), byteToInt)
			newS += multipliedChar
			i++
		default:
			newS += string(s[i])
		}
	}
	return newS, nil
}

// * "a4bc2d5e" => "aaaabccddddde"
