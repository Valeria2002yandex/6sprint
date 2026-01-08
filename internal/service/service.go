package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var converter = morse.NewConverter(morse.DefaultMorse)

func Conversion(text string) (string, error) {

	s := strings.TrimSpace(text)
	if s == "" {
		return "", errors.New("input string is empty")
	}

	hasLetters := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			hasLetters = true
			break
		}
	}

	var result string
	if hasLetters {
		result = converter.ToMorse(strings.ToUpper(s))
		if result == "" {
			return "", errors.New("failed to convert text to morse")
		}
	} else {

		result = converter.ToText(s)
	}

	if result == "" {
		return "", errors.New("invalid morse code")
	}

	return result, nil
}
