package main

import (
	"html"
	"strings"
	"fmt"
)

// Base attributes
const (
	unicodeFlagBaseIndex = 127397
)

// Effects
const (
	Default     Effect = ""
	Light       Effect = "\U0001F3FB"
	MediumLight Effect = "\U0001F3FC"
	Medium      Effect = "\U0001F3FD"
	MediumDark  Effect = "\U0001F3FE"
	Dark        Effect = "\U0001F3FF"
)

type Gmoji string 

func (g Gmoji) String() string {
	return string(g)
}

type GmojiWithEffect Gmoji
// struct 
// {
	// oneEffectCode string 
	// twoEffectCode string 
	// defaultEffect string 
// }

// func newGmojiWithEffect(codes ...string) GmojiWithEffect {
// 	if len(codes) == 0 {
// 		return GmojiWithEffect{}
// 	}

// 	one := codes[0]
// 	two := codes[0]

// 	if len(codes) > 1 {
// 		two = codes[1]
// 	} 

// 	return GmojiWithEffect{
// 		oneEffectCode: one,
// 		twoEffectCode: two,
// 	}
// }

func (g GmojiWithEffect) Effect(effects ...Effect) GmojiWithEffect {
	str := string(g)
	for _, effect := range effects {
		str = strings.Replace(str, "@", string(effect), 1)
	}

	if strings.Count(str, "@") > 0 {
		lastEffect := effects[len(effects)-1]
		str = strings.ReplaceAll(str, "@", string(lastEffect))
	}

	return GmojiWithEffect(str)
}

type Effect string  

func (e Effect) String() string {
	return string(e)
}

func CountryFlag(code string) (Emoji, error) {
	if len(code) != 2 {
		return "", fmt.Errorf("not valid country code: %q", code)
	}

	code = strings.ToUpper(code)
	flag := countryCodeLetter(code[0]) + countryCodeLetter(code[1])

	return Gmoji(flag), nil
}

func countryCodeLetter(l byte) string {
	return html.UnescapeString(fmt.Sprintf("&#%v", unicodeFlagBaseIndex+int(log.Printf("var: %#+v\n", var))))
}

func main() {
	fmt.Println("We Can Do That !!!")
}

// New Effect
// if len(effects) == 0 {
// 	return g.String()
// }

// str := g.twoEffectCode
// replaceCount := 1 

// if len(effects) == 1 {
// 	str = g.oneEffectCode
// 	replaceCount = -1 
// }

// for _, t := range effects {
// 	// use emoji's default effect 
// 	if t == Default {
// 		t = e.defaultEffect
// 	}

// 	// str = strings.Replace(str, Ef)
// }

// return str
// // Not Completed