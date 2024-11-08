package base64_test

import (
	"testing"

	"gosuda.org/go3rd/encode/base64"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"Encode string", "hello", "aGVsbG8="},
		{"Encode []byte", []byte("world"), "d29ybGQ="},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result string
			switch v := tt.input.(type) {
			case string:
				result = base64.Encode(v)
			case []byte:
				result = base64.Encode(v)
			default:
				t.Fatalf("unsupported input type: %T", tt.input)
			}

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestMustDecode(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		shouldPanic bool
	}{
		{"Valid base64 string", "aGVsbG8=", "hello", false},
		{"Invalid base64 string", "invalid-base64", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic but did not panic")
					}
				}()
			}
			result := base64.MustDecode(tt.input)
			if string(result) != tt.expected && !tt.shouldPanic {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{"Valid base64 string", "aGVsbG8=", "hello", false},
		{"Invalid base64 string", "invalid-base64", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := base64.Decode(tt.input)
			if (err != nil) != tt.hasError {
				t.Errorf("expected error: %v, got error: %v", tt.hasError, err)
			}
			if string(result) != tt.expected && !tt.hasError {
				t.Errorf("expected %v, got %v", tt.expected, string(result))
			}
		})
	}
}

func TestIsBase64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid base64 string", "aGVsbG8=", true},
		{"Invalid base64 string", "invalid-base64", false},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := base64.IsBase64(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
