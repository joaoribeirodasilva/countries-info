package names

import (
	"github.com/joaoribeirodasilva/countries_info/iso639"
)

type LocalName string

var localNames = map[iso639.ID]LocalName{
	iso639.CODE_None:                "",
	iso639.CODE_Abkhazian:           "аҧсуа бызшәа",
	iso639.CODE_Afar:                "ʿAfár af",
	iso639.CODE_Afrikaans:           "Afrikaans",
	iso639.CODE_Akan:                "Akan",
	iso639.CODE_Albanian:            "Shqip",
	iso639.CODE_Amharic:             "ኣማርኛ",
	iso639.CODE_Arabic:              "العربية",
	iso639.CODE_Aragonese:           "Fabla",
	iso639.CODE_Armenian:            "Հայերէն",
	iso639.CODE_Assamese:            "অসমীয়া",
	iso639.CODE_Avaric:              "магIарул мац",
	iso639.CODE_Avestan:             "𐬎𐬞𐬀𐬯𐬙𐬀𐬎𐬎𐬀𐬐𐬀𐬉𐬥𐬀",
	iso639.CODE_Aymara:              "aymar",
	iso639.CODE_Azerbaijani:         "Azərbaycan",
	iso639.CODE_Bambara:             "Bamanankan",
	iso639.CODE_Bashkir:             "башҡорт теле",
	iso639.CODE_Basque:              "Euskara",
	iso639.CODE_Belarusian:          "Беларуская мова",
	iso639.CODE_Bengali:             "বাংলা",
	iso639.CODE_Bislama:             "Bislama",
	iso639.CODE_Bosnian:             "Bosanski",
	iso639.CODE_Breton:              "Ar Brezhoneg",
	iso639.CODE_Bulgarian:           "български",
	iso639.CODE_Burmese:             "ဗမာစကား",
	iso639.CODE_Catalan:             "català",
	iso639.CODE_Valencian:           "català",
	iso639.CODE_Chamorro:            "Chamoru",
	iso639.CODE_Chechen:             "Нохчийн",
	iso639.CODE_Chichewa:            "Chicheŵa",
	iso639.CODE_Chewa:               "Chicheŵa",
	iso639.CODE_Nyanja:              "Chicheŵa",
	iso639.CODE_Chinese:             "中文",
	iso639.CODE_Church_Slavonic:     "Церковнославѧ́нскїй Ѧ҆зы́къ",
	iso639.CODE_Old_Slavonic:        "Церковнославѧ́нскїй Ѧ҆зы́къ",
	iso639.CODE_Old_Church_Slavonic: "Церковнославѧ́нскїй Ѧ҆зы́къ",
	iso639.CODE_Chuvash:             "Çăvašla",
	iso639.CODE_Cornish:             "Kernewek",
	iso639.CODE_Corsican:            "Corsu",
	iso639.CODE_Cree:                "ᓀᐦᐃᔭᐍᐏᐣ",
	iso639.CODE_Croatian:            "Hrvatski",
	iso639.CODE_Czech:               "čeština",
	iso639.CODE_Danish:              "Dansk",
	iso639.CODE_Divehi:              "ދިވެހި",
	iso639.CODE_Dhivehi:             "ދިވެހި",
	iso639.CODE_Maldivian:           "ދިވެހި",
	iso639.CODE_Dutch:               "Nederlands",
	iso639.CODE_Flemish:             "Flemish",
	iso639.CODE_Dzongkha:            "རྫོང་ཁ",
	iso639.CODE_English:             "English",
	iso639.CODE_Esperanto:           "Esperanto",
	iso639.CODE_Estonian:            "Eesti",
	iso639.CODE_Ewe:                 "Eʋegbe",
	iso639.CODE_Faroese:             "Føroyskt",
	iso639.CODE_Fijian:              "Vakaviti",
	iso639.CODE_Finnish:             "Suomi",
	iso639.CODE_French:              "Français",
	iso639.CODE_Western_Frisian:     "Frysk",
	iso639.CODE_Fulah:               "Fulfulde",
	iso639.CODE_Gaelic:              "Gaeilge",
	iso639.CODE_Scottish_Gaelic:     "Gàidhlig",
	iso639.CODE_Galician:            "Galego",
	iso639.CODE_Ganda:               "LùGáànda",
	iso639.CODE_Georgian:            "ქართული",
	iso639.CODE_German:              "Deutsch",
	iso639.CODE_Greek:               "ελληνικά",
	iso639.CODE_Kalaallisut:         "Kalaallisut",
	iso639.CODE_Greenlandic:         "Kalaallisut",
	iso639.CODE_Guarani:             "Avañe’ẽ",
	iso639.CODE_Gujarati:            "ગુજરાતી",
	iso639.CODE_Haitian:             "Ayisyen",
	iso639.CODE_Haitian_Creole:      "Kreyòl Ayisyen",
	iso639.CODE_Hausa:               "حَوْسَ",
	iso639.CODE_Hebrew:              "עברית",
	iso639.CODE_Herero:              "Otjiherero",
	iso639.CODE_Hindi:               "हिन्दी",
	iso639.CODE_Hiri_Motu:           "Motu",
	iso639.CODE_Hungarian:           "Magyar",
	iso639.CODE_Icelandic:           "Íslenska",
	iso639.CODE_Ido:                 "Ido",
	iso639.CODE_Igbo:                "Igbo",
	iso639.CODE_Indonesian:          "Indonesia",
	iso639.CODE_Interlingua:         "Interlingua",
	iso639.CODE_Interlingue:         "Interlingue",
	iso639.CODE_Occidental:          "Interlingue",
	iso639.CODE_Inuktitut:           "ᐃᓄᒃᑎᑐᑦ",
	iso639.CODE_Inupiaq:             "Inupiatun",
	iso639.CODE_Irish:               "Gaeilge",
	iso639.CODE_Italian:             "Italiano",
	iso639.CODE_Japanese:            "日本語",
	iso639.CODE_Javanese:            "Jawa",
	iso639.CODE_Kannada:             "ಕನ್ನಡ",
	iso639.CODE_Kanuri:              "Kanuri",
	iso639.CODE_Kashmiri:            "कॉशुर",
	iso639.CODE_Kazakh:              "Қазақ",
	iso639.CODE_Central_Khmer:       "ភាសាខ្មែរ",
	iso639.CODE_Kikuyu:              "Gĩkũyũ",
	iso639.CODE_Gikuyu:              "Gĩkũyũ",
	iso639.CODE_Kinyarwanda:         "Ikinyarwanda",
	iso639.CODE_Kirghiz:             "Ikinyarwanda",
	iso639.CODE_Kyrgyz:              "Ikinyarwanda",
	iso639.CODE_Komi:                "коми",
	iso639.CODE_Kongo:               "Kikongo",
	iso639.CODE_Korean:              "한국어",
	iso639.CODE_Kuanyama:            "Kuanyama",
	iso639.CODE_Kwanyama:            "Kuanyama",
	iso639.CODE_Kurdish:             "Kurdí",
	iso639.CODE_Lao:                 "ພາສາລາວ",
	iso639.CODE_Latin:               "Latino",
	iso639.CODE_Latvian:             "Latviešu",
	iso639.CODE_Limburgan:           "Lèmburgs",
	iso639.CODE_Limburger:           "Lèmburgs",
	iso639.CODE_Limburgish:          "Lèmburgs",
	iso639.CODE_Lingala:             "Lingála",
	iso639.CODE_Lithuanian:          "Lietuvių",
	iso639.CODE_Luba_Katanga:        "Luba-Katanga",
	iso639.CODE_Luxembourgish:       "Lëtzebuergesch",
	iso639.CODE_Letzeburgesch:       "Lëtzebuergesch",
	iso639.CODE_Macedonian:          "македонски",
	iso639.CODE_Malagasy:            "Malagasy",
	iso639.CODE_Malay:               "Melayu",
	iso639.CODE_Malayalam:           "മലയാളം",
	iso639.CODE_Maltese:             "Malti",
	iso639.CODE_Manx:                "Gaelg",
	iso639.CODE_Maori:               "Māori",
	iso639.CODE_Marathi:             "मराठी",
	iso639.CODE_Marshallese:         "M̧ajeļ",
	iso639.CODE_Mongolian:           "монгол",
	iso639.CODE_Nauru:               "Naoero",
	iso639.CODE_Navajo:              "Diné Bizaad",
	iso639.CODE_Navaho:              "Diné Bizaad",
	iso639.CODE_North_Ndebele:       "Mthwakazi Ndebele",
	iso639.CODE_South_Ndebele:       "Transvaal Ndebele",
	iso639.CODE_Ndonga:              "Ndonga",
	iso639.CODE_Nepali:              "नेपाली",
	iso639.CODE_Norwegian:           "Norsk",
	iso639.CODE_Norwegian_Bokmal:    "Bokmål",
	iso639.CODE_Norwegian_Nynorsk:   "Nynorsk",
	iso639.CODE_Sichuan_Yi:          "ꆇꉙ",
	iso639.CODE_Nuosu:               "ꆇꉙ",
	iso639.CODE_Occitan:             "Occitan",
	iso639.CODE_Ojibwa:              "ᐊᓂᔑᓇᐯᒧᐏᐣ",
	iso639.CODE_Oriya:               "ଓଡ଼ିଆ",
	iso639.CODE_Oromo:               "Oromo",
	iso639.CODE_Ossetian:            "ӕвзаг",
	iso639.CODE_Ossetic:             "ӕвзаг",
	iso639.CODE_Pali:                "पालि",
	iso639.CODE_Pashto:              "پښتو",
	iso639.CODE_Pushto:              "پښتو",
	iso639.CODE_Persian:             "فارسى",
	iso639.CODE_Polish:              "Polski",
	iso639.CODE_Portuguese:          "Português",
	iso639.CODE_Punjabi:             "ਪੰਜਾਬੀ",
	iso639.CODE_Panjabi:             "ਪੰਜਾਬੀ",
	iso639.CODE_Quechua:             "Qhichwa",
	iso639.CODE_Romanian:            "Român",
	iso639.CODE_Moldavian:           "молдовеняскэ",
	iso639.CODE_Moldovan:            "молдовеняскэ",
	iso639.CODE_Romansh:             "Rumantsch",
	iso639.CODE_Rundi:               "Ikirundi",
	iso639.CODE_Russian:             "Русский",
	iso639.CODE_Northern_Sami:       "Davvisámegiella",
	iso639.CODE_Samoan:              "Samoa",
	iso639.CODE_Sango:               "Sängö",
	iso639.CODE_Sanskrit:            "संस्कृतम्",
	iso639.CODE_Sardinian:           "Sardu",
	iso639.CODE_Serbian:             "српски",
	iso639.CODE_Shona:               "chiShona",
	iso639.CODE_Sindhi:              "سنڌي",
	iso639.CODE_Sinhala:             "සිංහල",
	iso639.CODE_Sinhalese:           "සිංහල",
	iso639.CODE_Slovak:              "Slovenčina",
	iso639.CODE_Slovenian:           "Slovenščina",
	iso639.CODE_Somali:              "Soomaali",
	iso639.CODE_Southern_Sotho:      "seSotho",
	iso639.CODE_Spanish:             "Español",
	iso639.CODE_Castilian:           "Español",
	iso639.CODE_Sundanese:           "Sunda",
	iso639.CODE_Swahili:             "Kiswahili",
	iso639.CODE_Swati:               "siSwati",
	iso639.CODE_Swedish:             "Svenska",
	iso639.CODE_Tagalog:             "Tagalog",
	iso639.CODE_Tahitian:            "Tahiti",
	iso639.CODE_Tajik:               "тоҷики",
	iso639.CODE_Tamil:               "தமிழ்",
	iso639.CODE_Tatar:               "татарча",
	iso639.CODE_Telugu:              "తెలుగు",
	iso639.CODE_Thai:                "ภาษาไทย",
	iso639.CODE_Tibetan:             "བོད་སྐད་",
	iso639.CODE_Tigrinya:            "ትግርኛ",
	iso639.CODE_Tonga:               "chiTonga",
	iso639.CODE_Tsonga:              "xiTsonga",
	iso639.CODE_Tswana:              "Setswana",
	iso639.CODE_Turkish:             "Türkçe",
	iso639.CODE_Turkmen:             "түркmенче",
	iso639.CODE_Twi:                 "twi",
	iso639.CODE_Uighur:              "Уйғур",
	iso639.CODE_Uyghur:              "Уйғур",
	iso639.CODE_Ukrainian:           "Українська",
	iso639.CODE_Urdu:                "اردو",
	iso639.CODE_Uzbek:               "o’zbek",
	iso639.CODE_Venda:               "tshiVenḓa",
	iso639.CODE_Vietnamese:          "Tiếng Việt",
	iso639.CODE_Volapuk:             "Volapük",
	iso639.CODE_Walloon:             "Walon",
	iso639.CODE_Welsh:               "Cymraeg",
	iso639.CODE_Wolof:               "Wollof",
	iso639.CODE_Xhosa:               "isiXhosa",
	iso639.CODE_Yiddish:             "ײִדיש",
	iso639.CODE_Yoruba:              "Yorùbá",
	iso639.CODE_Zhuang:              "Vaƅcueŋƅ",
	iso639.CODE_Chuang:              "Vaƅcueŋƅ",
	iso639.CODE_Zulu:                "isiZulu",
}

func GetCode(id iso639.ID) LocalName {

	return localNames[id]
}
