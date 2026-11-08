package isocode_test

import (
	"encoding/json"
	"fmt"

	"github.com/cristalhq/isocode/country"
	"github.com/cristalhq/isocode/currency"
	"github.com/cristalhq/isocode/language"
)

func Example_country() {
	code, err := country.Parse("ua")
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(code)
	if err != nil {
		panic(err)
	}

	fmt.Println(code.Name(), code.Flag())
	fmt.Println(string(data))

	// Output:
	// Ukraine 🇺🇦
	// "UA"
}

func Example_currency() {
	code, err := currency.Parse("UAH")
	if err != nil {
		panic(err)
	}

	info, ok := currency.Lookup(code)
	if !ok {
		panic("missing currency metadata")
	}

	data, err := json.Marshal(code)
	if err != nil {
		panic(err)
	}

	fmt.Println(info.Name)
	fmt.Println(string(data))

	// Output:
	// Hryvnia
	// "UAH"
}

func Example_language() {
	code, err := language.Parse("ukr")
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(code)
	if err != nil {
		panic(err)
	}

	fmt.Println(code.Name(), code.Alpha3())
	fmt.Println(string(data))

	// Output:
	// Ukrainian ukr
	// "uk"
}
