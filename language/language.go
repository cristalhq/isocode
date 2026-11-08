// Package language provides ISO 639 language codes.
package language

import (
	"errors"
	"strings"
)

// Code is a canonical lowercase ISO 639 language code. The zero value is invalid.
type Code string

// Info is the metadata for a language.
type Info struct {
	Code   Code
	Alpha2 string // Empty when no two-letter code is assigned.
	Alpha3 string // The terminology code when two variants are assigned.
	Name   string // English name.
}

var (
	// ErrInvalidFormat indicates input other than two or three ASCII letters.
	ErrInvalidFormat = errors.New("language: invalid code format")

	// ErrUnknownCode indicates a code absent from the language table.
	ErrUnknownCode = errors.New("language: unknown code")
)

// Parse accepts two- or three-letter language codes, case insensitively.
func Parse(s string) (Code, error) {
	if len(s) != 2 && len(s) != 3 {
		return "", ErrInvalidFormat
	}

	for i := 0; i < len(s); i++ {
		b := s[i]
		if !('A' <= b && b <= 'Z' || 'a' <= b && b <= 'z') {
			return "", ErrInvalidFormat
		}
	}

	code, ok := byText[strings.ToLower(s)]
	if !ok {
		return "", ErrUnknownCode
	}
	return code, nil
}

// Lookup returns the metadata for a canonical language code.
func Lookup(code Code) (Info, bool) {
	info, ok := table[code]
	return info, ok
}

// All returns a fresh slice of supported codes in canonical code order.
func All() []Code { return append([]Code(nil), codes[:]...) }

// String returns the underlying language code.
func (c Code) String() string { return string(c) }

// IsValid reports whether c is a canonical code in the table.
func (c Code) IsValid() bool {
	_, ok := Lookup(c)
	return ok
}

// Alpha2 returns the two-letter code, or an empty string if unavailable.
func (c Code) Alpha2() string { return table[c].Alpha2 }

// Alpha3 returns the three-letter code, or an empty string if c is invalid.
func (c Code) Alpha3() string { return table[c].Alpha3 }

// Name returns the English name, or an empty string if c is invalid.
func (c Code) Name() string { return table[c].Name }

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

var byText = func() map[string]Code {
	m := make(map[string]Code, 2*len(table)+len(aliases))
	for code, info := range table {
		m[string(code)] = code
		m[info.Alpha3] = code
	}
	for alias, code := range aliases {
		m[alias] = code
	}
	return m
}()
