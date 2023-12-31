// source: https://en.wikipedia.org/wiki/ISO_3166-2
package iso3162_2

import (
	"strings"

	"github.com/joaoribeirodasilva/countries_info/iso3162"
)

// Country code
type ISO3162_2 string

var codes = map[iso3162.ID]ISO3162_2{
	iso3162.ID_None:                         "",
	iso3162.ID_Andorra:                      "AD",
	iso3162.ID_United_Arab_Emirates:         "AE",
	iso3162.ID_Afghanistan:                  "AF",
	iso3162.ID_Antigua_And_Barbuda:          "AG",
	iso3162.ID_Anguilla:                     "AI",
	iso3162.ID_Albania:                      "AL",
	iso3162.ID_Armenia:                      "AM",
	iso3162.ID_Angola:                       "AO",
	iso3162.ID_Antarctica:                   "AQ",
	iso3162.ID_Argentina:                    "AR",
	iso3162.ID_American_Samoa:               "AS",
	iso3162.ID_Austria:                      "AT",
	iso3162.ID_Australia:                    "AU",
	iso3162.ID_Aruba:                        "AW",
	iso3162.ID_Aland_Islands:                "AX",
	iso3162.ID_Azerbaijan:                   "AZ",
	iso3162.ID_Bosnia_And_Herzegovina:       "BA",
	iso3162.ID_Barbados:                     "BB",
	iso3162.ID_Bangladesh:                   "BD",
	iso3162.ID_Belgium:                      "BE",
	iso3162.ID_Burkina_Faso:                 "BF",
	iso3162.ID_Bulgaria:                     "BG",
	iso3162.ID_Bahrain:                      "BH",
	iso3162.ID_Burundi:                      "BI",
	iso3162.ID_Benin:                        "BJ",
	iso3162.ID_Saint_Barthelemy:             "BL",
	iso3162.ID_Bermuda:                      "BM",
	iso3162.ID_Brunei_Darussalam:            "BN",
	iso3162.ID_Bolivia:                      "BO",
	iso3162.ID_Bonaire:                      "BQ",
	iso3162.ID_Brazil:                       "BR",
	iso3162.ID_Bahamas:                      "BS",
	iso3162.ID_Bhutan:                       "BT",
	iso3162.ID_Bouvet_Island:                "BV",
	iso3162.ID_Botswana:                     "BW",
	iso3162.ID_Belarus:                      "BY",
	iso3162.ID_Belize:                       "BZ",
	iso3162.ID_Canada:                       "CA",
	iso3162.ID_Cocos_Islands:                "CC",
	iso3162.ID_Congo_Democratic_Republic_Of: "CD",
	iso3162.ID_Central_African_Republic:     "CF",
	iso3162.ID_Congo:                        "CG",
	iso3162.ID_Switzerland:                  "CH",
	iso3162.ID_Cote_d_Ivoire:                "CI",
	iso3162.ID_Cook_Islands:                 "CK",
	iso3162.ID_Chile:                        "CL",
	iso3162.ID_Cameroon:                     "CM",
	iso3162.ID_China:                        "CN",
	iso3162.ID_Colombia:                     "CO",
	iso3162.ID_Costa_Rica:                   "CR",
	iso3162.ID_Cuba:                         "CU",
	iso3162.ID_Cabo_Verde:                   "CV",
	iso3162.ID_Curacao:                      "CW",
	iso3162.ID_Christmas_Island:             "CX",
	iso3162.ID_Cyprus:                       "CY",
	iso3162.ID_Czechia:                      "CZ",
	iso3162.ID_Germany:                      "DE",
	iso3162.ID_Djibouti:                     "DJ",
	iso3162.ID_Denmark:                      "DK",
	iso3162.ID_Dominica:                     "DM",
	iso3162.ID_Dominican_Republic:           "DO",
	iso3162.ID_Algeria:                      "DZ",
	iso3162.ID_Ecuador:                      "EC",
	iso3162.ID_Estonia:                      "EE",
	iso3162.ID_Egypt:                        "EG",
	iso3162.ID_Western_Sahara:               "EH",
	iso3162.ID_Eritrea:                      "ER",
	iso3162.ID_Spain:                        "ES",
	iso3162.ID_Ethiopia:                     "ET",
	iso3162.ID_Finland:                      "FI",
	iso3162.ID_Fiji:                         "FJ",
	iso3162.ID_Falkland_Islands:             "FK",
	iso3162.ID_Micronesia:                   "FM",
	iso3162.ID_Faroe_Islands:                "FO",
	iso3162.ID_France:                       "FR",
	iso3162.ID_Gabon:                        "GA",
	iso3162.ID_United_Kingdom:               "GB",
	iso3162.ID_Grenada:                      "GD",
	iso3162.ID_Georgia:                      "GE",
	iso3162.ID_French_Guiana:                "GF",
	iso3162.ID_Guernsey:                     "GG",
	iso3162.ID_Ghana:                        "GH",
	iso3162.ID_Gibraltar:                    "GI",
	iso3162.ID_Greenland:                    "GL",
	iso3162.ID_Gambia:                       "GM",
	iso3162.ID_Guinea:                       "GN",
	iso3162.ID_Guadeloupe:                   "GP",
	iso3162.ID_Equatorial_Guinea:            "GQ",
	iso3162.ID_Greece:                       "GR",
	iso3162.ID_South_Georgia_And_South_Sandwich_Islands: "GS",
	iso3162.ID_Guatemala:                         "GT",
	iso3162.ID_Guam:                              "GU",
	iso3162.ID_Guinea_Bissau:                     "GW",
	iso3162.ID_Guyana:                            "GY",
	iso3162.ID_Hong_Kong:                         "HK",
	iso3162.ID_Heard_Island_And_McDonald_Islands: "HM",
	iso3162.ID_Honduras:                          "HN",
	iso3162.ID_Croatia:                           "HR",
	iso3162.ID_Haiti:                             "HT",
	iso3162.ID_Hungary:                           "HU",
	iso3162.ID_Indonesia:                         "ID",
	iso3162.ID_Ireland:                           "IE",
	iso3162.ID_Israel:                            "IL",
	iso3162.ID_Isle_Of_Man:                       "IM",
	iso3162.ID_India:                             "IN",
	iso3162.ID_British_Indian_Ocean_Territory:    "IO",
	iso3162.ID_Iraq:                              "IQ",
	iso3162.ID_Iran:                              "IR",
	iso3162.ID_Iceland:                           "IS",
	iso3162.ID_Italy:                             "IT",
	iso3162.ID_Jersey:                            "JE",
	iso3162.ID_Jamaica:                           "JM",
	iso3162.ID_Jordan:                            "JO",
	iso3162.ID_Japan:                             "JP",
	iso3162.ID_Kenya:                             "KE",
	iso3162.ID_Kyrgyzstan:                        "KG",
	iso3162.ID_Cambodia:                          "KH",
	iso3162.ID_Kiribati:                          "KI",
	iso3162.ID_Comoros:                           "KM",
	iso3162.ID_Saint_Kitts_And_Nevis:             "KN",
	iso3162.ID_North_Korea:                       "KP",
	iso3162.ID_South_Korea:                       "KR",
	iso3162.ID_Kuwait:                            "KW",
	iso3162.ID_Cayman_Islands:                    "KY",
	iso3162.ID_Kazakhstan:                        "KZ",
	iso3162.ID_Lao:                               "LA",
	iso3162.ID_Lebanon:                           "LB",
	iso3162.ID_Saint_Lucia:                       "LC",
	iso3162.ID_Liechtenstein:                     "LI",
	iso3162.ID_Sri_Lanka:                         "LK",
	iso3162.ID_Liberia:                           "LR",
	iso3162.ID_Lesotho:                           "LS",
	iso3162.ID_Lithuania:                         "LT",
	iso3162.ID_Luxembourg:                        "LU",
	iso3162.ID_Latvia:                            "LV",
	iso3162.ID_Libya:                             "LY",
	iso3162.ID_Morocco:                           "MA",
	iso3162.ID_Monaco:                            "MC",
	iso3162.ID_Moldova:                           "MD",
	iso3162.ID_Montenegro:                        "ME",
	iso3162.ID_Saint_Martin_French:               "MF",
	iso3162.ID_Madagascar:                        "MG",
	iso3162.ID_Marshall_Islands:                  "MH",
	iso3162.ID_North_Macedonia:                   "MK",
	iso3162.ID_Mali:                              "ML",
	iso3162.ID_Myanmar:                           "MM",
	iso3162.ID_Mongolia:                          "MN",
	iso3162.ID_Macao:                             "MO",
	iso3162.ID_Northern_Mariana_Islands:          "MP",
	iso3162.ID_Martinique:                        "MQ",
	iso3162.ID_Mauritania:                        "MR",
	iso3162.ID_Montserrat:                        "MS",
	iso3162.ID_Malta:                             "MT",
	iso3162.ID_Mauritius:                         "MU",
	iso3162.ID_Maldives:                          "MV",
	iso3162.ID_Malawi:                            "MW",
	iso3162.ID_Mexico:                            "MX",
	iso3162.ID_Malaysia:                          "MY",
	iso3162.ID_Mozambique:                        "MZ",
	iso3162.ID_Namibia:                           "NA",
	iso3162.ID_New_Caledonia:                     "NC",
	iso3162.ID_Niger:                             "NE",
	iso3162.ID_Norfolk_Island:                    "NF",
	iso3162.ID_Nigeria:                           "NG",
	iso3162.ID_Nicaragua:                         "NI",
	iso3162.ID_Netherlands:                       "NL",
	iso3162.ID_Norway:                            "NO",
	iso3162.ID_Nepal:                             "NP",
	iso3162.ID_Nauru:                             "NR",
	iso3162.ID_Niue:                              "NU",
	iso3162.ID_New_Zealand:                       "NZ",
	iso3162.ID_Oman:                              "OM",
	iso3162.ID_Panama:                            "PA",
	iso3162.ID_Peru:                              "PE",
	iso3162.ID_French_Polynesia:                  "PF",
	iso3162.ID_Papua_New_Guinea:                  "PG",
	iso3162.ID_Philippines:                       "PH",
	iso3162.ID_Pakistan:                          "PK",
	iso3162.ID_Poland:                            "PL",
	iso3162.ID_Saint_Pierre_And_Miquelon:         "PM",
	iso3162.ID_Pitcairn:                          "PN",
	iso3162.ID_Puerto_Rico:                       "PR",
	iso3162.ID_Palestine:                         "PS",
	iso3162.ID_Portugal:                          "PT",
	iso3162.ID_Palau:                             "PW",
	iso3162.ID_Paraguay:                          "PY",
	iso3162.ID_Qatar:                             "QA",
	iso3162.ID_Russia:                            "RU",
	iso3162.ID_Rwanda:                            "RW",
	iso3162.ID_Saudi_Arabia:                      "SA",
	iso3162.ID_Solomon_Islands:                   "SB",
	iso3162.ID_Seychelles:                        "SC",
	iso3162.ID_Sudan:                             "SD",
	iso3162.ID_Sweden:                            "SE",
	iso3162.ID_Singapore:                         "SG",
	iso3162.ID_Saint_Helena_Ascension_And_Tristan_Da_Cunha: "SH",
	iso3162.ID_Slovenia:                             "SL",
	iso3162.ID_Svalbard_And_Jan_Mayen:               "SJ",
	iso3162.ID_Slovakia:                             "SK",
	iso3162.ID_Sierra_Leone:                         "SL",
	iso3162.ID_San_Marino:                           "SM",
	iso3162.ID_Senegal:                              "SN",
	iso3162.ID_Somalia:                              "SO",
	iso3162.ID_Suriname:                             "SR",
	iso3162.ID_South_Sudan:                          "SS",
	iso3162.ID_Sao_Tome_And_Principe:                "ST",
	iso3162.ID_El_Salvador:                          "SV",
	iso3162.ID_Sint_Maarten_Dutch:                   "SX",
	iso3162.ID_Syria:                                "SY",
	iso3162.ID_Eswatini:                             "SZ",
	iso3162.ID_Turks_And_Caicos_Islands:             "TC",
	iso3162.ID_Chad:                                 "TD",
	iso3162.ID_French_Southern_Territories:          "TF",
	iso3162.ID_Togo:                                 "TG",
	iso3162.ID_Thailand:                             "TH",
	iso3162.ID_Tajikistan:                           "TJ",
	iso3162.ID_Tokelau:                              "TK",
	iso3162.ID_Timor_Leste:                          "TL",
	iso3162.ID_Turkmenistan:                         "TM",
	iso3162.ID_Tunisia:                              "TN",
	iso3162.ID_Tonga:                                "TO",
	iso3162.ID_Turkey:                               "TR",
	iso3162.ID_Trinidad_And_Tobago:                  "TT",
	iso3162.ID_Tuvalu:                               "TV",
	iso3162.ID_Taiwan:                               "TW",
	iso3162.ID_Tanzania:                             "UA",
	iso3162.ID_Ukraine:                              "SB",
	iso3162.ID_United_States_Minor_Outlying_Islands: "UM",
	iso3162.ID_United_States:                        "US",
	iso3162.ID_Uruguay:                              "UY",
	iso3162.ID_Uzbekistan:                           "UZ",
	iso3162.ID_Holy_See:                             "VA",
	iso3162.ID_Vatican:                              "VA",
	iso3162.ID_Saint_Vincent_And_The_Grenadines:     "VC",
	iso3162.ID_Venezuela:                            "VE",
	iso3162.ID_Virgin_Islands_British:               "VG",
	iso3162.ID_Virgin_Islands_US:                    "VI",
	iso3162.ID_Vietnam:                              "VN",
	iso3162.ID_Vanuatu:                              "VU",
	iso3162.ID_Wallis_And_Futuna:                    "WF",
	iso3162.ID_Samoa:                                "WS",
	iso3162.ID_Yemen:                                "YE",
	iso3162.ID_Mayotte:                              "YT",
	iso3162.ID_South_Africa:                         "ZA",
	iso3162.ID_Zambia:                               "ZM",
	iso3162.ID_Zimbabwe:                             "ZW",
}

var ids = map[ISO3162_2]iso3162.ID{
	"":   iso3162.ID_None,
	"AD": iso3162.ID_Andorra,
	"AE": iso3162.ID_United_Arab_Emirates,
	"AF": iso3162.ID_Afghanistan,
	"AG": iso3162.ID_Antigua_And_Barbuda,
	"AI": iso3162.ID_Anguilla,
	"AL": iso3162.ID_Albania,
	"AM": iso3162.ID_Armenia,
	"AO": iso3162.ID_Angola,
	"AQ": iso3162.ID_Antarctica,
	"AR": iso3162.ID_Argentina,
	"AS": iso3162.ID_American_Samoa,
	"AT": iso3162.ID_Austria,
	"AU": iso3162.ID_Australia,
	"AW": iso3162.ID_Aruba,
	"AX": iso3162.ID_Aland_Islands,
	"AZ": iso3162.ID_Azerbaijan,
	"BA": iso3162.ID_Bosnia_And_Herzegovina,
	"BB": iso3162.ID_Barbados,
	"BD": iso3162.ID_Bangladesh,
	"BE": iso3162.ID_Belgium,
	"BF": iso3162.ID_Burkina_Faso,
	"BG": iso3162.ID_Bulgaria,
	"BH": iso3162.ID_Bahrain,
	"BI": iso3162.ID_Burundi,
	"BJ": iso3162.ID_Benin,
	"BL": iso3162.ID_Saint_Barthelemy,
	"BM": iso3162.ID_Bermuda,
	"BN": iso3162.ID_Brunei_Darussalam,
	"BO": iso3162.ID_Bolivia,
	"BQ": iso3162.ID_Bonaire,
	"BR": iso3162.ID_Brazil,
	"BS": iso3162.ID_Bahamas,
	"BT": iso3162.ID_Bhutan,
	"BV": iso3162.ID_Bouvet_Island,
	"BW": iso3162.ID_Botswana,
	"BY": iso3162.ID_Belarus,
	"BZ": iso3162.ID_Belize,
	"CA": iso3162.ID_Canada,
	"CC": iso3162.ID_Cocos_Islands,
	"CD": iso3162.ID_Congo_Democratic_Republic_Of,
	"CF": iso3162.ID_Central_African_Republic,
	"CG": iso3162.ID_Congo,
	"CH": iso3162.ID_Switzerland,
	"CI": iso3162.ID_Cote_d_Ivoire,
	"CK": iso3162.ID_Cook_Islands,
	"CL": iso3162.ID_Chile,
	"CM": iso3162.ID_Cameroon,
	"CN": iso3162.ID_China,
	"CO": iso3162.ID_Colombia,
	"CR": iso3162.ID_Costa_Rica,
	"CU": iso3162.ID_Cuba,
	"CV": iso3162.ID_Cabo_Verde,
	"CW": iso3162.ID_Curacao,
	"CX": iso3162.ID_Christmas_Island,
	"CY": iso3162.ID_Cyprus,
	"CZ": iso3162.ID_Czechia,
	"DE": iso3162.ID_Germany,
	"DJ": iso3162.ID_Djibouti,
	"DK": iso3162.ID_Denmark,
	"DM": iso3162.ID_Dominica,
	"DO": iso3162.ID_Dominican_Republic,
	"DZ": iso3162.ID_Algeria,
	"EC": iso3162.ID_Ecuador,
	"EE": iso3162.ID_Estonia,
	"EG": iso3162.ID_Egypt,
	"EH": iso3162.ID_Western_Sahara,
	"ER": iso3162.ID_Eritrea,
	"ES": iso3162.ID_Spain,
	"ET": iso3162.ID_Ethiopia,
	"FI": iso3162.ID_Finland,
	"FJ": iso3162.ID_Fiji,
	"FK": iso3162.ID_Falkland_Islands,
	"FM": iso3162.ID_Micronesia,
	"FO": iso3162.ID_Faroe_Islands,
	"FR": iso3162.ID_France,
	"GA": iso3162.ID_Gabon,
	"GB": iso3162.ID_United_Kingdom,
	"GD": iso3162.ID_Grenada,
	"GE": iso3162.ID_Georgia,
	"GF": iso3162.ID_French_Guiana,
	"GG": iso3162.ID_Guernsey,
	"GH": iso3162.ID_Ghana,
	"GI": iso3162.ID_Gibraltar,
	"GL": iso3162.ID_Greenland,
	"GM": iso3162.ID_Gambia,
	"GN": iso3162.ID_Guinea,
	"GP": iso3162.ID_Guadeloupe,
	"GQ": iso3162.ID_Equatorial_Guinea,
	"GR": iso3162.ID_Greece,
	"GS": iso3162.ID_South_Georgia_And_South_Sandwich_Islands,
	"GT": iso3162.ID_Guatemala,
	"GU": iso3162.ID_Guam,
	"GW": iso3162.ID_Guinea_Bissau,
	"GY": iso3162.ID_Guyana,
	"HK": iso3162.ID_Hong_Kong,
	"HM": iso3162.ID_Heard_Island_And_McDonald_Islands,
	"HN": iso3162.ID_Honduras,
	"HR": iso3162.ID_Croatia,
	"HT": iso3162.ID_Haiti,
	"HU": iso3162.ID_Hungary,
	"ID": iso3162.ID_Indonesia,
	"IE": iso3162.ID_Ireland,
	"IL": iso3162.ID_Israel,
	"IM": iso3162.ID_Isle_Of_Man,
	"N":  iso3162.ID_India,
	"IO": iso3162.ID_British_Indian_Ocean_Territory,
	"IQ": iso3162.ID_Iraq,
	"IR": iso3162.ID_Iran,
	"IS": iso3162.ID_Iceland,
	"IT": iso3162.ID_Italy,
	"JE": iso3162.ID_Jersey,
	"JM": iso3162.ID_Jamaica,
	"JO": iso3162.ID_Jordan,
	"JP": iso3162.ID_Japan,
	"KE": iso3162.ID_Kenya,
	"KG": iso3162.ID_Kyrgyzstan,
	"KH": iso3162.ID_Cambodia,
	"KI": iso3162.ID_Kiribati,
	"KM": iso3162.ID_Comoros,
	"KN": iso3162.ID_Saint_Kitts_And_Nevis,
	"KP": iso3162.ID_North_Korea,
	"KR": iso3162.ID_South_Korea,
	"KW": iso3162.ID_Kuwait,
	"KY": iso3162.ID_Cayman_Islands,
	"KZ": iso3162.ID_Kazakhstan,
	"LA": iso3162.ID_Lao,
	"LB": iso3162.ID_Lebanon,
	"LC": iso3162.ID_Saint_Lucia,
	"LI": iso3162.ID_Liechtenstein,
	"LK": iso3162.ID_Sri_Lanka,
	"LR": iso3162.ID_Liberia,
	"LS": iso3162.ID_Lesotho,
	"LT": iso3162.ID_Lithuania,
	"LU": iso3162.ID_Luxembourg,
	"LV": iso3162.ID_Latvia,
	"LY": iso3162.ID_Libya,
	"MA": iso3162.ID_Morocco,
	"MC": iso3162.ID_Monaco,
	"MD": iso3162.ID_Moldova,
	"ME": iso3162.ID_Montenegro,
	"MF": iso3162.ID_Saint_Martin_French,
	"MG": iso3162.ID_Madagascar,
	"MH": iso3162.ID_Marshall_Islands,
	"MK": iso3162.ID_North_Macedonia,
	"ML": iso3162.ID_Mali,
	"MM": iso3162.ID_Myanmar,
	"MN": iso3162.ID_Mongolia,
	"MO": iso3162.ID_Macao,
	"MP": iso3162.ID_Northern_Mariana_Islands,
	"MQ": iso3162.ID_Martinique,
	"MR": iso3162.ID_Mauritania,
	"MS": iso3162.ID_Montserrat,
	"MT": iso3162.ID_Malta,
	"MU": iso3162.ID_Mauritius,
	"MV": iso3162.ID_Maldives,
	"MW": iso3162.ID_Malawi,
	"MX": iso3162.ID_Mexico,
	"MY": iso3162.ID_Malaysia,
	"MZ": iso3162.ID_Mozambique,
	"NA": iso3162.ID_Namibia,
	"NC": iso3162.ID_New_Caledonia,
	"NE": iso3162.ID_Niger,
	"NF": iso3162.ID_Norfolk_Island,
	"NG": iso3162.ID_Nigeria,
	"NI": iso3162.ID_Nicaragua,
	"NL": iso3162.ID_Netherlands,
	"NO": iso3162.ID_Norway,
	"NP": iso3162.ID_Nepal,
	"NR": iso3162.ID_Nauru,
	"NU": iso3162.ID_Niue,
	"NZ": iso3162.ID_New_Zealand,
	"OM": iso3162.ID_Oman,
	"PA": iso3162.ID_Panama,
	"PE": iso3162.ID_Peru,
	"PF": iso3162.ID_French_Polynesia,
	"PG": iso3162.ID_Papua_New_Guinea,
	"PH": iso3162.ID_Philippines,
	"PK": iso3162.ID_Pakistan,
	"PL": iso3162.ID_Poland,
	"PM": iso3162.ID_Saint_Pierre_And_Miquelon,
	"PN": iso3162.ID_Pitcairn,
	"PR": iso3162.ID_Puerto_Rico,
	"PS": iso3162.ID_Palestine,
	"PT": iso3162.ID_Portugal,
	"PW": iso3162.ID_Palau,
	"PY": iso3162.ID_Paraguay,
	"QA": iso3162.ID_Qatar,
	"RU": iso3162.ID_Russia,
	"RW": iso3162.ID_Rwanda,
	"SA": iso3162.ID_Saudi_Arabia,
	"SB": iso3162.ID_Solomon_Islands,
	"SC": iso3162.ID_Seychelles,
	"SD": iso3162.ID_Sudan,
	"SE": iso3162.ID_Sweden,
	"SG": iso3162.ID_Singapore,
}

func GetID(code string) iso3162.ID {

	val, ok := ids[ISO3162_2(strings.ToLower(code))]
	if !ok {
		return ids[""]
	}
	return val
}

func GetCode(id iso3162.ID) ISO3162_2 {

	return codes[id]
}
