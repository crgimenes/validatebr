package main

import (
	"fmt"

	"github.com/crgimenes/validatebr"
)

func main() {
	cnpj := "19.JA2.KO8/Z001-51"
	ok := validatebr.CNPJAlphanumeric(cnpj)
	fmt.Printf("CNPJ %s is valid? %v\n", cnpj, ok)

	cnpj = "19.JA2.KO8/Z001-52"
	ok = validatebr.CNPJAlphanumeric(cnpj)
	fmt.Printf("CNPJ %s is valid? %v\n", cnpj, ok)
}
