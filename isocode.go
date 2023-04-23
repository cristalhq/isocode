package isocode

import (
	"errors"
	"fmt"
	"strings"
)

// FromString takes string representation of ISO 3166-1 Alpha2 or Alpha3 country code and returns a Country.
func FromString(code string) (Country, error) {
	var m map[string]Country

	switch len(code) {
	case 2:
		m = alpha2ToCountry
	case 3:
		m = alpha3ToCountry
	default:
		return 0, ErrUnsupportedFormat
	}

	c, ok := m[strings.ToUpper(code)]
	if !ok {
		return 0, ErrUnknownCode
	}
	return c, nil
}

// Country represents an ISO 3166 numeric country code.
type Country uint16

// String returns an Alpha2 string representation of the code.
func (c Country) String() string { return countries[c].Alpha2 }

// Alpha2 returns an Alpha2 string representation of the code.
func (c Country) Alpha2() string { return countries[c].Alpha2 }

// Alpha3 returns an Alpha3 string representation of the code.
func (c Country) Alpha3() string { return countries[c].Alpha3 }

// Name returns a country name related to Country.
func (c Country) Name() string { return countries[c].Name }

// Numeric returns a numeric code of the country.
func (c Country) Numeric() string { return fmt.Sprintf("%03d", c) }

// Flag returns an emoji flag for the country code.
func (c Country) Flag() string { return countries[c].Flag }

func (c Country) Int() int { return int(c) }

func (c Country) MarshalText() ([]byte, error) { return []byte(c.Alpha2()), nil }

func (c *Country) UnmarshalText(b []byte) error {
	code, err := FromString(string(b))
	if err != nil {
		return err
	}
	*c = code
	return nil
}

// countryDetails represents detailed information related to country code.
type countryDetails struct {
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
	Flag   string `json:"flag"`
	Name   string `json:"name"`

	// TODO:
	// sub  Country
}

var (
	ErrUnsupportedFormat = errors.New("invalid string representation of the code")

	ErrUnknownCode = errors.New("unknown country code")
)
