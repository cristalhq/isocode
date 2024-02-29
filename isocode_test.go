package isocode

import "testing"

func Test(t *testing.T) {
	code := FR

	b, err := code.MarshalText()
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != code.Alpha2() {
		t.Fatalf("have %s want %s", string(b), code.Alpha2())
	}

	var x Country
	if err := x.UnmarshalText(b); err != nil {
		t.Fatal(err)
	}

	if x != code {
		t.Fatalf("have %s want %s", x, code)
	}
}

func TestSanity(t *testing.T) {
	countriesLen := len(countries)
	alpha2ToCountryLen := len(alpha2ToCountry)
	alpha3ToCountryLen := len(alpha3ToCountry)

	if countriesLen != alpha2ToCountryLen {
		t.Fatalf("\nhave: %v\nwant: %v", countriesLen, alpha2ToCountryLen)
	}
	if countriesLen != alpha3ToCountryLen {
		t.Fatalf("\nhave: %v\nwant: %v", countriesLen, alpha3ToCountryLen)
	}

	currenciesLen := len(currencies)
	currenciesMapLen := len(currenciesMap)

	if currenciesLen != currenciesMapLen {
		t.Fatalf("\nhave: %v\nwant: %v", countriesLen, alpha2ToCountryLen)
	}
}
