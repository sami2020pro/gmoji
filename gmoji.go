package gmoji

import (
	"fmt"
	"regexp"
	"strings"
)

type Gmoji string

type All Gmoji

func replace(str string) string {
	match := regexp.MustCompile(`:\w*:`)

	rawGmoji := match.FindString(str)
	bakeGmoji := strings.Trim(rawGmoji, ":")

	newstr := string(strings.Replace(str, match.FindString(str), fmt.Sprint("%s", string(All)), -1))

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
