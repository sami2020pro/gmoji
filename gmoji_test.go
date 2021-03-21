package gmoji_test

import (
	"testing"
	"github.com/sami2020pro/gmoji"
)

// gmoji.go - Parse function
func TestParse(t *testing.T) {
	x := gmoji.Parse("Emoji aliases are :Plus:")
	if x != "Emoji aliases are \uf067" {
		t.Errorf("Error - Parse")
	}
}
