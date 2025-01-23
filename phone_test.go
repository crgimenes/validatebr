package validatebr_test

import (
	"testing"

	"github.com/crgimenes/validatebr"
)

func BenchmarkPhoneWithBrazilianAreaCode(b *testing.B) {
	phones := []string{
		"11987654321",   // valid
		"5511987654321", // valid with country code
		"20123456789",   // invalid DDD
		"11111111111",   // repetitive
		"1234567890",    // less digits
	}

	for _, phone := range phones {
		b.Run(phone, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				validatebr.PhoneWithBrazilianAreaCode(phone)
			}
		})
	}
}

func BenchmarkPhoneWithBrazilianAreaCodeMap(b *testing.B) {
	phones := []string{
		"11987654321",   // valid
		"5511987654321", // valid with country code
		"20123456789",   // invalid DDD
		"11111111111",   // repetitive
		"1234567890",    // less digits
	}

	for _, phone := range phones {
		b.Run(phone, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				validatebr.PhoneWithBrazilianAreaCodeArray(phone)
			}
		})
	}
}
