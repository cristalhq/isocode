package country_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/cristalhq/isocode/country"
)

func TestCodes(t *testing.T) {
	codes := country.All()
	if len(codes) != 249 {
		t.Fatalf("got %d codes, want 249", len(codes))
	}

	for i, code := range codes {
		info, ok := country.Lookup(code)
		if !ok || info.Code != code || info.Name == "" || info.Flag == "" {
			t.Fatalf("invalid metadata: %+v", info)
		}

		if i > 0 && codes[i-1].Alpha2() >= info.Alpha2 {
			t.Fatalf("codes out of alpha-2 order at %s", code)
		}

		for _, input := range []string{info.Alpha2, info.Alpha3, info.Numeric} {
			got, err := country.Parse(input)
			if err != nil || got != code {
				t.Errorf("Parse(%q) = %q, %v; want %q", input, got, err, code)
			}
		}
	}
}

func TestParseAndJSON(t *testing.T) {
	for input, want := range map[string]error{
		"04":  country.ErrInvalidFormat,
		"P1":  country.ErrInvalidFormat,
		"PL ": country.ErrInvalidFormat,
		"ZZ":  country.ErrUnknownCode,
		"000": country.ErrUnknownCode,
	} {
		if _, err := country.Parse(input); !errors.Is(err, want) {
			t.Errorf("Parse(%q) returned %v; want %v", input, err, want)
		}
	}

	var code country.Code
	err := json.Unmarshal([]byte(`"ukr"`), &code)
	if err != nil || code != country.UA {
		t.Fatalf("decode Ukraine: %q, %v", code, err)
	}

	if code.Numeric() != "804" || code.Int() != 804 || code.Flag() != "🇺🇦" {
		t.Fatalf("unexpected Ukraine metadata")
	}

	err = json.Unmarshal([]byte(`"ZZ"`), &code)
	if !errors.Is(err, country.ErrUnknownCode) || code != country.UA {
		t.Fatalf("decode invalid code: %q, %v", code, err)
	}

	data, err := json.Marshal(code)
	if err != nil || string(data) != `"UA"` {
		t.Fatalf("encode Ukraine: %s, %v", data, err)
	}

	_, err = json.Marshal(country.Code(65535))
	if !errors.Is(err, country.ErrUnknownCode) {
		t.Fatalf("encode invalid code: %v", err)
	}
}
