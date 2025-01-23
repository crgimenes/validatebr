package validatebr_test

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/crgimenes/validatebr"
)

func TestIsEmailValid(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "Valid email with standard format",
			email:    "user@example.com",
			expected: true,
		},
		{
			name:     "Valid email with subdomain",
			email:    "user@mail.example.co.uk",
			expected: true,
		},
		{
			name:     "Valid email with plus sign",
			email:    "user+alias@example.com",
			expected: true,
		},
		{
			name:     "Valid email with numbers",
			email:    "user123@example123.com",
			expected: true,
		},
		{
			name:     "Invalid email missing @ symbol",
			email:    "userexample.com",
			expected: false,
		},
		{
			name:     "Invalid email with multiple @ symbols",
			email:    "user@@example.com",
			expected: false,
		},
		{
			name:     "Invalid email with invalid domain",
			email:    "user@example",
			expected: false,
		},
		{
			name:     "Invalid email with special characters",
			email:    "user!#$%&'*+/=?^_`{|}~@example.com",
			expected: false, // Depending on regex, adjust expected
		},
		{
			name:     "Invalid email with spaces",
			email:    "user @example.com",
			expected: false,
		},
		{
			name:     "Empty email string",
			email:    "",
			expected: false,
		},
		{
			name:     "Email with uppercase letters",
			email:    "USER@EXAMPLE.COM",
			expected: true,
		},
		{
			name:     "Email with leading and trailing spaces",
			email:    "  user@example.com  ",
			expected: false, // Assuming spaces are not trimmed
		},
		{
			name:     "Email with invalid TLD",
			email:    "user@example.c",
			expected: false,
		},
		{
			name:     "Email with long TLD",
			email:    "user@example.technology",
			expected: true,
		},
		{
			name:     "Email with dash in domain",
			email:    "user@ex-ample.com",
			expected: true,
		},
		{
			name:     "Email with underscore in local part",
			email:    "user_name@example.com",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validatebr.IsEmailValid(tt.email)
			if result != tt.expected {
				t.Errorf("IsEmailValid(%q) = %v; want %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestIsCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected bool
	}{
		{
			name:     "Valid CNPJ with standard format",
			cnpj:     "12.345.678/0001-95",
			expected: true,
		},
		{
			name:     "Valid CNPJ without punctuation",
			cnpj:     "12345678000195",
			expected: true,
		},
		{
			name:     "Valid CNPJ with mixed punctuation",
			cnpj:     "12.345678/0001-95",
			expected: true,
		},
		{
			name:     "Invalid CNPJ with incorrect length",
			cnpj:     "12.345.678/0001-9",
			expected: false,
		},
		{
			name:     "Invalid CNPJ with letters",
			cnpj:     "12.345.678/0001-9A",
			expected: false,
		},
		{
			name:     "Invalid CNPJ with special characters",
			cnpj:     "12.345.678/0001-!5",
			expected: false,
		},
		{
			name:     "Invalid CNPJ with missing digits",
			cnpj:     "12.345.678/000-95",
			expected: false,
		},
		{
			name:     "Empty CNPJ string",
			cnpj:     "",
			expected: false,
		},
		{
			name:     "CNPJ with extra characters",
			cnpj:     "12.345.678/0001-95abc",
			expected: false,
		},
		{
			name:     "CNPJ with spaces",
			cnpj:     "12.345.678/0001-95 ",
			expected: false, // Assuming spaces are not trimmed
		},
		{
			name:     "Valid CNPJ with leading and trailing spaces",
			cnpj:     " 12.345.678/0001-95 ",
			expected: false, // Assuming spaces are not trimmed
		},
		{
			name:     "Invalid CNPJ with incorrect punctuation placement",
			cnpj:     "123.456.78/0001-95",
			expected: false,
		},
		{
			name:     "Valid CNPJ with all punctuation",
			cnpj:     "00.000.000/0000-00",
			expected: true, // This is formatted correctly but not a valid CNPJ, the function only checks for format
		},
		{
			name:     "Valid CNPJ with maximum digits",
			cnpj:     "99.999.999/9999-99",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validatebr.IsCNPJ(tt.cnpj)
			if result != tt.expected {
				t.Errorf("IsCNPJ(%q) = %v; want %v", tt.cnpj, result, tt.expected)
			}
		})
	}
}

func TestIsCNPJAlpha(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid alphanumeric CNPJ format (uppercase letters)",
			input:    "AB.CDE.FGH/0001-12",
			expected: true,
		},
		{
			name:     "Valid alphanumeric CNPJ format (mixed case)",
			input:    "aB.cDe.fGh/0001-12",
			expected: true,
		},
		{
			name:     "Invalid format (less characters)",
			input:    "AB.CDE.FGH/0001-1",
			expected: false,
		},
		{
			name:     "Invalid format (non-alphanumeric)",
			input:    "AB.CDE.FH@/0001-12",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: false,
		},
		{
			name:     "Missing digits in DV",
			input:    "AB.CDE.FGH/0001-AA",
			expected: false,
		},
		{
			name:     "Spaces not trimmed",
			input:    " AB.CDE.FGH/0001-12 ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validatebr.IsCNPJAlpha(tt.input)
			if result != tt.expected {
				t.Errorf("IsCNPJAlpha(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCNPJAlphanumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid alphanumeric CNPJ",
			input:    "19.JA2.KO8/Z001-51",
			expected: true,
		},
		{
			name:     "Invalid alphanumeric CNPJ",
			input:    "19.JA2.KO8/Z001-52",
			expected: false,
		},
		{
			name:     "Invalid length",
			input:    "12.AB3.ZQ7/123-5",
			expected: false,
		},
		{
			name:     "All characters repetitive",
			input:    "AAAAAAAAAAAAAA",
			expected: false,
		},
		{
			name:     "Contains non-alphanumeric characters",
			input:    "12.AB?.KO8/Z001-51",
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
			result := validatebr.CNPJAlphanumeric(tt.input)
			if result != tt.expected {
				t.Errorf("CNPJAlphanumeric(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsCPF(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid CPF format with punctuation",
			input:    "123.456.789-09",
			expected: true,
		},
		{
			name:     "Valid CPF format without punctuation",
			input:    "12345678909",
			expected: true,
		},
		{
			name:     "Invalid format (less digits)",
			input:    "123.456.789-0",
			expected: false,
		},
		{
			name:     "Invalid format (letters)",
			input:    "123.456.789-0A",
			expected: false,
		},
		{
			name:     "Invalid format (missing part)",
			input:    "123.456-09",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: false,
		},
		{
			name:     "Spaces not trimmed",
			input:    " 123.456.789-09 ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validatebr.IsCPF(tt.input)
			if result != tt.expected {
				t.Errorf("IsCPF(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCPF(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid CPF",
			input:    "529.982.247-25",
			expected: true,
		},
		{
			name:     "Invalid CPF (digits mismatch)",
			input:    "529.982.247-24",
			expected: false,
		},
		{
			name:     "All repetitive digits",
			input:    "111.111.111-11",
			expected: false,
		},
		{
			name:     "Invalid length",
			input:    "123.456.789-000",
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
			result := validatebr.CPF(tt.input)
			if result != tt.expected {
				t.Errorf("CPF(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid CNPJ",
			input:    "12.345.678/0001-95",
			expected: true,
		},
		{
			name:     "Invalid CNPJ",
			input:    "12.345.678/0001-96",
			expected: false,
		},
		{
			name:     "All repetitive digits",
			input:    "11.111.111/1111-11",
			expected: false,
		},
		{
			name:     "Invalid length",
			input:    "12.345.678/0001-9",
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
			result := validatebr.CNPJ(tt.input)
			if result != tt.expected {
				t.Errorf("CNPJ(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveNonDigits(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Digits only",
			input:    "1234567890",
			expected: "1234567890",
		},
		{
			name:     "Mixed alphanumeric",
			input:    "abc123xyz456",
			expected: "123456",
		},
		{
			name:     "Special characters",
			input:    "12#3@4!56",
			expected: "123456",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Spaces included",
			input:    " 1 2 3 ",
			expected: "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validatebr.RemoveNonDigits(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveNonDigits(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveNonAlphaNum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Alphanumeric only",
			input:    "ABC123xyz",
			expected: "ABC123xyz",
		},
		{
			name:     "With punctuation",
			input:    "ABC-123!xyz?",
			expected: "ABC123xyz",
		},
		{
			name:     "Spaces included",
			input:    "A B C 1 2 3",
			expected: "ABC123",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Symbols only",
			input:    "!@#$%^&*()",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validatebr.RemoveNonAlphaNum(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveNonAlphaNum(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsRepetitive(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "All same digits",
			input:    "1111111111",
			expected: true,
		},
		{
			name:     "All same letters",
			input:    "aaaaaa",
			expected: true,
		},
		{
			name:     "Mixed characters",
			input:    "aaab",
			expected: false,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: true,
		},
		{
			name:     "Empty string (edge case)",
			input:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validatebr.IsRepetitive(tt.input)
			if result != tt.expected {
				t.Errorf("IsRepetitive(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
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

func TestPixKeyType(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr error
		wantTypes []string
	}{
		{
			name:      "Valid email",
			input:     "user@example.com",
			expectErr: nil,
			wantTypes: []string{"EMAIL"},
		},
		{
			name:      "Valid CNPJ",
			input:     "12.345.678/0001-95",
			expectErr: nil,
			wantTypes: []string{"CNPJ"},
		},
		{
			name:      "Valid CPF",
			input:     "529.982.247-25",
			expectErr: nil,
			wantTypes: []string{"CPF"},
		},
		{
			name:      "Valid phone number",
			input:     "5511987654321",
			expectErr: nil,
			wantTypes: []string{"PHONE"},
		},
		{
			name:      "Valid EVP (UUID-like length 36)",
			input:     "123e4567-e89b-12d3-a456-426614174000",
			expectErr: nil,
			wantTypes: []string{"EVP"},
		},
		{
			name:      "Multiple valid possibilities (e.g., phone + CPF format)",
			input:     "11987654374", // if it also matches CPF, depends on numbers
			expectErr: nil,
			wantTypes: []string{"CPF", "PHONE"},
		},
		{
			name:      "Invalid input",
			input:     "invalid_input",
			expectErr: validatebr.ErrInvalidPixType,
			wantTypes: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			types, err := validatebr.PixKeyType(tt.input)

			slices.Sort(types)
			slices.Sort(tt.wantTypes)

			if !errors.Is(err, tt.expectErr) {
				t.Fatalf("PixKeyType(%q) error = %v; want %v", tt.input, err, tt.expectErr)
			}
			if tt.expectErr == nil {
				if len(types) != len(tt.wantTypes) {
					t.Errorf("PixKeyType(%q) returned types = %v; want %v", tt.input, types, tt.wantTypes)
				} else {
					for i := range types {
						if types[i] != tt.wantTypes[i] {
							t.Errorf("PixKeyType(%q) = %v; want %v", tt.input, types, tt.wantTypes)
							break
						}
					}
				}
			}
		})
	}
}

// ExampleIsEmailValid demonstrates how to validate emails using IsEmailValid.
func ExampleIsEmailValid() {
	emails := []string{
		"user@example.com",
		"userexample.com", // missing '@'
	}

	for _, e := range emails {
		fmt.Printf("%s -> %v\n", e, validatebr.IsEmailValid(e))
	}

	// Output:
	// user@example.com -> true
	// userexample.com -> false
}

// ExampleIsCNPJ demonstrates how to check if a string matches a CNPJ format.
func ExampleIsCNPJ() {
	cnpjs := []string{
		"12.345.678/0001-95",
		"12345678000195",
		"12.345.678/0001-9",
	}

	for _, c := range cnpjs {
		fmt.Printf("%s -> %v\n", c, validatebr.IsCNPJ(c))
	}

	// Output:
	// 12.345.678/0001-95 -> true
	// 12345678000195 -> true
	// 12.345.678/0001-9 -> false
}

// ExampleIsCNPJAlpha demonstrates how to check if a string matches a CNPJ format
// allowing letters in place of digits.
func ExampleIsCNPJAlpha() {
	inputs := []string{
		"AB.CDE.FGH/0001-12",
		"aB.cDe.fGh/0001-12",
		"AB.CDE.FH@/0001-12", // invalid character '@'
	}

	for _, in := range inputs {
		fmt.Printf("%s -> %v\n", in, validatebr.IsCNPJAlpha(in))
	}

	// Output:
	// AB.CDE.FGH/0001-12 -> true
	// aB.cDe.fGh/0001-12 -> true
	// AB.CDE.FH@/0001-12 -> false
}

// ExampleCNPJAlphanumeric demonstrates how to check for CNPJ in alphanumeric form.
func ExampleCNPJAlphanumeric() {
	inputs := []string{
		"19.JA2.KO8/Z001-51",
		"19.JA2.KO8/Z001-52", // invalid
	}

	for _, in := range inputs {
		fmt.Printf("%s -> %v\n", in, validatebr.CNPJAlphanumeric(in))
	}

	// Output:
	// 19.JA2.KO8/Z001-51 -> true
	// 19.JA2.KO8/Z001-52 -> false
}

// ExampleIsCPF demonstrates how to check if a string has a valid CPF format.
func ExampleIsCPF() {
	cpfs := []string{
		"123.456.789-09",
		"12345678909",
		" 123.456.789-09 ", // spaces not trimmed
	}

	for _, c := range cpfs {
		fmt.Printf("%s -> %v\n", c, validatebr.IsCPF(c))
	}

	// Output:
	// 123.456.789-09 -> true
	// 12345678909 -> true
	//  123.456.789-09  -> false
}

// ExampleCPF demonstrates full CPF validation, including digits check.
func ExampleCPF() {
	cpfs := []string{
		"529.982.247-25",
		"529.982.247-24", // digit mismatch
	}

	for _, c := range cpfs {
		fmt.Printf("%s -> %v\n", c, validatebr.CPF(c))
	}

	// Output:
	// 529.982.247-25 -> true
	// 529.982.247-24 -> false
}

// ExampleCNPJ demonstrates full CNPJ validation, including digit checks.
func ExampleCNPJ() {
	cnpjs := []string{
		"12.345.678/0001-95",
		"12.345.678/0001-96", // digit mismatch
	}

	for _, c := range cnpjs {
		fmt.Printf("%s -> %v\n", c, validatebr.CNPJ(c))
	}

	// Output:
	// 12.345.678/0001-95 -> true
	// 12.345.678/0001-96 -> false
}

// ExampleRemoveNonDigits demonstrates how to remove all non-digit characters from a string.
func ExampleRemoveNonDigits() {
	input := "abc123xyz456"
	output := validatebr.RemoveNonDigits(input)
	fmt.Printf("Before: %s\nAfter:  %s\n", input, output)

	// Output:
	// Before: abc123xyz456
	// After:  123456
}

// ExampleRemoveNonAlphaNum demonstrates how to remove all non-alphanumeric characters from a string.
func ExampleRemoveNonAlphaNum() {
	input := "ABC-123!xyz?"
	output := validatebr.RemoveNonAlphaNum(input)
	fmt.Printf("Before: %s\nAfter:  %s\n", input, output)

	// Output:
	// Before: ABC-123!xyz?
	// After:  ABC123xyz
}

// ExampleIsRepetitive demonstrates how to detect if all characters in a string are the same.
func ExampleIsRepetitive() {
	inputs := []string{
		"1111111",
		"aaaaaaa",
		"abc",
	}

	for _, in := range inputs {
		fmt.Printf("%s -> %v\n", in, validatebr.IsRepetitive(in))
	}

	// Output:
	// 1111111 -> true
	// aaaaaaa -> true
	// abc -> false
}

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

// ExamplePixKeyType demonstrates how to identify one or more valid PIX key types from a given input.
func ExamplePixKeyType() {
	keys := []string{
		"user@example.com",                     // EMAIL
		"12.345.678/0001-95",                   // CNPJ
		"529.982.247-25",                       // CPF
		"5511987654321",                        // PHONE
		"123e4567-e89b-12d3-a456-426614174000", // EVP
		"11987654374",                          // possibly PHONE or CPF if it matches
		"invalid_input",                        // invalid
	}

	for _, k := range keys {
		keyTypes, err := validatebr.PixKeyType(k)
		if err != nil {
			fmt.Printf("%s -> error: %v\n", k, err)
			continue
		}
		fmt.Printf("%s -> %v\n", k, keyTypes)
	}

	// Output:
	// user@example.com -> [EMAIL]
	// 12.345.678/0001-95 -> [CNPJ]
	// 529.982.247-25 -> [CPF]
	// 5511987654321 -> [PHONE]
	// 123e4567-e89b-12d3-a456-426614174000 -> [EVP]
	// 11987654374 -> [CPF PHONE]
	// invalid_input -> error: invalid pix type
}
