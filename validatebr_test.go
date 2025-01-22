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
