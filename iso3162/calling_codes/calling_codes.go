// source: https://en.wikipedia.org/wiki/List_of_country_calling_codes
package calling_codes

import (
	"github.com/joaoribeirodasilva/countries_info/iso3162"
)

type CallingCode string

var codes = map[iso3162.ID]CallingCode{
	iso3162.ID_None:                         "",
	iso3162.ID_Andorra:                      "+376",
	iso3162.ID_United_Arab_Emirates:         "+971",
	iso3162.ID_Afghanistan:                  "+93",
	iso3162.ID_Antigua_And_Barbuda:          "+1(268)",
	iso3162.ID_Anguilla:                     "+1(264)",
	iso3162.ID_Albania:                      "+355",
	iso3162.ID_Armenia:                      "+374",
	iso3162.ID_Angola:                       "+244",
	iso3162.ID_Antarctica:                   "",
	iso3162.ID_Argentina:                    "+54",
	iso3162.ID_American_Samoa:               "+1(684)",
	iso3162.ID_Austria:                      "+43",
	iso3162.ID_Australia:                    "+61",
	iso3162.ID_Aruba:                        "+297",
	iso3162.ID_Aland_Islands:                "+358(18)",
	iso3162.ID_Azerbaijan:                   "+994",
	iso3162.ID_Bosnia_And_Herzegovina:       "+387",
	iso3162.ID_Barbados:                     "+1(246)",
	iso3162.ID_Bangladesh:                   "+880",
	iso3162.ID_Belgium:                      "+32",
	iso3162.ID_Burkina_Faso:                 "+226",
	iso3162.ID_Bulgaria:                     "+359",
	iso3162.ID_Bahrain:                      "+973",
	iso3162.ID_Burundi:                      "+257",
	iso3162.ID_Benin:                        "+229",
	iso3162.ID_Saint_Barthelemy:             "+590",
	iso3162.ID_Bermuda:                      "+1(441)",
	iso3162.ID_Brunei_Darussalam:            "+673",
	iso3162.ID_Bolivia:                      "+591",
	iso3162.ID_Bonaire:                      "+599(7)",
	iso3162.ID_Brazil:                       "+55",
	iso3162.ID_Bahamas:                      "+1(242)",
	iso3162.ID_Bhutan:                       "+957",
	iso3162.ID_Bouvet_Island:                "",
	iso3162.ID_Botswana:                     "+267",
	iso3162.ID_Belarus:                      "+375",
	iso3162.ID_Belize:                       "+501",
	iso3162.ID_Canada:                       "+1",
	iso3162.ID_Cocos_Islands:                "+682",
	iso3162.ID_Congo_Democratic_Republic_Of: "+243",
	iso3162.ID_Central_African_Republic:     "+236",
	iso3162.ID_Congo:                        "+242",
	iso3162.ID_Switzerland:                  "+41",
	iso3162.ID_Cote_d_Ivoire:                "+225",
	iso3162.ID_Cook_Islands:                 "+682",
	iso3162.ID_Chile:                        "+56",
	iso3162.ID_Cameroon:                     "+237",
	iso3162.ID_China:                        "+86",
	iso3162.ID_Colombia:                     "+57",
	iso3162.ID_Costa_Rica:                   "+506",
	iso3162.ID_Cuba:                         "+53",
	iso3162.ID_Cabo_Verde:                   "+238",
	iso3162.ID_Curacao:                      "+599(9)",
	iso3162.ID_Christmas_Island:             "+61(89164)",
	iso3162.ID_Cyprus:                       "+357",
	iso3162.ID_Czechia:                      "+420",
	iso3162.ID_Germany:                      "+49",
	iso3162.ID_Djibouti:                     "+253",
	iso3162.ID_Denmark:                      "+45",
	iso3162.ID_Dominica:                     "+1(767)",
	iso3162.ID_Dominican_Republic:           "+1(809,829,849)",
	iso3162.ID_Algeria:                      "+213",
	iso3162.ID_Ecuador:                      "+593",
	iso3162.ID_Estonia:                      "+372",
	iso3162.ID_Egypt:                        "+20",
	iso3162.ID_Western_Sahara:               "+212",
	iso3162.ID_Eritrea:                      "+291",
	iso3162.ID_Spain:                        "+34",
	iso3162.ID_Ethiopia:                     "+251",
	iso3162.ID_Finland:                      "+358",
	iso3162.ID_Fiji:                         "+679",
	iso3162.ID_Falkland_Islands:             "+500",
	iso3162.ID_Micronesia:                   "+691",
	iso3162.ID_Faroe_Islands:                "+298",
	iso3162.ID_France:                       "+33",
	iso3162.ID_Gabon:                        "+241",
	iso3162.ID_United_Kingdom:               "+44",
	iso3162.ID_Grenada:                      "+1(473)",
	iso3162.ID_Georgia:                      "+995",
	iso3162.ID_French_Guiana:                "+594",
	iso3162.ID_Guernsey:                     "+44(1481,7781,7839,7911)",
	iso3162.ID_Ghana:                        "+233",
	iso3162.ID_Gibraltar:                    "+350",
	iso3162.ID_Greenland:                    "+299",
	iso3162.ID_Gambia:                       "+220",
	iso3162.ID_Guinea:                       "+224",
	iso3162.ID_Guadeloupe:                   "+590",
	iso3162.ID_Equatorial_Guinea:            "+240",
	iso3162.ID_Greece:                       "+30",
	iso3162.ID_South_Georgia_And_South_Sandwich_Islands: "+500",
	iso3162.ID_Guatemala:                         "+501",
	iso3162.ID_Guam:                              "+1(671)",
	iso3162.ID_Guinea_Bissau:                     "+245",
	iso3162.ID_Guyana:                            "+592",
	iso3162.ID_Hong_Kong:                         "+852",
	iso3162.ID_Heard_Island_And_McDonald_Islands: "",
	iso3162.ID_Honduras:                          "+504",
	iso3162.ID_Croatia:                           "+385",
	iso3162.ID_Haiti:                             "+509",
	iso3162.ID_Hungary:                           "+36",
	iso3162.ID_Indonesia:                         "+62",
	iso3162.ID_Ireland:                           "+353",
	iso3162.ID_Israel:                            "+972",
	iso3162.ID_Isle_Of_Man:                       "+44(1624,7524,7624,7924)",
	iso3162.ID_India:                             "+91",
	iso3162.ID_British_Indian_Ocean_Territory:    "",
	iso3162.ID_Iraq:                              "+964",
	iso3162.ID_Iran:                              "+98",
	iso3162.ID_Iceland:                           "+354",
	iso3162.ID_Italy:                             "+39",
	iso3162.ID_Jersey:                            "+44(1534)",
	iso3162.ID_Jamaica:                           "+1(658,876)",
	iso3162.ID_Jordan:                            "+962",
	iso3162.ID_Japan:                             "+81",
	iso3162.ID_Kenya:                             "+254",
	iso3162.ID_Kyrgyzstan:                        "+996",
	iso3162.ID_Cambodia:                          "+855",
	iso3162.ID_Kiribati:                          "+686",
	iso3162.ID_Comoros:                           "+269",
	iso3162.ID_Saint_Kitts_And_Nevis:             "+1(869)",
	iso3162.ID_North_Korea:                       "+850",
	iso3162.ID_South_Korea:                       "+82",
	iso3162.ID_Kuwait:                            "+965",
	iso3162.ID_Cayman_Islands:                    "+1(345)",
	iso3162.ID_Kazakhstan:                        "+7(6,7)",
	iso3162.ID_Lao:                               "+856",
	iso3162.ID_Lebanon:                           "+961",
	iso3162.ID_Saint_Lucia:                       "+1(758)",
	iso3162.ID_Liechtenstein:                     "+423",
	iso3162.ID_Sri_Lanka:                         "+94",
	iso3162.ID_Liberia:                           "+231",
	iso3162.ID_Lesotho:                           "+266",
	iso3162.ID_Lithuania:                         "+370",
	iso3162.ID_Luxembourg:                        "+352",
	iso3162.ID_Latvia:                            "+371",
	iso3162.ID_Libya:                             "+218",
	iso3162.ID_Morocco:                           "+212",
	iso3162.ID_Monaco:                            "+377",
	iso3162.ID_Moldova:                           "+373",
	iso3162.ID_Montenegro:                        "+382",
	iso3162.ID_Saint_Martin_French:               "+590",
	iso3162.ID_Madagascar:                        "+261",
	iso3162.ID_Marshall_Islands:                  "+692",
	iso3162.ID_North_Macedonia:                   "+389",
	iso3162.ID_Mali:                              "+223",
	iso3162.ID_Myanmar:                           "+95",
	iso3162.ID_Mongolia:                          "+976",
	iso3162.ID_Macao:                             "+853",
	iso3162.ID_Northern_Mariana_Islands:          "+1(670)",
	iso3162.ID_Martinique:                        "+596",
	iso3162.ID_Mauritania:                        "+222",
	iso3162.ID_Montserrat:                        "+1(664)",
	iso3162.ID_Malta:                             "+356",
	iso3162.ID_Mauritius:                         "+230",
	iso3162.ID_Maldives:                          "+960",
	iso3162.ID_Malawi:                            "+265",
	iso3162.ID_Mexico:                            "+52",
	iso3162.ID_Malaysia:                          "+60",
	iso3162.ID_Mozambique:                        "+258",
	iso3162.ID_Namibia:                           "+264",
	iso3162.ID_New_Caledonia:                     "+687",
	iso3162.ID_Niger:                             "+227",
	iso3162.ID_Norfolk_Island:                    "+672(3)",
	iso3162.ID_Nigeria:                           "+234",
	iso3162.ID_Nicaragua:                         "+505",
	iso3162.ID_Netherlands:                       "+31",
	iso3162.ID_Norway:                            "+47",
	iso3162.ID_Nepal:                             "+977",
	iso3162.ID_Nauru:                             "+674",
	iso3162.ID_Niue:                              "+683",
	iso3162.ID_New_Zealand:                       "+64",
	iso3162.ID_Oman:                              "+968",
	iso3162.ID_Panama:                            "+507",
	iso3162.ID_Peru:                              "+51",
	iso3162.ID_French_Polynesia:                  "+689",
	iso3162.ID_Papua_New_Guinea:                  "+675",
	iso3162.ID_Philippines:                       "+63",
	iso3162.ID_Pakistan:                          "+92",
	iso3162.ID_Poland:                            "+48",
	iso3162.ID_Saint_Pierre_And_Miquelon:         "+508",
	iso3162.ID_Pitcairn:                          "+64",
	iso3162.ID_Puerto_Rico:                       "+1(787,930)",
	iso3162.ID_Palestine:                         "+970",
	iso3162.ID_Portugal:                          "+351",
	iso3162.ID_Palau:                             "+680",
	iso3162.ID_Paraguay:                          "+595",
	iso3162.ID_Qatar:                             "+974",
	iso3162.ID_Russia:                            "+7",
	iso3162.ID_Rwanda:                            "+250",
	iso3162.ID_Saudi_Arabia:                      "+966",
	iso3162.ID_Solomon_Islands:                   "+677",
	iso3162.ID_Seychelles:                        "+248",
	iso3162.ID_Sudan:                             "+249",
	iso3162.ID_Sweden:                            "+46",
	iso3162.ID_Singapore:                         "+65",
	iso3162.ID_Saint_Helena_Ascension_And_Tristan_Da_Cunha: "+290",
	iso3162.ID_Slovenia:                             "+386",
	iso3162.ID_Svalbard_And_Jan_Mayen:               "+47(79)",
	iso3162.ID_Slovakia:                             "+421",
	iso3162.ID_Sierra_Leone:                         "+232",
	iso3162.ID_San_Marino:                           "+378",
	iso3162.ID_Senegal:                              "+221",
	iso3162.ID_Somalia:                              "+252",
	iso3162.ID_Suriname:                             "+597",
	iso3162.ID_South_Sudan:                          "+211",
	iso3162.ID_Sao_Tome_And_Principe:                "+239",
	iso3162.ID_El_Salvador:                          "+503",
	iso3162.ID_Sint_Maarten_Dutch:                   "+1(721)",
	iso3162.ID_Syria:                                "+963",
	iso3162.ID_Eswatini:                             "+268",
	iso3162.ID_Turks_And_Caicos_Islands:             "+1(649)",
	iso3162.ID_Chad:                                 "+235",
	iso3162.ID_French_Southern_Territories:          "",
	iso3162.ID_Togo:                                 "+228",
	iso3162.ID_Thailand:                             "+66",
	iso3162.ID_Tajikistan:                           "+992",
	iso3162.ID_Tokelau:                              "+690",
	iso3162.ID_Timor_Leste:                          "+670",
	iso3162.ID_Turkmenistan:                         "+993",
	iso3162.ID_Tunisia:                              "+216",
	iso3162.ID_Tonga:                                "+676",
	iso3162.ID_Turkey:                               "+90",
	iso3162.ID_Trinidad_And_Tobago:                  "+1(868)",
	iso3162.ID_Tuvalu:                               "+688",
	iso3162.ID_Taiwan:                               "+886",
	iso3162.ID_Tanzania:                             "+255",
	iso3162.ID_Ukraine:                              "+380",
	iso3162.ID_United_States_Minor_Outlying_Islands: "",
	iso3162.ID_United_States:                        "+1",
	iso3162.ID_Uruguay:                              "+598",
	iso3162.ID_Uzbekistan:                           "+998",
	iso3162.ID_Holy_See:                             "+39(06698)",
	iso3162.ID_Vatican:                              "+39(06698)",
	iso3162.ID_Saint_Vincent_And_The_Grenadines:     "+1(784)",
	iso3162.ID_Venezuela:                            "+58",
	iso3162.ID_Virgin_Islands_British:               "",
	iso3162.ID_Virgin_Islands_US:                    "",
	iso3162.ID_Vietnam:                              "+84",
	iso3162.ID_Vanuatu:                              "",
	iso3162.ID_Wallis_And_Futuna:                    "+681",
	iso3162.ID_Samoa:                                "+685",
	iso3162.ID_Yemen:                                "+967",
	iso3162.ID_Mayotte:                              "+262(269,639)",
	iso3162.ID_South_Africa:                         "+27",
	iso3162.ID_Zambia:                               "+260",
	iso3162.ID_Zimbabwe:                             "+263",
}

func GetCallingCode(id iso3162.ID) CallingCode {

	return codes[id]
}
