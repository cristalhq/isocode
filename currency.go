package isocode

import (
	"errors"
	"fmt"
	"strings"
)

// AsCurrency takes string representation of ISO 4217 of currency code and returns a Currency.
func AsCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return 0, ErrUnsupportedCurrencyFormat
	}

	c, ok := currenciesMap[strings.ToUpper(code)]
	if !ok {
		return 0, ErrUnknownCurrencyCode
	}
	return c, nil
}

// Currency represents an ISO currency code.
type Currency int16

// String returns string representation of the code.
func (c Currency) String() string { return currencies[c].Code }

// Name returns a currency name.
func (c Currency) Name() string { return currencies[c].Name }

func (c Currency) Int() int { return int(c) }

// Numeric returns a numeric code of the currency.
func (c Currency) Numeric() string { return fmt.Sprintf("%03d", c) }

// Flag returns an emoji flag for the currency code.
func (c Currency) Flag() string { return currencies[c].Flag }

func (c Currency) MarshalText() ([]byte, error) { return []byte(c.String()), nil }

func (c *Currency) UnmarshalText(b []byte) error {
	code, err := AsCurrency(string(b))
	if err != nil {
		return err
	}
	*c = code
	return nil
}

type currencyDetails struct {
	Code    string `json:"code"`
	Numeric int16  `json:"num"`
	Digits  int8   `json:"digits"`
	Name    string `json:"name"`
	Flag    string `json:"flag"`
}

var (
	ErrUnsupportedCurrencyFormat = errors.New("invalid string representation of the currency")

	ErrUnknownCurrencyCode = errors.New("unknown currency code")
)
