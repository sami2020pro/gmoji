package gmoji

import (
	"fmt"
	"strings"
	"unicode"
)

type Gmoji string

func Parse(input string) string {
	var matched strings.Builder
	var output strings.Builder

	for _, r := range input {
		if r != ':' {
			if matched.Len() == 0 {
				output.WriteRune(r)
				continue
			}

			matched.WriteRune(r)

			if unicode.IsSpace(r) {
				output.WriteString(matched.String())
				matched.Reset()
			}
			continue
		}

		if matched.Len() == 0 {
			matched.WriteRune(r)
			continue
		}

		match := matched.String()
		alias := match + ":"

		if code, ok := find(alias); ok {
			output.WriteString(code)
			matched.Reset()
			continue
		}

		output.WriteString(match)

		matched.Reset()
		matched.WriteRune(r)

	}

	if matched.Len() != 0 {
		output.WriteString(matched.String())
		matched.Reset()
	}

	return output.String()
}

func find(alias string) (string, bool) {
	if code, ok := gmojiMap[alias]; ok {
		return code, true
	}

	return "", false
}

func Print(str string) {
	fmt.Print(Parse(str))
}

func Println(str string) {
	fmt.Println(Parse(str))
}

/*
O------------------------------------O
O License: MIT | Copyright 2020-2021 O
O------------------------------------O
*/
