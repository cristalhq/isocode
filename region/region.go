package region

import (
	"github.com/cristalhq/isocode/country"
	"github.com/cristalhq/isocode/currency"
)

// Currencies returns the ISO 4217 codes listed for c,
// sorted by currency code. It returns nil if there are no matches.
func Currencies(c country.Code) []currency.Code

// Countries returns the countries and territories listed for c,
// sorted by alpha-2 code. It returns nil if there are no matches.
func Countries(c currency.Code) []country.Code
