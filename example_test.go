package isocode_test

import (
	"encoding/json"
	"fmt"

	"github.com/cristalhq/isocode"
)

func ExampleCountry() {
	brave, err := isocode.AsCountry("ua")
	if err != nil {
		panic(err)
	}

	fmt.Println(brave.Name(), brave.Flag())
	fmt.Println(toJSON(brave))

	// Output:
	// Ukraine 🇺🇦
	// "UA"
}

func ExampleCurrency() {
	trust, err := isocode.AsCurrency("USD")
	if err != nil {
		panic(err)
	}

	fmt.Println(trust.Name())
	fmt.Println(toJSON(trust))

	// Output:
	// United States dollar
	// "USD"
}

func toJSON(x any) string {
	b, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	return string(b)
}
