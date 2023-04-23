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
