package currency_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/cristalhq/isocode/currency"
)

func TestCodes(t *testing.T) {
	codes := currency.All()
	if len(codes) < 170 {
		t.Fatalf("unexpectedly short code list: %d", len(codes))
	}

	for i, code := range codes {
		if i > 0 && codes[i-1] >= code {
			t.Fatalf("codes out of order at %s", code)
		}

		info, ok := currency.Lookup(code)
		if !ok || info.Code != code || info.Name == "" {
			t.Fatalf("invalid metadata for %s: %+v", code, info)
		}

		for _, s := range []string{code.String(), info.Numeric} {
			got, err := currency.Parse(s)
			if err != nil || got != code {
				t.Errorf("Parse(%q) = %q, %v; want %q", s, got, err, code)
			}
		}
	}

	codes[0] = ""
	if currency.All()[0] == "" {
		t.Fatal("All exposed shared storage")
	}
}

func TestParseAndJSON(t *testing.T) {
	for input, want := range map[string]error{
		"":    currency.ErrInvalidFormat,
		"1LN": currency.ErrInvalidFormat,
		"PL ": currency.ErrInvalidFormat,
		"ZZZ": currency.ErrUnknownCode,
		"000": currency.ErrUnknownCode,
	} {
		got, err := currency.Parse(input)
		if got != "" || !errors.Is(err, want) {
			t.Errorf("Parse(%q) = %q, %v; want %v", input, got, err, want)
		}
	}

	for code, units := range map[currency.Code]int{
		currency.PLN: 2,
		currency.JPY: 0,
		currency.KWD: 3,
		currency.XAU: -1,
	} {
		info, ok := currency.Lookup(code)
		if !ok || info.MinorUnit != units {
			t.Errorf("Lookup(%q) = %+v, %v; want %d minor unit digits", code, info, ok, units)
		}
	}

	var code currency.Code
	err := json.Unmarshal([]byte(`"pln"`), &code)
	if err != nil || code != currency.PLN {
		t.Fatalf("decode PLN: %q, %v", code, err)
	}

	err = json.Unmarshal([]byte(`"ZZZ"`), &code)
	if !errors.Is(err, currency.ErrUnknownCode) || code != currency.PLN {
		t.Fatalf("decode invalid code: %q, %v", code, err)
	}

	data, err := json.Marshal(code)
	if err != nil || string(data) != `"PLN"` {
		t.Fatalf("encode PLN: %q, %v", data, err)
	}

	_, err = json.Marshal(currency.Code(""))
	if !errors.Is(err, currency.ErrUnknownCode) {
		t.Fatalf("encode zero code: %v", err)
	}
}
