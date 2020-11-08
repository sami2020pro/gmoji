package gmoji

import (
	"fmt"
	"regexp"
	"strings"
)

type Gmoji string

func (g Gmoji) String() string {
	return string(g)
}

func replace(str string) string {
	match := regexp.MustCompile(`:\w*:`)

	rawGmoji := match.FindString(str)
	bakeGmoji := strings.Trim(rawGmoji, ":")

	newstr := string(strings.Replace(str, match.FindString(str), gmoji.Glass, -1))

	return newstr
}

func Print(str string) {
	fmt.Print(replace(str))
}

func Println(str string) {
	fmt.Println(replace(str))
}

/*
O-------------------------------O
O License: MIT 				    O
O-------------------------------O
*/
