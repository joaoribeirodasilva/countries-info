// [Source]https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes

package iso639

import "strings"

type Code int
type Name string

const (
	CODE_None Code = iota
	CODE_Abkhazian
	CODE_Afar
	CODE_Afrikaans
	CODE_Akan
	CODE_Albanian
	CODE_Amharic
	CODE_Arabic
	CODE_Aragonese
	CODE_Armenian
	CODE_Assamese
	CODE_Avaric
	CODE_Avestan
	CODE_Aymara
	CODE_Azerbaijani
	CODE_Bambara
	CODE_Bashkir
	CODE_Basque
	CODE_Belarusian
	CODE_Bengali
	CODE_Bislama
	CODE_Bosnian
	CODE_Breton
	CODE_Bulgarian
	CODE_Burmese
	CODE_Catalan
	CODE_Valencian
	CODE_Chamorro
	CODE_Chechen
	CODE_Chichewa
	CODE_Chewa
	CODE_Nyanja
	CODE_Chinese
	CODE_Church_Slavonic
	CODE_Old_Slavonic
	CODE_Old_Church_Slavonic
	CODE_Chuvash
	CODE_Cornish
	CODE_Corsican
	CODE_Cree
	CODE_Croatian
	CODE_Czech
	CODE_Danish
	CODE_Divehi
	CODE_Dhivehi
	CODE_Maldivian
	CODE_Dutch
	CODE_Flemish
	CODE_Dzongkha
	CODE_English
	CODE_Esperanto
	CODE_Estonian
	CODE_Ewe
	CODE_Faroese
	CODE_Fijian
	CODE_Finnish
	CODE_French
	CODE_Fulah
	CODE_Gaelic
	CODE_Scottish_Gaelic
	CODE_Galician
	CODE_Ganda
	CODE_Georgian
	CODE_German
	CODE_Greek
	CODE_Kalaallisut
	CODE_Greenlandic
	CODE_Guarani
	CODE_Gujarati
	CODE_Haitian
	CODE_Haitian_Creole
	CODE_Hausa
	CODE_Hebrew
	CODE_Herero
	CODE_Hindi
	CODE_Hiri_Motu
	CODE_Hungarian
	CODE_Icelandic
	CODE_Ido
	CODE_Igbo
	CODE_Indonesian
	CODE_Interlingua
	CODE_Interlingue
	CODE_Occidental
	CODE_Inuktitut
	CODE_Inupiaq
	CODE_Irish
	CODE_Italian
	CODE_Japanese
	CODE_Javanese
	CODE_Kannada
	CODE_Kanuri
	CODE_Kashmiri
	CODE_Kazakh
	CODE_Central_Khmer
	CODE_Kikuyu
	CODE_Gikuyu
	CODE_Kinyarwanda
	CODE_Kirghiz
	CODE_Kyrgyz
	CODE_Komi
	CODE_Kongo
	CODE_Korean
	CODE_Kuanyama
	CODE_Kwanyama
	CODE_Kurdish
	CODE_Lao
	CODE_Latin
	CODE_Latvian
	CODE_Limburgan
	CODE_Limburger
	CODE_Limburgish
	CODE_Lingala
	CODE_Lithuanian
	CODE_Luba_Katanga
	CODE_Luxembourgish
	CODE_Letzeburgesch
	CODE_Macedonian
	CODE_Malagasy
	CODE_Malay
	CODE_Malayalam
	CODE_Maltese
	CODE_Manx
	CODE_Maori
	CODE_Marathi
	CODE_Marshallese
	CODE_Mongolian
	CODE_Nauru
	CODE_Navajo
	CODE_Navaho
	CODE_North_Ndebele
	CODE_South_Ndebele
	CODE_Ndonga
	CODE_Nepali
	CODE_Norwegian
	CODE_Norwegian_Bokmal
	CODE_Norwegian_Nynorsk
	CODE_Sichuan_Yi
	CODE_Nuosu
	CODE_Occitan
	CODE_Ojibwa
	CODE_Oriya
	CODE_Oromo
	CODE_Ossetian
	CODE_Ossetic
	CODE_Pali
	CODE_Pashto
	CODE_Pushto
	CODE_Persian
	CODE_Polish
	CODE_Portuguese
	CODE_Punjabi
	CODE_Panjabi
	CODE_Quechua
	CODE_Romanian
	CODE_Moldavian
	CODE_Moldovan
	CODE_Romansh
	CODE_Rundi
	CODE_Russian
	CODE_Northern_Sami
	CODE_Samoan
	CODE_Sango
	CODE_Sanskrit
	CODE_Sardinian
	CODE_Serbian
	CODE_Shona
	CODE_Sindhi
	CODE_Sinhala
	CODE_Sinhalese
	CODE_Slovak
	CODE_Slovenian
	CODE_Somali
	CODE_Southern_Sotho
	CODE_Spanish
	CODE_Castilian
	CODE_Sundanese
	CODE_Swahili
	CODE_Swati
	CODE_Swedish
	CODE_Tagalog
	CODE_Tahitian
	CODE_Tajik
	CODE_Tamil
	CODE_Tatar
	CODE_Telugu
	CODE_Thai
	CODE_Tibetan
	CODE_Tigrinya
	CODE_Tonga
	CODE_Tsonga
	CODE_Tswana
	CODE_Turkish
	CODE_Turkmen
	CODE_Twi
	CODE_Uighur
	CODE_Uyghur
	CODE_Ukrainian
	CODE_Urdu
	CODE_Uzbek
	CODE_Venda
	CODE_Vietnamese
	CODE_Volapuk
	CODE_Walloon
	CODE_Welsh
	CODE_Wolof
	CODE_Xhosa
	CODE_Yiddish
	CODE_Yoruba
	CODE_Zhuang
	CODE_Chuang
	CODE_Zulu
)

var codes = map[Name]Code{
	"":                CODE_None,
	"Abkhazian":       CODE_Abkhazian,
	"Afar":            CODE_Afar,
	"Afrikaans":       CODE_Afrikaans,
	"Akan":            CODE_Akan,
	"Albanian":        CODE_Albanian,
	"Amharic":         CODE_Amharic,
	"Arabic":          CODE_Arabic,
	"Aragonese":       CODE_Aragonese,
	"Armenian":        CODE_Armenian,
	"Assamese":        CODE_Assamese,
	"Avaric":          CODE_Avaric,
	"Avestan":         CODE_Avestan,
	"Aymara":          CODE_Aymara,
	"Azerbaijani":     CODE_Azerbaijani,
	"Bambara":         CODE_Bambara,
	"Bashkir":         CODE_Bashkir,
	"Basque":          CODE_Basque,
	"Belarusian":      CODE_Belarusian,
	"Bengali":         CODE_Bengali,
	"Bislama":         CODE_Bislama,
	"Bosnian":         CODE_Bosnian,
	"Breton":          CODE_Breton,
	"Bulgarian":       CODE_Bulgarian,
	"Burmese":         CODE_Burmese,
	"Catalan":         CODE_Catalan,
	"Chamorro":        CODE_Chamorro,
	"Chechen":         CODE_Chechen,
	"Chichewa":        CODE_Chichewa,
	"Chinese":         CODE_Chinese,
	"Church Slavonic": CODE_Church_Slavonic,
	"Chuvash":         CODE_Chuvash,
	"Cornish":         CODE_Cornish,
	"Corsican":        CODE_Corsican,
	"Cree":            CODE_Cree,
	"Croatian":        CODE_Croatian,
	"Czech":           CODE_Czech,
	"Danish":          CODE_Danish,
	"Maldivian":       CODE_Maldivian,
	"Dutch":           CODE_Dutch,
	"Flemish":         CODE_Flemish,
	"Dzongkha":        CODE_Dzongkha,
	"English":         CODE_English,
	"Esperanto":       CODE_Esperanto,
	"Estonian":        CODE_Estonian,
	"Ewe":             CODE_Ewe,
	"Faroese":         CODE_Faroese,
	"Fijian":          CODE_Fijian,
	"Finnish":         CODE_Finnish,
	"French":          CODE_French,
	"Fulah":           CODE_Fulah,
	"Gaelic":          CODE_Gaelic,
	"Galician":        CODE_Galician,
	"Ganda":           CODE_Ganda,
	"Georgian":        CODE_Georgian,
	"German":          CODE_German,
	"Greek":           CODE_Greek,
	"Kalaallisut":     CODE_Kalaallisut,
	"Guarani":         CODE_Guarani,
	"Gujarati":        CODE_Gujarati,
	"Haitian":         CODE_Haitian,
	"Hausa":           CODE_Hausa,
	"Hebrew":          CODE_Hebrew,
	"Herero":          CODE_Herero,
	"Hindi":           CODE_Hindi,
	"Hiri Motu":       CODE_Hiri_Motu,
	"Hungarian":       CODE_Hungarian,
	"Icelandic":       CODE_Icelandic,
	"Igbo":            CODE_Ido,
	"Indonesian":      CODE_Indonesian,
}

var names = map[Code]Name{
	CODE_None:                "None",
	CODE_Abkhazian:           "Abkhazian",
	CODE_Afar:                "Afar",
	CODE_Afrikaans:           "Afrikaans",
	CODE_Akan:                "Akan",
	CODE_Albanian:            "Albanian",
	CODE_Amharic:             "Amharic",
	CODE_Arabic:              "Arabic",
	CODE_Aragonese:           "Aragonese",
	CODE_Armenian:            "Armenian",
	CODE_Assamese:            "Assamese",
	CODE_Avaric:              "Avaric",
	CODE_Avestan:             "Avestan",
	CODE_Aymara:              "Aymara",
	CODE_Azerbaijani:         "Azerbaijani",
	CODE_Bambara:             "Bambara",
	CODE_Bashkir:             "Bashkir",
	CODE_Basque:              "Basque",
	CODE_Belarusian:          "Belarusian",
	CODE_Bengali:             "Bengali",
	CODE_Bislama:             "Bislama",
	CODE_Bosnian:             "Bosnian",
	CODE_Breton:              "Breton",
	CODE_Bulgarian:           "Bulgarian",
	CODE_Burmese:             "Burmese",
	CODE_Catalan:             "Catalan",
	CODE_Valencian:           "Catalan",
	CODE_Chamorro:            "Chamorro",
	CODE_Chechen:             "Chechen",
	CODE_Chichewa:            "Chichewa",
	CODE_Chewa:               "Chichewa",
	CODE_Nyanja:              "Chichewa",
	CODE_Chinese:             "Chinese",
	CODE_Church_Slavonic:     "Church Slavonic",
	CODE_Old_Slavonic:        "Old Slavonic",
	CODE_Old_Church_Slavonic: "Old Church Slavonic",
	CODE_Chuvash:             "Chuvash",
	CODE_Cornish:             "Cornish",
	CODE_Corsican:            "Corsican",
	CODE_Cree:                "Cree",
	CODE_Croatian:            "Croatian",
	CODE_Czech:               "Czech",
	CODE_Danish:              "Danish",
	CODE_Divehi:              "Divehi",
	CODE_Dhivehi:             "Dhivehi",
	CODE_Maldivian:           "Maldivian",
	CODE_Dutch:               "Dutch",
	CODE_Flemish:             "Flemish",
	CODE_Dzongkha:            "Dzongkha",
	CODE_English:             "English",
	CODE_Esperanto:           "Esperanto",
	CODE_Estonian:            "Estonian",
	CODE_Ewe:                 "Ewe",
	CODE_Faroese:             "Faroese",
	CODE_Fijian:              "Fijian",
	CODE_Finnish:             "Finnish",
	CODE_French:              "French",
	CODE_Fulah:               "Fulah",
	CODE_Gaelic:              "Gaelic",
	CODE_Scottish_Gaelic:     "Scottish Gaelic",
	CODE_Galician:            "Galician",
	CODE_Ganda:               "Ganda",
	CODE_Georgian:            "Georgian",
	CODE_German:              "German",
	CODE_Greek:               "Greek",
	CODE_Kalaallisut:         "Kalaallisut",
	CODE_Greenlandic:         "Greenlandic",
	CODE_Guarani:             "Guarani",
	CODE_Gujarati:            "Gujarati",
	CODE_Haitian:             "Haitian",
	CODE_Haitian_Creole:      "Haitian",
	CODE_Hausa:               "Hausa",
	CODE_Hebrew:              "Hebrew",
	CODE_Herero:              "Herero",
	CODE_Hindi:               "Hindi",
	CODE_Hiri_Motu:           "Hiri Motu",
	CODE_Hungarian:           "Hungarian",
	CODE_Icelandic:           "Icelandic",
	CODE_Ido:                 "Ido",
	CODE_Igbo:                "Igbo",
	CODE_Indonesian:          "Indonesian",
	CODE_Interlingua:         "Interlingua",
	CODE_Interlingue:         "Interlingue",
	CODE_Occidental:          "Interlingue",
	CODE_Inuktitut:           "Inuktitut",
	CODE_Inupiaq:             "Inupiaq",
	CODE_Irish:               "Irish",
	CODE_Italian:             "Italian",
	CODE_Japanese:            "Japanese",
	CODE_Javanese:            "Javanese",
	CODE_Kannada:             "Kannada",
	CODE_Kanuri:              "Kanuri",
	CODE_Kashmiri:            "Kashmiri",
	CODE_Kazakh:              "Kazakh",
	CODE_Central_Khmer:       "Central Khmer",
	CODE_Kikuyu:              "Kikuyu",
	CODE_Gikuyu:              "Gikuyu",
	CODE_Kinyarwanda:         "Kinyarwanda",
	CODE_Kirghiz:             "Kirghiz",
	CODE_Kyrgyz:              "Kirghiz",
	CODE_Komi:                "Komi",
	CODE_Kongo:               "Kongo",
	CODE_Korean:              "Korean",
	CODE_Kuanyama:            "Kuanyama",
	CODE_Kwanyama:            "Kuanyama",
	CODE_Kurdish:             "Kurdish",
	CODE_Lao:                 "Lao",
	CODE_Latin:               "Latin",
	CODE_Latvian:             "Latvian",
	CODE_Limburgan:           "Limburgan",
	CODE_Limburger:           "Limburgan",
	CODE_Limburgish:          "Limburgan",
	CODE_Lingala:             "Lingala",
	CODE_Lithuanian:          "Lithuanian",
	CODE_Luba_Katanga:        "Luba-Katanga",
	CODE_Luxembourgish:       "Luxembourgish",
	CODE_Letzeburgesch:       "Luxembourgish",
	CODE_Macedonian:          "Macedonian",
	CODE_Malagasy:            "Malagasy",
	CODE_Malay:               "Malay",
	CODE_Malayalam:           "Malayalam",
	CODE_Maltese:             "Maltese",
	CODE_Manx:                "Manx",
	CODE_Maori:               "Maori",
	CODE_Marathi:             "Marathi",
	CODE_Marshallese:         "Marshallese",
	CODE_Mongolian:           "Mongolian",
	CODE_Nauru:               "Nauru",
	CODE_Navajo:              "Navajo",
	CODE_Navaho:              "Navaho",
	CODE_North_Ndebele:       "North Ndebele",
	CODE_South_Ndebele:       "South Ndebele",
	CODE_Ndonga:              "Ndonga",
	CODE_Nepali:              "Nepali",
	CODE_Norwegian:           "Norwegian",
	CODE_Norwegian_Bokmal:    "Norwegian Bokmål",
	CODE_Norwegian_Nynorsk:   "Norwegian Nynorsk",
	CODE_Sichuan_Yi:          "Sichuan Yi",
	CODE_Nuosu:               "Nuosu",
	CODE_Occitan:             "Occitan",
	CODE_Ojibwa:              "Ojibwa",
	CODE_Oriya:               "Oriya",
	CODE_Oromo:               "Oromo",
	CODE_Ossetian:            "Ossetian",
	CODE_Ossetic:             "Ossetic",
	CODE_Pali:                "Pali",
	CODE_Pashto:              "Pashto",
	CODE_Pushto:              "Pashto",
	CODE_Persian:             "Persian",
	CODE_Polish:              "Polish",
	CODE_Portuguese:          "Portuguese",
	CODE_Punjabi:             "Punjabi",
	CODE_Panjabi:             "Punjabi",
	CODE_Quechua:             "Quechua",
	CODE_Romanian:            "Romanian",
	CODE_Moldavian:           "Romanian",
	CODE_Moldovan:            "Romanian",
	CODE_Romansh:             "Romansh",
	CODE_Rundi:               "Rundi",
	CODE_Russian:             "Russian",
	CODE_Northern_Sami:       "Northern Sami",
	CODE_Samoan:              "Samoan",
	CODE_Sango:               "Sango",
	CODE_Sanskrit:            "Sanskrit",
	CODE_Sardinian:           "Sardinian",
	CODE_Serbian:             "Serbian",
	CODE_Shona:               "Shona",
	CODE_Sindhi:              "Sindhi",
	CODE_Sinhala:             "Sinhala",
	CODE_Sinhalese:           "Sinhalese",
	CODE_Slovak:              "Slovak",
	CODE_Slovenian:           "Slovenian",
	CODE_Somali:              "Somali",
	CODE_Southern_Sotho:      "Southern Sotho",
	CODE_Spanish:             "Spanish",
	CODE_Castilian:           "Spanish",
	CODE_Sundanese:           "Sundanese",
	CODE_Swahili:             "Swahili",
	CODE_Swati:               "Swati",
	CODE_Swedish:             "Swedish",
	CODE_Tagalog:             "Tagalog",
	CODE_Tahitian:            "Tahitian",
	CODE_Tajik:               "Tajik",
	CODE_Tamil:               "Tamil",
	CODE_Tatar:               "Tatar",
	CODE_Telugu:              "Telugu",
	CODE_Thai:                "Thai",
	CODE_Tibetan:             "Tibetan",
	CODE_Tigrinya:            "Tigrinya",
	CODE_Tonga:               "Tonga",
	CODE_Tsonga:              "Tsonga",
	CODE_Tswana:              "Tswana",
	CODE_Turkish:             "Turkish",
	CODE_Turkmen:             "Turkmen",
	CODE_Twi:                 "Twi",
	CODE_Uighur:              "Uighur",
	CODE_Uyghur:              "Uighur",
	CODE_Ukrainian:           "Ukrainian",
	CODE_Urdu:                "Urdu",
	CODE_Uzbek:               "Uzbek",
	CODE_Venda:               "Venda",
	CODE_Vietnamese:          "Vietnamese",
	CODE_Volapuk:             "Volapük",
	CODE_Walloon:             "Walloon",
	CODE_Welsh:               "Welsh",
	CODE_Wolof:               "Wolof",
	CODE_Xhosa:               "Xhosa",
	CODE_Yiddish:             "Yiddish",
	CODE_Yoruba:              "Yoruba",
	CODE_Zhuang:              "Zhuang",
	CODE_Chuang:              "Zhuang",
	CODE_Zulu:                "Zulu",
}

func GetCode(name string) Code {

	return codes[Name(strings.ToLower(name))]
}

func GetName(code Code) Name {

	return names[code]
}
