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
		if unicode.IsDigit(rune(s[i])) {
			return "", ErrInvalidString
		} else if i < len(s)-1 && unicode.IsDigit(rune(s[i+1])) {
			byteToInt, _ := strconv.Atoi(string(s[i+1]))
			multipliedChar := strings.Repeat(string(s[i]), byteToInt)
			newS += multipliedChar
			i++
		} else {
			newS += string(s[i])
		}
	}
	return newS, nil
}


// * "a4bc2d5e" => "aaaabccddddde"
