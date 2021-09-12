package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const BACKSLASH = 92 // '\'

type slice struct {
	value  rune
	repeat int
}

func Unpack(str string) (string, error) {
	if !validate(str) {
		return "", ErrInvalidString
	}

	if len(str) == 0 {
		return "", nil
	}

	builder := strings.Builder{}

	for _, s := range parse(str) {
		builder.WriteString(strings.Repeat(string(s.value), s.repeat))
	}

	return builder.String(), nil
}

func parse(str string) []slice {
	var slices []slice

	runes := []rune(str)
	count := len(runes) - 1

	for i := 0; i < count; i++ {
		val := runes[i]
		repeat := 1

		if runes[i] != BACKSLASH {
			if unicode.IsDigit(runes[i+1]) {
				repeat, _ = strconv.Atoi(string(runes[i+1]))
				i++
			}
		}

		if runes[i] == BACKSLASH {
			val = runes[i+1]
			repeat = 1
			i++

			if i+1 <= count && unicode.IsDigit(runes[i+1]) {
				repeat, _ = strconv.Atoi(string(runes[i+1]))
				i++
			}
		}

		slices = append(slices, slice{
			value:  val,
			repeat: repeat,
		})

		if count-i == 1 {
			slices = append(slices, slice{
				value:  runes[count],
				repeat: 1,
			})
		}
	}

	return slices
}

func validate(str string) bool {
	re := regexp.MustCompile(`^\d|(\\[a-z]|\\[A-Z])|(\\\d{3,}|([a-z]|[A-Z])\d{2,})`)
	result := re.FindAllString(str, -1)

	return len(result) == 0
}
