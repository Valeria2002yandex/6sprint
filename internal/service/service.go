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

	isMorse := true
	for _, r := range s {
		if unicode.IsLetter(r) {
			isMorse = false
			break
		}
	}

	if isMorse {
		result := converter.ToText(s)
		if result == "" {
			return "", errors.New("invalid morse code")
		}
		return result, nil
	}

	result := converter.ToMorse(s)
	if result == "" {
		return "", errors.New("invalid morse code")
	}
	return result, nil
}
