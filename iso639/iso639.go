// [Source]https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
package iso639

import "strings"

type Code int
type Names string
const (
	CODE_None                Code = iota
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

var codes = map[string]Names{
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

var names = map[int]string{
	CODE_None                Names = "None",
	CODE_Abkhazian           Names = "Abkhazian",
	CODE_Afar                Names = "Afar",
	CODE_Afrikaans           Names = "Afrikaans",
	CODE_Akan                Names = "Akan",
	CODE_Albanian            Names = "Albanian",
	CODE_Amharic             Names = "Amharic",
	CODE_Arabic              Names = "Arabic",
	CODE_Aragonese           Names = "Aragonese",
	CODE_Armenian            Names = "Armenian",
	CODE_Assamese            Names = "Assamese",
	CODE_Avaric              Names = "Avaric",
	CODE_Avestan             Names = "Avestan",
	CODE_Aymara              Names = "Aymara",
	CODE_Azerbaijani         Names = "Azerbaijani",
	CODE_Bambara             Names = "Bambara",
	CODE_Bashkir             Names = "Bashkir",
	CODE_Basque              Names = "Basque",
	CODE_Belarusian          Names = "Belarusian",
	CODE_Bengali             Names = "Bengali",
	CODE_Bislama             Names = "Bislama",
	CODE_Bosnian             Names = "Bosnian",
	CODE_Breton              Names = "Breton",
	CODE_Bulgarian           Names = "Bulgarian",
	CODE_Burmese             Names = "Burmese",
	CODE_Catalan             Names = "Catalan",
	CODE_Valencian           Names = "Catalan",
	CODE_Chamorro            Names = "Chamorro",
	CODE_Chechen             Names = "Chechen",
	CODE_Chichewa            Names = "Chichewa",
	CODE_Chewa               Names = "Chichewa",
	CODE_Nyanja              Names = "Chichewa",
	CODE_Chinese             Names = "Chinese",
	CODE_Church_Slavonic     Names = "Church Slavonic",
	CODE_Old_Slavonic        Names = "Old Slavonic",
	CODE_Old_Church_Slavonic Names = "Old Church Slavonic",
	CODE_Chuvash             Names = "Chuvash",
	CODE_Cornish             Names = "Cornish",
	CODE_Corsican            Names = "Corsican",
	CODE_Cree                Names = "Cree"
	CODE_Croatian            Names = "Croatian"
	CODE_Czech               Names = "Czech"
	CODE_Danish              Names = "Danish"
	CODE_Divehi              Names = "Divehi"
	CODE_Dhivehi             Names = "Dhivehi"
	CODE_Maldivian           Names = "Maldivian"
	CODE_Dutch               Names = "Dutch"
	CODE_Flemish             Names = "Flemish"
	CODE_Dzongkha            Names = "Dzongkha"
	CODE_English             Names = "English"
	CODE_Esperanto           Names = "Esperanto"
	CODE_Estonian            Names = "Estonian"
	CODE_Ewe                 Names = "Ewe"
	CODE_Faroese             Names = "Faroese"
	CODE_Fijian              Names = "Fijian"
	CODE_Finnish             Names = "Finnish"
	CODE_French              Names = "French"
	CODE_Fulah               Names = "Fulah"
	CODE_Gaelic              Names = "Gaelic"
	CODE_Scottish_Gaelic     Names = "Scottish Gaelic"
	CODE_Galician            Names = "Galician"
	CODE_Ganda               Names = "Ganda"
	CODE_Georgian            Names = "Georgian"
	CODE_German              Names = "German"
	CODE_Greek               Names = "Greek"
	CODE_Kalaallisut         Names = "Kalaallisut"
	CODE_Greenlandic         Names = "Greenlandic"
	CODE_Guarani             Names = "Guarani"
	CODE_Gujarati            Names = "Gujarati"
	CODE_Haitian             Names = "Haitian"
	CODE_Haitian_Creole      Names = "Haitian"
	CODE_Hausa               Names = "Hausa"
	CODE_Hebrew              Names = "Hebrew"
	CODE_Herero              Names = "Herero"
	CODE_Hindi               Names = "Hindi"
	CODE_Hiri_Motu           Names = "Hiri Motu"
	CODE_Hungarian           Names = "Hungarian"
	CODE_Icelandic           Names = "Icelandic"
	CODE_Ido                 Names = "Ido"
	CODE_Igbo                Names = "Igbo"
	CODE_Indonesian          Names = "Indonesian"
	CODE_Interlingua         Names = "Interlingua"
	CODE_Interlingue         Names = "Interlingue"
	CODE_Occidental          Names = "Interlingue"
	CODE_Inuktitut           Names = "Inuktitut"
	CODE_Inupiaq             Names = "Inupiaq"
	CODE_Irish               Names = "Irish"
	CODE_Italian             Names = "Italian"
	CODE_Japanese            Names = "Japanese"
	CODE_Javanese            Names = "Javanese"
	CODE_Kannada             Names = "Kannada"
	CODE_Kanuri              Names = "Kanuri"
	CODE_Kashmiri            Names = "Kashmiri"
	CODE_Kazakh              Names = "Kazakh"
	CODE_Central_Khmer       Names = "Central Khmer"
	CODE_Kikuyu              Names = "Kikuyu"
	CODE_Gikuyu              Names = "Gikuyu"
	CODE_Kinyarwanda         Names = "Kinyarwanda"
	CODE_Kirghiz             Names = "Kirghiz"
	CODE_Kyrgyz              Names = "Kirghiz"
	CODE_Komi                Names = "Komi"
	CODE_Kongo               Names = "Kongo"
	CODE_Korean              Names = "Korean"
	CODE_Kuanyama            Names = "Kuanyama"
	CODE_Kwanyama            Names = "Kuanyama"
	CODE_Kurdish             Names = "Kurdish"
	CODE_Lao                 Names = "Lao"
	CODE_Latin               Names = "Latin"
	CODE_Latvian             Names = "Latvian"
	CODE_Limburgan           Names = "Limburgan"
	CODE_Limburger           Names = "Limburgan"
	CODE_Limburgish          Names = "Limburgan"
	CODE_Lingala             Names = "Lingala"
	CODE_Lithuanian          Names = "Lithuanian"
	CODE_Luba_Katanga        Names = "Luba-Katanga"
	CODE_Luxembourgish       Names = "Luxembourgish"
	CODE_Letzeburgesch       Names = "Luxembourgish"
	CODE_Macedonian          Names = "Macedonian"
	CODE_Malagasy            Names = "Malagasy"
	CODE_Malay               Names = "Malay"
	CODE_Malayalam           Names = "Malayalam"
	CODE_Maltese             Names = "Maltese"
	CODE_Manx                Names = "Manx"
	CODE_Maori               Names = "Maori"
	CODE_Marathi             Names = "Marathi"
	CODE_Marshallese         Names = "Marshallese"
	CODE_Mongolian           Names = "Mongolian"
	CODE_Nauru               Names = "Nauru"
	CODE_Navajo              Names = "Navajo"
	CODE_Navaho              Names = "Navaho"
	CODE_North_Ndebele       Names = "North Ndebele"
	CODE_South_Ndebele       Names = "South Ndebele"
	CODE_Ndonga              Names = "Ndonga"
	CODE_Nepali              Names = "Nepali"
	CODE_Norwegian           Names = "Norwegian"
	CODE_Norwegian_Bokmal    Names = "Norwegian Bokmål"
	CODE_Norwegian_Nynorsk   Names = "Norwegian Nynorsk"
	CODE_Sichuan_Yi          Names = "Sichuan Yi"
	CODE_Nuosu               Names = "Nuosu"
	CODE_Occitan             Names = "Occitan"
	CODE_Ojibwa              Names = "Ojibwa"
	CODE_Oriya               Names = "Oriya"
	CODE_Oromo               Names = "Oromo"
	CODE_Ossetian            Names = "Ossetian"
	CODE_Ossetic             Names = "Ossetic"
	CODE_Pali                Names = "Pali"
	CODE_Pashto              Names = "Pashto"
	CODE_Pushto              Names = "Pashto"
	CODE_Persian             Names = "Persian"
	CODE_Polish              Names = "Polish"
	CODE_Portuguese          Names = "Portuguese"
	CODE_Punjabi             Names = "Punjabi"
	CODE_Panjabi             Names = "Punjabi"
	CODE_Quechua             Names = "Quechua"
	CODE_Romanian            Names = "Romanian"
	CODE_Moldavian           Names = "Romanian"
	CODE_Moldovan            Names = "Romanian"
	CODE_Romansh             Names = "Romansh"
	CODE_Rundi               Names = "Rundi"
	CODE_Russian             Names = "Russian"
	CODE_Northern_Sami       Names = "Northern Sami"
	CODE_Samoan              Names = "Samoan"
	CODE_Sango               Names = "Sango"
	CODE_Sanskrit            Names = "Sanskrit"
	CODE_Sardinian           Names = "Sardinian"
	CODE_Serbian             Names = "Serbian"
	CODE_Shona               Names = "Shona"
	CODE_Sindhi              Names = "Sindhi"
	CODE_Sinhala             Names = "Sinhala"
	CODE_Sinhalese           Names = "Sinhalese"
	CODE_Slovak              Names = "Slovak"
	CODE_Slovenian           Names = "Slovenian"
	CODE_Somali              Names = "Somali"
	CODE_Southern_Sotho      Names = "Southern Sotho"
	CODE_Spanish             Names = "Spanish"
	CODE_Castilian           Names = "Spanish"
	CODE_Sundanese           Names = "Sundanese"
	CODE_Swahili             Names = "Swahili"
	CODE_Swati               Names = "Swati"
	CODE_Swedish             Names = "Swedish"
	CODE_Tagalog             Names = "Tagalog"
	CODE_Tahitian            Names = "Tahitian"
	CODE_Tajik               Names = "Tajik"
	CODE_Tamil               Names = "Tamil"
	CODE_Tatar               Names = "Tatar"
	CODE_Telugu              Names = "Telugu"
	CODE_Thai                Names = "Thai"
	CODE_Tibetan             Names = "Tibetan"
	CODE_Tigrinya            Names = "Tigrinya"
	CODE_Tonga               Names = "Tonga"
	CODE_Tsonga              Names = "Tsonga"
	CODE_Tswana              Names = "Tswana"
	CODE_Turkish             Names = "Turkish"
	CODE_Turkmen             Names = "Turkmen"
	CODE_Twi                 Names = "Twi"
	CODE_Uighur              Names = "Uighur"
	CODE_Uyghur              Names = "Uighur"
	CODE_Ukrainian           Names = "Ukrainian"
	CODE_Urdu                Names = "Urdu"
	CODE_Uzbek               Names = "Uzbek"
	CODE_Venda               Names = "Venda"
	CODE_Vietnamese          Names = "Vietnamese"
	CODE_Volapuk             Names = "Volapük"
	CODE_Walloon             Names = "Walloon"
	CODE_Welsh               Names = "Welsh"
	CODE_Wolof               Names = "Wolof"
	CODE_Xhosa               Names = "Xhosa"
	CODE_Yiddish             Names = "Yiddish"
	CODE_Yoruba              Names = "Yoruba"
	CODE_Zhuang              Names = "Zhuang"
	CODE_Chuang              Names = "Zhuang"
	CODE_Zulu                Names = "Zulu"	
} 

func GetCode(Names string) Code {

	return codes[strings.ToLower(Names)]
}

func GetName(code Code) Names {

	return names[Code]
}
