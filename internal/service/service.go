package service

import (
	"errors"

	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var converter = morse.NewConverter(morse.DefaultMorse)

func Conversion(text string) (string, error) {

	str := strings.TrimSpace(text)
	if str == "" {
		return "", errors.New("input string is empty")
	}

	isMorse := strings.IndexFunc(str, func(r rune) bool {
		return r != '.' && r != '-' && r != ' ' && r != '/' && r != '\t' && r != '\n' && r != '\r'
	}) == -1

	if isMorse {
		result := converter.ToText(str)
		if result == "" || strings.Contains(result, "?") {
			return "", errors.New("invalid morse code")
		}
		return result, nil
	} else {

		result := converter.ToMorse(str)
		if result == "" {
			return "", errors.New("failed to convert text to morse")
		}
		return result, nil
	}
}
