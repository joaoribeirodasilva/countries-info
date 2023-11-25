// [Source]https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
package languages

import "strings"

type Language string

// TODO: change to numeric codes
const (
	LANG_None                Language = "None"
	LANG_Abkhazian           Language = "Abkhazian"
	LANG_Afar                Language = "Afar"
	LANG_Afrikaans           Language = "Afrikaans"
	LANG_Akan                Language = "Akan"
	LANG_Albanian            Language = "Albanian"
	LANG_Amharic             Language = "Amharic"
	LANG_Arabic              Language = "Arabic"
	LANG_Aragonese           Language = "Aragonese"
	LANG_Armenian            Language = "Armenian"
	LANG_Assamese            Language = "Assamese"
	LANG_Avaric              Language = "Avaric"
	LANG_Avestan             Language = "Avestan"
	LANG_Aymara              Language = "Aymara"
	LANG_Azerbaijani         Language = "Azerbaijani"
	LANG_Bambara             Language = "Bambara"
	LANG_Bashkir             Language = "Bashkir"
	LANG_Basque              Language = "Basque"
	LANG_Belarusian          Language = "Belarusian"
	LANG_Bengali             Language = "Bengali"
	LANG_Bislama             Language = "Bislama"
	LANG_Bosnian             Language = "Bosnian"
	LANG_Breton              Language = "Breton"
	LANG_Bulgarian           Language = "Bulgarian"
	LANG_Burmese             Language = "Burmese"
	LANG_Catalan             Language = "Catalan"
	LANG_Valencian           Language = "Catalan"
	LANG_Chamorro            Language = "Chamorro"
	LANG_Chechen             Language = "Chechen"
	LANG_Chichewa            Language = "Chichewa"
	LANG_Chewa               Language = "Chichewa"
	LANG_Nyanja              Language = "Chichewa"
	LANG_Chinese             Language = "Chinese"
	LANG_Church_Slavonic     Language = "Church Slavonic"
	LANG_Old_Slavonic        Language = "Old Slavonic"
	LANG_Old_Church_Slavonic Language = "Old Church Slavonic"
	LANG_Chuvash             Language = "Chuvash"
	LANG_Cornish             Language = "Cornish"
	LANG_Corsican            Language = "Corsican"
	LANG_Cree                Language = "Cree"
	LANG_Croatian            Language = "Croatian"
	LANG_Czech               Language = "Czech"
	LANG_Danish              Language = "Danish"
	LANG_Divehi              Language = "Divehi"
	LANG_Dhivehi             Language = "Dhivehi"
	LANG_Maldivian           Language = "Maldivian"
	LANG_Dutch               Language = "Dutch"
	LANG_Flemish             Language = "Flemish"
	LANG_Dzongkha            Language = "Dzongkha"
	LANG_English             Language = "English"
	LANG_Esperanto           Language = "Esperanto"
	LANG_Estonian            Language = "Estonian"
	LANG_Ewe                 Language = "Ewe"
	LANG_Faroese             Language = "Faroese"
	LANG_Fijian              Language = "Fijian"
	LANG_Finnish             Language = "Finnish"
	LANG_French              Language = "French"
	LANG_Fulah               Language = "Fulah"
	LANG_Gaelic              Language = "Gaelic"
	LANG_Scottish_Gaelic     Language = "Scottish Gaelic"
	LANG_Galician            Language = "Galician"
	LANG_Ganda               Language = "Ganda"
	LANG_Georgian            Language = "Georgian"
	LANG_German              Language = "German"
	LANG_Greek               Language = "Greek"
	LANG_Kalaallisut         Language = "Kalaallisut"
	LANG_Greenlandic         Language = "Greenlandic"
	LANG_Guarani             Language = "Guarani"
	LANG_Gujarati            Language = "Gujarati"
	LANG_Haitian             Language = "Haitian"
	LANG_Haitian_Creole      Language = "Haitian"
	LANG_Hausa               Language = "Hausa"
	LANG_Hebrew              Language = "Hebrew"
	LANG_Herero              Language = "Herero"
	LANG_Hindi               Language = "Hindi"
	LANG_Hiri_Motu           Language = "Hiri Motu"
	LANG_Hungarian           Language = "Hungarian"
	LANG_Icelandic           Language = "Icelandic"
	LANG_Ido                 Language = "Ido"
	LANG_Igbo                Language = "Igbo"
	LANG_Indonesian          Language = "Indonesian"
	LANG_Interlingua         Language = "Interlingua"
	LANG_Interlingue         Language = "Interlingue"
	LANG_Occidental          Language = "Interlingue"
	LANG_Inuktitut           Language = "Inuktitut"
	LANG_Inupiaq             Language = "Inupiaq"
	LANG_Irish               Language = "Irish"
	LANG_Italian             Language = "Italian"
	LANG_Japanese            Language = "Japanese"
	LANG_Javanese            Language = "Javanese"
	LANG_Kannada             Language = "Kannada"
	LANG_Kanuri              Language = "Kanuri"
	LANG_Kashmiri            Language = "Kashmiri"
	LANG_Kazakh              Language = "Kazakh"
	LANG_Central_Khmer       Language = "Central Khmer"
	LANG_Kikuyu              Language = "Kikuyu"
	LANG_Gikuyu              Language = "Gikuyu"
	LANG_Kinyarwanda         Language = "Kinyarwanda"
	LANG_Kirghiz             Language = "Kirghiz"
	LANG_Kyrgyz              Language = "Kirghiz"
	LANG_Komi                Language = "Komi"
	LANG_Kongo               Language = "Kongo"
	LANG_Korean              Language = "Korean"
	LANG_Kuanyama            Language = "Kuanyama"
	LANG_Kwanyama            Language = "Kuanyama"
	LANG_Kurdish             Language = "Kurdish"
	LANG_Lao                 Language = "Lao"
	LANG_Latin               Language = "Latin"
	LANG_Latvian             Language = "Latvian"
	LANG_Limburgan           Language = "Limburgan"
	LANG_Limburger           Language = "Limburgan"
	LANG_Limburgish          Language = "Limburgan"
	LANG_Lingala             Language = "Lingala"
	LANG_Lithuanian          Language = "Lithuanian"
	LANG_Luba_Katanga        Language = "Luba-Katanga"
	LANG_Luxembourgish       Language = "Luxembourgish"
	LANG_Letzeburgesch       Language = "Luxembourgish"
	LANG_Macedonian          Language = "Macedonian"
	LANG_Malagasy            Language = "Malagasy"
	LANG_Malay               Language = "Malay"
	LANG_Malayalam           Language = "Malayalam"
	LANG_Maltese             Language = "Maltese"
	LANG_Manx                Language = "Manx"
	LANG_Maori               Language = "Maori"
	LANG_Marathi             Language = "Marathi"
	LANG_Marshallese         Language = "Marshallese"
	LANG_Mongolian           Language = "Mongolian"
	LANG_Nauru               Language = "Nauru"
	LANG_Navajo              Language = "Navajo"
	LANG_Navaho              Language = "Navaho"
	LANG_North_Ndebele       Language = "North Ndebele"
	LANG_South_Ndebele       Language = "South Ndebele"
	LANG_Ndonga              Language = "Ndonga"
	LANG_Nepali              Language = "Nepali"
	LANG_Norwegian           Language = "Norwegian"
	LANG_Norwegian_Bokmal    Language = "Norwegian Bokmål"
	LANG_Norwegian_Nynorsk   Language = "Norwegian Nynorsk"
	LANG_Sichuan_Yi          Language = "Sichuan Yi"
	LANG_Nuosu               Language = "Nuosu"
	LANG_Occitan             Language = "Occitan"
	LANG_Ojibwa              Language = "Ojibwa"
	LANG_Oriya               Language = "Oriya"
	LANG_Oromo               Language = "Oromo"
	LANG_Ossetian            Language = "Ossetian"
	LANG_Ossetic             Language = "Ossetic"
	LANG_Pali                Language = "Pali"
	LANG_Pashto              Language = "Pashto"
	LANG_Pushto              Language = "Pashto"
	LANG_Persian             Language = "Persian"
	LANG_Polish              Language = "Polish"
	LANG_Portuguese          Language = "Portuguese"
	LANG_Punjabi             Language = "Punjabi"
	LANG_Panjabi             Language = "Punjabi"
	LANG_Quechua             Language = "Quechua"
	LANG_Romanian            Language = "Romanian"
	LANG_Moldavian           Language = "Romanian"
	LANG_Moldovan            Language = "Romanian"
	LANG_Romansh             Language = "Romansh"
	LANG_Rundi               Language = "Rundi"
	LANG_Russian             Language = "Russian"
	LANG_Northern_Sami       Language = "Northern Sami"
	LANG_Samoan              Language = "Samoan"
	LANG_Sango               Language = "Sango"
	LANG_Sanskrit            Language = "Sanskrit"
	LANG_Sardinian           Language = "Sardinian"
	LANG_Serbian             Language = "Serbian"
	LANG_Shona               Language = "Shona"
	LANG_Sindhi              Language = "Sindhi"
	LANG_Sinhala             Language = "Sinhala"
	LANG_Sinhalese           Language = "Sinhalese"
	LANG_Slovak              Language = "Slovak"
	LANG_Slovenian           Language = "Slovenian"
	LANG_Somali              Language = "Somali"
	LANG_Southern_Sotho      Language = "Southern Sotho"
	LANG_Spanish             Language = "Spanish"
	LANG_Castilian           Language = "Spanish"
	LANG_Sundanese           Language = "Sundanese"
	LANG_Swahili             Language = "Swahili"
	LANG_Swati               Language = "Swati"
	LANG_Swedish             Language = "Swedish"
	LANG_Tagalog             Language = "Tagalog"
	LANG_Tahitian            Language = "Tahitian"
	LANG_Tajik               Language = "Tajik"
	LANG_Tamil               Language = "Tamil"
	LANG_Tatar               Language = "Tatar"
	LANG_Telugu              Language = "Telugu"
	LANG_Thai                Language = "Thai"
	LANG_Tibetan             Language = "Tibetan"
	LANG_Tigrinya            Language = "Tigrinya"
	LANG_Tonga               Language = "Tonga"
	LANG_Tsonga              Language = "Tsonga"
	LANG_Tswana              Language = "Tswana"
	LANG_Turkish             Language = "Turkish"
	LANG_Turkmen             Language = "Turkmen"
	LANG_Twi                 Language = "Twi"
	LANG_Uighur              Language = "Uighur"
	LANG_Uyghur              Language = "Uighur"
	LANG_Ukrainian           Language = "Ukrainian"
	LANG_Urdu                Language = "Urdu"
	LANG_Uzbek               Language = "Uzbek"
	LANG_Venda               Language = "Venda"
	LANG_Vietnamese          Language = "Vietnamese"
	LANG_Volapuk             Language = "Volapük"
	LANG_Walloon             Language = "Walloon"
	LANG_Welsh               Language = "Welsh"
	LANG_Wolof               Language = "Wolof"
	LANG_Xhosa               Language = "Xhosa"
	LANG_Yiddish             Language = "Yiddish"
	LANG_Yoruba              Language = "Yoruba"
	LANG_Zhuang              Language = "Zhuang"
	LANG_Chuang              Language = "Zhuang"
	LANG_Zulu                Language = "Zulu"
)

var language = map[string]Language{
	"":                LANG_None,
	"Abkhazian":       LANG_Abkhazian,
	"Afar":            LANG_Afar,
	"Afrikaans":       LANG_Afrikaans,
	"Akan":            LANG_Akan,
	"Albanian":        LANG_Albanian,
	"Amharic":         LANG_Amharic,
	"Arabic":          LANG_Arabic,
	"Aragonese":       LANG_Aragonese,
	"Armenian":        LANG_Armenian,
	"Assamese":        LANG_Assamese,
	"Avaric":          LANG_Avaric,
	"Avestan":         LANG_Avestan,
	"Aymara":          LANG_Aymara,
	"Azerbaijani":     LANG_Azerbaijani,
	"Bambara":         LANG_Bambara,
	"Bashkir":         LANG_Bashkir,
	"Basque":          LANG_Basque,
	"Belarusian":      LANG_Belarusian,
	"Bengali":         LANG_Bengali,
	"Bislama":         LANG_Bislama,
	"Bosnian":         LANG_Bosnian,
	"Breton":          LANG_Breton,
	"Bulgarian":       LANG_Bulgarian,
	"Burmese":         LANG_Burmese,
	"Catalan":         LANG_Catalan,
	"Chamorro":        LANG_Chamorro,
	"Chechen":         LANG_Chechen,
	"Chichewa":        LANG_Chichewa,
	"Chinese":         LANG_Chinese,
	"Church Slavonic": LANG_Church_Slavonic,
	"Chuvash":         LANG_Chuvash,
	"Cornish":         LANG_Cornish,
	"Corsican":        LANG_Corsican,
	"Cree":            LANG_Cree,
	"Croatian":        LANG_Croatian,
	"Czech":           LANG_Czech,
	"Danish":          LANG_Danish,
	"Maldivian":       LANG_Maldivian,
	"Dutch":           LANG_Dutch,
	"Flemish":         LANG_Flemish,
	"Dzongkha":        LANG_Dzongkha,
	"English":         LANG_English,
	"Esperanto":       LANG_Esperanto,
	"Estonian":        LANG_Estonian,
	"Ewe":             LANG_Ewe,
	"Faroese":         LANG_Faroese,
	"Fijian":          LANG_Fijian,
	"Finnish":         LANG_Finnish,
	"French":          LANG_French,
	"Fulah":           LANG_Fulah,
	"Gaelic":          LANG_Gaelic,
	"Galician":        LANG_Galician,
	"Ganda":           LANG_Ganda,
	"Georgian":        LANG_Georgian,
	"German":          LANG_German,
	"Greek":           LANG_Greek,
	"Kalaallisut":     LANG_Kalaallisut,
	"Guarani":         LANG_Guarani,
	"Gujarati":        LANG_Gujarati,
	"Haitian":         LANG_Haitian,
	"Hausa":           LANG_Hausa,
	"Hebrew":          LANG_Hebrew,
	"Herero":          LANG_Herero,
	"Hindi":           LANG_Hindi,
	"Hiri Motu":       LANG_Hiri_Motu,
	"Hungarian":       LANG_Hungarian,
	"Icelandic":       LANG_Icelandic,
	"Igbo":            LANG_Ido,
	"Indonesian":      LANG_Indonesian,
	"lang":            LANG,
	"lang":            LANG,
	"lang":            LANG,
	"lang":            LANG,
}

func GetLanguage(language string) Language {

	return languages[strings.ToLower(language)]
}
