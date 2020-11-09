package gmoji

import (
	"fmt"
	"regexp"
	"strings"
)

type Gmoji string

func Parse(str string) string {
	match := regexp.MustCompile(`:\w*:`)
	gmoji := match.FindString(str)
	strings.Trim(gmoji, ":")

	for k, v := range gmojiMap {
		if gmoji == k {
			fmt.Printf("%s\n", v)
		}
	}
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
