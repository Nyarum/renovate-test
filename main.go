package main

import (
	"fmt"

	"golang.org/x/text/language"

	"golang.org/x/text/currency"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	fmt.Println(currency.FromTag(language.Amharic))

}
