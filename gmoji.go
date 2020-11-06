package gmoji

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

func CountryFlag(code string) (Gmoji, error) {
	if len(code) != 2 {
		return "", fmt.Errorf("not valid country code: %q", code)
	}

	code = strings.ToUpper(code)
	flag := countryCodeLetter(code[0]) + countryCodeLetter(code[1])

	return Gmoji(flag), nil
}

func countryCodeLetter(l byte) string {
	return html.UnescapeString(fmt.Sprintf("&#%v;", unicodeFlagBaseIndex+int(l)))
}
