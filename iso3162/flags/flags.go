package flags

import (
	"github.com/joaoribeirodasilva/countries_info/iso639/names"
	"github.com/joaoribeirodasilva/countries_info/iso3162"
)

type Flag string

var codes = map[iso3162.ID]Flag{
	iso3162.ID_None:"",
	iso3162.ID_Andorra: ""
}