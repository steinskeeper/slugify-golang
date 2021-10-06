package slugger

import (
	"fmt"
	"strings"
	"unicode"
)

// Params represents params pass to slugger
type Params struct {
	text      string
	separator string
}

func main() {
	fmt.Println(slugger(Params{text: "Kudos to everyone! here"}))
}

func slugger(p Params) string {
	if p.separator == "" {
		p.separator = "-"
	}

	splitter := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	textLower := strings.ToLower(p.text)
	splits := strings.FieldsFunc(textLower, splitter)
	output := strings.Join(splits, p.separator)

	return output
}