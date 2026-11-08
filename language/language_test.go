package language_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/cristalhq/isocode/language"
)

func TestAll(t *testing.T) {
	codes := language.All()
	if len(codes) < 350 {
		t.Fatalf("unexpectedly short code list: %d", len(codes))
	}

	for i, code := range codes {
		if i > 0 && codes[i-1] >= code {
			t.Fatalf("codes out of order at %s", code)
		}
		info, ok := language.Lookup(code)
		if !ok || info.Code != code || info.Name == "" {
			t.Fatalf("invalid metadata for %s: %+v", code, info)
		}

		for _, input := range []string{code.String(), info.Alpha3} {
			got, err := language.Parse(input)
			if err != nil || got != code {
				t.Errorf("Parse(%q) = %q, %v; want %q", input, got, err, code)
			}
		}
	}

	codes[0] = ""
	if language.All()[0] == "" {
		t.Fatal("All exposed shared storage")
	}
}

func TestParse(t *testing.T) {
	for input, want := range map[string]language.Code{
		"PL":  language.PL,
		"pol": language.PL,
		"deu": language.DE,
		"GER": language.DE,
		"sco": language.SCO,
		"zho": language.ZH,
	} {
		got, err := language.Parse(input)
		if err != nil || got != want {
			t.Errorf("Parse(%q) = %q, %v; want %q", input, got, err, want)
		}
	}

	for input, want := range map[string]error{
		"pt-BR": language.ErrInvalidFormat,
		"p1":    language.ErrInvalidFormat,
		"zz":    language.ErrUnknownCode,
		"qaa":   language.ErrUnknownCode,
		"gem":   language.ErrUnknownCode,
		"mul":   language.ErrUnknownCode,
	} {
		if _, err := language.Parse(input); !errors.Is(err, want) {
			t.Errorf("Parse(%q) returned %v; want %v", input, err, want)
		}
	}

	if language.SCO.Alpha2() != "" || language.SCO.Alpha3() != "sco" || language.SCO.Name() != "Scots" {
		t.Fatal("unexpected Scots metadata")
	}
}

func TestJSON(t *testing.T) {
	var code language.Code
	err := json.Unmarshal([]byte(`"GER"`), &code)
	if err != nil || code != language.DE {
		t.Fatalf("decode German: %q, %v", code, err)
	}

	err = json.Unmarshal([]byte(`"zz"`), &code)
	if !errors.Is(err, language.ErrUnknownCode) || code != language.DE {
		t.Fatalf("decode invalid code: %q, %v", code, err)
	}

	data, err := json.Marshal(code)
	if err != nil || string(data) != `"de"` {
		t.Fatalf("encode German: %s, %v", data, err)
	}

	_, err = json.Marshal(language.Code(""))
	if !errors.Is(err, language.ErrUnknownCode) {
		t.Fatalf("encode invalid code: %v", err)
	}
}
