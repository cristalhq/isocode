package isocode_test

import (
	"fmt"

	"github.com/cristalhq/isocode"
)

func Example() {
	brave, err := isocode.FromString("ua")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s", brave.Name(), brave.Flag())

	// Output:
	// Ukraine 🇺🇦
}
