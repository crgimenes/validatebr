package validatebr_test

import (
	"fmt"
	"testing"

	"github.com/crgimenes/validatebr"
)

// ExamplePhoneWithBrazilianAreaCode demonstrates how to validate a Brazilian phone number.
func ExamplePhoneWithBrazilianAreaCode() {
	phones := []string{
		"11987654321",   // valid, 11 digits
		"5511987654321", // valid, country code + DDD + number
		"20123456789",   // invalid DDD
		"11111111111",   // repetitive
	}

	for _, p := range phones {
		fmt.Printf("%s -> %v\n", p, validatebr.PhoneWithBrazilianAreaCode(p))
	}

	// Output:
	// 11987654321 -> true
	// 5511987654321 -> true
	// 20123456789 -> false
	// 11111111111 -> false
}

func TestPhoneWithBrazilianAreaCode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid phone with 11 digits",
			input:    "11987654321",
			expected: true,
		},
		{
			name:     "Valid phone with area code included (13 digits)",
			input:    "5511987654321",
			expected: true,
		},
		{
			name:     "Invalid phone with invalid DDD",
			input:    "20123456789",
			expected: false,
		},
		{
			name:     "Repetitive digits",
			input:    "11111111111",
			expected: false,
		},
		{
			name:     "Less digits than required",
			input:    "1234567890",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validatebr.PhoneWithBrazilianAreaCode(tt.input)
			if result != tt.expected {
				t.Errorf("PhoneWithBrazilianAreaCode(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkPhoneWithBrazilianAreaCodeWithMap(b *testing.B) {
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
				validatebr.PhoneWithBrazilianAreaCodeMap(phone)
			}
		})
	}
}

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
