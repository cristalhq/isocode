// Package currency provides ISO 4217 currency and fund codes.
package currency

import (
	"errors"
	"strings"
)

// Code is an uppercase ISO 4217 alphabetic code. The zero value is invalid.
type Code string

// Info is the ISO 4217 metadata for a code.
type Info struct {
	Code      Code
	Numeric   string // Three digits, including leading zeros.
	Name      string
	MinorUnit int // Decimal places, or -1 if not applicable.
}

var (
	// ErrInvalidFormat indicates input other than three ASCII letters or digits.
	ErrInvalidFormat = errors.New("currency: invalid code format")

	// ErrUnknownCode indicates a code absent from the ISO 4217 table.
	ErrUnknownCode = errors.New("currency: unknown code")
)

// Parse accepts a three-letter code (case insensitive) or a three-digit code.
func Parse(s string) (Code, error) {
	if len(s) != 3 {
		return "", ErrInvalidFormat
	}

	var letters, digits int
	for i := 0; i < len(s); i++ {
		switch b := s[i]; {
		case 'A' <= b && b <= 'Z', 'a' <= b && b <= 'z':
			letters++
		case '0' <= b && b <= '9':
			digits++
		default:
			return "", ErrInvalidFormat
		}
	}

	switch {
	case letters == 3:
		code := Code(strings.ToUpper(s))
		if code.IsValid() {
			return code, nil
		}
	case digits == 3:
		if code, ok := byNumeric[s]; ok {
			return code, nil
		}
	default:
		return "", ErrInvalidFormat
	}
	return "", ErrUnknownCode
}

// Lookup returns the metadata for a canonical alphabetic code.
func Lookup(code Code) (Info, bool) {
	info, ok := table[code]
	return info, ok
}

// All returns a fresh slice of supported codes in alphabetic order.
func All() []Code { return append([]Code(nil), codes[:]...) }

// String returns the underlying alphabetic code.
func (c Code) String() string { return string(c) }

// IsValid reports whether c is a canonical code in the table.
func (c Code) IsValid() bool {
	_, ok := Lookup(c)
	return ok
}

func (c Code) MarshalText() ([]byte, error) {
	if !c.IsValid() {
		return nil, ErrUnknownCode
	}
	return []byte(c), nil
}

func (c *Code) UnmarshalText(text []byte) error {
	code, err := Parse(string(text))
	if err != nil {
		return err
	}
	*c = code
	return nil
}

var byNumeric = func() map[string]Code {
	m := make(map[string]Code, len(table))
	for code, info := range table {
		m[info.Numeric] = code
	}
	return m
}()
