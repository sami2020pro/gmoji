package gmoji

import (
	"fmt"
	"regexp"
	"strings"
)

type Gmoji string

func Parse(str string) string {
	var newstr string
	
	match := regexp.MustCompile(`:\w*:`)
	gmoji := match.FindString(str)
	strings.Trim(gmoji, ":")

	for k, v := range gmojiMap {
		if gmoji == k {
			newstr = fmt.Sprint("%s", v)
		}
	}

	return string(newstr)
}

func Print(str string) {
	fmt.Print(Parse(str))
}

func Println(str string) {
	fmt.Println(Parse(str))
}

/*
O-------------------------------O
O License: MIT 				    O
O-------------------------------O
*/
