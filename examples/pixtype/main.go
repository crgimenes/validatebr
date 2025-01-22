package main

import "github.com/crgimenes/validatebr"

func main() {
	pixKey := "5511999999999"
	types, err := validatebr.PixKeyType(pixKey)
	if err != nil {
		panic(err)
	}

	for _, v := range types {
		println(v)
	}
}
