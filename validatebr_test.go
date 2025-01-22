package validatebr_test

import (
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
