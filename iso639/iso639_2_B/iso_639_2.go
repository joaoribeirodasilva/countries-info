// [Source]https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
package iso_639_2

import "strings"

type ISO639_2_B string

const (
	ISO639_2_B_None      ISO639_2_B = ""
	ISO639_2_B_Abkhazian ISO639_2_B = "abk"
)

var codes = map[string]ISO639_2_B{
	"":    ISO639_2_B_None,
	"abk": ISO639_2_B_Abkhazian,
}

func GetCode(language string) ISO639_2_B {

	return codes[strings.ToLower(language)]
}
