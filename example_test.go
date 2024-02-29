package isocode_test

import (
	"fmt"

	"github.com/cristalhq/isocode"
)

func ExampleCountry() {
	brave, err := isocode.AsCountry("ua")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s", brave.Name(), brave.Flag())

	// Output:
	// Ukraine 🇺🇦
}
