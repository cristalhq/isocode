// Package country provides ISO 3166-1 country and territory codes.
package country

import (
	"errors"
	"strings"
)

// Code is an ISO 3166-1 numeric country code. The zero value is invalid.
type Code uint16

// Info is the metadata for a country or territory.
type Info struct {
	Code    Code
	Alpha2  string
	Alpha3  string
	Numeric string // Three digits, including leading zeros.
	Name    string
	Flag    string // Emoji symbol.
}

var (
	// ErrInvalidFormat indicates input other than two or three ASCII letters or three digits.
	ErrInvalidFormat = errors.New("country: invalid code format")

	// ErrUnknownCode indicates a code absent from the ISO 3166-1 table.
	ErrUnknownCode = errors.New("country: unknown code")
)

// Parse accepts alpha-2, alpha-3 (case insensitive), or three-digit numeric codes.
func Parse(s string) (Code, error) {
	if len(s) != 2 && len(s) != 3 {
		return 0, ErrInvalidFormat
	}

	var letters, digits int
	for i := 0; i < len(s); i++ {
		switch b := s[i]; {
		case 'A' <= b && b <= 'Z', 'a' <= b && b <= 'z':
			letters++
		case '0' <= b && b <= '9':
			digits++
		default:
			return 0, ErrInvalidFormat
		}
	}
	if letters != len(s) && digits != 3 {
		return 0, ErrInvalidFormat
	}

	code, ok := byText[strings.ToUpper(s)]
	if !ok {
		return 0, ErrUnknownCode
	}
	return code, nil
}

// Lookup returns metadata for code.
func Lookup(code Code) (Info, bool) {
	info, ok := table[code]
	return info, ok
}

// All returns a fresh slice of supported codes in alpha-2 order.
func All() []Code { return append([]Code(nil), codes[:]...) }

// IsValid reports whether c is present in the table.
func (c Code) IsValid() bool {
	_, ok := Lookup(c)
	return ok
}

// String returns the alpha-2 code, or an empty string if c is invalid.
func (c Code) String() string { return c.Alpha2() }

// Alpha2 returns the two-letter code, or an empty string if c is invalid.
func (c Code) Alpha2() string { return table[c].Alpha2 }

// Alpha3 returns the three-letter code, or an empty string if c is invalid.
func (c Code) Alpha3() string { return table[c].Alpha3 }

// Numeric returns the three-digit code, or an empty string if c is invalid.
func (c Code) Numeric() string { return table[c].Numeric }

// Name returns the country name, or an empty string if c is invalid.
func (c Code) Name() string { return table[c].Name }

// Flag returns the flag emoji sequence, or an empty string if c is invalid.
func (c Code) Flag() string { return table[c].Flag }

// Int returns the underlying numeric value.
func (c Code) Int() int { return int(c) }

func (c Code) MarshalText() ([]byte, error) {
	s := c.Alpha2()
	if s == "" {
		return nil, ErrUnknownCode
	}
	return []byte(s), nil
}

func (c *Code) UnmarshalText(text []byte) error {
	code, err := Parse(string(text))
	if err != nil {
		return err
	}
	*c = code
	return nil
}

var byText = func() map[string]Code {
	m := make(map[string]Code, 3*len(table))
	for code, info := range table {
		m[info.Alpha2] = code
		m[info.Alpha3] = code
		m[info.Numeric] = code
	}
	return m
}()
