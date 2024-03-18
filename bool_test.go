package assert

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	tests := []struct {
		name     string
		haystack []string
	}{
		{
			name:     "should be empty",
			haystack: nil,
		},
		{
			name:     "should also be empty",
			haystack: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Empty(t, tt.haystack)
		})
	}
}

func TestFalse(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{
			name:     "should be false",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			False(t, tt.expected)
		})
	}
}

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		haystack []string
	}{
		{
			name:     "should not be empty",
			haystack: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NotEmpty(t, tt.haystack)
		})
	}
}

func TestTrue(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{
			name:     "should be true",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			True(t, tt.expected)
		})
	}
}
