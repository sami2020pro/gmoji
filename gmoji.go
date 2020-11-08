package gmoji

import (
	"fmt"
	"strings"
)

var (
	match := regexp.MustCompile(`:\w*:`)

	rawGmoji := match.FindString(str)
	bakeGmoji := strings.Trim(rawGmoji, ":")
)

type Gmoji string

func replace(str string) string {
	newstr := string(strings.Replace(str, match.FindString(str), fmt.Sprint("%s", string(bakeGmoji)), -1))

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
