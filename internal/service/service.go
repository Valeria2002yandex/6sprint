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

	isText := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			isText = true
			break
		}
	}

	if isText {
		result := converter.ToMorse(s)
		if result == "" {
			return "", errors.New("invalid morse code")
		}
		return result, nil
	}

	result := converter.ToText(s)
	if result == "" {
		return "", errors.New("invalid morse code")
	}
	return result, nil
}
